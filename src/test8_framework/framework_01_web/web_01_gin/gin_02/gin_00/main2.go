package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http://192.168.201.192:8090/hello?id=123&name=zhangsan
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 从查询参数获取 id 和 name
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		// 设置响应头
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		str := fmt.Sprintf("我是 hello 页面，现在请求的id = %v, name = %v\n", id, name)

		fmt.Printf(str)

		fmt.Fprintf(w, str)
	})

	// 启动 HTTP 服务
	fmt.Println("Server is running on :8090 ...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
