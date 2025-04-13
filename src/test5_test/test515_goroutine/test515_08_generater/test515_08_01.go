package main

import (
	"fmt"
)

func main01() {
	/*
			生成器是指当被调用时返回一个序列中下一个值的函数，
				generateInteger() => 0
				generateInteger() => 1
				generateInteger() => 2
				....

			生成器每次返回的是序列中下一个值而非整个序列；这种特性也称之为惰性求值：只在你需要时进行求值，同时保留相关变量资源（内存和 CPU）：
		这是一项在需要时对表达式进行求值的技术。例如，生成一个无限数量的偶数序列：要产生这样一个序列并且在一个一个的使用可能会很困难，
		而且内存会溢出！但是一个含有通道和 go 协程的函数能轻易实现这个需求。


	*/

	// 4、使用工厂方法 初始化resmue
	var gi *GenerateInteger = getInstance()
	// 5、使用
	fmt.Println(gi.generateInteger())
	fmt.Println(gi.generateInteger())
	fmt.Println(gi.generateInteger())
	fmt.Println(gi.generateInteger())
	fmt.Println(gi.generateInteger())
	fmt.Println(gi.generateInteger())

}

// 1、定义一个全局通道
type GenerateInteger struct {
	resume chan int
}

// 2、定义一个 工厂来创建一个resume的实例
func getInstance() *GenerateInteger {

	var gi GenerateInteger = GenerateInteger{
		resume: make(chan int),
	}

	// 定义初始值
	count := 0

	// 开启一个协程，每次向通道内发送这次加+1的数据
	// 死循环的go协程，用于无限数据的累加
	go func() {
		for {
			// 只要通道里面的数据被取走了，那么计数器就加一
			gi.resume <- count
			count++
		}
	}()

	return &gi
}

// 3、定义每次调用的方法
func (this *GenerateInteger) generateInteger() int {
	return <-this.resume
}
