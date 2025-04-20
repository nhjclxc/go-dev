package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Redis 的 发布/订阅（Pub/Sub）机制
/*
Redis 的 Pub/Sub 是一种 消息通信模型，主要包含三个角色：
	发布者（Publisher）：向某个频道发送消息。Publish 方法进行发布。
	订阅者（Subscriber）：订阅某个频道，监听该频道的消息。Subscribe 方法进行订阅
	频道（Channel）：消息的“通道”，消息通过频道进行传递。
订阅者订阅频道后，只要有发布者往该频道发送消息，订阅者就会接收到该消息。
优点是松耦合，缺点是 Redis 不会存储消息，只有订阅者在线时才能接收到消息。
*/
func main() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	//2. 订阅频道
	pubsub := redisClient.Subscribe("go-channel")

	// 等待确认订阅成功（重要）
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// 开启协程监听消息
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			fmt.Printf("收到消息：频道=%s，消息=%s\n", msg.Channel, msg.Payload)
		}
	}()

	// 3. 发布消息
	// 发布一条消息
	err2 := redisClient.Publish("go-channel", "Hello from publisher!").Err()
	if err2 != nil {
		panic(err2)
	}
	time.Sleep(5 * time.Second)

	//4. 取消订阅
	// 停止订阅
	err = pubsub.Close()
	if err != nil {
		panic(err)
	}

	//模式订阅（Pattern Subscribe）
	//可以订阅模式匹配的频道：
	//例如 news.sports, news.tech 等都会被接收。
	pubsubs := redisClient.PSubscribe("news.*")
	fmt.Println(pubsubs) // 对订阅者进行处理

	// 多频道订阅
	//监听多个频道同时接收消息。
	pubsubss := redisClient.Subscribe("chan1", "chan2", "chan3")
	fmt.Println(pubsubss) // 对订阅者进行处理

	// 问题			说明
	//消息不会持久化	如果订阅者不在线，发布的消息会丢失。
	//订阅是阻塞的	Receive() 是阻塞式的，建议用协程处理。
	//订阅不支持事务	订阅之后该连接不能执行其他 Redis 命令。
	//使用专用连接	订阅和普通命令不能共享一个连接。
	//使用 Channel() 非阻塞读取	可以使用 for msg := range ch 的形式非阻塞接收消息。

	// 消息不会被持久化，消息无需持久，只要“现在”能看到。只能实现消息的实时转发

}
