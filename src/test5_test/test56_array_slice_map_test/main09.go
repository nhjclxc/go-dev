package main

import "fmt"

/*
## 🧠 题目九：找出只出现一次的元素（其他元素都出现两次）

### 🔹题目描述：

给定一个整数切片，其中除一个元素外，其余都出现两次，找出这个只出现一次的元素。

> 要求：时间复杂度 O(n)，空间复杂度 O(1) 或 O(n)

 */

func FindOne1(nums []int) int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for key, val := range countMap {
		if val == 1 {
			return key
		}
	}
	return 0
}

func FindOne(nums []int) int {
	result := 0
	for _, num := range nums {
		temp := result
		result ^= num
		fmt.Printf("num = %d, result = %d, result ^ num = %d \n", num, temp, result)
	}
	return result
}

/*
| 表达式         | 结果                  | 说明               |
| ----------- | ------------------- | ---------------- |
| `a ^ a = 0` | 相同为 0               | 相同的数字异或为 0       |
| `a ^ 0 = a` | 零为单位元               | 任何数字与 0 异或仍然是它自己 |
| 异或满足交换律     | `a^b = b^a`         | 计算顺序无关           |
| 异或满足结合律     | `(a^b)^c = a^(b^c)` | 多个数异或可以任意组合      |


1,2,3,4,5,4,3,2,1

=1^2^3^4^5^4^3^2^1
=(1^1)^(2^2)^(3^3)^(4^4)^5
=0^0^0^0^5
=0^5
=5

 */