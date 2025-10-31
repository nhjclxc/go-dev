package mqttcore

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

// MqttConfig 配置类
type MqttConfig struct {
	Broker    string   `json:"Broker"`
	ClientId  string   `json:"ClientId"`
	Username  string   `json:"Username"`
	Password  string   `json:"Password"`
	SubTopics []string `json:"SubTopics"`
}

// MqttClient mqtt客户端
type MqttClient struct {
	client mqtt.Client
	cfg    *MqttConfig
}

func NewMqttClient(cfg *MqttConfig) *MqttClient {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientId + time.Now().String())
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	// 定义收到消息的回调函数
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message on [%s]: %s\n", msg.Topic(), msg.Payload())
	})

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if !token.WaitTimeout(3 * time.Second) {
		log.Println("订阅超时")
	} else if token.Error() != nil {
		log.Printf("订阅失败: %v", token.Error())
	}

	for _, topic := range cfg.SubTopics {
		token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
			// 接收到了消息
			log.Printf("Received message from topic [%s]: %s", msg.Topic(), string(msg.Payload()))

			HandleMessage(msg.Topic(), msg.Payload())

		})
		token.Wait()
		if token.Error() != nil {
			log.Printf("Subscription to [%s] failed: %v", topic, token.Error())
		} else {
			log.Printf("Subscribed to topic: %s", topic)
		}
	}

	return &MqttClient{
		client: client,
		cfg:    cfg,
	}
}

func (c *MqttClient) PublishQos0(topic string, payload interface{}) error {
	token := c.client.Publish(topic, 1, false, payload)
	token.WaitTimeout(3 * time.Second)
	return token.Error()
}

func (c *MqttClient) Publish(topic string, qos byte, payload interface{}) error {
	token := c.client.Publish(topic, qos, false, payload)
	token.WaitTimeout(3 * time.Second)
	return token.Error()
}
