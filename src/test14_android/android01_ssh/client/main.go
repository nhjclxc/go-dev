package main

import (
	"flag"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	// 解析命令行参数
	serverAddr := flag.String("server", "", "SSH server address (e.g., 192.168.1.100:2222)")
	flag.Parse()

	if *serverAddr == "" {
		log.Println("Usage: go_base_project-client -server <server_ip:port>")
		log.Println("Example: go_base_project-client -server 192.168.1.100:2222")
		os.Exit(1)
	}

	// SSH客户端配置
	config := &ssh.ClientConfig{
		User: "adb-client",
		Auth: []ssh.AuthMethod{
			ssh.Password(""), // 简化demo，不需要密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 简化demo，忽略host key验证
	}

	// 连接到公网SSH服务器
	log.Printf("Connecting to SSH server at %s", *serverAddr)

	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	// 进行SSH握手
	sshConn, chans, reqs, err := ssh.NewClientConn(conn, *serverAddr, config)
	if err != nil {
		log.Fatal("Failed to create SSH connection:", err)
	}
	defer sshConn.Close()

	sshClient := ssh.NewClient(sshConn, chans, reqs)
	defer sshClient.Close()

	log.Println("Connected to SSH server")

	// 注册并处理服务器发起的forwarded-tcpip通道
	forwardedChans := sshClient.HandleChannelOpen("forwarded-tcpip")
	if forwardedChans == nil {
		log.Fatal("Failed to register forwarded-tcpip handler")
	}
	go handleChannels(forwardedChans)

	log.Println("Requesting reverse tunnel...")

	// 发送tcpip-forward请求
	payload := ssh.Marshal(&struct {
		Address string
		Port    uint32
	}{
		Address: "0.0.0.0",
		Port:    6666,
	})

	ok, response, err := sshClient.SendRequest("tcpip-forward", true, payload)
	if err != nil {
		log.Fatal("Failed to request port forwarding:", err)
	}
	if !ok {
		log.Fatal("Port forwarding request denied by server")
	}

	log.Println("Reverse tunnel established, server is now listening on :6666")
	log.Printf("Users can connect with: adb connect %s", (*serverAddr)[:len(*serverAddr)-5]+":6666")
	log.Printf("Response: %v", response)

	// 保持连接
	select {}
}

func handleChannels(chans <-chan ssh.NewChannel) {
	// 处理从服务器来的通道请求
	for newChannel := range chans {
		if newChannel.ChannelType() != "forwarded-tcpip" {
			log.Printf("Rejecting unknown channel type: %s", newChannel.ChannelType())
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		log.Println("Received forwarded-tcpip channel from server")
		go handleChannel(newChannel)
	}
}

func handleChannel(newChannel ssh.NewChannel) {
	channel, requests, err := newChannel.Accept()
	if err != nil {
		log.Println("Failed to accept channel:", err)
		return
	}
	defer channel.Close()
	go ssh.DiscardRequests(requests)

	log.Println("Channel accepted, connecting to local ADB...")

	// 连接到本地 Android 设备的 ADB 端口
	localADBAddr := "127.0.0.1:5555"
	localConn, err := net.Dial("tcp", localADBAddr)
	if err != nil {
		log.Println("Failed to connect to local ADB:", err)
		return
	}
	defer localConn.Close()

	log.Println("Connected to local ADB, forwarding traffic")

	// 双向转发数据
	done := make(chan struct{}, 2)

	go func() {
		io.Copy(localConn, channel)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(channel, localConn)
		done <- struct{}{}
	}()

	<-done
	log.Println("Connection closed")
}

/*
https://www.yuque.com/nhjclxc/clq5m5/wu374b142oi76444?singleDoc# 《Android盒子跑Go程序》


// 设置环境
export NDK=/opt/homebrew/share/android-ndk
export CC_FOR_TARGET=$NDK/toolchains/llvm/prebuilt/darwin-arm64/bin/armv7a-linux-androideabi34-clang
export CGO_ENABLED=1
export PATH=$NDK/toolchains/llvm/prebuilt/darwin-x86_64/bin:$PATH
export CC=armv7a-linux-androideabi34-clang
export CXX=armv7a-linux-androideabi34-clang++

// 编译
GOOS=android GOARCH=arm CGO_ENABLED=1 go build -o go_base_project-client2 main.go

*/

/*

1、1000台服务器，每个文件只暴露一个文件地址出来，
2、不同地区的用户访问同一个文件地址时，根据用户所在地/ISP厂商等条件进行调度实际的服务器给用户
3、被调度到的服务器如果没有对应文件，那么返回302回原
4、下载对应文件

*/
