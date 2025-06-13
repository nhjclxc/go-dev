package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_15_rabbitmq/common/rabbitmq"
	"go_zero_15_rabbitmq/internal/svc"
)

// rabbitmq 消息处理类

type RabbitMqMessageHandler struct {
	svcCtx   *svc.ServiceContext
	RabbitMQ *rabbitmq.RabbitMQ
}

func NewRabbitMqMessageHandler(rabbitMQ *rabbitmq.RabbitMQ) *RabbitMqMessageHandler {
	return &RabbitMqMessageHandler{
		RabbitMQ: rabbitMQ,
	}
}

//func (RabbitMqMessageHandler) MessageHandler(rabbitMqClient *rabbitmq.RabbitMQ) {
func (rabbitMqClient RabbitMqMessageHandler) MessageHandler() {


	fmt.Printf("消费者注册开始！！！\n\n")

	// ------------------ 一、简单模式 ---单生产者，单消费者---一条消息只能被一个人消费-------
	rabbitMqClient.RabbitMQ.ConsumeSimple("simpleQueue", "ConsumeSimple1", func(message string) {
		logx.Info(fmt.Sprintf("简单模式 queueName = %s, consumeName = %s, 接收到消息，message = %s", "simpleQueue", "ConsumeSimple22", message))
	})
	rabbitMqClient.RabbitMQ.ConsumeSimple("simpleQueue", "ConsumeSimple22", func(message string) {
		// 模拟单生产者多消费者
		// 1.2 Work Queue(⼯作队列)
		logx.Info(fmt.Sprintf("简单模式 queueName = %s, consumeName = %s, 接收到消息，message = %s", "simpleQueue", "ConsumeSimple22", message))
	})

	// ------------------ 二、发布订阅模式 ---生产者-交换机-队列-消费者---一个消息可以被多个消费者消费-------
	rabbitMqClient.RabbitMQ.RecieveSub("exchangePublishPub", "exchangeQueue1", "RecieveSubConsume11", func(message string) {
		logx.Info(fmt.Sprintf(" 发布订阅模式 exchange = %s, queue = %s, consumeName = %s, 接收到消息，message = %s", "exchangePublishPub", "exchangeQueue1", "RecieveSubConsume11", message))
	})
	rabbitMqClient.RabbitMQ.RecieveSub("exchangePublishPub", "exchangeQueue1", "RecieveSubConsume12", func(message string) {
		logx.Info(fmt.Sprintf(" 发布订阅模式 exchange = %s, queue = %s, consumeName = %s, 接收到消息，message = %s", "exchangePublishPub", "exchangeQueue1", "RecieveSubConsume12", message))
	})

	rabbitMqClient.RabbitMQ.RecieveSub("exchangePublishPub", "exchangeQueue2", "RecieveSubConsume21", func(message string) {
		logx.Info(fmt.Sprintf(" 发布订阅模式 exchange = %s, queue = %s, consumeName = %s, 接收到消息，message = %s", "exchangePublishPub", "exchangeQueue2", "RecieveSubConsume21", message))
	})
	rabbitMqClient.RabbitMQ.RecieveSub("exchangePublishPub", "exchangeQueue2", "RecieveSubConsume22", func(message string) {
		logx.Info(fmt.Sprintf(" 发布订阅模式 exchange = %s, queue = %s, consumeName = %s, 接收到消息，message = %s", "exchangePublishPub", "exchangeQueue2", "RecieveSubConsume22", message))

	})


	// ------------------ 三、路由模式
	rabbitMqClient.RabbitMQ.RecieveRouting(
		"RouterExchange",
		[]string{"RouterExchangeRouterKey1"},
		"RouterExchangeQueue1",
		func(routingKey, queueName, message string) {
			fmt.Printf(" 路由模式 exchange = %s, routerKey = %s, queueName = %s, 接收到消息，message = %s \n\n",
				"RouterExchange", routingKey, queueName, message)
	})
	rabbitMqClient.RabbitMQ.RecieveRouting(
		"RouterExchange",
		[]string{"RouterExchangeRouterKey2.1", "RouterExchangeRouterKey2.2"},
		"RouterExchangeQueue22",
		func(routingKey, queueName, message string) {
			fmt.Printf(" 路由模式 exchange = %s, routerKey = %s, queueName = %s, 接收到消息，message = %s \n\n",
				"RouterExchange", routingKey, queueName, message)
	})


	// ------------------ 四、话题模式
	// // 注意,在接收消息的时候才能使用通配符,发送的时候不能使用
	rabbitMqClient.RabbitMQ.RecieveTopic(
		"TopicExchange",
		[]string{"usa.#"},
		"TopicExchangeQueue-usa.#",
		func(routingKey, queueName, message string) {
			fmt.Printf(" 话题模式 exchange = %s, routerKey = %s, queueName = %s, 接收到消息，message = %s \n\n",
				"TopicExchange", routingKey, queueName, message)
		})
	rabbitMqClient.RabbitMQ.RecieveTopic(
		"TopicExchange",
		[]string{"europe.#"},
		"TopicExchangeQueue-europe.#",
		func(routingKey, queueName, message string) {
			fmt.Printf(" 话题模式 exchange = %s, routerKey = %s, queueName = %s, 接收到消息，message = %s \n\n",
				"TopicExchange", routingKey, queueName, message)
		})
	rabbitMqClient.RabbitMQ.RecieveTopic(
		"TopicExchange",
		[]string{"#.news"},
		"TopicExchangeQueue-#.news",
		func(routingKey, queueName, message string) {
			fmt.Printf(" 话题模式 exchange = %s, routerKey = %s, queueName = %s, 接收到消息，message = %s \n\n",
				"TopicExchange", routingKey, queueName, message)
		})
	rabbitMqClient.RabbitMQ.RecieveTopic(
		"TopicExchange",
		[]string{"#.weather"},
		"TopicExchangeQueue-#.weather",
		func(routingKey, queueName, message string) {
			fmt.Printf(" 话题模式 exchange = %s, routerKey = %s, queueName = %s, 接收到消息，message = %s \n\n",
				"TopicExchange", routingKey, queueName, message)
		})


	// ------------------ 五、RPC（远程过程调用）通信模式
	rabbitMqClient.RabbitMQ.RpcServer(context.Background(), "RpcServerQueue", func(args []byte) ([]byte) {
		fmt.Printf(" RPC通信模式服务器 queueName = %s, 接收到消息，message = %s \n\n",
			"RpcServerQueue", string(args))

		var restul map[string]any = map[string]any{
			"code": 200,
			"data": "RpcServer Return",
		}
		restulByte, err := json.Marshal(restul)
		if err != nil {
		}

		return restulByte
	})

	fmt.Printf("\n消费者注册结束！！！\n\n")

}
