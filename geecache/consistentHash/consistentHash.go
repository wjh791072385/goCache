package consistentHash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash //可自定义哈希函数
	replicas int
	keys     []int          //sorted
	hashMap  map[int]string //键为hash值，值为结点地址
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			//将下标i也转化为字符串，加入hash值的计算中，保证虚拟结点的唯一性
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	//排序
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	// Binary search for appropriate replica.
	idx := sort.SearchInts(m.keys, hash)

	//当hash值大于所有m.keys[i]时，idx = len(m.keys)，因此要取余
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
