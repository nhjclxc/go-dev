
## 🧩 一、Melody 是什么？

**Melody** 是一个 Go 的 WebSocket 管理库，作者是 [olahol](https://github.com/olahol)。它对底层的 WebSocket 实现进行了封装，提供了更高级的接口来处理 WebSocket 连接、广播消息、私聊等常见功能。

**目标人群：**
适合希望在 Go 项目中内嵌轻量级 WebSocket 实时通信能力的开发者。

---

## 📦 二、安装方式

```bash
go get github.com/olahol/melody
```

---

## 🧰 三、核心功能和特性

| 功能点                | 说明                          |
| ------------------ | --------------------------- |
| `HandleConnect`    | 处理新连接建立时的事件                 |
| `HandleDisconnect` | 处理连接断开事件                    |
| `HandleMessage`    | 收到消息时的回调                    |
| `Broadcast`        | 向所有连接广播消息                   |
| `BroadcastOthers`  | 向除了发送者以外的连接广播               |
| `Send` / `Write`   | 点对点发送消息                     |
| `Close`            | 主动断开连接                      |
| `Sessions`         | 获取当前所有会话                    |
| `Session.Keys`     | 可为每个 Session 添加自定义属性（如用户ID） |

---

## 📘 四、完整示例：聊天室（基于 Gin）

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"net/http"
)

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
		println("New anonymous_user connected")
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
```

打开浏览器访问 [http://localhost:5000，就能进入聊天室页面。](http://localhost:5000，就能进入聊天室页面。)

---

## 🧠 五、高级功能讲解

### 1️⃣ 每个连接的“上下文存储”

```go
m.HandleConnect(func(s *melody.Session) {
	s.Set("uid", "user123")
})

m.HandleMessage(func(s *melody.Session, msg []byte) {
	uid, _ := s.Get("uid")
	log.Printf("User %s sent: %s", uid, string(msg))
})
```

可以存储任意键值对，比如用户ID、Token、IP等。

---

### 2️⃣ 私聊（点对点发送）

```go
m.HandleMessage(func(s *melody.Session, msg []byte) {
	// 给自己回一条
	s.Write([]byte("you said: " + string(msg)))
})
```

---

### 3️⃣ 分组广播

没有原生分组 API，但可以用 `Session.Set` 实现：

```go
// 设置组别
s.Set("room", "roomA")

// 广播到 roomA
m.IterSessions(func(sess *melody.Session) bool {
	room, _ := sess.Get("room")
	if room == "roomA" {
		sess.Write([]byte("RoomA Broadcast"))
	}
	return true
})
```

---

### 4️⃣ 断开连接 / 检查连接状态

```go
if s.IsClosed() {
	// 连接已经断开
}
s.Close() // 主动断开
```

---

## 🚧 六、使用注意事项

| 问题                | 说明                      |
| ----------------- | ----------------------- |
| 不支持分布式广播          | Melody 是进程内方案，无法多节点共享连接 |
| 不支持客户端重连机制        | 需要前端手动实现                |
| 不支持消息持久化          | 一般用于轻量实时通知类服务           |
| 可以结合 Redis 实现简易广播 | 结合 `pub/sub` 实现进程间消息同步  |

---

## 🧩 七、适合的场景

* 内网即时通知系统（如 OA 审批）
* 简单聊天室
* 设备连接实时状态监控
* 小型实时协作工具（如白板、看板）

---

## ✅ 总结

| 特性         | 说明                     |
| ---------- | ---------------------- |
| 轻量级        | 单文件集成，使用简单             |
| 高性能        | 基于 Go 原生高并发能力          |
| 灵活可扩展      | 支持自定义连接数据、钩子函数         |
| 不适合大规模公网推送 | 建议结合其他方案（如 Centrifugo） |

---

是否需要我帮你写一个集成 JWT 鉴权的 WebSocket 示例？或者做一个多房间聊天室的改进版本？
