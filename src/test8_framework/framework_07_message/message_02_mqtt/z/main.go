package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Broker 地址（tcp://host:port）
	broker := "tcp://localhost:1883" // 或 tcp://1.2.3.4:1883
	clientID := "go-client-001"
	username := "admin"  // 如果启用了认证
	password := "public" // 如果启用了认证
	topic := "test/topic"

	// 定义消息处理回调
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("收到消息: [%s] %s\n", msg.Topic(), string(msg.Payload()))
	}

	// 配置 MQTT 客户端选项
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messageHandler)
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("✅ 已连接到 EMQX")

		// 订阅主题
		if token := c.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
			fmt.Println("订阅失败:", token.Error())
		} else {
			fmt.Println("📡 已订阅主题:", topic)
		}
	}
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Println("❌ 连接断开:", err)
	}

	// 创建客户端并连接
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 发布测试消息
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("Hello MQTT %d", i)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(1 * time.Second)
	}

	// 持续运行等待消息
	select {}
}
