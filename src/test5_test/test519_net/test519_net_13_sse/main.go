package main

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置必要的头部
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 如果涉及跨域

	// 获取 http.Flusher 接口
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// 每秒发送一条消息
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 10; i++ {
		select {
		case t := <-ticker.C:
			msg := fmt.Sprintf("data: 当前时间是 %s\n\n", t.Format(time.RFC3339))
			_, _ = fmt.Fprintf(w, msg)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("客户端断开连接")
			return
		}
	}
}

func main() {
	http.HandleFunc("/sse", sseHandler)
	fmt.Println("SSE 服务启动：http://localhost:8080/sse")
	http.ListenAndServe(":8080", nil)
}

