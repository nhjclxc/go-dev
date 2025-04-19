package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	/*
		编写一个网页服务器监听端口 8080，有如下处理函数：
			当请求 http://localhost:8080/hello/Name 时，响应：hello Name（Name 需是一个合法的姓，比如 Chris 或者 Madeleine）
			当请求 http://localhost:8080/shouthello/Name 时，响应：hello NAME
	*/

	http.HandleFunc("/hello/", helloServer)
	http.HandleFunc("/shouthello/", shouthelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

	// 启动项目后，浏览器地址栏访问：http://localhost:8080/world
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside helloServer handler start")

	// 向w写数据，即响应数据
	fmt.Fprintf(w, "Hello111,"+req.URL.Path[1:])
	fmt.Println(req.URL.Path)
	fmt.Println("Inside helloServer handler end")
}
func shouthelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside shouthelloServer handler start")

	// 向w写数据，即响应数据
	fmt.Fprintf(w, "Hello222,"+req.URL.Path[1:])
	fmt.Println(req.URL.Path)
	fmt.Println("Inside shouthelloServer handler end")
}
