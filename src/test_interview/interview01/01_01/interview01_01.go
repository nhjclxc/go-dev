package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	/*
		在go中有一个ticker定时任务每隔3s执行一次，但是每一次处理任务需要5s。
		请问做过定时任务会发生什么？能否正常执行任务？该怎么优化
	*/

	// 问题对应的代码：

	//ticker := time.NewTicker(time.Second * time.Duration(3))
	//defer ticker.Stop()
	//
	//counter := 0
	//// 按固定间隔执行任务
	////for c := range ticker.C {
	//for range ticker.C {
	//	fmt.Println(strconv.Itoa(counter) + " 开始执行 = " + time.Now().String())
	//	time.Sleep(time.Second * time.Duration(5))
	//	fmt.Println(strconv.Itoa(counter) + " 结束执行 = " + time.Now().String())
	//	fmt.Println()
	//	counter = counter + 1
	//}

	/*
		执行结构如下：
		【
			从以下结果中可以看出，定时任务实际变为了每隔5秒执行一次，会引发以下 定时任务阻塞与堆积问题。
				由此我们可以推断出一个结论：定时任务的实际执行间隔 = Min(定时任务定义的时间间隔, 定时任务的实际执行耗时)
		】

		0 开始执行 = 2025-06-29 20:39:34.7319152 +0800 CST m=+3.001414201
		0 结束执行 = 2025-06-29 20:39:39.7415189 +0800 CST m=+8.011017901

		1 开始执行 = 2025-06-29 20:39:39.7418758 +0800 CST m=+8.011374801
		1 结束执行 = 2025-06-29 20:39:44.7430375 +0800 CST m=+13.012536501

		2 开始执行 = 2025-06-29 20:39:44.7430375 +0800 CST m=+13.012536501
		2 结束执行 = 2025-06-29 20:39:49.743787 +0800 CST m=+18.013286001

		3 开始执行 = 2025-06-29 20:39:49.743787 +0800 CST m=+18.013286001
		3 结束执行 = 2025-06-29 20:39:54.7442723 +0800 CST m=+23.013771301

		4 开始执行 = 2025-06-29 20:39:54.7442723 +0800 CST m=+23.013771301
		4 结束执行 = 2025-06-29 20:39:59.7449568 +0800 CST m=+28.014455801

		5 开始执行 = 2025-06-29 20:39:59.7466433 +0800 CST m=+28.016142301
		5 结束执行 = 2025-06-29 20:40:04.7475683 +0800 CST m=+33.017067301

		6 开始执行 = 2025-06-29 20:40:04.7475683 +0800 CST m=+33.017067301

	*/

	// ✅ 如何优化？
	//
	//optimize1()
	//
	//optimize2()

	//optimize3()

	optimize5()

}

func optimize5() {
	/*
	“✅ 方案三：任务入队 + 单独消费者处理”，也不是最理想的，因为定时任务每隔三秒就会生产一个任务，
	而虽然使用了缓冲队列，但是缓冲队列也会有被填满的时候，这时候该怎么办？
	 */
	// ✅ 最优方案：任务入队 + 多消费者并行处理

	// 1、定义一个任务队列
	var taskQueue chan int = make(chan int, 100)

	// 2、处理任务队列里面的任务
	go doTask(taskQueue, 1)
	go doTask(taskQueue, 2)

	// 3、定时任务开启任务的执行

	ticker := time.NewTicker(time.Second * time.Duration(3))
	defer ticker.Stop()
	taskId := 0

	// 按固定间隔执行任务
	for range ticker.C {
		// 往任务队列里面塞数据

		taskQueue <- taskId
		taskId = taskId + 1
	}

}

func doTask(taskQueue chan int, workerId int) {
	for taskId := range taskQueue {
		fmt.Println(strconv.Itoa(workerId) + ": len(taskQueue) = " + strconv.Itoa(len(taskQueue)))
		fmt.Println(strconv.Itoa(taskId) + " 开始执行 = " + time.Now().String())
		time.Sleep(time.Second * time.Duration(5))
		fmt.Println(strconv.Itoa(taskId) + " 结束执行 = " + time.Now().String())
		fmt.Println()
	}
}

func optimize3() {
	// ✅ 方案三：任务入队 + 单独消费者处理
	//如果任务不可并发执行，也可以：
	//每 3 秒将任务放入一个 channel（任务队列）
	//单独一个 goroutine 持续处理 channel 中的任务（串行消费）


	// 1、定义一个任务队列
	var taskQueue chan int = make(chan int, 100)

	// 2、处理任务队列里面的任务
	go func() {
		//taskId := <- taskQueue
		for taskId :=  range taskQueue {
			fmt.Println("len(taskQueue) = " + strconv.Itoa(len(taskQueue)))
			fmt.Println(strconv.Itoa(taskId) + " 开始执行 = " + time.Now().String())
			time.Sleep(time.Second * time.Duration(5))
			fmt.Println(strconv.Itoa(taskId) + " 结束执行 = " + time.Now().String())
			fmt.Println()
		}
	}()

	// 3、定时任务开启任务的执行

	ticker := time.NewTicker(time.Second * time.Duration(3))
	defer ticker.Stop()
	taskId := 0

	// 按固定间隔执行任务
	for range ticker.C {
		// 往任务队列里面塞数据

		taskQueue <- taskId

		taskId = taskId + 1
	}


	/*
	特点：
		定时生成任务（3 秒）
		消费任务是串行执行（5 秒一个）
		保证任务 不丢失，可根据需要扩展为多协程消费
	 */
}

func optimize2() {
	// ✅ 方案二：带并发限制的任务处理池（推荐）  使用一个带缓冲 channel 控制最大并发数：

	const maxConcurrent = 5
	sem := make(chan struct{}, maxConcurrent) // 控制并发上限

	ticker := time.NewTicker(time.Second * time.Duration(3))
	defer ticker.Stop()
	counter := 0
	// 按固定间隔执行任务
	for range ticker.C {

		// 往 sem 通道里面放数据，如果放进去了就说明并发数小于5个，让其执行本次的并发任务
		// 如果没放进去就说明目前任务的并发数已经超过了5个，阻塞任务执行，待有位置了在放进入
		sem <- struct{}{} // 获取令牌

		// 异步执行实际的处理任务
		go func() {

			defer func() { <-sem }() // 释放令牌，消耗一个缓冲区的资源

			fmt.Println(strconv.Itoa(counter) + " 开始执行 = " + time.Now().String())
			time.Sleep(time.Second * time.Duration(5))
			fmt.Println(strconv.Itoa(counter) + " 结束执行 = " + time.Now().String())
			fmt.Println()
			counter = counter + 1
		}()
	}

	/*
	优点：
		保证每 3 秒都能调度新任务
		避免 goroutine 无限制增长
	 */



}

func optimize1() {
	// ✅ 方案一：任务异步执行（启动新 goroutine 处理任务）
	ticker := time.NewTicker(time.Second * time.Duration(3))
	defer ticker.Stop()
	counter := 0
	// 按固定间隔执行任务
	for range ticker.C {

		// 异步执行实际的处理任务
		go func() {
			fmt.Println(strconv.Itoa(counter) + " 开始执行 = " + time.Now().String())
			time.Sleep(time.Second * time.Duration(5))
			fmt.Println(strconv.Itoa(counter) + " 结束执行 = " + time.Now().String())
			fmt.Println()
			counter = counter + 1
		}()
	}

	/*
		优点：
			每 3 秒都能响应一次 ticker 事件
			每个任务单独在一个 goroutine 中处理，不互相阻塞

		缺点：
			若任务执行时间过长、内存占用大或无法重入，可能造成 goroutine 数量无限增长
			需要考虑 任务并发安全、限流
	*/

}
