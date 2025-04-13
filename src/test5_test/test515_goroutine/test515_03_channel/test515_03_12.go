package main

import "fmt"

func main12() {

	// 两个协程交替打印1-10之间的奇偶数
	var num = 1

	var lockChan chan bool = make(chan bool)
	var exitChan chan bool = make(chan bool)

	// 协程1，打印奇数
	go func() {
		for true {
			if num <= 10 {
				if num%2 == 1 {
					fmt.Println("goroutine-1：", num)
					num++
					lockChan <- true
				}
			} else {
				break
			}
		}
		exitChan <- true
	}()

	// 协程2，打印偶数
	go func() {
		for true {
			_, ok := <-lockChan
			if ok && num <= 10 {
				if num%2 == 0 {
					fmt.Println("goroutine-2：", num)
					num++
					//<-lockChan
				}
			} else {
				break
			}
		}
		exitChan <- true
	}()

	// 阻塞主协程
	<-exitChan

}
