

使用 go mod init client 和 go mod init server 分别初始化客户端和服务器

编译proto文件：


注意：protobuf协议仅限于golang程序之间的通信使用


# 1、protoc 相关插件安装

## 安装 protoc-gen-go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

## 安装 protoc-gen-go-grpc 插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


# 2、定义 proto 消息文件

详细代码定义看：[UserRequest.proto](proto%2FUserRequest.proto)

# 3、编译 proto 消息文件

在proto文件夹下面，进入cmd输入编译命令：```protoc --go_out=. --go-grpc_out=. proto/UserRequest.proto```

注意：*.pb.go文件生成之后不要轻易修改它。如果要修改应当从*.proto源文件夹中修改，然后重新编译输出

## 3.1、输出的两个文件介绍

### 3.1.1 [UserRequest.pb.go](proto%2Fgrpc%2Fuserinfo%2FUserRequest.pb.go)

这个文件定义了消息结构体


### 3.1.2 [UserRequest_grpc.pb.go](proto%2Fgrpc%2Fuserinfo%2FUserRequest_grpc.pb.go)

这个文件定义了远程rpc调用的接口定义。

注意：提供微服务接口的结构体必须匿名嵌入这个文件的UserInfoServiceServer结构体，详细看[main.go](server%2Fmain.go)中UserInfoService结构体的实现，这样UserInfoService提供的远程调用接口才能被发现


# 项目结构

[protobuf_01_user_rpc]()这个项目总共分为两个小项目（微服务），分别是：提供微服务接口的[server](server)和使用微服务接口的[client](client)。

因此，在[server](server)和[client](client) 都应当使用 ```go mod init ``` 来初始化这个项目。



```markdown
目录结构
├── proto
│   ├── UserRequest.proto  # protoc协议的消息实体文件
│   └── pb                 # protoc命令编译输出路径
│       └── grpc
│             └── userinfo
│                   ├── UserRequest.pb.go       # protoc命令编译输出的消息体文件，定义远程调用grpc的消息体
│                   └── UserRequest_grpc.pb.go  # protoc命令编译输出的消息体文件，定义远程调用grpc的接口
├── server
│   ├── main.go           # 服务端入口文件
│   └── grpc
│         └── userinfo
│                ├── UserRequest.pb.go       # protoc命令编译输出的消息体文件，定义远程调用grpc的消息体【由UserRequest.proto编译而来】
│                └── UserRequest_grpc.pb.go  # protoc命令编译输出的消息体文件，定义远程调用grpc的接口【由UserRequest.proto编译而来】
└── client
    ├── main.go           # 客户端入口文件
    └── grpc
          └── userinfo
                 ├── UserRequest.pb.go       # protoc命令编译输出的消息体文件，定义远程调用grpc的消息体【由UserRequest.proto编译而来】
                 └── UserRequest_grpc.pb.go  # protoc命令编译输出的消息体文件，定义远程调用grpc的接口【由UserRequest.proto编译而来】
```
