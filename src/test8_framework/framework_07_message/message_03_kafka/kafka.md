下面是一个**用 Go 实现 Kafka 消息队列**（生产者 + 消费者）的完整示例与讲解，包括：

* Kafka 的核心概念
* Go 客户端库推荐
* 实际代码（生产者、消费者、消息结构、优雅退出）
* 在项目中如何封装

---

## 一、Kafka 核心概念简述

| 名称                 | 含义                           |
| ------------------ | ---------------------------- |
| **Topic**          | 消息主题（类似于消息队列的名字）             |
| **Producer**       | 消息生产者，负责发送消息到指定 Topic        |
| **Consumer**       | 消息消费者，从 Topic 读取消息           |
| **Partition**      | 每个 Topic 分为多个分区，用于并行和高吞吐     |
| **Offset**         | 每条消息在分区中的顺序号（类似游标）           |
| **Consumer Group** | 消费组，Kafka 保证组内的每个分区只被一个消费者消费 |

---

## 二、Go Kafka 客户端推荐

✅ **主流推荐库：**

1. [Shopify/sarama](https://github.com/Shopify/sarama)（经典稳定，应用广泛）
2. [segmentio/kafka-go](https://github.com/segmentio/kafka-go)（更简洁、Go 风格更好）

这里我们选 **`segmentio/kafka-go`** 来演示。

---

## 三、安装依赖

```bash
go get github.com/segmentio/kafka-go
```

---

## 四、生产者示例（Producer）

```go
package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	topic := "test_topic"
	broker := "localhost:9092"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello Kafka! message #%d", i)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(msg),
			},
		)
		if err != nil {
			fmt.Println("写入消息失败:", err)
			break
		}
		fmt.Println("已发送:", msg)
		time.Sleep(time.Second)
	}
}
```

### 🔍 说明

* `WriterConfig` 配置生产者
* `LeastBytes` 负载均衡策略：将消息发往最空闲的分区
* `writer.WriteMessages` 发送一条或多条消息

---

## 五、消费者示例（Consumer）

```go
package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	topic := "test_topic"
	broker := "localhost:9092"
	groupID := "group_1"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,  // 10KB
		MaxBytes: 10e6,  // 10MB
	})

	defer reader.Close()

	fmt.Println("开始消费消息...")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("读取消息出错:", err)
			break
		}
		fmt.Printf("读取到消息: topic=%s partition=%d offset=%d key=%s value=%s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		time.Sleep(500 * time.Millisecond)
	}
}
```

### 🔍 说明

* `GroupID` 指定消费者组
* Kafka 自动管理 offset（断点续读）
* 每个分区同组中只会被一个消费者读取

---

## 六、优雅退出（Graceful Shutdown）

```go
ctx, cancel := context.WithCancel(context.Background())
go func() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    fmt.Println("收到中断信号，准备退出...")
    cancel()
}()

for {
    m, err := reader.ReadMessage(ctx)
    if err != nil {
        if errors.Is(err, context.Canceled) {
            fmt.Println("退出消费者循环")
            break
        }
        fmt.Println("读取错误:", err)
        continue
    }
    fmt.Println("消费:", string(m.Value))
}
```

---

## 七、项目中如何封装 Kafka 模块

推荐封装如下目录结构：

```
/internal
  /kafka
    producer.go
    consumer.go
```

示例封装：

```go
// producer.go
package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  brokers,
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		}),
	}
}

func (p *Producer) Send(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
```

使用：

```go
p := kafka.NewProducer([]string{"localhost:9092"}, "test_topic")
defer p.Close()

_ = p.Send(context.Background(), []byte("key1"), []byte("hello world"))
```

---

## 八、在容器中使用 Kafka（本地调试）

使用 Docker 启动一个单节点 Kafka：

```bash
docker run -d --name kafka -p 9092:9092 \
  -e KAFKA_BROKER_ID=1 \
  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  confluentinc/cp-kafka
```

---

## 九、可扩展方向

* ✅ 支持 JSON 消息结构（序列化/反序列化）
* ✅ 封装中间件（重试、死信队列）
* ✅ Kafka 消费者监控（lag 延迟监控）
* ✅ 分布式追踪（结合 Jaeger）

---

如果你告诉我你的使用场景（比如：“延迟处理任务”、“实时日志流”、“后台指令下发”等），
我可以帮你进一步设计一个**Kafka + Go 的完整业务架构方案**（含并发消费、重试、持久化）。

是否要我帮你基于你的业务场景定制一个？
