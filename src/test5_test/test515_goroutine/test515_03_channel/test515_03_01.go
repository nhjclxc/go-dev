package main

import (
	"fmt"
)

// 14.2 协程间的信道
// https://denganliang.github.io/the-way-to-go_ZH_CN/14.2.html
func main1() {

	// 14.2.1 概念

	// 协程是独立执行的，他们之间没有通信。他们必须通信才会变得更有用：彼此之间发送和接收信息并且协调/同步他们的工作。
	//协程可以使用共享变量来通信，但是很不提倡这样做，因为这种方式给所有的共享内存的多线程都带来了困难。

	// 而 Go 有一种特殊的类型，通道（channel），就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信，从而避开所有由共享内存导致的陷阱；
	//这种通过通道进行通信的方式保证了同步性。数据在通道中进行传递：在任何给定时间，一个数据被设计为只有一个协程可以对其访问，所以不会发生数据竞争。
	//数据的所有权（可以读写数据的能力）也因此被传递。

	// 通道服务于通信的两个目的：值的交换，同步的，保证了两个计算（协程）任何时候都是可知状态。

	// 通常使用这样的格式来声明通道：var identifier chan datatype，如：var intChan chan int，表示intChan是int类型的chan通道
	//未初始化的通道的值是 nil。
	// 所以通道只能传输一种类型的数据，比如 chan int 或者 chan string，所有的类型都可以用于通道，空接口 interface{} 也可以，甚至可以（有时非常有用）创建通道的通道。

	//通道实际上是类型化消息的队列：使数据得以传输。它是先进先出(FIFO) 的结构所以可以保证发送给他们的元素的顺序。
	//通道也是引用类型，所以我们使用 make() 函数来给它分配内存。这里先声明了一个字符串通道 strChan，然后创建了它（实例化）：
	var strChan chan string
	fmt.Println(strChan)
	strChan = make(chan string)
	fmt.Println(strChan)

	//var strChan2 chan string = make(chan string)
	//strChan3 := make(chan string)

	// 通道的通道，后面的【chan int】是整体
	//chanOfChans := make(chan chan int)

	// 函数通道
	//funcChan := make(chan func())

	// 所以通道是第一类对象：可以存储在变量中，作为函数的参数传递，从函数返回以及通过通道发送它们自身。
	//另外它们是类型化的，允许类型检查，比如尝试使用整数通道发送一个指针。
}
