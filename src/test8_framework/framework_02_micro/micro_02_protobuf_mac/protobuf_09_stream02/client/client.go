package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	stream02 "protobuf_09_stream02/proto"
	"strconv"
	"sync"
	"time"
)

func main() {
	// 1、创建连接
	// 2、创建客户端
	// 3、调用接口

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	_ = ctx

	// 1、创建连接
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("创建grpc客户端失败：%s \n", err)
		return
	}
	defer conn.Close()

	// 2、创建客户端
	client := stream02.NewStreamServiceClient(conn)

	// 3、调用接口
	hello, err := client.SayHello(ctx, &stream02.HelloRequest{Name: "stream02.Client.SayHello"})
	if err != nil {
		log.Printf("stream02.Client.SayHello 调用接口失败：%s \n", err)
		return
	}
	fmt.Printf("调用SayHello接口成功，返回的消息：%s \n", hello.Message)

	ListHello(err, client, ctx)

	RecordHello(err, client, ctx)

	Chat(err, client, ctx)

	fmt.Println("客户端运行完毕！！！")
}

func ListHello(err error, client stream02.StreamServiceClient, ctx context.Context) {
	fmt.Printf("------------------------------ \n")
	// 3.2、调用ListHello流式返回接口
	listHello, err := client.ListHello(ctx, &stream02.HelloRequest{Name: "stream02.Client.ListHello"})
	if err != nil {
		log.Printf("stream02.Client.ListHello 调用接口失败：%s \n", err)
		return
	}

	for {
		recv, err := listHello.Recv()
		if err == io.EOF {
			// 服务端发送完毕
			log.Println("服务端已关闭流 (EOF)")
			break
		}
		if err != nil {
			log.Printf("stream02.Client.ListHello 接收消息失败：%s \n", err)
			return
		}

		log.Printf("收到服务端消息: %v\n", recv)
	}
	return
}

func RecordHello(err error, client stream02.StreamServiceClient, ctx context.Context) {
	fmt.Printf("------------------------------ \n")
	// 3.3、调用 RecordHello 请求参数流式接口

	// 调用客户端流方法
	stream, err := client.RecordHello(ctx)
	if err != nil {
		log.Fatalf("调用 RecordHello 失败: %v", err)
	}

	// 连续发送多条请求
	for i := 1; i <= 3; i++ {
		req := &stream02.HelloRequest{
			Name: fmt.Sprintf("client.RecordHello. message-%d", i),
		}

		fmt.Println("客户端发送:", req.Name)

		if err := stream.Send(req); err != nil {
			log.Fatalf("发送失败: %v", err)
		}

		time.Sleep(time.Second)
	}

	// 发送完毕，客户端告诉服务端不再发送
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("接收响应失败: %v", err)
	}

	// 这是服务端最终返回的唯一响应
	fmt.Println("服务端最终响应:", resp.Message)
}

func Chat(err error, client stream02.StreamServiceClient, ctx context.Context) {
	fmt.Printf("------------------------------ \n")
	// 3.4、调用 Chat 请求参数流式请求和返回接口
	chat, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("调用 Chat 失败: %v", err)
		return
	}

	wg := sync.WaitGroup{}
	// 启动一个子协程接收消息，这个接收消息和下面的发送消息不可以写在一起，因为接收消息是阻塞的
	wg.Add(1)
	go func() {
		for {
			recv, err := chat.Recv()
			if err == io.EOF {
				log.Println("<服务端关闭>")
				break
			}
			if err != nil {
				log.Printf("<<接收消息发送错误>> %s\n", err)
				return
			}
			log.Printf("<客户端接收到服务端的消息>%s\n", recv.Message)
		}
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		// 模拟本次业务处理
		time.Sleep(time.Millisecond * 100)
		chatMsg := "【client：" + strconv.Itoa(i) + "】, "
		err = chat.Send(&stream02.HelloRequest{Name: chatMsg})
		if err != nil {
			log.Printf("<<客户端给服务端发送消息失败>> %s\n", err)
			return
		}
	}
	// 消息发送完毕关闭客户端
	err = chat.CloseSend()
	if err != nil {
		log.Printf("<<客户端通道关闭失败>> %s\n", err)
		return
	}

	wg.Wait()
}
