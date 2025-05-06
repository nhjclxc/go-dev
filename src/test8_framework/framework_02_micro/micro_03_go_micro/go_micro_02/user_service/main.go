package main

import (
	"context"
	"fmt"
	"go-micro.dev/v5"
	"log"
	user_micro "user_service/user/micro"
)

type UserService struct {
	user_micro.UserServiceServer
}

var userService UserService = UserService{}

func (this *UserService) GetUserByName(context context.Context, userRequest *user_micro.UserRequest) (*user_micro.UserResponse, error) {
	fmt.Printf("UserService.GetUserByName：接收到调用者的请求数据：%#v 。\n", userRequest)

	response := user_micro.UserResponse{
		Name: "wqqq",
	}
	return &response, nil
}

func (this *UserService) InsertUser(context context.Context, userRequest *user_micro.UserRequest) (*user_micro.UserResponse, error) {
	fmt.Printf("UserService.InsertUser：接收到调用者的请求数据：%#v 。\n", userRequest)

	response := user_micro.UserResponse{
		Name: "wqqq",
	}
	return &response, nil
}


// go get go-micro.dev/v5
//
func main() {
	service := micro.NewService(
		micro.Name("user.service"),
	)

	service.Init()

	server := service.Server()

	server.NewHandler(userService.GetUserByName)
	server.NewHandler(userService.InsertUser)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
