package main

import (
	"fmt"
	"time"
)

// 定时器
// https://topgoer.com/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/%E5%AE%9A%E6%97%B6%E5%99%A8.html
func main() {

	//timer01()

	timer02()

}

func timer02() {

	// Ticker：每隔多久执行一次

	// 1.获取ticker对象
	ticker := time.NewTicker(2 * time.Second)
	i := 0

	for {
		i++
		fmt.Println(<-ticker.C)
		if i == 5 {
			//停止
			ticker.Stop()
			return
		}
	}

}

func timer01() {
	// 1.timer基本使用

	// timer只能响应1次
	// timer实现延时的功能

	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	// 等待2秒以后触发
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	// 重置定时器
	timer1.Reset(5 * time.Second)

	// 停止定时器
	timer1.Stop()

}
