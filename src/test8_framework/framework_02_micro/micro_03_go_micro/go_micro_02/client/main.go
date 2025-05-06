package main

import (
	user_micro "client/user/micro"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

// go get google.golang.org/grpc
// go get github.com/golang/protobuf/proto
func main() {

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 1.连接服务端
	clientConn, err := grpc.NewClient("127.0.0.1:56399", grpc.WithInsecure(), grpc.WithBlock())
	//clientConn, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("微服务连接失败！！！", err)
		return
	}
	defer clientConn.Close()

	// 2.实例gRPC客户端
	userServiceClient := user_micro.NewUserServiceClient(clientConn)

	// 3.组装请求参数
	userRequest := user_micro.UserRequest{Name: "666"}

	// 4.调用接口
	userResponse, err2 := userServiceClient.GetUserByName(ctx, &userRequest)
	if err2 != nil {
		fmt.Println("UserService微服务GetUser接口调用失败！！！", err2)
		return
	}
	fmt.Printf("GetUser接口调用成功：%v \n", userResponse)
	fmt.Printf("GetUser接口调用成功：%#v \n", userResponse)

}
