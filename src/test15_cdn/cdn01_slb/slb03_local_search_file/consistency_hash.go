package slb03_local_search_file

import (
	"context"
	"fmt"
	"hash/crc32"
	"math/rand/v2"
	"sort"
	"strconv"
	"sync"
	"time"
)

type FileData struct {
	filename string
	size     int64
	// ...
}

// Server 节点服务器定义，存储着当前节点的服务器信息
type Server struct {
	Name     string
	IP       string
	FileList []*FileData // 模拟服务器中的文件列表
	// 其他信息...
}

func NewServer(name, ip string) *Server {
	return &Server{
		Name:     name,
		IP:       ip,
		FileList: make([]*FileData, 0),
	}
}
func (s *Server) SaveFile(filename string) bool {
	fd := FileData{filename: filename, size: int64(rand.Int())}
	s.FileList = append(s.FileList, &fd)
	return true
}

func (s *Server) FindFile(filename string) (*FileData, error) {
	for _, file := range s.FileList {
		if filename == file.filename {
			return file, nil
		}
	}
	return nil, fmt.Errorf("未找到该文件")
}

// HashRing 节点层级的hash环结构
type HashRing struct {
	nodes       []int           // 存储的是虚拟节点hash值，这个是存储每一个虚拟节点hash结果的，用于存取文件时快速查找范围
	nodeMap     map[int]*Server // 存储的是nodes对应的服务器，key是服务器的hash结果，value是服务器名称
	virtualNode int             // 虚拟节点数量，即每个服务器分配的虚拟节点数
	mu          sync.Mutex      // 互斥锁
}

func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		nodeMap:     make(map[int]*Server),
		virtualNode: replicas,
	}
}

// AddServer 添加服务器
//
// Params：
//
//	-node：服务器名称
func (h *HashRing) AddServer(server *Server) {
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
func (h *HashRing) GetServerByFilename(filename string) (*Server, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
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

type Node struct {
	ctx       context.Context
	ServerMap map[string]*Server
	HashRing  *HashRing

	// 节点内部兄弟节点同步文件的通道
	fileSyncQueue chan *SyncTask
	wg            sync.WaitGroup // 等待所有worker结束
}
type SyncTask struct {
	originServer *Server
	fileData     *FileData
	targetServer *Server
}

func NewNode(ctx context.Context, replicas int) *Node {
	node := Node{
		ctx:           ctx,
		ServerMap:     make(map[string]*Server),
		HashRing:      NewHashRing(replicas),
		fileSyncQueue: make(chan *SyncTask, 10),
	}
	node.startSyncTask()
	return &node
}
func (n *Node) AddServer(server *Server) {
	n.HashRing.AddServer(server)
	n.ServerMap[server.Name] = server
}
func (n *Node) FindServerByFilename(filename string) (*Server, error) {
	return n.HashRing.GetServerByFilename(filename)
}
func (n *Node) RemoveServer(serverName string) {
	n.HashRing.RemoveServer(serverName)
}

func (n *Node) FindFileOnOtherServer(file string, originServer *Server) (*Server, *FileData, error) {
	for _, server := range n.ServerMap {
		if originServer.Name == server.Name {
			continue
		}

		fileData, err := server.FindFile(file)
		if err != nil {
			continue
		}
		return server, fileData, err
	}

	return nil, nil, nil
}

// GetFileByFilename 将 查找目标服务器，目标服务器查文件，目标服务器没有对应文件则在node内部查找功能整合
func (n *Node) GetFileByFilename(filename string) (*Server, *FileData, bool, error) {
	// 1、查找目标服务器
	server1, err := n.FindServerByFilename(filename)
	if err != nil {
		fmt.Println("server1 find file -> err", err.Error())
	}
	fmt.Println("server1 find file  ->", server1.Name)

	// 2、目标服务器查文件
	findFile, err := server1.FindFile(filename)
	if err != nil {
		fmt.Println("server1 find file error", err)
	}
	fmt.Println("server1 find file", findFile)

	// 3、目标服务器没有对应文件则在node内部查找
	returnSource := false
	server2, findFile2, err := n.FindFileOnOtherServer(filename, server1)
	if err != nil {
		fmt.Println("server2 find file err", err)
		fmt.Println("node内未找到该文件，触发302回源")
		returnSource = true
		return nil, nil, returnSource, err
	}
	fmt.Println("server2 find file", server2.Name, findFile2.filename, findFile2.size)

	// 在其他服务器里面找到了对应的文件，
	// 则先返回该文件，之后在将文件同步到目前该文件应该落在的那个服务器里面，并且同时在当前服务器删除该文件
	n.asyncFile(server2, findFile2, server1)

	return server2, findFile2, returnSource, err
}

// asyncFile 在服务器内部异步同步文件
func (n *Node) asyncFile(originServer *Server, filedata *FileData, targetServer *Server) {
	//异步方式1:开一个协程去写
	//go n.syncFile(originServer, filedata, targetServer)

	// 异步方式2:使用chan完成
	n.fileSyncQueue <- &SyncTask{
		originServer: originServer,
		fileData:     filedata,
		targetServer: targetServer,
	}
}

// syncFile 在服务器内部异步同步文件
func (n *Node) syncFile(originServer *Server, filedata *FileData, targetServer *Server) {
	// 模拟内网两个服务器之间同步文件
	fmt.Println("targetServer文件同步中.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("targetServer文件同步中..")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("targetServer文件同步中...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("targetServer文件同步完成")

	// 在当前服务器内删除该文件
	fmt.Println("originServer删除文件中.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("originServer文件同步完成")

}

func (n *Node) startSyncTask() {
	fmt.Println("开始异步同步任务...")
	go func() {
		for {
			select {
			case <-n.ctx.Done():
				fmt.Println("程序被关闭，退出异步同步任务")
				return
			case syncTask := <-n.fileSyncQueue:
				n.syncFile(syncTask.originServer, syncTask.fileData, syncTask.targetServer)
			}
		}
	}()
}

// 启动异步任务worker
func (n *Node) startSyncTask2() {
	fmt.Println("开始异步同步任务...")

	n.wg.Add(1)
	go func() {
		defer n.wg.Done()
		for {
			select {
			case <-n.ctx.Done():
				fmt.Println("收到关闭信号，worker退出")
				return
			case syncTask, ok := <-n.fileSyncQueue:
				if !ok {
					fmt.Println("任务队列已关闭，worker退出")
					return
				}
				if syncTask == nil {
					continue
				}
				n.asyncFile2(syncTask.originServer, syncTask.fileData, syncTask.targetServer)
			}
		}
	}()
}

// 提交异步任务
func (n *Node) asyncFile2(originServer *Server, filedata *FileData, targetServer *Server) {
	select {
	case <-n.ctx.Done():
		fmt.Println("系统已关闭，拒绝新任务")
		return
	case n.fileSyncQueue <- &SyncTask{
		originServer: originServer,
		fileData:     filedata,
		targetServer: targetServer,
	}:
		fmt.Printf("[提交任务] %s → %s 文件: %s\n", originServer.Name, targetServer.Name, filedata.filename)
	}
}

// 模拟文件同步操作
func (n *Node) syncFile2(originServer *Server, filedata *FileData, targetServer *Server) {
	fmt.Printf("[同步中] %s -> %s 文件: %s\n", originServer.Name, targetServer.Name, filedata.filename)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("[同步完成] 文件 %s 已同步到 %s\n", filedata.filename, targetServer.Name)
}

// Stop 停止节点：安全关闭通道+等待任务完成
func (n *Node) Stop() {
	fmt.Println("Node.Stop(): 正在关闭...")
	close(n.fileSyncQueue) // 关闭任务通道
	n.wg.Wait()            // 等待所有worker退出
	fmt.Println("Node已安全关闭 ✅")
}
