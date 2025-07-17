package main

import "fmt"

// 官方教程
// https://go.dev/doc/tutorial/generics

// Go 从 1.18 开始正式支持泛型（Generics），你可以使用类型参数来编写更通用的函数或结构体。
func main01() {
	fmt.Println(ToSlice[int](1, 2))        // [1 2]
	fmt.Println(ToSlice[string]("a", "b")) // [a b]

}

// ToSlice 定义一个泛型函数，接受任何类型的两个参数并返回它们组成的切片
func ToSlice[T any](a, b T) []T {
	return []T{a, b}
}

/*
泛型函数的定义：
	func 方法名[T any](参数列表) 返回数据类型 { }
	其中，	[T any] 表示 T 是任意类型
			这个 T 就是表示泛型 T 的类型
			any 是 interface{} 的别名

例如：
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

*/
