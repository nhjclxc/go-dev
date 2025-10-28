package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	proto "protobuf_08_stream/proto"
	"time"
)

type server struct {
	proto.UnimplementedStreamServiceServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("SayHello %s, message ", req.Name)
	return &proto.HelloResponse{Message: req.Name + "你好！！！"}, nil
}
func (s *server) ListHello(req *proto.HelloRequest, stream proto.StreamService_ListHelloServer) error {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello %s, message #%d", req.Name, i+1)
		if err := stream.Send(&proto.HelloResponse{Message: msg}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // gRPC 服务端口
	if err != nil {
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}
	grpcServer := grpc.NewServer()
	proto.RegisterStreamServiceServer(grpcServer, &server{})
	fmt.Println("🚀 gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
