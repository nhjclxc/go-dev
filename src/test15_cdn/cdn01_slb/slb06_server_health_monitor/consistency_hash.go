package slb06_server_health_monitor

import (
	"context"
)

//
// 每一个节点node实现节点内部每一台服务器的健康监测，并且服务器恢复上线后自动加入
//

// Server 节点服务器定义，存储着当前节点的服务器信息
type Server struct {
	Name string
	IP   string
	// 其他信息...
}

func NewServer(name, ip string) *Server {
	return &Server{
		Name: name,
		IP:   ip,
	}
}

type Node struct {
	ctx       context.Context
	ServerMap map[string]*Server
}

func NewNode() *Node {
	node := Node{
		ServerMap: make(map[string]*Server),
	}
	return &node
}
