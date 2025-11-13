package slb03_local_search_file

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"
)

// 背景：cdn的node1节点里面的某个服务器node1-1掉线---> 下载文件  ---> node1-1上线
// 即：当node1-1服务器掉线后源站向每一个cdn节点推送了一个video.mp4文件，这时被映射到了node1-3，此时用户访问video.mp4可以在node1-3获取到，但是如果当node1-1有回复上线的时候后，用户请求video.mp4文件经过hash映射到了node1-1，此时node1-1没有这个文件而node1-3却有这个文件。那么这种情况就会触发302回源，但是当前节点node1里面有服务器node1-3有这个文件，即我有一个想法能不能先在node1内网里面的所有服务器里面查找是否有该文件，如果没有再回源

// 模拟节点内部查找冗余文件，节点内部冗余与回源策略
// 模拟处理方法：
// 		如果用户请求video.mp4结果hash后落到node1-1，但是node1-1服务器没有文件，此时先不要去302回源，
//		而是先去节点node1内部的其他服务器上查找文件，如果其他服务器上都没有再去执行302回源。
//		如果其他节点（如node1-3）上有该文件，那么先将该文件返回给用户，再将node1-3的该文件同步给node1-1，最后在node1-3上删除该文件

func TestLocalSearchFile(t *testing.T) {

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

	ip, err := GetInternalIP()
	fmt.Println("GetInternalIP", ip, err)
}

// 获取本机内网IP
func GetInternalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}
		if ip4 := ipnet.IP.To4(); ip4 != nil {
			if strings.HasPrefix(ip4.String(), "10.") ||
				strings.HasPrefix(ip4.String(), "172.") ||
				strings.HasPrefix(ip4.String(), "192.168.") {
				return ip4.String(), nil
			}
		}
	}
	return "", fmt.Errorf("未找到内网IP")
}
