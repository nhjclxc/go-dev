package main

// 官网：https://protobuf.com.cn/overview/
// 官网教程：https://protobuf.com.cn/getting-started/gotutorial/、https://protobuf.com.cn/reference/go/go-generated/
// 官方包：https://pkg.go.dev/google.golang.org/protobuf/proto

// 环境配置
// 1、去 github.com/golang/protobuf 下载最新的 protoc.exe，并将其路径放到系统环境变量里面
// 		检查是否配置成功：在cmd里面执行：protoc --version
// 2、下载 protobuf 的 golang 插件(注意这个插件是一个exe文件)，
//		命令：go install google.golang.org/protobuf/protoc-gen-go@latest
// 		在 GOPATH 下面将生成一个 protoc-gen-go.exe 可执行文件

// 3、编译 protoc 文件
//		protoc --go_out=pb addressbook.proto

/* https://protobuf.com.cn/getting-started/gotutorial/
使用 protoc 编译 addressbook.proto 的基本命令如下：

	protoc --proto_path=IMPORT_PATH --cpp_out=OUTPUT_DIR addressbook.proto
	其中：
		--proto_path（或 -I）指定 .proto 文件的搜索路径；
		--cpp_out 表示将生成 C++ 代码，换成 --java_out、--go_out、--python_out 等可以生成对应语言的代码；
		OUTPUT_DIR 是生成的代码存放目录；
		addressbook.proto 是你要编译的 .proto 文件。
 */


// go get google.golang.org/protobuf/proto
func main() {


}



