package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
	"testing"
)

// hash环结构
type HashRing struct {
	nodes       []int          // 存储的是虚拟节点hash值，这个是存储每一个虚拟节点hash结果的，用于存取文件时快速查找范围
	nodeMap     map[int]string // 存储的是nodes对应的服务器，key是服务器的hash结果，value是服务器名称
	virtualNode int            // 虚拟节点数量，即每个服务器分配的虚拟节点数
	mu          sync.Mutex     // 互斥锁
}

func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		nodeMap:     make(map[int]string),
		virtualNode: replicas,
	}
}

// AddServer 添加服务器
//
// Params：
//
//	-node：服务器名称
func (h *HashRing) AddServer(node string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := 0; i < h.virtualNode; i++ {
		// 对 服务器名称#虚拟节点编号进行hash，即hash(服务器名称#虚拟节点编号)，即hash(serverName#virtualNodeIndex)
		// 得到这个服务器的hash结果，那么小于这个hash值的文件将落到这个服务器上
		key := int(crc32.ChecksumIEEE([]byte(node + "#" + strconv.Itoa(i))))
		h.nodes = append(h.nodes, key)
		h.nodeMap[key] = node // 标识这个hash值对应的服务器
	}
	// 将每一个服务器中每一个虚拟节点的hash进行排序，便于文件存取的时候查找hash范围
	sort.Ints(h.nodes)
}

// RemoveServer 添加服务器
//
// Params：
//
//	-node：服务器名称
func (h *HashRing) RemoveServer(node string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := 0; i < h.virtualNode; i++ {
		key := int(crc32.ChecksumIEEE([]byte(node + "#" + strconv.Itoa(i))))
		delete(h.nodeMap, key)

		// 查找 key 在 h.nodes 中的索引
		idx := sort.SearchInts(h.nodes, key)
		if idx < len(h.nodes) && h.nodes[idx] == key {
			// 删除 idx
			h.nodes = append(h.nodes[:idx], h.nodes[idx+1:]...)
		}
	}
}

// GetServerByFilename 通过文件名获取这个文件要落到哪一个服务器
//
// Params:
//
//   - filename: 文件名称
//
// Returns:
//
//   - string: 服务器名称
func (h *HashRing) GetServerByFilename(filename string) string {
	if len(h.nodes) == 0 {
		return ""
	}
	hash := int(crc32.ChecksumIEEE([]byte(filename)))
	idx := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i] >= hash
	})
	if idx == len(h.nodes) {
		idx = 0
	}
	return h.nodeMap[h.nodes[idx]]
}

func Test111(t *testing.T) {
	ring := NewHashRing(100)
	ring.AddServer("node1-1")
	ring.AddServer("node1-2")
	ring.AddServer("node1-3")

	fmt.Println("file.txt ->", ring.GetServerByFilename("file.txt"))

	// 模拟 node1-1 掉线
	ring.RemoveServer("node1-1")

	fmt.Println("file.txt after node1-1 down ->", ring.GetServerByFilename("file.txt"))

	// 模拟 node1-1 恢复
	ring.AddServer("node1-1")
	fmt.Println("file.txt after node1-1 down ->", ring.GetServerByFilename("file.txt"))

}

// 模拟服务器hash槽分配和文件如何hash到指定的hash槽
// 	hash方法有：
// 		1.传统取模方法
// 		2.一致性hash方法
// 		3.一致性hash➕虚拟节点
// 对于1.传统取模hash的算法存在很大弊端，如果服务器数量变化整个hash结果就变了，所以不适用于cdn的hash
// 对于2.一致性hash方法，一致性hash是将所有服务器hash到(0,2³²-1]范围内（哈希函数输出 32 位整数）
//		但是对于一致性hash方法而言，每当新增或删减节点内的服务器时，会导致每一个服务器负责的hash范围不平均，因此最常用的是一致性hash➕虚拟节点的方法
// 对于3.一致性hash➕虚拟节点，每一个服务器又被分为多个虚拟节点
// 		node1-1服务器有node1-1#0，node1-1#1，node1-1#2...多个虚拟节点
//		这样每一个新增或删减服务器时，影响的区域就小了，可以做到尽力每一个服务器负责的区域平均。

// 一致性hash是：在环上顺时针找到第一个比文件哈希值大的服务器节点，把文件分配给那台服务器

/*
假设现在一个节点里面有4个机器，每一个机器使用两个虚拟节点，hash还的范围是0到80。
hash(node1-1#1)=10属于(0,10]
hash(node1-2#1)=20属于(10,20]
hash(node1-3#1)=30属于(20,30]
hash(node1-4#1)=40属于(30,40]
hash(node1-1#2)=50属于(40,50]
hash(node1-2#2)=60属于(50,60]
hash(node1-3#2)=70属于(60,70]
hash(node1-4#2)=80属于(70,80]

如果此时hash(file.txt) = 66，那么就是分配hash(node1-3#2)的服务器
*/

// 注意对于服务器的hsah是：hash(服务器名称#虚拟节点编号)，即hash(serverName#virtualNodeIndex)
// 注意对于文件的hsah是：hash(文件名)，即hash(filename)
