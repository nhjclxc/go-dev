package main

import "fmt"

// 闭包在go里面就是匿名函数
func main() {

	//test510_01_01()

	// defer的一个面试体题
	test510_01_02()

}

func test510_01_02() {

	/// 变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的。
	//这可用于在返回语句之后修改返回的 error 时使用。
	fmt.Println(deferTest())
}

func deferTest() (res int) {
	res = 1
	defer func() {
		res++
	}()
	return
}

// 匿名函数
func test510_01_01() {

	// 匿名函数
	res := func(num1, num2 int) int {
		return num1 + num2
	}(1, 2)
	println(res)

	// 将匿名函数赋值给一个变量，将这个变量作为这个匿名函数的函数名称
	add := func(num1, num2 int) int {
		return num1 + num2
	}
	println(add(3, 2))

	fmt.Println(1e-2)
	fmt.Println(1e-1)
	fmt.Println(1e0)
	fmt.Println(1e1)
	fmt.Println(1e2)
	fmt.Println(1e3)

}
