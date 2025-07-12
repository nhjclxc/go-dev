package main

/*

## 🧠 题目十八：多级嵌套 map 数据扁平化

### 🔹题目描述：

给定如下嵌套结构：

```go
data := map[string]map[string]int{
    "2025": {
        "Jan": 100,
        "Feb": 200,
    },
    "2026": {
        "Jan": 150,
    },
}
```

实现函数 `Flatten(data map[string]map[string]int) map[string]int`，返回：

```go
map[string]int{
    "2025-Jan": 100,
    "2025-Feb": 200,
    "2026-Jan": 150,
}
```

 */

func Flatten(data map[string]map[string]int) map[string]int {
	tempMap := make(map[string]int)

	for key1, val1 := range data {
		for key2, val2 := range val1 {
			tempMap[key1 + "-" + key2] = val2
		}
	}
	return tempMap
}