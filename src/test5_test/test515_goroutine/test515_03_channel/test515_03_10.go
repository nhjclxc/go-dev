package main

import (
	"fmt"
	"math/rand"
)

// 14.2.11 通道的方向
func main10() {
	/*
			通道的三种类型
				类型			可做的操作		描述				示例
				双向通道		发送 & 接收		可发送也可接收	chan int
				只发送通道	只能发送（写入）	只能发送数据		chan<- int， <- 在chan后面表示只能向通道中发送数据
				只接收通道	只能接收（读取）	只能接收数据		<-chan int， <- 在chan前面表示只能从通道中取出数据
			注意：只有能发送的通道才能使用close()函数来关闭通道

		例如，规定生产者只能使用发送的通道，消费组只能使用发送的通道
	*/

	test515_03_10_01()
}

func test515_03_10_01() {
	//写一个典型的 生产者 - 消费者模型，用 Go 的 channel + goroutine + 方向通道 来实现。
	/*
	   	🧩 场景描述
	      生产者：负责生成数据（比如随机整数）
	      消费者：负责读取数据并处理
	      通道（channel）：用作缓冲区
	      用方向通道明确职责：
	   	   生产者函数：chan<- int → 只发送
	   	   消费者函数：<-chan int → 只接收
	*/

	// 定义一个双向通道
	var intChan chan int = make(chan int, 3)

	// 使用一个阻塞通道来控制程序的关闭，从而无需使用sleep方法
	var boolChan chan bool = make(chan bool)

	// 定义生产者 producer
	// 定义生产者只能使用只发送通道
	go func(onlySendIntChan chan<- int) {
		// 为了能够使用for ... range 来读取数据，这里显示的使用close来关闭通道
		defer close(onlySendIntChan)

		for i := 0; i < 10; i++ {
			onlySendIntChan <- rand.Intn(100)
		}
	}(intChan)

	// 定义消费者 consumer
	// 定义消费者的只读通道
	go func(readOnlyIntChan <-chan int) {
		for val := range readOnlyIntChan {
			fmt.Println("consumer: ", val)
		}
		if v, ok := <-readOnlyIntChan; ok {
			fmt.Println(v)
		}
		boolChan <- true
	}(intChan)

	//time.Sleep(5 * time.Second)

	// 表示从通道中读取数据，在消费组里面如果没有向通道中传入数据的话，主协程就一直卡在这里等待，从而优化了主协程的开关
	<-boolChan

}
