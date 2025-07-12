package main

import "sort"

/*
	## 🧠 题目一：切片去重并保持原顺序

	### 🔹题目描述：

	实现一个函数 `DedupSlice`，接收一个字符串切片 `[]string`，返回一个去重后的新切片，**并保留第一次出现的顺序**。

	### ✅ 示例：

	```go
	输入: []string{"a", "b", "a", "c", "b"}
	输出: []string{"a", "b", "c"}
	```
*/
func DedupSlice1(strs []string) []string {

	if strs == nil || len(strs) <= 0 {
		return nil
	}

	tempMap := make(map[string]int)
	index := 1
	for _, str := range strs {
		if tempMap[str] == 0 {
			tempMap[str] = index
			index++
		}
	}

	res := make([]string, len(tempMap), len(tempMap))
	for key, val := range tempMap {
		res[val -1] = key
	}
	return res
}

func DedupSlice(strs []string) []string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(strs))

	for _, s := range strs {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}

	return result
}




/*
## 🧠 题目二：统计切片中每个元素出现的次数

### 🔹题目描述：

实现一个函数 `CountFrequency`，接收一个 `[]int` 切片，返回一个 `map[int]int`，表示每个整数出现的频率。

### ✅ 示例：

```go
输入: []int{1, 2, 2, 3, 1, 1}
输出: map[int]int{1: 3, 2: 2, 3: 1}
```
 */

func CountFrequency(ints []int) map[int]int {
	countMap := make(map[int]int, len(ints))

	for _, val := range ints {
		if countMap[val] == 0 {
			countMap[val] = 1
		} else {
			countMap[val] = countMap[val] + 1
		}
	}
	return countMap
}



/*
## 🧠 题目三：切片过滤器（支持自定义条件）

### 🔹题目描述：

实现一个通用函数 `FilterInts`，接收一个整数切片 `[]int` 和一个函数 `func(int) bool`，返回所有满足条件的整数组成的新切片。

### ✅ 示例：

```go
FilterInts([]int{1, 2, 3, 4, 5}, func(i int) bool {
    return i%2 == 0
})
输出: []int{2, 4}
```
 */

func FilterInts(ints []int, filter func(int) bool) []int {
	res := make([]int, 0)
	for _, val := range ints {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}



/*
## 🧠 题目四：两个切片求交集（高性能）

### 🔹题目描述：

实现函数 `Intersect`，接收两个整数切片 `a` 和 `b`，返回它们的交集。结果无需排序，但不能包含重复值。

> 使用 `map[int]struct{}` 实现高效查找。

### ✅ 示例：

```go
输入: a = []int{1, 2, 3, 4}, b = []int{3, 4, 5, 6}
输出: []int{3, 4}
```
 */

// 交集
func Intersect(a, b []int) []int {
	tempAMap := make(map[int]bool, len(a))
	tempBMap := make(map[int]bool, len(b))
	resMap := make(map[int]bool)

	for _, val := range a {
		tempAMap[val] = true
	}
	for _, val := range b {
		tempBMap[val] = true
	}


	tempSli := a
	tempSli = append(tempSli, b...)

	for _, val := range tempSli {
		if tempAMap[val] && tempBMap[val] {
			resMap[val] = true
		}
	}

	return keys[int, bool](resMap)
}

// 以下方法实现两个切片的并集
func Union(a, b []int) []int {
	tempMap := make(map[int]bool)

	tempSli := a
	tempSli = append(tempSli, b...)

	for _, val := range tempSli {
		if !tempMap[val] {
			tempMap[val] = true
		}
	}

	return keys[int, bool](tempMap)
}

// 泛型函数 keys[K, V]，用于从任意 map[K]V 中提取 []K 类型的 key 切片。
// K comparable：Go 中 map 的 key 必须是可比较的（comparable）。
// V any：值的类型可以是任意的。
func keys[K comparable, V any](tempMap map[K]V) []K {
	//make([]K, 0, len(temp))：预分配容量提高性能。
	res := make([]K, 0, len(tempMap))

	for key, _ := range tempMap {
		res = append(res, key)
	}

	return res
}



/*
## 🧠 题目五：Top K 高频元素

### 🔹题目描述：

实现一个函数 `TopKFrequent`，接收一个字符串切片 `[]string` 和一个整数 `k`，返回出现频率最高的前 K 个字符串。

> 提示：可使用 `map[string]int` 统计频率，结合排序。

### ✅ 示例：

```go
输入: ["a", "b", "a", "c", "b", "a"], k=2
输出: ["a", "b"]
```
 */

func TopKFrequent(strs []string, k int) []string {
	// 1. 统计频率
	freqMap := make(map[string]int)
	for _,val := range strs {
		//tempMap[val] = tempMap[val] + 1
		freqMap[val]++
	}

	// 2. 转为切片方便排序
	type Pair struct {
		Key   string
		Count int
	}
	pairs := make([]Pair, 0, len(freqMap))
	for key, count := range freqMap {
		pairs = append(pairs, Pair{key, count})
	}

	// 3. 排序（按 Count 降序）
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	// 4. 取前 k 个
	result := make([]string, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		result = append(result, pairs[i].Key)
	}

	return result
}