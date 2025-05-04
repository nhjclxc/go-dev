package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"testing"
	"time"
)

/*
// EventBus
// go get github.com/asaskevich/EventBus
// import "github.com/asaskevich/EventBus"
// 接口文档：https://pkg.go.dev/github.com/asaskevich/EventBus


// https://pkg.go.dev/github.com/asaskevich/EventBus#EventBus

Implemented methods

	New()
		New 返回带有空处理程序的新 EventBus。

	Subscribe(topic string, handler interface{}) error
		订阅主题。如果 handler 不是函数，则返回错误。

	SubscribeOnce(topic string, handler interface{}) error
		订阅一次主题。处理程序执行后将被移除。如果 handler 不是函数，则返回错误。

	Unsubscribe(topic string, handler interface{}) error
		删除主题定义的回调。如果主题没有订阅回调，则返回错误。

	HasCallback(topic string) bool
		如果存在任何订阅该主题的回调，则返回 true。

	Publish(topic string, args ...interface{})
		发布会执行主题定义的回调函数。任何附加参数都将传递给回调函数。【注意：传参必须和回调函数的参数列表一样】

	SubscribeAsync(topic string, handler interface{}, transactional bool)
		使用异步回调订阅主题。如果 handler 不是函数，则返回错误。
		事务性确定主题的后续回调是串行运行（true）还是并发运行（false）

	SubscribeOnceAsync(topic string, args ...interface{})
		SubscribeOnceAsync 的工作方式与 SubscribeOnce 类似，只是回调是异步执行的

	WaitAsync()
		WaitAsync 等待所有异步回调完成。


*/
// 基本使用
func TestBase1(t *testing.T) {

	// 创建总线对象
	bus := EventBus.New()

	// 订阅对象
	bus.Subscribe("main:calculator", calculator)

	// 发布消息
	bus.Publish("main:calculator", 20, 40)

	// 取消订阅
	bus.Unsubscribe("main:calculator", calculator)

	// 发布消息
	bus.Publish("main:calculator", 20, 40)
}

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func TestBase2(t *testing.T) {

	// 创建总线对象
	bus := EventBus.New()

	// 注册监听（订阅）
	err := bus.Subscribe("event:name", func(name ...string) {
		fmt.Println("Hello,", name)
	})
	if err != nil {
		fmt.Println("订阅错误！！！")
		return
	}

	// 发布事件（发布者）
	bus.Publish("event:name", "Alice")
	bus.Publish("event:name", "Alice", "Alice22222")

}

// 订阅一次测试
func TestOnce(t *testing.T) {

	// 创建总线对象
	bus := EventBus.New()

	// 订阅
	bus.SubscribeOnce("event:once", func(args ...any) {
		fmt.Println("event:once：", args)
	})

	// 发布事件
	bus.Publish("event:once", 12, 35, "qazxsw", 123.369)
	bus.Publish("event:once", 12, 35)

}

// 异步调用
func TestAsync(t *testing.T) {

	// 使用 PublishAsync 方法会异步调用监听器（handler）：

	// 创建总线对象
	bus := EventBus.New()

	err := bus.SubscribeAsync("async:event", func(msg string) {
		fmt.Println("Received async:", msg)
	}, false)
	if err != nil {
		return
	}

	bus.Publish("async:event", "Hello async")

	// 等待消息总线处理消息
	time.Sleep(3 * time.Second)

}
