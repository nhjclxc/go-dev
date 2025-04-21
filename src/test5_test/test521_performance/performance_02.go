package main

import "fmt"

// https://denganliang.github.io/the-way-to-go_ZH_CN/18.2.html
// 18.2 数组和切片
func main02() {
	//创建：仅创建
	arr1 := new([5]int)
	slice1 := make([]int, 5)
	fmt.Println(arr1)
	fmt.Println(slice1)

	//初始化：创建并初始化
	arr12 := []int{1, 2, 3, 4, 5, 6}
	arrKeyValue := [6]int{1, 2, 3, 4, 5, 6}
	var slice12 []int = arr1[1:2]
	fmt.Println(arr12)
	fmt.Println(arrKeyValue)
	fmt.Println(slice12)

	//（1）如何截断数组或者切片的最后一个元素：
	fmt.Println(arr12)
	fmt.Println(arr12[:len(arr12)-1])

	// 18.3 映射
	//创建：
	map11 := make(map[string]int)

	//初始化：
	map12 := map[string]int{"one": 1, "two": 2}

	//2）如何在一个映射中检测键 key1 是否存在：
	//返回值：键 key1 对应的值或者 0，true 或者 false
	val1, isPresent := map11["key1"]
	fmt.Println(val1)
	fmt.Println(isPresent)

	//（3）如何在映射中删除一个键：
	delete(map12, "key1")

	// 当结构体的命名以大写字母开头时，该结构体在包外可见。 通常情况下，为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体（参考例 10.2）：
}
