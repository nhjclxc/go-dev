package main

import (
	"fmt"
	"time"
)

func main() {

	// 新模式：通道模式
	// 使用通道进行同步：使用一个通道接受需要处理的任务，一个通道接受处理完成的任务（及其结果）。worker 在协程中启动，其数量 N 应该根据任务数量进行调整。

	// 首先，将任务一个发送给处理通道，
	// 接着，在处理通道接收到数据后进行数据处理，
	// 然后，当处理通道将数据处理完毕之后，将该数据发送到完成通道里面

	// 定义发送通道和处理通道
	sender := make(chan interface{})
	done := make(chan interface{})

	// 启动数据后处理协程
	go func() {
		for data := range done {
			fmt.Println("done =", data)
		}
	}()

	// 启动多个处理协程
	for i := 0; i < 5; i++ {
		go func() {
			for data := range sender {
				fmt.Println("sender =", data)
				// 模拟处理
				time.Sleep(500 * time.Millisecond)
				done <- data
			}
		}()
	}

	// 发送任务
	go func() {
		for i := 0; i < 5; i++ {
			sender <- i
		}
		// 所有任务发完后关闭 sender 通道，通知处理协程退出
		close(sender)
	}()

	// 等待处理完（这里只是简单 sleep，为演示方便）
	time.Sleep(3 * time.Second)
	fmt.Println("所有数据处理完毕！！！")
}

func Worker(in, out chan *Task) {
	for {
		t := <-in
		//process(t)
		out <- t
	}
}
