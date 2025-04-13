package main

import (
	"fmt"
	"time"
)

// 14.2.3 通道阻塞
func main3() {
	/*
		默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束（也就是要将通道里面的那个数据取出去了，发送者才能接着发送下一个数据，接收者则一个一个数据接收）。
		可以想象一个无缓冲的通道在没有空间来保存数据的时候：必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。
		所以通道的发送/接收操作在对方准备好之前是阻塞的：
			1）对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果 ch 中的数据无人接收，就无法再给通道传入其他数据：
				新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。
			2）对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

	*/

	// 定义一个全局变量来演示上面的效果
	var count int = 0

	// 通道变量
	var intChan chan int = make(chan int)

	// 等待组
	//var wg sync.WaitGroup

	// 发送者
	go func() {
		//defer wg.Done()
		for count = 1; count < 10; count++ {
			intChan <- count
			fmt.Println("sned count = ", count)
		}
	}()

	// 接收者
	go func() {
		//defer wg.Done()
		intVal := 0
		// 必须使用死循环去接收数据
		for true {
			intVal = <-intChan
			fmt.Println("receive count = ", count, ", intVal = ", intVal)
		}
	}()

	fmt.Println("main.start")
	//wg.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("main.end")
	/*
	   main.start
	   receive count =  1 , intVal =  1
	   sned count =  1
	   sned count =  2
	   receive count =  2 , intVal =  2
	   receive count =  3 , intVal =  3
	   sned count =  3
	   sned count =  4
	   receive count =  4 , intVal =  4
	   receive count =  5 , intVal =  5
	   sned count =  5
	   sned count =  6
	   receive count =  6 , intVal =  6
	   receive count =  7 , intVal =  7
	   sned count =  7
	   sned count =  8
	   receive count =  8 , intVal =  8
	   receive count =  9 , intVal =  9
	   sned count =  9
	   main.end

	*/

	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)

	/*
		sending 10
		(15 s later):
		received 10
		sent 10
	*/
}
