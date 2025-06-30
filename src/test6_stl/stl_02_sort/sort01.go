package main

import (
	"fmt"
	"sort"
)

func main01() {
	// 1.1 升序排序：[]int
	nums := []int{5, 2, 9, 1, 6}
	sort.Ints(nums)
	fmt.Println(nums) // [1 2 5 6 9]

	// 反转slice
	nums2 := reversedSlice(nums)
	fmt.Println(nums2) // [9 6 5 2 1]

	fmt.Println(nums) // [1 2 5 6 9]

	reverseSlice(nums)
	fmt.Println(nums) // [9 6 5 2 1]

	sort.Ints(nums)
	isSorted := sort.IntsAreSorted(nums)
	fmt.Println(isSorted)


	// 将 IntSlice 转换为 sort.Interface，然后反转排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	fmt.Println(nums) // [4 3 2 1]

}

//  方法一：就地反转（in-place）
func reverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


// 方法二：返回新切片（不修改原始数据）
func reversedSlice[T any](s []T) []T {
	res := make([]T, len(s))
	for i := range s {
		res[len(s)-1-i] = s[i]
	}
	return res
}
