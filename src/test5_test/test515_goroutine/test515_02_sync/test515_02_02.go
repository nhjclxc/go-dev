package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex.Lock() 和 sync.Mutex.TryLock()的区别
// sync.Mutex.Lock() 和 sync.Mutex.TryLock() 的主要区别在于 是否会阻塞 以及 锁的获取方式。
func main2() {

	// 测试 sync.Mutex.Lock()
	//test515_02_02_01()

	// 测试 sync.Mutex.TryLock()
	//test515_02_02_02()

	//test515_02_02_03()

	test515_02_02_04()

}

func test515_02_02_01() {
	// 1. sync.Mutex.Lock()
	//阻塞式加锁：如果锁已被其他 goroutine 持有，则当前 goroutine 会一直阻塞，直到锁可用。
	//适用于必须获取锁的情况，例如修改共享数据时需要确保互斥访问。

	// 声明一个互斥锁变量
	var mu sync.Mutex

	// 开启一个协程模拟执行
	go func() {
		mu.Lock()
		fmt.Println("Goroutine 1: 获取到锁，开始执行")
		time.Sleep(2 * time.Second) // 模拟长时间持有锁
		mu.Unlock()
	}()

	// 主协程
	time.Sleep(500 * time.Millisecond) // 确保 Goroutine 1 先获取到锁

	fmt.Println("Main: 尝试获取锁")
	mu.Lock() // 这里会阻塞，直到 Goroutine 1 释放锁  可以看到，mu.Lock() 在 main 函数中会 阻塞，直到 Goroutine 1 释放锁。
	fmt.Println("Main: 获取到锁")
	mu.Unlock()
	// Goroutine 1: 获取到锁，开始执行
	//Main: 尝试获取锁
	//Main: 获取到锁
}

func test515_02_02_02() {
	// 2. sync.Mutex.TryLock()
	//非阻塞式加锁：如果锁已被其他 goroutine 持有，它会立即返回 false，不会阻塞。
	//返回 true 代表成功获取锁，返回 false 代表锁已被占用。
	//适用于尝试获取锁但不希望阻塞的情况，例如避免死锁或者实现超时控制。
	//Go 1.18 及以上版本 才支持 TryLock() 方法。

	var mu sync.Mutex

	go func() {
		mu.Lock()
		fmt.Println("Goroutine 1: 获取到锁")
		time.Sleep(2 * time.Second)
		mu.Unlock()
	}()

	time.Sleep(500 * time.Millisecond) // 确保 Goroutine 1 先获取到锁

	fmt.Println("Main: 尝试获取锁")
	if mu.TryLock() {
		fmt.Println("Main: 成功获取到锁")
		// 如果返回true，表示获取成功，则再获取成功之后要释放
		mu.Unlock()
	} else {
		fmt.Println("Main: 获取锁失败，锁已被其他 goroutine 持有")
	}
}

/*
3. Lock() vs TryLock() 总结
方法				是否阻塞		是否一定获取到锁		适用场景
Lock()			是			是（最终）			需要确保获取锁的情况
TryLock()		否			不一定				不希望等待锁，避免死锁，或者尝试执行非关键任务
如果你的场景需要 强制等待并获取锁，使用 Lock()；
如果你的场景是 尽量获取锁但不会等待（例如任务超时控制、非关键任务），使用 TryLock()。
*/

func test515_02_02_03() {

	var mu sync.Mutex

	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			fmt.Println("goroutine 1.count = ", i)
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			fmt.Println("goroutine 2.count = ", i)
			mu.Unlock()
		}
	}()

	for i := 0; i < 10; i++ {
		mu.Lock()
		fmt.Println("goroutine main.count = ", i)
		mu.Unlock()
	}

	// 休眠10s防止提前退出
	time.Sleep(10 * time.Second)

}

func test515_02_02_04() {

	var mu sync.Mutex

	go func() {
		for i := 0; i < 10; i++ {
			if mu.TryLock() {
				fmt.Println("goroutine 1.count = ", i)
				mu.Unlock()
			} else {
				fmt.Println("goroutine 1.TryLock = false")
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if mu.TryLock() {
				fmt.Println("goroutine 2.count = ", i)
				mu.Unlock()
			} else {
				fmt.Println("goroutine 2.TryLock = false")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		if mu.TryLock() {
			fmt.Println("goroutine main.count = ", i)
			mu.Unlock()
		} else {
			fmt.Println("goroutine main.TryLock = false")
		}
	}

	// 休眠10s防止提前退出
	time.Sleep(10 * time.Second)

}
