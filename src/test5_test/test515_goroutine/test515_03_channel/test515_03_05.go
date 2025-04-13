package main

import (
	"fmt"
	"time"
)

// 14.2.5 同步通道-使用带缓冲的通道
func main5() {
	/*
		一个无缓冲通道，一次通信只能向通道内放入 1 个元素，有时显得很局限。
		因此，我们给通道提供了一个缓存，可以在扩展的 make 命令中设置它的容量，如下：
		buf := 100
		ch1 := make(chan string, buf)
		这里的 buf 是通道可以同时容纳的元素（这里是 string）个数

		在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。

		缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 cap() 函数可以返回缓冲区的容量。

		如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是 0 或者未设置，通信仅在收发双方准备好的情况下才可以成功。


		使用ch :=make(chan type, value)创建一个通道时，判断通道是同步的还是异步的主要取决于 value 元素
			value == 0 -> synchronous（同步阻塞的）, unbuffered （阻塞）
			value > 0 -> asynchronous（异步非阻塞的）, buffered（非阻塞）
		若使用通道的缓冲，你的程序会在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性(scalable)。
		在设计算法时首先考虑使用无缓冲通道，只在不确定的情况下使用缓冲。



	*/

	c := make(chan int, 5)
	//c := make(chan int)
	go func() {
		//time.Sleep(1 * time.Second)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10 // 现在通道c有缓存，不必等待放入的10被消费，即可立即执行吓一跳语句
	fmt.Println("sent", 10)
	// sending 10
	// sent 10

	//time.Sleep(3 * time.Second)
	// sending 10
	//sent 10
	//received 10

	go func() {
		//time.Sleep(1 * time.Second)
		//xx := make([]int, 5)
		for true {
			// 会从缓冲区中按照先进先出的顺序一个一个消费
			xx := <-c
			fmt.Println("received", xx)
		}
	}()
	c <- 101
	c <- 102
	c <- 103
	c <- 104
	c <- 105

	time.Sleep(3 * time.Second)
}
