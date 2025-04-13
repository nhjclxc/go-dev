package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 一个发送协程，一个接收协程，在此基础上，接收协程接收到异常数据的时候不能停止工作
	// 因此，在接收协程这里要写一个panic和recover机制

	// 定义一个双向通道
	var dataChan chan int = make(chan int)
	var exitCHan chan bool = make(chan bool)

	// 发送协程
	go func(sender chan<- int) {
		for i := 0; i < 10; i++ {
			sender <- i
		}
		close(sender)
	}(dataChan)

	// 接收协程
	go func(receiver <-chan int) {
		for val := range receiver {
			// 开启一个子协程去处理每一个数据
			go doWork(val)
		}
		exitCHan <- true
	}(dataChan)

	<-exitCHan

}

func doWork(val int) {
	// 处理每隔数据的异常
	defer func() {
		if err := recover(); err != nil {
			// 一般这个时候就是记日志，或做一些异常处理
			fmt.Println("doWork.err", err)
		}
	}()

	// 如果 doWork0(work) 发生 panic()，错误会被记录且协程会退出并释放，而其他协程不受影响。

	// 处理每一个数据
	doWork0(val)
}
func doWork0(val int) {

	if val%2 != 0 {
		panic("数据不能被2整除，val = " + strconv.Itoa(val))
	}
	fmt.Println("receiver: ", val, "处理完毕！！！")
}
