package main

import (
	"bytes"
	"sync"
)

// 9.3 锁和 sync 包
// https://denganliang.github.io/the-way-to-go_ZH_CN/09.3.html
func main1() {
	/*
	   	在一些复杂的程序中，通常通过不同线程执行不同应用来实现程序的并发。当不同线程要使用同一个变量时，
	   	经常会出现一个问题：无法预知变量被不同线程修改的顺序！（这通常被称为资源竞争，指不同线程对同一变量使用的竞争）显然这无法让人容忍，那我们该如何解决这个问题呢？

	      经典的做法是一次只能让一个线程对共享变量进行操作。当变量被一个线程改变时（临界区），我们为它上锁，直到这个线程执行完成并解锁后，其他线程才能访问它。

	      特别是我们之前章节学习的 map 类型是不存在锁的机制来实现这种效果（出于对性能的考虑），所以 map 类型是非线程安全的。当并行访问一个共享的 map 类型的数据，map 数据将会出错。

	      在 Go 语言中这种锁的机制是通过 sync 包中 Mutex 来实现的。sync 来源于 “synchronized” 一词，这意味着线程将有序的对同一变量进行访问。

	      sync.Mutex 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
	*/

}

type Info struct {
	// 声明一个互斥锁变量 mu
	mu   sync.Mutex
	name string
	// ... other fields, e.g.: Str string
}

// 如果一个函数想要改变这个变量可以这样写:
func update(info *Info, name string) {
	// 1、先上锁，利用info对象里面下互斥锁，先把这个对象的访问权限锁起来，不允许其他线程（协程）访问
	info.mu.Lock()
	// 2、再修改
	info.name = name
	// 3、最后解锁，表示本线程对改变量已经完成了访问，现在可以让其他线程去访问了
	info.mu.Unlock()
}

// 还有一个很有用的例子是通过 Mutex 来实现一个可以上锁的共享缓冲器:
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}
