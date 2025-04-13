package main

import "fmt"

// 14.2.9 用带缓冲通道实现一个信号量
func main9() {
	/*
	   	信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题，比如没有实现信号量的 sync 的 Go 包，使用带缓冲的通道可以轻松实现：

	      带缓冲通道的容量和要同步的资源容量相同
	      通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
	      容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）
	      不用管通道中存放的是什么，只关注长度；因此我们创建了一个长度可变但容量为 0（字节）的通道：
	*/
	// 将可用资源的数量 N 来初始化信号量 semaphore：
	//N := 5
	//sem := make(semaphore, N)

	//test01()

	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)

	<-done

}

// integer producer:
func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
		fmt.Println("numGen")
	}
	// 显示的关闭通道，为的是 可以使用 for ... range来输出数据
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	//for true {
	//	num := <-in

	// for ... range 用到通道上可以向通道获取数据
	// 它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭。
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func test01() {
	//用这种习惯用法写一个程序，开启一个协程来计算 2 个整数的和并等待计算结果并打印出来。
	sem := make(semaphore, 2)

	sem.Lock()
	go func(num1, num2 int) {
		sem <- (num1 + num2)
	}(1, 2)
	sem.Unlock()
	fmt.Println(sem)

}

type Empty interface{}
type semaphore chan Empty

// acquire n resources，写
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources，读
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

//可以用来实现一个互斥的例子：
/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
