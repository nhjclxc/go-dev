package main

import (
	"fmt"
	"sync"
	"time"
)

/*
自己实现

			(wg *WaitGroup) Add(delta int)		计数器+delta
			(wg *WaitGroup) Done()				计数器-1
			(wg *WaitGroup) Wait()				阻塞直到计数器变为0
*/

func main04() {

	var mwg MyWaitGroup

	// 设置协程数量
	mwg.Add(10)

	fmt.Println("开启协程执行")
	// 开10个协程计算
	for i := 0; i < 10; i++ {
		go func(goroutineNum int) {
			time.Sleep(time.Second * 1)
			fmt.Println("goroutineNum = ", goroutineNum)

			// 每一个协程完成之后计算减一
			mwg.Done()
		}(i)
	}

	// 等待所有协程执行完毕
	mwg.Wait()
	fmt.Println("所有协程执行完毕")
}

type MyWaitGroup struct {
	count   int
	counter int
	intChan chan int
	mu      sync.Mutex
}

func (this *MyWaitGroup) Add(count int) {
	this.count = count
	this.counter = 0
	this.intChan = make(chan int, count)
}

func (this *MyWaitGroup) Done() {
	this.mu.Lock()
	this.intChan <- 1
	this.counter++
	if this.counter == this.count {
		close(this.intChan)
	}
	this.mu.Unlock()
}

func (this *MyWaitGroup) Wait() {
	for {
		if _, ok := <-this.intChan; ok {
			if this.count == this.counter {
				break
			}
		}
	}
}

type SafeWaitGroup struct {
	mu    sync.Mutex
	count int
	cond  *sync.Cond
}

func NewSafeWaitGroup() *SafeWaitGroup {
	wg := &SafeWaitGroup{}
	wg.cond = sync.NewCond(&wg.mu)
	return wg
}

func (wg *SafeWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.count += delta
}

func (wg *SafeWaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.count--
	if wg.count == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *SafeWaitGroup) Wait() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	for wg.count > 0 {
		wg.cond.Wait()
	}
}
