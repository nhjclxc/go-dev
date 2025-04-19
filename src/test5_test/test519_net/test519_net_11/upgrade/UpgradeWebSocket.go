package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// 消息结构体
type Message struct {
	Type    string `json:"type"`    // register / broadcast / private
	Target  string `json:"target"`  // 目标用户名（私聊用）
	Content string `json:"content"` // 消息内容
}

// 用户结构体
type User struct {
	ID       string
	Username string
	Conn     *websocket.Conn
}

// 线程安全的用户管理器
type UserManager struct {
	mu                    sync.RWMutex
	users                 map[string]*User     // map[username]*User
	offlineChatHistoryMap map[string][]Message // map[username][]Message，key是用户名，val是消息列表
}

func NewUserManager() *UserManager {
	return &UserManager{
		users:                 make(map[string]*User),
		offlineChatHistoryMap: make(map[string][]Message),
	}
}

func (um *UserManager) Add(username string, user *User) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.users[username] = user
}

func (um *UserManager) Remove(username string) {
	um.mu.Lock()
	defer um.mu.Unlock()
	delete(um.users, username)
}

// Get 获取某个用户
func (um *UserManager) Get(username string) (*User, bool) {
	um.mu.RLock()
	defer um.mu.RUnlock()
	user, exists := um.users[username]
	return user, exists
}

// Broadcast 广播所有在线用户
func (um *UserManager) Broadcast(sender string, msg string) {
	um.mu.RLock()
	defer um.mu.RUnlock()
	for username, user := range um.users {
		if username != sender {
			user.Conn.WriteJSON(Message{
				Type:    "broadcast",
				Content: fmt.Sprintf("[%s]: %s", sender, msg),
			})
		}
	}
}

// KickUser 踢掉指定用户名的连接
func (um *UserManager) KickUser(username string) {
	um.mu.Lock()
	defer um.mu.Unlock()

	if client, ok := um.users[username]; ok {
		_ = client.Conn.Close()    // 主动关闭连接
		delete(um.users, username) // 从管理器移除
		fmt.Printf("⚠️ 用户 %s 被踢下线\n", username)
	}
}

// AddOfflineChatHistory 不在线用户消息加入离线记录
func (um *UserManager) AddOfflineChatHistory(username string, message Message) {
	um.mu.RLock()
	defer um.mu.RUnlock()

	//messages := append(um.offlineChatHistoryMap[username], message)
	//um.offlineChatHistoryMap[username] = messages

	um.offlineChatHistoryMap[username] = append(um.offlineChatHistoryMap[username], message)
}

// PushOfflineChatHistory 用户上线的时候推送离线消息离线记录
func (um *UserManager) PushOfflineChatHistory(username string, conn *websocket.Conn) {
	um.mu.RLock()
	defer um.mu.RUnlock()
	historySlice := um.offlineChatHistoryMap[username]
	if len(historySlice) <= 0 {
		return
	}

	for _, message := range historySlice {
		conn.WriteJSON(message)
	}

	// 推送完成之后，情况离线消息
	um.offlineChatHistoryMap[username] = make([]Message, 0)

}

// 创建 websocket 对象
var ws = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

// 初始化用户池
var userManager = NewUserManager()

// 新客户端链接处理方法
// w：客户端响应对象
// r：客户端响请求对象
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 获取该客户端的链接对象
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级失败:", err)
		return
	}
	defer conn.Close()

	// 初始化连接
	id := uuid.New().String()
	var username string
	registered := false
	isAdmin := false

	fmt.Println("新客户端连接，等待注册...")

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			if registered {
				fmt.Printf("用户 %s 断开连接\n", username)
				userManager.Remove(username)
			} else {
				fmt.Printf("未注册连接断开: %s\n", id)
			}
			break
		}

		switch msg.Type {
		case "register":
			username = msg.Content
			userManager.Add(username, &User{
				ID:       id,
				Username: username,
				Conn:     conn,
			})
			registered = true
			fmt.Printf("用户注册: %s\n", username)
			conn.WriteJSON(Message{
				Type:    "system",
				Content: fmt.Sprintf("欢迎你，%s！", username),
			})

			// 告诉前端你是管理员（或者不是）
			isAdmin = msg.Content == "admin"
			conn.WriteJSON(map[string]interface{}{
				"type":    "role",
				"isAdmin": isAdmin,
			})

			// 推送离线消息
			userManager.PushOfflineChatHistory(username, conn)

		case "broadcast":
			if !registered {
				conn.WriteJSON(Message{Type: "error", Content: "请先注册用户名"})
				continue
			}
			fmt.Printf("广播来自 %s: %s\n", username, msg.Content)
			userManager.Broadcast(username, msg.Content)

		case "private":
			if !registered {
				conn.WriteJSON(Message{Type: "error", Content: "请先注册用户名"})
				continue
			}
			targetUser, ok := userManager.Get(msg.Target)
			if !ok {
				// 用户不在线，将词条消息加入离线消息，待用户上线之后发送给他
				userManager.AddOfflineChatHistory(msg.Target, Message{
					Type:    "private",
					Content: fmt.Sprintf("[离线消息 - 私聊 - %s]: %s", msg.Target, msg.Content),
				})
				conn.WriteJSON(Message{Type: "error", Content: "目标用户不存在，已加入离线消息"})
				continue
			}
			targetUser.Conn.WriteJSON(Message{
				Type:    "private",
				Content: fmt.Sprintf("[私聊 - %s]: %s", msg.Target, msg.Content),
			})

		case "kick":
			// 踢人操作
			// 只有admin用户才具备踢人操作
			if isAdmin {

				// 检测 admi 是否登录注册
				if !registered {
					conn.WriteJSON(Message{Type: "error", Content: "请先注册用户名"})
					continue
				}

				// 检测被踢的那个人是否存在
				targetUser, exists := userManager.Get(msg.Target)
				if !exists {
					// 用户不存在，回复管理员，找不到指定用户
					conn.WriteJSON(Message{Type: "system", Content: fmt.Sprintf("[系统错误 - %s]: %s 不在线踢除失败！！！", username, msg.Target)})
					continue
				}

				// 踢之前给那个人发条消息，说他被踢了
				targetUser.Conn.WriteJSON(Message{Type: "system", Content: "由于您违法了社区规定，你已被踢出聊天！！！"})

				// 踢那个人
				userManager.KickUser(msg.Target)

				// 踢成功了和管理说一声，或者广播更具影响力
				conn.WriteJSON(Message{Type: "system", Content: fmt.Sprintf("[系统响应 - %s]: %s 踢除成功！！！", username, msg.Target)})
			}

		default:
			conn.WriteJSON(Message{Type: "error", Content: "未知消息类型"})
		}
	}
}

/*
实现一个支持用户管理、广播、私聊的 WebSocket 管理器（用 map 管理连接 + 简单协议判断消息类型）。

🎯 功能目标

	每个连接分配唯一 ID（UUID）
	用户主动注册用户名（支持昵称）
	支持广播（发给所有人）
	支持私聊（发给指定用户）
	支持断开连接自动移除

💡 消息协议（简单 JSON）：客户端发送消息格式如下（类型+目标+内容）：

	{
	  "type": "register",     // "register" | "broadcast" | "private"
	  "target": "user123",    // 私聊目标用户名（type = private 时使用）
	  "content": "你好世界！"  // 消息内容
	}
*/
func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket 服务器启动：http://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
