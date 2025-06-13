package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go_zero_15_rabbitmq/common/rabbitmq"
	"time"

	"go_zero_15_rabbitmq/internal/svc"
	"go_zero_15_rabbitmq/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RabbitmqApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	RabbitMQ *rabbitmq.RabbitMQ
}

// 获取用户信息
func NewRabbitmqApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RabbitmqApiLogic {
	return &RabbitmqApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		RabbitMQ: svcCtx.RabbitMQ,
	}
}

func (l *RabbitmqApiLogic) RabbitmqApi(req *types.RabbitmqApiReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n RabbitmqApi = %#v \n\n\n", req)


	// ------------------ 一、简单模式 ---单生产者，单消费者---一条消息只能被一个人消费-------
	l.RabbitMQ.PublishSimple("simpleQueue", req.Msg)

	// ------------------ 二、发布订阅模式 ---生产者-交换机-队列-消费者---一个消息可以被多个消费者消费-------
	l.RabbitMQ.PublishPub("exchangePublishPub", "exchangeQueue1", "exchangeQueue1"+req.Msg)
	l.RabbitMQ.PublishPub("exchangePublishPub", "exchangeQueue2", "exchangeQueue2"+req.Msg)

	fmt.Printf("200 success !!!\n\n\n")

	return nil
}


func (l *RabbitmqApiLogic) RabbitmqApiSimple(req *types.RabbitmqApiReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n RabbitmqApiSimple = %#v \n\n\n", req)


	// ------------------ 一、简单模式 ---单生产者，单消费者---一条消息只能被一个人消费-------
	l.RabbitMQ.PublishSimple("simpleQueue", req.Msg)


	fmt.Printf("\n200 success !!!\n\n")

	return nil
}


func (l *RabbitmqApiLogic) RabbitmqApiPublish(req *types.RabbitmqApiReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n RabbitmqApiPublish = %#v \n\n\n", req)


	// ------------------ 二、发布订阅模式 ---生产者-交换机-队列-消费者---一个消息可以被多个消费者消费-------
	l.RabbitMQ.PublishPub("exchangePublishPub", "exchangeQueue1", "exchangeQueue1"+req.Msg)
	l.RabbitMQ.PublishPub("exchangePublishPub", "exchangeQueue2", "exchangeQueue2"+req.Msg)

	fmt.Printf("\n200 success !!!\n\n")

	return nil
}


func (l *RabbitmqApiLogic) RabbitmqApiRouter(req *types.RabbitmqApiReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n RabbitmqApiRouter = %#v \n\n\n", req)

	// ------------------ 三、路由模式
	l.RabbitMQ.PublishRouting("RouterExchange", "RouterExchangeRouterKey1", "RouterExchangeRouterKey1 - " + req.Msg)
	l.RabbitMQ.PublishRouting("RouterExchange", "RouterExchangeRouterKey2.1", "RouterExchangeRouterKey2.1 - " + req.Msg)
	l.RabbitMQ.PublishRouting("RouterExchange", "RouterExchangeRouterKey2.2", "RouterExchangeRouterKey2.2 - " + req.Msg)

	fmt.Printf("\n200 success !!!\n\n")

	return nil
}


func (l *RabbitmqApiLogic) RabbitmqApiTopic(req *types.RabbitmqApiReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n RabbitmqApiTopic = %#v \n\n\n", req)


	// ------------------ 四、话题模式
	// // 注意,在接收消息的时候才能使用通配符,发送的时候不能使用
	l.RabbitMQ.PublishTopic("TopicExchange", "usa.news", "usa.news - " + req.Msg)
	l.RabbitMQ.PublishTopic("TopicExchange", "usa.weather", "usa.weather - " + req.Msg)
	l.RabbitMQ.PublishTopic("TopicExchange", "europe.news", "europe.news - " + req.Msg)
	l.RabbitMQ.PublishTopic("TopicExchange", "europe.weather", "europe.weather - " + req.Msg)


	fmt.Printf("\n200 success !!!\n\n")

	return nil
}


func (l *RabbitmqApiLogic) RabbitmqApiRPC(req *types.RabbitmqApiReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n RabbitmqApiRPC = %#v \n\n\n", req)


	// ------------------ 五、RPC通信模式
	var args map[string]any = map[string]any{
		"code": 200,
		"data": "RabbitmqApiRPC",
	}
	argsByte, err := json.Marshal(args)
	if err != nil {
		return err
	}

	response, err := l.RabbitMQ.RpcClientCall(context.Background(), "RpcServerQueue", argsByte, 5*time.Second)
	if err != nil {
		return err
	}

	fmt.Printf("\n\n RPC远程调用消息返回成功!!! RabbitmqApiRPCl.RabbitMQ.RpcCall = %s \n\n\n", string(response))

	fmt.Printf("\n200 success !!!\n\n")

	return nil
}
