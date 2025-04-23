package main

import (
	"fmt"
	"sync"
	"time"
)

// sync包
func main03() {
	/*
			在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：

				方法名								功能
			(wg * WaitGroup) Add(delta int)		计数器+delta
			(wg *WaitGroup) Done()				计数器-1
			(wg *WaitGroup) Wait()				阻塞直到计数器变为0

		sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。
		每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。


	*/

	// Go 中的大多数结构体类型（如 sync.WaitGroup）在 零值状态 下就可以安全使用，不需要手动初始化。
	// 这行代码虽然没显式写 wg = sync.WaitGroup{}，但 wg 已经是个 合法的 WaitGroup 实例，它内部字段都已被初始化为对应零值。
	var wg sync.WaitGroup

	// 设置协程数量
	wg.Add(10)

	fmt.Println("开启协程执行")
	// 开10个协程计算
	for i := 0; i < 10; i++ {
		go func(goroutineNum int) {
			time.Sleep(time.Second * 1)
			fmt.Println("goroutineNum = ", goroutineNum)

			// 每一个协程完成之后计算减一
			wg.Done()
		}(i)
	}

	// 等待所有协程执行完毕
	wg.Wait()
	fmt.Println("所有协程执行完毕")

}
