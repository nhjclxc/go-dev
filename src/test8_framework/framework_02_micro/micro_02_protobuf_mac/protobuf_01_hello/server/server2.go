package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	pb "protobuf_01_hello/proto"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("接收[%s]的到请求 \n", req.Name)
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func (s *server) SayMorning(ctx context.Context, req *pb.MorningRequest) (*pb.MorningResponse, error) {
	log.Printf("Morning call: %v", req.GetName())
	return &pb.MorningResponse{Greeting: "Good morning, " + req.GetName() + "!"}, nil
}

// 即提供了grpc服务，又提供了http服务

func main() {
	go func() {
		lis, err := net.Listen("tcp", ":50051") // gRPC 服务端口
		if err != nil {
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
		}
		grpcServer := grpc.NewServer()
		pb.RegisterHelloServiceServer(grpcServer, &server{})
		fmt.Println("🚀 gRPC server listening on :50051")
		grpcServer.Serve(lis)
	}()

	go func() {
		//  http://localhost:8080/hello
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello HTTP!"))
		})
		log.Println("🚀 🚀 HTTP server listening on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	// select {} 保持主程序持续运行中，当不消耗cpu资源
	select {}
}
