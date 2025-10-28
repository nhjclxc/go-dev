https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/



非常棒的问题 👍
下面给你一份 **在 macOS 上完整搭建 Go + Protobuf + gRPC 环境** 的指南（一步一步可直接复制运行），是业界标准配置。

---

## 🧩 一、前置环境检查

先确认你已经装好了以下工具 👇

```bash
go version
brew --version
```

如果没有 Homebrew，请先安装：

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

---

## 🧱 二、安装 Protocol Buffers 编译器（`protoc`）

```bash
brew install protobuf
```

验证：

```bash
protoc --version
# 输出类似：libprotoc 25.1
```

---

## ⚙️ 三、安装 Go 的 Protobuf 插件

Go 使用 **protoc-gen-go** 和 **protoc-gen-go-grpc** 两个插件来生成代码。

执行以下命令：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## 🧭 四、配置环境变量（关键！）

确保 Go 的安装路径下的可执行文件可以被 `protoc` 找到：

```bash
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```

验证是否可用：

```bash
which protoc-gen-go
which protoc-gen-go-grpc
```

都能输出路径即可 ✅

---

## 📘 五、编写一个示例 `.proto` 文件

创建目录结构：

```bash
mkdir -p ~/go/src/demo/proto
cd ~/go/src/demo/proto
```

创建文件 `hello.proto`：

```proto
syntax = "proto3";

package proto;

option go_package = "./;proto";

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
```

---

## 🔧 六、使用 protoc 生成 Go 代码

在当前目录执行：

```bash
protoc --go_out=. --go-grpc_out=. hello.proto
```

生成两个文件：

```
hello.pb.go
hello_grpc.pb.go
```

---

## 🚀 七、编写服务端与客户端（可运行）

### 1️⃣ 服务端（server.go）

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "demo/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	fmt.Println("🚀 gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

---

### 2️⃣ 客户端（client.go）

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "demo/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Response:", r.Message)
}
```

---

## ✅ 八、运行验证

在两个终端窗口中：

**终端 1：**

```bash
go run server.go
```

**终端 2：**

```bash
go run client.go
```

输出：

```
Response: Hello Go Developer
```

✅ 表示 Go + Protobuf + gRPC 环境搭建成功！

---

是否希望我帮你写一个 **自动化脚本（如 setup.sh）**，一键安装并配置完整的 Go + Protobuf 环境？这在新机器上会非常方便。




