package main

import "fmt"

// 🧩 5. 泛型类型实参（类型推断）
func main() {

	PrintSlice[int]([]int{1, 2, 3})
	PrintSlice([]int{1, 2, 3}) // 自动推断 T=int
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
