package main

import "fmt"

// 展开操作符
func main() {
	// 直接传多个数字
	fmt.Println("Sum:", Sum(1, 2, 3, 4)) // 输出: Sum: 10

	// 也可以用切片调用，加上 ...
	values := []int{5, 10, 15}
	// ... 来告诉编译器：“请把这个切片拆开一个个地传入”。
	fmt.Println("Sum:", Sum(values...)) // 输出: Sum: 30

	fmt.Println("Sum:", Sum([]int{5, 10, 15}...)) // 输出: Sum: 30
	fmt.Println("Sum:", Sum(5, 10, 15))           // 输出: Sum: 30
}

// 函数定义：接收任意数量的 int 参数
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
