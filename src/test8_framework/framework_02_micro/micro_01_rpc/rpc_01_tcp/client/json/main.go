package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


type JsonRpcRequest struct {
	Id int
	Name string
	Age int
}
type JsonRpcResponse struct {
	Success bool
	Data any
	Msg string
}

func main() {

	// 连接微服务
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer conn.Close()

	// 获取客户端对象，建立基于json编码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var response JsonRpcResponse

	// 发送请求
	err2 := client.Call("JsonRpc.GetJsonData", JsonRpcRequest{
		Id:   666,
		Name: "你好名字",
		Age:  18,
	}, &response)
	if err2 != nil {
		return
	}

	fmt.Printf("客户端接收到的数据：%#v", response)


}
