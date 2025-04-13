package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	// 14.2.2 通信操作符 <-

	/*
		 这个操作符直观的表示了数据的传输：信息按照箭头的方向流动。
			流向通道（发送）
				ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）
			从通道流出（接收），三种方式：
				int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。
		<- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：
			// 一直向通道中读取数据，直到数据是1000的时候进行某种操作
			if <- ch != 1000{
				...
			}

		同一个操作符 <- 既用于发送也用于接收，但 Go 会根据操作对象弄明白该干什么 。虽非强制要求，但为了可读性通道的命名通常以 ch 开头或者包含 chan 。
		通道的发送和接收都是原子操作：它们总是互不干扰地完成。


		// 通道变量在 <- 左边，表示向通道内放数据，如：strChan <- "Hello"，将 “hello”放入通道
		// 通道变量在 <- 右边，表示向通道内取数据，如：data = <- strChan，将 strChan 通道内的数据取出并放入data变量

	*/

	// 1、声明一个通道
	var strChan chan string = make(chan string)

	// 2、启动一个协程，模拟数据发送者，sendData
	go func() {
		//func() {
		strChan <- "Hello"
		strChan <- ","
		strChan <- "Golang"
		strChan <- " "
		strChan <- "World"
	}()

	// 3、启动一个协程，模拟一个接收者，receiveData
	go func() {
		//func() {
		//time.Sleep(1 * time.Second)
		//time.Sleep(5 * time.Second)
		str := ""
		for {
			str = <-strChan
			fmt.Println("receiveData: " + str)

			if str == "," {
				fmt.Println("receiveData goroutine exit")
				return
			}
		}

	}()

	// 休眠防止程序提前退出
	time.Sleep(3 * time.Second)

	/*
		我们发现协程之间的同步非常重要：

		main() 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
		getData() 使用了无限循环：它随着 sendData() 的发送完成和 ch 变空也结束了。
		如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 panic：
		---- Error run E:/Go/Goboek/code examples/chapter 14/goroutine2.exe with code Crashed ---- Program exited with code -2147483645: panic: all goroutines are asleep-deadlock!
		为什么会这样？运行时 (runtime) 会检查所有的协程（像本例中只有一个）是否在等待着什么东西（可从某个通道读取或者写入某个通道），这意味着程序将无法继续执行。这是死锁 (deadlock) 的一种形式，而运行时 (runtime) 可以为我们检测到这种情况。

		注意：不要使用打印状态来表明通道的发送和接收顺序：由于打印状态和通道实际发生读写的时间延迟会导致和真实发生的顺序不同。



		练习 14.4：解释一下为什么如果在函数 getData() 的一开始插入 time.Sleep(2e9)，不会出现错误但也没有输出呢。
			答：此时主协程已经退出了，子协程无法进行

		go为什么主协程退出之后，子协程无法正常使用了？？？
			这是 Go 的一个非常核心的行为特性：当主协程（main() 函数）退出后，整个程序就会立即终止，不管其他子协程（goroutine）有没有执行完。
			🔍 原因解析：Go 中的 main() 函数就是程序的入口，main() 所在的 goroutine 是“主 goroutine”。当这个主 goroutine 执行完成，Go 的运行时（runtime）就会退出整个程序，包括所有还在运行中的子 goroutine。
			解决方法：
				✅ 方法 1：使用 sync.WaitGroup
				✅ 方法 2：time.Sleep()（仅适合测试）

	*/

	// 1、首先定义一个等待组
	var wg sync.WaitGroup
	wg.Add(1)

	// 开启一个子协程
	go func() {
		// 使用defer 结合wg.Done() 来表示这个函数介绍后也就是这个协程执行完毕的信号
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("子协程执行完毕")
	}()

	// 主协程等待所有子协程的完成
	wg.Wait() // 等待所有子 goroutine 执行完
	fmt.Println("主协程退出")

}
