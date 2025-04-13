package main

import (
	"fmt"
	"runtime"
	"time"
)

func task01(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, "执行第", i, "次")
		time.Sleep(time.Millisecond * 500) // 模拟任务执行
	}
}

func test515_01_01_01() {
	// 认识1. 并发（Concurrency）
	go task01("任务A") // 并发执行任务A
	go task01("任务B") // 并发执行任务B

	time.Sleep(time.Second * 3) // 防止主 goroutine 结束
	fmt.Println("主程序退出")
	// 可以看到，任务 A 和任务 B 在“交替”执行，而不是严格的串行执行，这就是 并发 的体现。
}

func task02(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d 执行第 %d 次\n", id, i)
		time.Sleep(time.Millisecond * 500)
	}
}
func test515_01_01_02() {
	runtime.GOMAXPROCS(2) // 设置最多使用 2 个 CPU 核心

	for i := 0; i < 4; i++ {
		go task02(i)
	}

	time.Sleep(time.Second * 3) // 等待 goroutine 执行完毕
	fmt.Println("主程序退出")
}

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}
func test515_01_01_03() {

	go printMessage("Hello")    // 创建 goroutine
	go printMessage("World")    // 创建 goroutine
	time.Sleep(time.Second * 2) // 等待 goroutine 运行
}

func main1() {

	// test515_01_01_01()

	//test515_01_01_02()

	test515_01_01_03()

}
