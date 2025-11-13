package main

import (
	"fmt"
	"os"
	"time"
)

// 14.4 使用 select 切换协程
func main1() {
	/*
	   	从不同的并发执行的协程中获取值可以通过关键字 select 来完成，它和 switch 控制语句非常相似（章节 5.3）也被称作通信开关；它的行为像是“你准备好了吗”的轮询机制；select 监听进入通道的数据，也可以是用通道发送值的时候。
	   	   select {
	   	   case u:= <- ch1:
	   			   ...
	   	   case v:= <- ch2:
	   			   ...
	   			   ...
	   	   default: // no value ready to be received
	   			   ...
	   	   }

	      default 语句是可选的；fallthrough 行为，和普通的 switch 相似，是不允许的。在任何一个 case 中执行 break 或者 return，select 就结束了。

	      select 做的就是：选择处理列出的多个通信情况中的一个。
	   	   如果都阻塞了，会等待直到其中一个可以处理
	   	   如果多个可以处理，随机选择一个
	   	   如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。
	      在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！如果没有 default，select 就会一直阻塞。
	*/

	//test515_02_12_01()
	//test515_02_12_02()
	test515_02_12_03()

}

func test515_02_12_03() {
}

func sender(intChan chan<- int, exit chan<- bool) {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	exit <- true
}

func test515_02_12_02() {

	var intChan chan int = make(chan int)

	var exitChan chan bool = make(chan bool)

	go sender(intChan, exitChan)

	for true {
		select {
		case val := <-intChan:
			fmt.Println("val = ", val)
		case <-exitChan:
			os.Exit(0)
		}
	}

}

func test515_02_12_01() {
	intChan1 := make(chan int)
	intChan2 := make(chan int)

	// 往 intChan1 生产数据
	go func() {
		for i := 0; i < 5; i++ {
			intChan1 <- i
		}
	}()

	// 往 intChan2 生产数据
	go func() {
		for i := 550; i < 555; i++ {
			intChan2 <- i
		}
	}()

	// 消费者
	go func() {
		for {
			select {
			//  如果 intChan1 一直没有关闭，不会一直阻塞而deadlock，会自动到下一个case匹配数据
			case v1 := <-intChan1:
				fmt.Println("intChan1.v1 = ", v1)
			case v2 := <-intChan2:
				fmt.Println("intChan2.v2 = ", v2)
				//default:
				//	fmt.Println("default")
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
