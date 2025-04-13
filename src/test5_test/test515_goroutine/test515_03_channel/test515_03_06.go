package main

import "fmt"

// 14.2.6 协程中用通道输出结果
func main6() {
	/*
		为了知道计算何时完成，可以通过信道回报。在例子 go sum(bigArray) 中，要这样写：

		ch := make(chan int)
		go sum(bigArray, ch) // bigArray puts the calculated sum on ch
		// .. do something else for a while
		sum := <- ch // wait for, and retrieve the sum

		也可以使用通道来达到同步的目的，这个很有效的用法在传统计算机中称为信号量 (semaphore)。或者换个方式：通过通道发送信号告知处理已经完成（在协程中）。

		在其他协程运行时让 main 程序无限阻塞的通常做法是在 main() 函数的最后放置一个 select {}。

		也可以使用通道让 main 程序等待协程完成，就是所谓的信号量模式，我们会在接下来的部分讨论。
	*/

	// 以下：使用通道来返回子协程的计算结果

	var intChan chan int = make(chan int)
	var intArr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func(intArr2 []int, intChan2 chan int) {
		sum2 := 0
		for _, val := range intArr {
			sum2 += val
		}

		// 向通道里面传数据，以达到利用通道返回数据的目的
		intChan2 <- sum2

	}(intArr, intChan)

	// 等待通道完成，接收通道内部的数据
	sum := <-intChan
	fmt.Println("sum = ", sum)
	//fmt.Println("sum = ", <-intChan)

}
