package main

import "fmt"

// 类型 T 方法集包含全部 receiver T 方法。
func main1() {

	t := T{}
	tp := &t
	fmt.Println(t)
	fmt.Println(tp)

	// t的所有方法集
	t.funcT()
	t.funcTP()

	// tp的所有方法集
	tp.funcT()
	tp.funcTP()

}

// 声明一个类型 T
type T struct{}

// 声明类型 T 的一个方法
func (self T) funcT() {
	fmt.Println("T 的所有方法集")
}

// 声明类型 *T 的一个方法
func (self *T) funcTP() {
	fmt.Println("*****T 的所有方法集")
}
