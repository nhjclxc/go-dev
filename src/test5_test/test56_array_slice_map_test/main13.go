package main


/*

## 🧠 题目十三：按属性分组结构体切片（Group By）

### 🔹题目描述：

给定如下结构体：

```go
type Order struct {
    OrderID  int
    UserID   int
    Amount   float64
}
```

实现函数 `GroupOrdersByUser(orders []Order) map[int][]Order`，将订单按 `UserID` 分组。

 */

type Order struct {
	ID   int
	Name string
}


func GroupOrdersByUser(orders []Order) map[int][]Order {
	groupMap := make(map[int][]Order)
	for _, order := range orders {
		if val, exists := groupMap[order.ID]; exists {
			temp := append(val, order)
			groupMap[order.ID] = temp
		} else {
			temp := make([]Order, 0)
			temp = append(temp, order)
			groupMap[order.ID] = temp
		}
	}
	return groupMap
}