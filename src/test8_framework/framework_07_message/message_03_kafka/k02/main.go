package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"net"
)

func main() {
	// SASL/PLAIN 认证配置
	mechanism := plain.Mechanism{
		Username: "test",
		Password: "test123",
	}
	_ = mechanism

	dialer := &kafka.Dialer{
		//SASLMechanism: mechanism,
		// 可选：TLSConfig 也可以加，如果启用 SASL_SSL
	}

	kafka.Client{
		Addr:      net.Addr("192.168.203.182:9092"),
		Timeout:   0,
		Transport: nil,
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"192.168.203.182:9092"},
		Topic:   "test-topic2",
		Dialer:  dialer,
	})

	// 发送消息
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key1"),
			Value: []byte("Hello SASL Kafka!"),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("消息发送成功")
	writer.Close()

	// 消费消息
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"192.168.203.182:9092"},
		Topic:   "test-topic2",
		Dialer:  dialer,
	})

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("收到消息: %s = %s\n", string(msg.Key), string(msg.Value))
	reader.Close()
}
