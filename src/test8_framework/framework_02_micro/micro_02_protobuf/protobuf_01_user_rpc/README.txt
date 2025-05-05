

目录结构
   - proto
        - UserRequest.proto protoc协议的消息实体文件
            - pb【protoc命令编译输出路径】
                - pb.q
                    - UserRequest.pb.go【protoc命令编译输出的消息体文件，得到这个文件之后，将这个文件放入server/dto和client/dto下面作为服务器与客户端之间通讯的消息体】
   - server
        - main.go 服务端入口文件
        - dto
            - UserRequest.pb.go 服务端接收到客户端的请求实体【由UserRequest.proto编译而来】
   - client
        - main.go 客户端入口文件
        - dto
            - UserRequest.pb.go 客户端像服务端请求的请求实体【由UserRequest.proto编译而来】

使用 go mod init client 和 go mod init server 分别初始化客户端和服务器

编译proto文件：protoc --go_out=. --go-grpc_out=. proto/UserRequest.proto


注意：protobuf协议仅限于golang程序之间的通信使用



# 安装 protoc-gen-go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 安装 protoc-gen-go-grpc 插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

