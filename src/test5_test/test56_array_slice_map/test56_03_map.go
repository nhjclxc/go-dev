package main

import (
	"fmt"
	"sort"
)

func main3() {
	/*
		Go 语言映射（Map）详解
		map 是 Go 语言内置的数据结构，用于存储 键值对（key-value）。与数组和切片不同，map 允许使用 非整数索引，常用于快速查找和存储数据。

		（1）使用 make 创建 Map，make 是 Go 推荐的创建 map 的方式：
			m := make(map[string]int) // key 为 string，value 为 int

		（2）使用字面量初始化
			m := map[string]int{
				"apple":  5,
				"banana": 10,
			}
			fmt.Println(m) // map[apple:5 banana:10]

		3）声明但不初始化
			var m map[string]int
			fmt.Println(m == nil) // true（nil map 不能直接赋值）







	*/

	//test56_03_01()

	test56_03_02()

}

// 在 Go 语言中，如果 map 的键和值是可交换的（即值是唯一的），可以通过遍历原 map，创建一个新 map 来实现键值对调。
func test56_03_02() {

	// 适用于 map[string]int 变成 map[int]string：
	// 原始 map
	original := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// 交换键值
	swapped := make(map[int]string)
	for key, value := range original {
		swapped[value] = key
	}

	fmt.Println(swapped) // map[1:apple 2:banana 3:cherry]

	//// 2. 处理值重复的情况
	////如果原 map 的值可能重复，可以使用 map[int][]string 存储多个键：
	//
	//original := map[string]int{
	//	"apple":  1,
	//	"banana": 2,
	//	"grape":  1, // 值重复
	//}
	//
	//swapped := make(map[int][]string)
	//for key, value := range original {
	//	swapped[value] = append(swapped[value], key)
	//}
	//
	//fmt.Println(swapped) // map[1:[apple grape] 2:[banana]]
	//

}

func test56_03_01() {
	// 2.4 Map 的排序
	//Go 语言 map 不保证 key 的顺序，如果需要排序：
	m := map[string]int{"banana": 2, "apple": 1, "cherry": 3}

	// 提取所有 key
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 排序 key
	sort.Strings(keys)

	// 按排序后的 key 访问 map
	for _, k := range keys {
		fmt.Println(k, m[k])
	}

}

// 如果要比较两个 map，只能手写遍历：
func mapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
