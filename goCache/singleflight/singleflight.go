package singleflight

import "sync"

// call 代表正在进行中，或已经结束的请求。使用 sync.WaitGroup 锁避免重入
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group 管理不同 key 的请求(call)。
type Group struct {
	//mu 是保护 Group 的成员变量 m 不被并发读写而加上的锁
	mu sync.Mutex

	// 这里必须保存*call指针类型，因为要确保多个重复阻塞的请求得到的g.m[key]指向同一地址
	// 那么当第一个请求设置g.m[key]的值后，阻塞的请求返回的结果也是设置之后的值
	m map[string]*call
}

// Do 针对相同的 key，只创建一个实体，无论 Do 被调用多少次，函数 fn 都只会被调用一次，等待 fn 调用结束了，返回返回值或错误。
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	// 如果key请求已经存在
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()

		//阻塞完成之后返回
		//返回的c.val  c.err是由第一个请求更新过后的
		return c.val, c.err
	}

	// 如果key请求不存在
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	//释放阻塞在c.wg.wait()处的重复请求
	c.wg.Done()

	//请求已完成，从Group.m中删除该请求
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
