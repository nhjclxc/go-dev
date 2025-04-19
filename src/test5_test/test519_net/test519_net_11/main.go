package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
)

// 升级 HTTP 连接为 WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域
		return true
	},
}

// 可选：用于管理连接
var clients = make(map[string]*websocket.Conn)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 升级连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级失败:", err)
		return
	}
	defer conn.Close()

	//fmt.Println("客户端已连接")
	// 为连接分配唯一 ID
	id := uuid.New().String()
	clients[id] = conn
	defer delete(clients, id)

	fmt.Printf("客户端已连接，ID: %s\n", id)

	for {
		// 读取消息
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取错误:", err)
			break
		}

		fmt.Printf("收到 [%s] 的消息: %s\n", id, message)
		//fmt.Printf("收到消息: %s\n", message)

		// 回复客户端
		err = conn.WriteMessage(mt, append([]byte("服务端回复: "), message...))
		if err != nil {
			fmt.Println("发送错误:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("WebSocket 服务运行在: http://localhost:8080/ws")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("服务器启动失败: " + err.Error())
	}
}
