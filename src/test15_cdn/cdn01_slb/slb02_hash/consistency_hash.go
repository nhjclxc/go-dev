package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// ServerInfo 节点服务器定义，存储着当前节点的服务器信息
type ServerInfo struct {
	Name string
	IP   string
	// 其他信息...
}

// HashRing 节点层级的hash环结构
type HashRing struct {
	nodes       []int               // 存储的是虚拟节点hash值，这个是存储每一个虚拟节点hash结果的，用于存取文件时快速查找范围
	nodeMap     map[int]*ServerInfo // 存储的是nodes对应的服务器，key是服务器的hash结果，value是服务器名称
	virtualNode int                 // 虚拟节点数量，即每个服务器分配的虚拟节点数
	mu          sync.Mutex          // 互斥锁
}

func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		nodeMap:     make(map[int]*ServerInfo),
		virtualNode: replicas,
	}
}

// AddServer 添加服务器
//
// Params：
//
//	-node：服务器名称
func (h *HashRing) AddServer(server *ServerInfo) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := 0; i < h.virtualNode; i++ {
		// 对 服务器名称#虚拟节点编号进行hash，即hash(服务器名称#虚拟节点编号)，即hash(serverName#virtualNodeIndex)
		// 得到这个服务器的hash结果，那么小于这个hash值的文件将落到这个服务器上
		key := int(crc32.ChecksumIEEE([]byte(server.Name + "#" + strconv.Itoa(i))))
		h.nodes = append(h.nodes, key)
		h.nodeMap[key] = server // 标识这个hash值对应的服务器
	}
	// 将每一个服务器中每一个虚拟节点的hash进行排序，便于文件存取的时候查找hash范围
	sort.Ints(h.nodes)
}

// RemoveServer 添加服务器
//
// Params：
//
//	-node：服务器名称
func (h *HashRing) RemoveServer(serverName string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := 0; i < h.virtualNode; i++ {
		key := int(crc32.ChecksumIEEE([]byte(serverName + "#" + strconv.Itoa(i))))
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
func (h *HashRing) GetServerByFilename(filename string) (*ServerInfo, error) {
	if len(h.nodes) == 0 {
		return nil, fmt.Errorf("节点中不存在该服务器")
	}
	hash := int(crc32.ChecksumIEEE([]byte(filename)))
	idx := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i] >= hash
	})
	if idx == len(h.nodes) {
		idx = 0
	}
	return h.nodeMap[h.nodes[idx]], nil
}
