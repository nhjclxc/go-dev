package main

/*

## 🧠 题目十四：按年月聚合求和（嵌套 Map）

### 🔹题目描述：

订单结构体如下：

```go
type Order struct {
    Date   string  // 格式 "2025-07-11"
    Amount float64
}
```

实现函数 `SumAmountByMonth(orders []Order) map[string]float64`，返回每个月的总金额。

> ✅ 提示：可截取前 7 位 "yyyy-MM" 作为 key。

 */


type Order2 struct {
	Date   string  // 格式 "2025-07-11"
	Amount float64
}

func SumAmountByMonth(orders []Order2) map[string]float64 {
	tempMap := make(map[string]float64)
	for _, order := range orders {
		tempMap[order.Date[0:7]] += order.Amount
	}
	return tempMap
}
