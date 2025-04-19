package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler start")

	// 向w写数据，即响应数据
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
	fmt.Println(req.URL.Path)
	fmt.Println("Inside HelloServer handler end")
}

func main1() {

	//详细请参照 package http import "net/http"

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

	// 启动项目后，浏览器地址栏访问：http://localhost:8080/world
}
