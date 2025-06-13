package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)


// 定义回调函数类型，用于接收到消息之后进行回调
// 普通的消息回调
type MessageHandler func(message string)
// 带有RoutingKey的消息回调
type RoutingKeyMessageHandler func(routingKey, queueName, message string)
// 定义一个 RPC 消息处理函数类型：接收参数 args，返回响应 result
type RpcMessageHandler func(args []byte) ([]byte)


// RabbitMQ 封装结构体
type RabbitMQ struct {
	// *amqp.Connection 是 线程安全且可以复用的，推荐在应用中作为全局单例复用
	conn     *amqp.Connection
	// Connection 可以共用，但信道不能共用
	// RabbitMQ 的 每个 channel 都不能被多个 goroutine 并发读写。否则就会出现类似： UNEXPECTED_FRAME - expected content header for class 60, got non content header frame instead
	Exchange string
	Key      string
	MqUrl    string
}

// 创建新实例
func NewRabbitMQ(MqUrl string) *RabbitMQ {
	return &RabbitMQ{
		Exchange: "",
		Key:      "",
		MqUrl:    MqUrl,
	}
}

// 连接与通道初始化
func (r *RabbitMQ) connect() {
	var err0 error
	if r.conn == nil {
		r.conn, err0 = amqp.Dial(r.MqUrl)
		r.failOnErr(err0, "连接失败")
	}
}

// 释放资源
func (r *RabbitMQ) Destroy() {
	if r.conn != nil {
		_ = r.conn.Close()
	}
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}


// ------------------ 一、简单模式 ---单生产者，单消费者---一条消息只能被一个人消费-------

// 简单模式 发送
func (r *RabbitMQ) PublishSimple(queueName, message string) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")
	// 注意：只有生产者可以关闭通道，消费者不能关闭。消费者发出这条消息之后可以关闭通道，但是生产者的通道必须一直保持激活状态，这样才能接受消息。
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		r.failOnErr(err, "关闭通道失败")
	}(channel)

	_, err := channel.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	r.failOnErr(err, "声明队列失败")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 在简单模式下，你可以不手动声明交换机，RabbitMQ 会使用默认交换机（名称是空字符串 ""），然后通过 routingKey（也就是你的 QueueName）将消息直接路由到队列。
	err = channel.PublishWithContext(ctx,
		r.Exchange, // 这里是 ""
		queueName,  // routingKey，用于路由到目标队列
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "发送消息失败")
}

// 简单模式 接收
func (r *RabbitMQ) ConsumeSimple(queueName string, consumeName string, handler MessageHandler) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	err0 = channel.Qos(
		1,     // 每次只给消费者分发1条未ack的消息，处理完才能发下一条
		0,     // prefetch size 不限制
		false, // 是否全局设置（false表示每个consumer独立）
	)
	r.failOnErr(err0, "设置 Qos 失败")

	_, err := channel.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	r.failOnErr(err, "声明队列失败")

	msgs, err := channel.Consume(
		//queueName, "", true, false, false, false, nil,
		queueName, "", false, false, false, false, nil,
	)
	r.failOnErr(err, "消费失败")

	//forever := make(chan bool)
	go func() {
		for d := range msgs {
			//log.Printf("收到消息: %s", d.Body)
			handler(string(d.Body))

			// ✅ 手动确认
			err := d.Ack(false)
			if err != nil {
				log.Println("收到消息: ", err)
			}
		}
	}()
	logx.Info(fmt.Sprintf("注册 简单模式 queueName = %s, consumeName = %s, 等待消息\n", queueName, consumeName))
	//<-forever
}


// ------------------ 二、发布订阅模式 ---生产者-交换机-队列-消费者---一个消息可以被多个消费者消费-------

// 发布订阅模式 发送
func (r *RabbitMQ) PublishPub(exchange, queue, message string) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	// 注意：只有生产者可以关闭通道，消费者不能关闭。消费者发出这条消息之后可以关闭通道，但是生产者的通道必须一直保持激活状态，这样才能接受消息。
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		r.failOnErr(err, "关闭通道失败")
	}(channel)

	// kind string,       // 交换机类型：direct（直连交换机） / fanout（扇出交换机） / topic（主题交换机） / headers（头交换机）
	err := channel.ExchangeDeclare(
		// "direct" 是固定值交换机类型，不能修改
		exchange, "direct", true, false, false, false, nil,
	)
	r.failOnErr(err, "声明交换机失败")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx,
		exchange, r.Key, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "发送失败")
}

// 发布订阅模式 接收
func (r *RabbitMQ) RecieveSub(exchange, queue string, consumeName string, handler MessageHandler) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	err0 = channel.Qos(
		1,     // 每次只给消费者分发1条未ack的消息，处理完才能发下一条
		0,     // prefetch size 不限制
		false, // 是否全局设置（false表示每个consumer独立）
	)
	r.failOnErr(err0, "设置 Qos 失败")

	err := channel.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
	)
	r.failOnErr(err, "声明交换机失败")

	q, err := channel.QueueDeclare(
		queue, false, false, true, false, nil,
	)
	r.failOnErr(err, "声明队列失败")

	err = channel.QueueBind(q.Name, "", exchange, false, nil)
	r.failOnErr(err, "绑定失败")

	msgs, err := channel.Consume(
		q.Name, "", false, false, false, false, nil,
	)
	r.failOnErr(err, "消费失败")

	//forever := make(chan bool)
	go func() {
		for d := range msgs {
			//log.Printf("订阅者收到消息: %s", d.Body)
			handler(string(d.Body))

			// ✅ 手动确认
			err := d.Ack(false)
			if err != nil {
				log.Println("收到消息: ", err)
			}
		}
	}()
	logx.Info(fmt.Sprintf("注册 发布订阅模式 exchange = %s, queue = %s, consumeName = %s, 等待消息\n", exchange, queue, consumeName))
	//<-forever
}


// ------------------ 三、路由模式 ---先交给交换机---交换机根据routingKey交给指定的队列(routingKey 不可以使用通配符)---消费者去队列里面消费---------
// 队列与交换机的绑定，要指定一个 RoutingKey（路由key）
// 消息的发送方在向 Exchange 发送消息时，也必须指定消息的 RoutingKey
// Exchange 不再把消息交给每一个绑定的队列，而是根据消息的 Routing Key 进行判断，只有队列的 Routingkey 与消息的 Routing key 完全一致，才会接收到消息

// 路由模式 发送
func (r *RabbitMQ) PublishRouting(exchange, routerKey, message string) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	// 注意：只有生产者可以关闭通道，消费者不能关闭。消费者发出这条消息之后可以关闭通道，但是生产者的通道必须一直保持激活状态，这样才能接受消息。
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		r.failOnErr(err, "关闭通道失败")
	}(channel)

	err := channel.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
		//exchange, "topic", true, false, false, false, nil,
	)
	r.failOnErr(err, "声明交换机失败")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx,
		exchange, routerKey, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "发送失败")
}

// 路由模式 接收
func (r *RabbitMQ) RecieveRouting(exchange string, routerKeys []string, queueName string, handler RoutingKeyMessageHandler) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	err := channel.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
		//exchange, "topic", true, false, false, false, nil,
	)
	r.failOnErr(err, "声明交换机失败")

	q, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	r.failOnErr(err, "声明队列失败")

	// 支持多个 routingKey 绑定到该队列
	for _, key := range routerKeys {
		err = channel.QueueBind(q.Name, key, exchange, false, nil)
		r.failOnErr(err, fmt.Sprintf("绑定 routingKey [%s] 失败", key))
	}
	r.failOnErr(err, "绑定失败")

	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	r.failOnErr(err, "消费失败")

	go func() {
		for d := range msgs {
			handler(d.RoutingKey, q.Name, string(d.Body))
		}
	}()
	log.Println("路由模式 等待消息\n")
}


// ------------------ 四、话题模式 ---先交给交换机---交换机根据routingKey交给指定的队列(routingKey可以使用通配符)---消费者去队列里面消费---------
//想使用 通配符匹配，请使用 topic 类型交换机：
//| 通配符 | 说明                       |
//| ----- | -------------------------- |
//|  `*`  | 匹配一个单词                |
//|  `#`  | 匹配零个或多个单词（包含`.`） |
// 绑定 routing key：order.* 表示匹配如 order.create、order.update
// 绑定 routing key：order.# 表示匹配如 order.create、order.cancel.email、order.cancel.sms
// // 注意,在接收消息的时候才能使用通配符,发送的时候不能使用


// 话题模式 发送
func (r *RabbitMQ) PublishTopic(exchange, routerKey, message string) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	err := channel.ExchangeDeclare(
		exchange, "topic", true, false, false, false, nil,
	)
	r.failOnErr(err, "声明交换机失败")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx,
		exchange, routerKey, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "发送失败")
}

// 话题模式 接收
func (r *RabbitMQ) RecieveTopic(exchange string, routerKeys []string, queueName string, handler RoutingKeyMessageHandler) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	err := channel.ExchangeDeclare(
		exchange, "topic", true, false, false, false, nil,
	)
	r.failOnErr(err, "声明交换机失败")

	q, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	r.failOnErr(err, "声明队列失败")

	// 支持多个 routingKey 绑定到该队列
	for _, key := range routerKeys {
		err = channel.QueueBind(q.Name, key, exchange, false, nil)
		r.failOnErr(err, fmt.Sprintf("绑定 routingKey [%s] 失败", key))
	}
	r.failOnErr(err, "绑定失败")

	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	r.failOnErr(err, "消费失败")

	go func() {
		for d := range msgs {
			handler(d.RoutingKey, q.Name, string(d.Body))
		}
	}()
	log.Println("话题模式 等待消息\n")
}


// ------------------ 五、RPC通信模式 ------------------
// 🧩 基本原理
//RabbitMQ 并没有原生的 RPC 功能，但我们可以借助消息队列机制自己实现：
//客户端（Client） 发送请求到一个 RPC 请求队列。
//服务端（Server） 监听该队列并处理请求，处理完成后将结果发送到 客户端指定的回复队列（reply_to）。
//客户端通过一个唯一标识 correlation_id 区分是哪一个请求的响应。
//客户端监听自己的临时队列，接收响应。

func (r *RabbitMQ) RpcServer(ctx context.Context, queueName string, handler RpcMessageHandler) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")

	q, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	r.failOnErr(err, "声明队列失败")

	msgs, _ := channel.Consume(q.Name, "", false, false, false, false, nil)
	r.failOnErr(err, "消费失败")


	go func() {
		for d := range msgs {

			fmt.Printf("\n\nRPC接收到消息: %s\n", d)

			// 调用处理函数,一般就是业务代码的执行
			result := handler(d.Body)

			// 返回响应
			responseBody, _ := json.Marshal(result)
			err := channel.Publish(r.Exchange, d.ReplyTo, false, false,
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          responseBody,
				}, )
			// 手动确认
			channel.Ack(d.DeliveryTag, false)

			if err != nil {
				fmt.Printf("消息回复失败!!!")
			}
			fmt.Printf("消息回复成功!!!")
		}
	}()
	log.Println("RPC通信模式 等待消息\n")
}

func (r *RabbitMQ) RpcClientCall(ctx context.Context, queueName string, args []byte, timeout time.Duration) (res []byte, err error) {
	r.connect()

	channel, err0 := r.conn.Channel()
	r.failOnErr(err0, "打开通道失败")
	//defer ch.Close()

	// 创建临时队列作为响应
	replyQueue, _ := channel.QueueDeclare(queueName, true, false, false, false, nil)
	msgs, _ := channel.Consume(replyQueue.Name, "", true, false, false, false, nil)

	corrId := uuid.New().String()

	// 发送请求
	err = channel.Publish(r.Exchange, "RpcServerQueue", false, false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       replyQueue.Name,
			Body:          args,
		}, )
	if err != nil {
		fmt.Printf("消息回复失败!!!", err)
	}

	// 等待响应 + 超时机制
	resultCh := make(chan []byte)

	go func() {
		for d := range msgs {
			fmt.Printf("\n\n接收到消息: %s \n", string(d.Body))
			if d.CorrelationId == corrId {
				resultCh <- d.Body
				return
			}
		}
	}()

	select {
	case res := <-resultCh:
		return res, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("调用超时: 超过 %s 无响应", timeout)
	}


}

