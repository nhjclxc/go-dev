package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "protobuf_01_hello/proto"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 连接到服务器端口
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	//conn, err := grpc.Dial(
	//	"localhost:50051",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Response:", r.Message)

	resp2, err := client.SayMorning(ctx, &pb.MorningRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("could not greet morning: %v", err)
	}
	fmt.Println("Morning Response:", resp2.Greeting)

}
