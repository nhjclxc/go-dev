package main

import (
	"fmt"
	"net/rpc"
)

func main() {


	// 打开一个对话
	conn, err := rpc.Dial("http", "127.0.0.1:8090")
	if err != nil {
		return
	}
	defer conn.Close()

	request := "requestName"
	var response string

	// 远程调用
	_ = conn.Call("Http01.GetById", request, &response)

	fmt.Println("response = ", response)


}
