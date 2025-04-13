package main

import "fmt"

// 两个协程交替打印1-20之间的奇偶数
func main13() {
	var num = 1
	var lockAChan chan bool = make(chan bool)
	var lockBChan chan bool = make(chan bool)
	var exitChan chan bool = make(chan bool)

	go func() {
		for num <= 20 {
			<-lockBChan
			fmt.Println("goroutine-1：", num)
			num++
			lockAChan <- true
		}
		exitChan <- true
	}()

	go func() {
		for num <= 20 {
			lockBChan <- true
			<-lockAChan
			fmt.Println("goroutine-2：", num)
			num++
		}
		exitChan <- true
	}()
	<-exitChan
}
