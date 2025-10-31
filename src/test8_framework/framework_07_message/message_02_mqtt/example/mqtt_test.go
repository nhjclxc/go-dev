package example

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"message_02_mqtt/mqttcore"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {

	subscribeTopics := make([]*mqttcore.SubscribeTopic, 0)
	subscribeTopics = append(subscribeTopics, &mqttcore.SubscribeTopic{
		Topic: "/test/send",
		Callback: func(client mqtt.Client, msg mqtt.Message) {
			log.Printf("[Message - [/test/send]消息回调处理器正在处理消息] topic=%s payload=%s", msg.Topic(), msg.Payload())
		},
	})
	subscribeTopics = append(subscribeTopics, &mqttcore.SubscribeTopic{
		Topic: "/test/rec",
		Callback: func(client mqtt.Client, msg mqtt.Message) {
			log.Printf("[Message - [/test/rec]  消息回调处理器正在处理消息] topic=%s payload=%s", msg.Topic(), msg.Payload())
		},
	})

	// Broker 地址（tcp://host:port）
	mqttCfg := mqttcore.MqttConfig{
		Broker:          "tcp://localhost:1883",
		ClientId:        "go-client-",
		Username:        "admin",
		Password:        "public",
		SubscribeTopics: subscribeTopics,
	}

	client := mqttcore.NewMqttClient(&mqttCfg)

	_ = client

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go client.Run(ctx)

	// 持续运行等待消息
	select {}
}
func TestName2(t *testing.T) {

	// Broker 地址（tcp://host:port）
	mqttCfg := mqttcore.MqttConfig{
		Broker:   "tcp://localhost:1883",
		ClientId: "go-client-",
		Username: "admin",
		Password: "public",
	}

	client := mqttcore.NewMqttClient(&mqttCfg)

	for i := 0; i < 6; i++ {
		err := client.PublishWithAck("/test/send", 0, true, "who are you?"+strconv.Itoa(i))
		if err != nil {
			fmt.Println("q", err)
			return
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go client.Run(ctx)

	select {}
}
