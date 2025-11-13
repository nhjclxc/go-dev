package main

import (
	"fmt"
	"testing"
)

func TestConsistencyHash(t *testing.T) {

	ring := NewHashRing(100)
	ring.AddServer(&ServerInfo{Name: "node1-1", IP: "192.168.1.1"})
	ring.AddServer(&ServerInfo{Name: "node1-2", IP: "192.168.1.2"})
	ring.AddServer(&ServerInfo{Name: "node1-3", IP: "192.168.1.3"})

	server, err := ring.GetServerByFilename("file.txt")
	if err != nil {
		fmt.Println("file.txt -> err", err.Error())
	} else {
		fmt.Println("file.txt ->", server.Name)
	}

	// 模拟 node1-1 掉线
	ring.RemoveServer("node1-1")
	server, err = ring.GetServerByFilename("file.txt")
	if err != nil {
		fmt.Println("file.txt after node1-1 down ->err", err.Error())
	} else {
		fmt.Println("file.txt after node1-1 down ->", server.Name)
	}

	// 模拟 node1-1 恢复
	ring.AddServer(&ServerInfo{Name: "node1-1", IP: "192.168.1.1"})
	server, err = ring.GetServerByFilename("file.txt")
	if err != nil {
		fmt.Println("file.txt after node1-1 reOnline ->err", err.Error())
	} else {
		fmt.Println("file.txt after node1-1 reOnline ->", server.Name)
	}

}
