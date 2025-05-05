package main

import (
	grpc_userinfo "client/grpc/userinfo"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)


// go get google.golang.org/grpc
// go get github.com/golang/protobuf/proto
func main() {

	// 1.连接服务端
	// 2.实例gRPC客户端
	// 3.组装请求参数
	// 4. 调用接口


	// 1.连接服务端
	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. 实例化gRPC客户端
	client := grpc_userinfo.NewUserInfoServiceClient(conn)


	// 3.组装请求参数调用
	req := new(grpc_userinfo.UserRequest)
	req.Name = "张三123456"

	// 4. 调用接口
	response, err := client.GetUserInfoByName(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常  %s\n", err)
	}
	fmt.Printf("响应结果： %v\n", response)
	fmt.Printf("响应结果： %#v\n", response)

}
