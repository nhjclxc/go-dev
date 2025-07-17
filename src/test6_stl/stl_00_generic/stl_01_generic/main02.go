package main

import "fmt"

type Pair[T any, U any] struct {
	First  T
	Second U
}

func main02() {
	p := Pair[int, string]{First: 1, Second: "hello"}
	fmt.Println(p)
}

/*

泛型结构体定义：
	type 结构体名[类型参数列表] struct { }

type Box[T any] struct {
	Value T
}

*/
