package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestName2201(t *testing.T) {

	m := map[string][]int{
		"k1": []int{1, 2, 3, 4, 5, 6},
		"k2": []int{21, 2, 23, 6, 5},
		"k3": []int{31, 2, 33, 21, 6, 55},
		"k4": []int{111, 2, 6},
		"k5": {11, 2, 13, 212, 32, 6},
	}

	fmt.Println(IntersectAll(m))
	fmt.Println(IntersectAll2(m))

}

// IntersectAll 取map中value的[]int交集数据
func IntersectAll(m map[string][]int) []int {
	ret := make([]int, 0)
	if len(m) == 0 {
		return ret
	}

	count := make(map[int]int)
	for _, arr := range m {
		seen := make(map[int]bool)
		for _, v := range arr {
			if _, ok := seen[v]; !ok {
				seen[v] = true
				count[v]++
			}
		}
	}

	for k, v := range count {
		if v == len(m) {
			ret = append(ret, k)
		}
	}

	return ret
}

// IntersectAll 取 map 中所有 value 切片的交集
func IntersectAll2(m map[string][]int) []int {
	if len(m) == 0 {
		return nil
	}

	count := make(map[int]int)
	first := true

	for _, arr := range m {
		if first {
			// 第一个切片直接加入 count
			for _, v := range arr {
				count[v] = 1
			}
			first = false
			continue
		}

		// 后续切片只保留交集
		tmp := make(map[int]int)
		for _, v := range arr {
			if _, ok := count[v]; ok {
				tmp[v] = 1
			}
		}
		count = tmp
	}

	// 转为切片返回
	ret := make([]int, 0, len(count))
	for k := range count {
		ret = append(ret, k)
	}

	// 可选：保证有序输出
	sort.Ints(ret)
	return ret
}
