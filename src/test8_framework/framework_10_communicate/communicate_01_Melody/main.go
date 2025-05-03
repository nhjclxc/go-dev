package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)


// go get -u github.com/gin-gonic/gin
// go get -u github.com/olahol/melody
func main() {
	r := gin.Default()
	m := melody.New()

	// 前端 HTML 页面
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<body>
			<h2>Melody Chatroom</h2>
			<input id="msg" type="text" placeholder="Say something..."/>
			<button onclick="send()">Send</button>
			<ul id="chat"></ul>
			<script>
				var ws = new WebSocket("ws://" + location.host + "/ws");
				ws.onmessage = function(evt) {
					var li = document.createElement("li");
					li.innerText = evt.data;
					document.getElementById("chat").appendChild(li);
				};
				function send() {
					var msg = document.getElementById("msg").value;
					ws.send(msg);
				}
			</script>
		</body>
		</html>
		`))
	})

	// WebSocket 连接处理
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// 有人连接
	m.HandleConnect(func(s *melody.Session) {
		println("New user connected")
	})

	// 有人断开连接
	m.HandleDisconnect(func(s *melody.Session) {
		println("User disconnected")
	})

	// 有人发消息
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg) // 广播给所有人
	})

	r.Run(":5000")
}
