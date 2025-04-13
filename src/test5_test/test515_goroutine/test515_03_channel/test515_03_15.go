package main

import (
	"fmt"
	"time"
)

func main15() {

	// 开启一个协程，每隔一秒输出一个 hello, goroutine
	// 在主线程中，每隔一秒输出一个 hello, world，输出10此后退出程序

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello, goroutine ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("hello, world ", i)
		time.Sleep(1 * time.Second)
	}

}
