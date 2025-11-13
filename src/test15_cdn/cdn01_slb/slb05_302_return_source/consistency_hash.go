package main

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

type FileData struct {
	filename string
	Content  []byte
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
	fd := FileData{filename: filename, size: int64(rand.Int()), Content: []byte{1, 2, 3, 4, 5, 6}}
	s.FileList = append(s.FileList, &fd)
	return true
}
func (s *Server) SaveFile2(fileData *FileData) bool {
	s.FileList = append(s.FileList, fileData)
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
		return server1, nil, returnSource, err
	}
	fmt.Println("server2 find file", server2.Name, findFile2.filename, findFile2.size)

	// 在其他服务器里面找到了对应的文件，
	// 则先返回该文件，之后在将文件同步到目前该文件应该落在的那个服务器里面，并且同时在当前服务器删除该文件
	n.asyncFile(server2, findFile2, server1)

	// 当前文件应当在哪一个服务器server，
	// 找到的文件数据findFile
	// 是否要回源returnSource
	return server1, findFile2, returnSource, err
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

// ------------------- 回源处理 -------------------

// CDNHandler 处理用户请求
// 回源策略可选
// useRedirect = true → 302/307 重定向
// useRedirect = false → 代理回源并缓存到目标服务器
func (n *Node) CDNHandler(w http.ResponseWriter, r *http.Request, originURL string, useRedirect bool) {
	filename := r.URL.Path[1:] // 假设 URL = /file.mp4

	// 在整个节点内部查找
	targetServer, fileData, returnSourceFlag, err := n.GetFileByFilename(filename)
	if err != nil {
		http.Error(w, "No available server", http.StatusServiceUnavailable)
		return
	}

	if fileData != nil {
		// 文件命中
		w.Write(fileData.Content)
		return
	}
	if !returnSourceFlag {
		http.Error(w, "No returnSourceFlag==false server", http.StatusServiceUnavailable)
		return
	}

	// 3. 回源到源站
	if useRedirect {
		// 302/307 重定向
		http.Redirect(w, r, originURL, http.StatusTemporaryRedirect)
	} else {
		// 代理回源
		resp, err := http.Get(originURL)
		if err != nil {
			http.Error(w, "Failed to fetch from origin", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.StatusCode)

		// 小文件用 io.ReadAll
		// 读取原站返回的数据
		//content, _ := io.ReadAll(resp.Body)

		// 大文件用 io.TeeReader
		// 使用 TeeReader 异步写入缓存
		pr, pw := io.Pipe()
		tee := io.TeeReader(resp.Body, pw)

		// 1. 写回客户端
		//w.Write(content)
		io.Copy(w, tee)

		// 2. 异步保存到目标服务器
		content, _ := io.ReadAll(pr)
		go targetServer.SaveFile2(&FileData{filename: filename, Content: content})
	}
}

/*
回源的两种实现：
1️⃣ 302/307 重定向回源
2️⃣ 代理回源（Fetch / Reverse Proxy）


1️⃣ 302/307 重定向回源
✅ 概念
当 CDN 节点没有目标文件时，直接告诉客户端去源站获取。
	HTTP 状态码：
		302 Found（临时重定向）
		307 Temporary Redirect（保留请求方法，POST/PUT 不会变成 GET）
✅ 流程
	用户请求 CDN URL → CDN 节点查不到文件
	节点返回 302/307 响应，Location 指向源站 URL
	用户浏览器或客户端去源站请求文件
	【客户端 --> CDN --> 302/307 --> 客户端 --> 源站】


2️⃣ 代理回源（Fetch / Reverse Proxy）
✅ 概念
	当 CDN 节点没有目标文件时，节点自己去源站拉取文件，然后返回给客户端。
	对客户端而言，仍然只访问 CDN 节点。
✅ 流程
	用户请求 CDN URL → CDN 节点查不到文件
	节点向源站请求文件
	节点返回文件给客户端，并可缓存
	【客户端 --> CDN节点 --> CDN节点代理去访问源站资源 --> 源站返回资源 --> 资源返回到CDN节点 --> CDN节点异步同步数据到CDN节点并返回文件给客户端 --> 客户端】


| 特性           | 302/307 重定向   | 代理回源       |
| ------------ | ------------- | ---------- |
| **客户端访问源站**  | 是             | 否          |
| **CDN 节点流量** | 小             | 较大         |
| **实现复杂度**    | 简单            | 中等         |
| **缓存能力**     | 无（除非客户端缓存）    | 可缓存、节点内部同步 |
| **延迟**       | 高（客户端需直接访问源站） | 低（节点可以优化）  |
| **安全性**      | 源站 URL暴露      | 源站隐藏       |

*/

// ------------------- 回源处理 -------------------
// 明白，我们可以完全 边读边写，避免一次性读入内存。核心思路是：
// 用 io.TeeReader 或自定义 MultiWriter 同时写到客户端和一个本地缓存 Writer。
// 缓存写入可以直接是目标服务器的 SaveFile 方法（假设它能接收 io.Writer 或我们可以写临时文件再保存）。
// 这样即使是几十 GB 的文件，也只占少量内存缓冲（默认 32KB/64KB）。
// CDNHandler2 处理用户请求
func (n *Node) CDNHandler2(w http.ResponseWriter, r *http.Request, originURL string, useRedirect bool) {
	filename := r.URL.Path[1:] // 假设 URL = /file.mp4

	// 在整个节点内部查找
	targetServer, fileData, returnSourceFlag, err := n.GetFileByFilename(filename)
	if err != nil {
		http.Error(w, "No available server", http.StatusServiceUnavailable)
		return
	}

	if fileData != nil {
		// 文件命中
		w.Write(fileData.Content)
		return
	}
	if !returnSourceFlag {
		http.Error(w, "No returnSourceFlag==false server", http.StatusServiceUnavailable)
		return
	}

	// 3. 回源到源站
	if useRedirect {
		http.Redirect(w, r, originURL, http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get(originURL)
	if err != nil {
		http.Error(w, "Failed to fetch from origin", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 设置响应头
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)

	// 创建一个临时文件用于边缓存
	tmpFile, err := os.CreateTemp("", "cdn-cache-*")
	if err != nil {
		http.Error(w, "Failed to create cache file", http.StatusInternalServerError)
		return
	}
	defer tmpFile.Close()
	tmpFileName := tmpFile.Name()

	// MultiWriter 同时写到客户端和临时缓存文件
	writer := io.MultiWriter(w, tmpFile)

	// 边读边写
	buf := make([]byte, 8*1024*1024) // 8MB 缓冲
	for {
		nRead, readErr := resp.Body.Read(buf)
		if nRead > 0 {
			_, writeErr := writer.Write(buf[:nRead])
			if writeErr != nil {
				// 写客户端失败直接返回
				break
			}
		}
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				// 写到文件末尾了正常退出
				break
			} else {
				// 网络读取错误
				break
			}
		}
	}

	// 异步保存到目标服务器
	go func(tmpFileName, filename string) {
		data, err := os.ReadFile(tmpFileName)
		if err == nil {
			targetServer.SaveFile2(&FileData{
				filename: filename,
				Content:  data,
			})
		}
		os.Remove(tmpFileName)
	}(tmpFileName, filename)
}
