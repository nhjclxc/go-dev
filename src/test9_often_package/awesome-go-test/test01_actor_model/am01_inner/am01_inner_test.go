package am01_inner

import (
	"fmt"
	"testing"
	"time"
)

// 使用 goroutine + channel 实现简单的 Actor

// Actor 消息结构体
type Message struct {
	From    string
	Content string
}

// Actor 结构体
type Actor struct {
	Name    string
	Mailbox chan Message
}

// NewActor 创建一个新的 Actor
func NewActor(name string) *Actor {
	return &Actor{
		Name:    name,
		Mailbox: make(chan Message, 10), // 带缓冲的消息队列
	}
}
func (a *Actor) Start() {
	go func() {
		fmt.Printf("[%s] 等待消息到来！！！", a.Name)
		for msg := range a.Mailbox {
			fmt.Printf("[%s] 接到了来自 [%s] 的消息  \n", a.Name, msg.From)
			if msg.From != "" {
				fromActor, ok := actorMap[msg.From]
				if ok {
					fromActor.Mailbox <- Message{From: a.Name, Content: "收到你的消息: " + msg.Content + "\n"}
				}
			}
		}
	}()
}

var actorMap map[string]*Actor = make(map[string]*Actor)

func Test1(t *testing.T) {

	a1 := Actor{Name: "zhangsan", Mailbox: make(chan Message, 10)}
	a2 := Actor{Name: "lisi", Mailbox: make(chan Message, 10)}
	actorMap["zhangsan"] = &a1
	actorMap["lisi"] = &a2

	a1.Start()
	a2.Start()

	a1.Mailbox <- Message{From: "lisi", Content: "收到你的消息: 111"}

	time.Sleep(10 * time.Second)

}
