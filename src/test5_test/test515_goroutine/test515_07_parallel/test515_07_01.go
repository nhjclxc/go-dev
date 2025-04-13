package main

import (
	"fmt"
	"sync"
)

func main01() {

	/*
		旧模式：使用共享内存进行同步

		新模式：
			假设我们需要处理很多任务；一个 worker 处理一项任务。任务可以被定义为一个结构体（具体的细节在这里并不重要）：
			type Task struct {
				// some state
			}
		在旧模式中，由各个任务组成的任务池共享内存；为了同步各个 worker 以及避免资源竞争，我们需要对任务池（类似于Java的线程池）进行加锁保护：
			type Pool struct {
				Mu      sync.Mutex
				Tasks   []*Task
			}
		sync.Mutex（参见9.3）是互斥锁：它用来在代码中保护临界区资源：同一时间只有一个 go 协程 (goroutine) 可以进入该临界区。
			如果出现了同一时间多个 go 协程都进入了该临界区，则会产生竞争：Pool 结构就不能保证被正确更新。


	*/

	// 以下的处理是旧模式，新模式看test515_07_02.go的通道模式
	var exitCHan chan bool = make(chan bool)

	var goroutinePool Pool = Pool{
		Mu:       sync.Mutex{},
		taskList: []*Task{},
	}

	// 创建10个任务
	for i := 0; i < 10; i++ {
		goroutinePool.taskList = append(goroutinePool.taskList, &Task{val: i * 10})
	}

	count := 0
	// 开5个协程去互斥消费任务
	for i := 0; i < 10; i++ {
		go func(task *Task) {
			// 开始消费上锁
			goroutinePool.Mu.Lock()
			vallll := (*task).val
			count++
			fmt.Println("资源消费：vallll=", vallll, count)
			// 消费结束解锁
			goroutinePool.Mu.Unlock()
		}(goroutinePool.taskList[i])

		if count >= 9 {
			exitCHan <- true
		}
	}

	<-exitCHan
}

type Task struct {
	val int
}

type Pool struct {

	// 控制互斥资源
	Mu sync.Mutex

	// 协程池里面的所有任务
	taskList []*Task
}
