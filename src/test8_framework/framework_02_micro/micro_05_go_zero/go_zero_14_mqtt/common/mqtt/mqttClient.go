package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go_zero_14_mqtt/internal/config"
	"go_zero_14_mqtt/internal/message"
	"log"
	"time"
)

type Client struct {
	mqtt.Client
}

func NewMqttClient(c config.MqttConfig) *Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.Broker)
	opts.SetClientID(c.ClientId + time.Now().String())
	opts.SetUsername(c.Username)
	opts.SetPassword(c.Password)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message on [%s]: %s\n", msg.Topic(), msg.Payload())
	})

	cli := mqtt.NewClient(opts)
	token := cli.Connect()
	if !token.WaitTimeout(3 * time.Second) {
		log.Println("订阅超时")
	} else if token.Error() != nil {
		log.Printf("订阅失败: %v", token.Error())
	}


	for _, topic := range c.SubTopics {
		token := cli.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
			// 接收到了消息
			log.Printf("Received message from topic [%s]: %s", msg.Topic(), string(msg.Payload()))

			message.HandleMessage(msg.Topic(), msg.Payload())

		})
		token.Wait()
		if token.Error() != nil {
			log.Printf("Subscription to [%s] failed: %v", topic, token.Error())
		} else {
			log.Printf("Subscribed to topic: %s", topic)
		}
	}

	return &Client{Client: cli}
}

func (c *Client) PublishQos0(topic string, payload interface{}) error {
	token := c.Client.Publish(topic, 1, false, payload)
	token.WaitTimeout(3 * time.Second)
	return token.Error()
}

func (c *Client) Publish(topic string, qos byte,payload interface{}) error {
	token := c.Client.Publish(topic, qos, false, payload)
	token.WaitTimeout(3 * time.Second)
	return token.Error()
}
