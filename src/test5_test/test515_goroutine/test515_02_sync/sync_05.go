package main

import (
	"fmt"
	"sync"
)

// 1.1.2. sync.Once
func main05() {

	/*
		在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

		Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。

		sync.Once只有一个Do方法，其签名如下：

		func (o *Once) Do(f func()) {}
		注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。
	*/

	// 使用 sync.Once() 实现懒加载初始化文件或数据

	/*
		加载配置文件示例
		延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践。
		因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，
		而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。
		我们来看一个例子：
	*/
	// sync.Once 的作用
	//它确保 某段代码只会被执行一次，无论多少 goroutine 调用了 Do()。
	//是线程安全的，适合懒加载、只初始化一次的场景。

	var lazyLoading sync.Once

	var data string

	var wg sync.WaitGroup

	wg.Add(10)

	// 开10个协程取加载
	for i := 0; i < 10; i++ {
		go func(num int) {
			lazyLoading.Do(func() {
				// 模拟这里取加载配置文件
				// ...
				data = "This is profile !!!"
				fmt.Println("配置文件加载者 ", num)
			})

			fmt.Println("num = ", num, "拿到配置文件了，好开心，配置文件是：", data)

			wg.Done()
		}(i)
	}

	wg.Wait()
}
