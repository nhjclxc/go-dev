package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	pb "protobuf_03_test1/proto"
)

func main() {

	// 开启ctx会话
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	//连接服务器
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务器失败，err = %s \n", err)
		return
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	fmt.Printf("User客户端创建成功！！！\n")

	// 客户端发送请求
	in := pb.GetUserRequest{UserId: 666}
	user, err := client.GetUser(ctx, &in)
	if err != nil {
		fmt.Printf("GetUser氢气失败！！！err = %s \n", err)
		return
	}
	fmt.Printf("GetUser请求成功：code = %d, msg = %s, data = %v", user.Code, user.Msg, user.Data)

}
