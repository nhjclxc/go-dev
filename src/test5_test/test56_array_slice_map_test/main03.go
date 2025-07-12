package main

import (
	"sort"
)

/*

## 🧠 题目七：词频统计并按频率排序（升序）

### 🔹题目描述：

实现函数 `SortedWordCount(text string) []WordCount`，将一段英文字符串按单词统计出现频率，并按频率升序排序。

结构体如下：

```go
type WordCount struct {
    Word  string
    Count int
}
```

> 提示：使用 `map[string]int` 统计，再用 `sort.Slice` 排序。

 */


type WordCount struct {
	Word  string
	Count int
}

func SortedWordCount(text string) []WordCount {

	// 1、统计频率
	countMap := make(map[string]int)
	for _, val := range text {
		countMap[string(val)]++
	}

	// 2、转化为结构体
	var wordCountList []WordCount = make([]WordCount, 0, len(countMap))
	for key, val := range countMap {
		wordCountList = append(wordCountList, WordCount{Word: key, Count: val})
	}

	// 3、排序

	sort.Slice(wordCountList, func(i, j int) bool {
		return wordCountList[i].Count > wordCountList[j].Count
	})

	return wordCountList
}

