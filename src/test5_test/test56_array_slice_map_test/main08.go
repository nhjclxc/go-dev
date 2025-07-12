package main

/*

## 🧠 题目八：统计切片中每个元素第一次和最后一次出现的位置

### 🔹题目描述：

实现函数 `IndexRange(nums []int) map[int][2]int`，返回每个元素的 **首次与末次** 出现位置。

### ✅ 示例：

```go
输入: []int{3, 1, 2, 3, 1, 4}
输出: map[int][2]int{1: [1 4], 2: [2 2], 3: [0 3], 4: [5 5]}
```
 */

func IndexRange(nums []int) map[int][2]int {
	result  := make(map[int][2]int)

	for i, num := range nums {
		if val, exists := result[num]; exists {
			val[1] = i // 更新末次出现位置
			result[num] = val  // 更新回去，因为val是result里面出来的副本，上一行的操作不会修改原数据
		} else {
			result[num] = [2]int{i, i} // 首次出现，第一次出现也有可能是最后一次出现，因此先将两个值赋值为i
		}
	}

	return result
}
