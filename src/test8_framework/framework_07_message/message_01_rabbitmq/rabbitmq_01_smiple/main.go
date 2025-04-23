package main

func main() {

	// https://www.topgoer.cn/docs/golang/chapter10-6-3

	// 消息产生着§将消息放入队列
	//消息的消费者(consumer) 监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)
	//做simple简单模式之前首先我们新建一个Virtual Host并且给他分配一个用户名，用来隔离数据，根据自己需要自行创建

	/*
		kuteng-RabbitMQ
		RabbitMQ
			–rabitmq.go //这个是RabbitMQ的封装
		SimlpePublish
			–mainSimlpePublish.go //Publish 先启动
		SimpleRecieve
			–mainSimpleRecieve.go
	*/

	// 先启动，消息接收者去阻塞接收消息，再启动发布者发布消息
}
