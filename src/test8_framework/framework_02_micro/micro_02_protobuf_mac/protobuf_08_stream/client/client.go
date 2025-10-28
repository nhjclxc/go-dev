package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	proto "protobuf_08_stream/proto"
)

//func (s *StreamServer) RecordHello(stream proto.StreamService_RecordHelloServer) error {
//	var names []string
//	for {
//		req, err := stream.Recv()
//		if err == io.EOF {
//			// 所有消息接收完毕，返回结果
//			return stream.SendAndClose(&proto.HelloResponse{
//				Message: "Hello to " + strings.Join(names, ", "),
//			})
//		}
//		if err != nil {
//			return err
//		}
//		names = append(names, req.Name)
//	}
//}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewStreamServiceClient(conn)

	r, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Response:", r.Message)

}
