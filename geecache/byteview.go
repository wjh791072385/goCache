package geecache

// ByteView 一个只读的数据结构，用来表示缓存值
type ByteView struct {
	b []byte //b不对外暴露
}

// Len Lru.Cache中要求被缓存value需要实现Len()
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
// 仅仅返回拷贝，防止缓存值背外部修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
