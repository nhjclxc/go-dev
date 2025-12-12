package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
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

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"192.168.203.182:9092"},
		Topic:   "test-topic",
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
		Topic:   "test-topic",
		Dialer:  dialer,
	})

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("收到消息: %s = %s\n", string(msg.Key), string(msg.Value))
	reader.Close()
}

/*vim kafka_jaas.conf
KafkaServer {
  org.apache.kafka.common.security.plain.PlainLoginModule required
  username="admin"
  password="admin123"
  user_admin="admin"
  user_test="test";
};


mkdir kafka-logs
vim docker-compose.yaml

version: '3.8'

services:
  kafka:
    image: apache/kafka:3.5.1
    container_name: kafka
    ports:
      - "9092:9092"
      - "9093:9093"
    environment:
      KAFKA_KRAFT_MODE: "true"
      KAFKA_PROCESS_ROLES: "broker,controller"
      KAFKA_NODE_ID: 1
      KAFKA_CONTROLLER_QUORUM_VOTERS: "1@kafka:9093"
      KAFKA_LISTENERS: "SASL_PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9093"
      KAFKA_ADVERTISED_LISTENERS: "SASL_PLAINTEXT://localhost:9092"
      KAFKA_LOG_DIRS: "/kafka-logs"
      KAFKA_SASL_ENABLED_MECHANISMS: "PLAIN"
      KAFKA_SASL_MECHANISM_INTER_BROKER_PROTOCOL: "PLAIN"
      KAFKA_AUTHORIZER_CLASS_NAME: "kafka.security.auth.SimpleAclAuthorizer"
      KAFKA_ALLOW_EVERYONE_IF_NO_ACL_FOUND: "true"
      KAFKA_SUPER_USERS: "User:admin"
    volumes:
      - ./kafka-logs:/kafka-logs
      - ./kafka_jaas.conf:/etc/kafka/kafka_jaas.conf


docker-compose up -d
*/
