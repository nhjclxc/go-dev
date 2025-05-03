package main

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {

	// 1、连接微服务
	conn, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("微服务连接失败！！！", err)
		return
	}
	defer conn.Close()


	// 2、执行远程调用
	var request string = "client111-" + time.Now().String()
	var response string

	err2 := conn.Call("MicroRpcTcpServer.HelloWorld", request, &response)
	if err2 != nil {
		fmt.Println("客户端1远程调用微服务失败！！！", err2)
		return
	}

	// 输出远程调用结果
	fmt.Println(response)






}
