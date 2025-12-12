package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	stream02 "protobuf_09_stream02/proto"
	"time"
)

type stream02Server struct {
	stream02.UnimplementedStreamServiceServer
}

func (s *stream02Server) SayHello(ctx context.Context, req *stream02.HelloRequest) (*stream02.HelloResponse, error) {
	str := fmt.Sprintf("服务端SayHello接收到请求，数据：%s \n", req.Name)
	log.Printf(str)
	return &stream02.HelloResponse{Message: str}, nil
}

func (s *stream02Server) ListHello(req *stream02.HelloRequest, stream grpc.ServerStreamingServer[stream02.HelloResponse]) error {
	str := fmt.Sprintf("服务端SayHello接收到请求，数据：%s \n", req.Name)
	log.Printf(str)
	// 流式返回数据
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello %s, message #%d", req.Name, i+1)
		if err := stream.Send(&stream02.HelloResponse{Message: msg}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 200)
	}
	return nil
}

func (s *stream02Server) RecordHello(stream grpc.ClientStreamingServer[stream02.HelloRequest, stream02.HelloResponse]) error {

	var messages []string

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Printf("客户端消息发送完毕，断开连接")
			break
		}
		if err != nil {
			log.Printf("服务端发送错误：%s \n", err.Error())
			return err
		}

		str := fmt.Sprintf("服务端 RecordHello 接收到请求，数据：%s \n", recv.Name)
		log.Printf(str)
		messages = append(messages, recv.Name)
	}

	// 可以构造一个最终响应
	response := &stream02.HelloResponse{
		Message: fmt.Sprintf("收到 %d 条消息", len(messages)),
	}

	// SendAndClose 表示“发回一个响应并关闭流”
	return stream.SendAndClose(response)
}

func (s *stream02Server) Chat(stream grpc.BidiStreamingServer[stream02.HelloRequest, stream02.HelloResponse]) error {

	var chatCh chan string = make(chan string) // 每次调用创建独立通道

	// 读协程
	go func() {
		defer close(chatCh) // 读完就关闭消息通道，让写协程退出
		for {
			recv, err := stream.Recv()
			if err == io.EOF {
				log.Printf("<客户端消息发送完毕>")

				// 通知写协程关闭
				chatCh <- "close"
				break
			}
			if err != nil {
				// 将错误响应给客户端
				log.Printf("<<Chat接收消息错误>>: %s \n", err.Error())
				return
			}

			// 模拟业务处理
			time.Sleep(time.Millisecond * 200)
			str := fmt.Sprintf("<UNK> Chat <UNK> 接收到客户端消息：%s \n", recv.Name)
			fmt.Println(str)

			// 将消息推送至写协程
			chatCh <- recv.Name
		}
	}()

	// 开启客户端请求消息接收，直至客户端关闭
	// 写协程（写在主 goroutine 中）
	for msg := range chatCh { // msgCh 被关闭时自动退出
		reply := msg + "【server deal】"

		if err := stream.Send(&stream02.HelloResponse{Message: reply}); err != nil {
			log.Printf("<<服务端发送消息给客户端失败！>> %s\n", err.Error())
			return err
		}
	}
	//双向流不能有“最终返回”这件事！
	//Bidi Stream 不存在 unary 或 client_streaming 的“最终返回”概念。
	// 也就是说双向留不能在for之外调用send向客户端发送消息，因为退出for了就意味着通道已经关闭了，不能在写通道了

	// 收发都结束，结束 RPC，自动关闭 stream
	return nil
}

// chat方法使得客户端能够接收多个请求

func main() {
	// 1、创建lis监听器
	// 2、创建grpc服务
	// 3、注册grpc服务，实现类要实现相应的接口
	// 4、启动监听

	// 1、创建lis监听器
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("创建grpc监听器失败：%s \n", err)
		return
	}
	defer listen.Close()

	// 2、创建grpc服务
	gprcServer := grpc.NewServer()

	// 3、注册grpc服务，实现类要实现相应的接口
	stream02.RegisterStreamServiceServer(gprcServer, &stream02Server{})

	// 4、启动监听
	log.Printf("grpc服务启动成功！")
	gprcServer.Serve(listen)

}
