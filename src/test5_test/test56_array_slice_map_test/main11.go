package main

/*


## 🧠 题目十一：按字段去重结构体切片

### 🔹题目描述：

定义如下结构体：

```go
type User struct {
    ID   int
    Name string
}
```

实现函数 `DedupByID(users []User) []User`，按 `ID` 字段去重，保留第一次出现的记录。

 */
type User struct {
	ID   int
	Name string
}

func DedupByID(users []User) []User {
	tmepMap := make(map[int]bool)
	var res []User
	for _, val := range users {
		if _, exists := tmepMap[val.ID]; !exists {
			tmepMap[val.ID] = true
			res = append(res, val)
		}
	}
	return res
}