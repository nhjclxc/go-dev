package main

import "fmt"

// defer 在func里面和在for里面的区别
func main02() {

	example1()

	example2()

	// defer 仅在函数返回时才会执行，在循环内的结尾或其他一些有限范围的代码内不会执行。

}

func example1() {
	// B
	// A
	defer fmt.Println("A")
	fmt.Println("B")
}

func example2() {
	//每次循环都执行一个 defer，这些 defer 都注册在函数返回前执行， 而 不 是 循 环 结 束 时。
	//所以执行顺序是先进后出（LIFO）。

	// 2
	// 1
	// 0
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
