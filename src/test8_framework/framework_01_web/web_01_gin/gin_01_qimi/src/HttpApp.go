package main

import (
	"fmt"
	"net/http"
)

// /hello 接口的处理函数
// @param rw 接口响应数据
// @param rw 接口请求数据
func sayHello(rw http.ResponseWriter, r *http.Request) {

	fprintln, err := fmt.Fprintln(rw, "你好，我是http服务器")
	if err != nil {
		return
	}
	fmt.Println(fprintln)
}

func main0() {

	fmt.Println("Hello HttpApp.go")

	// 使用http定义一个接口
	http.HandleFunc("/hello", sayHello) // localhost:8080/hello
	http.HandleFunc("/hello/", sayHello)

	// 启动http服务
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
