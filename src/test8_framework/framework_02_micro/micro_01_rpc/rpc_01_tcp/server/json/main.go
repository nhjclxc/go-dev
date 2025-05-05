package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type JsonRpc struct {
}
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


func (this *JsonRpc) GetJsonData(request JsonRpcResponse, response *JsonRpcResponse)  error{

	fmt.Printf("接收到的客户端数据： %#v \n", request)

	*response = JsonRpcResponse{
		Success: true,
		Data: "sdertgh",
		Msg: "操作成功",
	}

	fmt.Printf("发送给客户端的数据： %#v \n", response)

	return nil
}


// 使用 json 方式编码数据传输的过程
func main() {


	// 注册rpc服务
	err := rpc.RegisterName("JsonRpc", new(JsonRpc))
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer listener.Close()

	for true {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
