package main

import (
	"context"
	"fmt"
)

func main() {

	s1 := NewServer("node1-1", "192.168.1.1")
	s2 := NewServer("node1-2", "192.168.1.2")
	s3 := NewServer("node1-3", "192.168.1.3")

	node := NewNode(context.Background(), 100)
	node.AddServer(s1)
	node.AddServer(s2)
	node.AddServer(s3)

	file := "file.txt"
	server, err := node.FindServerByFilename(file)
	if err != nil {
		fmt.Println("file.txt -> err", err.Error())
	}
	fmt.Println("file.txt ->", server.Name)
	server.SaveFile(file)

	// 模拟 node1-1 掉线，并且新增 video.mp4 文件
	node.RemoveServer("node1-1")
	file2 := "video.mp4"
	server2, err := node.FindServerByFilename(file2)
	if err != nil {
		fmt.Println("file.txt after node1-1 down -> err", err.Error())
	}
	fmt.Println("file.txt after node1-1 down ->", server2.Name)
	server2.SaveFile(file2)

	// 模拟 node1-1 恢复，并模拟有用户来获取video.mp4
	node.AddServer(s1)
	server3, err := node.FindServerByFilename(file2)
	if err != nil {
		fmt.Println("file.txt after node1-1 reOnline -> err", err.Error())
	}
	fmt.Println("file.txt after node1-1 reOnline ->", server3.Name)

	findFile, err := server3.FindFile(file2)
	if err == nil {
		fmt.Println("server3 find file", findFile)
	}
	fmt.Println("server3 find video.mp4 error", err)

	// 未找到该文件，在本节点内的其他节点查找
	server5, findFile2, err := node.FindFileOnOtherServer(file2, server3)
	if err != nil {
		fmt.Println("server5 find file err", err)
	}
	fmt.Println("server5 find file", server5.Name, findFile2.filename, findFile2.size)

}
