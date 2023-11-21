package consistenthash

import (
	"hash/crc32"
	"sort"
)

type HashFunc func(data []byte) uint32

type NodeMap struct {
	hashFunc       HashFunc
	nodeHashValues []int
	nodeHashMap    map[int]string
}

func NewNodeMap(fn HashFunc) *NodeMap {
	m := &NodeMap{
		hashFunc:    fn,
		nodeHashMap: make(map[int]string),
	}
	if m.hashFunc == nil {
		m.hashFunc = crc32.ChecksumIEEE
	}
	return m
}

func (m *NodeMap) IsEmpty() bool {
	return len(m.nodeHashValues) == 0
}

func (m *NodeMap) AddNode(keys ...string) {
	for _, key := range keys {
		if key == "" {
			continue
		}
		hash := int(m.hashFunc([]byte(key)))
		m.nodeHashValues = append(m.nodeHashValues, hash)
		m.nodeHashMap[hash] = key
	}
	sort.Ints(m.nodeHashValues)
}

func (m *NodeMap) PickNode(key string) string {
	if m.IsEmpty() {
		return ""
	}
	hash := int(m.hashFunc([]byte(key)))
	idx := sort.Search(len(m.nodeHashValues), func(i int) bool {
		return m.nodeHashValues[i] >= hash
	})
	if idx == len(m.nodeHashValues) {
		idx = 0
	}
	return m.nodeHashMap[m.nodeHashValues[idx]]
}
