当然可以！以下是几道针对 **Go 语言中切片（slice）与映射（map）高级应用** 的**练习题目**，适合加深你对这两种数据结构的理解与实战能力。

---

## 🧠 题目一：切片去重并保持原顺序

### 🔹题目描述：

实现一个函数 `DedupSlice`，接收一个字符串切片 `[]string`，返回一个去重后的新切片，**并保留第一次出现的顺序**。

### ✅ 示例：

```go
输入: []string{"a", "b", "a", "c", "b"}
输出: []string{"a", "b", "c"}
```

---

## 🧠 题目二：统计切片中每个元素出现的次数

### 🔹题目描述：

实现一个函数 `CountFrequency`，接收一个 `[]int` 切片，返回一个 `map[int]int`，表示每个整数出现的频率。

### ✅ 示例：

```go
输入: []int{1, 2, 2, 3, 1, 1}
输出: map[int]int{1: 3, 2: 2, 3: 1}
```

---

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

---

## 🧠 题目四：两个切片求交集（高性能）

### 🔹题目描述：

实现函数 `Intersect`，接收两个整数切片 `a` 和 `b`，返回它们的交集。结果无需排序，但不能包含重复值。

> 使用 `map[int]struct{}` 实现高效查找。

### ✅ 示例：

```go
输入: a = []int{1, 2, 3, 4}, b = []int{3, 4, 5, 6}
输出: []int{3, 4}
```

---

## 🧠 题目五：Top K 高频元素

### 🔹题目描述：

实现一个函数 `TopKFrequent`，接收一个字符串切片 `[]string` 和一个整数 `k`，返回出现频率最高的前 K 个字符串。

> 提示：可使用 `map[string]int` 统计频率，结合排序。

### ✅ 示例：

```go
输入: ["a", "b", "a", "c", "b", "a"], k=2
输出: ["a", "b"]
```

---

### 如果你希望：

* 我给出上述题目的**参考解法代码**
* 或者自定义更多贴近你项目的题目

请告诉我你最感兴趣的题目编号或领域方向（如 Web、算法、系统设计），我可以继续提供 **实战型题目 + 讲解 + 最佳实践代码**。





很好！以下是更多偏向实战 + 面试 + 高阶思维的 **Go 切片 + map 综合题目**，题型设计涵盖数据处理、统计分析、集合操作、数据结构模拟等。

---

## 🧠 题目六：模拟集合（Set）

### 🔹题目描述：

实现一个结构体 `StringSet`，内部使用 `map[string]struct{}` 实现集合，支持以下操作：

* `Add(s string)`
* `Remove(s string)`
* `Contains(s string) bool`
* `Values() []string`
* `Len() int`

> ✅ 要求：集合中不能有重复值。

---

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

---

## 🧠 题目八：统计切片中每个元素第一次和最后一次出现的位置

### 🔹题目描述：

实现函数 `IndexRange(nums []int) map[int][2]int`，返回每个元素的 **首次与末次** 出现位置。

### ✅ 示例：

```go
输入: []int{3, 1, 2, 3, 1, 4}
输出: map[int][2]int{1: [1 4], 2: [2 2], 3: [0 3], 4: [5 5]}
```

---

## 🧠 题目九：找出只出现一次的元素（其他元素都出现两次）

### 🔹题目描述：

给定一个整数切片，其中除一个元素外，其余都出现两次，找出这个只出现一次的元素。

> 要求：时间复杂度 O(n)，空间复杂度 O(1) 或 O(n)

---

cc
---

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

---

## 🧠 题目十二：分页切片数据

### 🔹题目描述：

实现函数 `Paginate[T any](data []T, page, pageSize int) []T`，返回切片的第 `page` 页数据（支持泛型）。

---

这些题目难度从中等到较难，非常适合用于：

* ✅ Go 的数据结构和算法练习
* ✅ 面试题训练
* ✅ 项目中常用切片和 map 的综合操作
* ✅ 泛型、排序、高阶函数的运用

---

🔧 **如你指定编号，我可以给出**：

* 高性能/简洁解法
* 带注释讲解
* 单元测试用例

是否需要我继续出一些关于多层嵌套 map、JSON 解析、结构体切片聚合等更复杂的数据处理题？




太棒了，咱们现在进入 **中高阶难度的 Go 切片 + map 综合实战题**，更加偏向真实业务场景、数据聚合、变换与结构体组合，适合深入理解 Go 在数据处理中的强大能力。

---

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

---

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

---

## 🧠 题目十五：结构体切片转成 map（主键索引）

### 🔹题目描述：

将如下结构体切片 `[]User` 转成 map：

```go
type User struct {
    ID   int
    Name string
}
```

实现函数 `ToUserMap(users []User) map[int]User`，使用 `ID` 作为 key。

---

## 🧠 题目十六：切片去重后计数并排序（组合问题）

### 🔹题目描述：

给定字符串切片 `[]string`，请返回每个去重元素的出现次数，并按频率**从高到低排序**。

> ✅ 要求返回 `[]struct{Word string; Count int}`，使用 `map + sort` 实现。

---

## 🧠 题目十七：查找最常出现的前 K 个组合项

### 🔹题目描述：

一个用户日志切片如下：

```go
type Log struct {
    UserID string
    Action string
}
```

实现函数 `TopKActions(logs []Log, k int) []string`，找出出现频率最多的 `UserID|Action` 组合项。

---

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

---

## 🧠 题目十九：结构体切片根据多个字段去重

### 🔹题目描述：

给定如下结构体：

```go
type File struct {
    Name string
    Size int64
}
```

实现函数 `DedupFiles(files []File) []File`，按 `Name+Size` 组合判断是否重复（用 map 实现）。

---

## 🧠 题目二十：根据 map 数据补全结构体字段

### 🔹题目描述：

已有结构体：

```go
type Product struct {
    ID   int
    Name string
}
```

和一个 map：

```go
priceMap := map[int]float64{
    1: 9.9,
    2: 15.5,
}
```

请创建新的结构体切片：

```go
type ProductWithPrice struct {
    ID    int
    Name  string
    Price float64
}
```

将 `[]Product` 转为 `[]ProductWithPrice`，其中价格字段来自 `priceMap`。

---

这些题目要求你熟练使用：

* `map` 构建索引、多层嵌套、组合 key
* `slice` 的过滤、聚合、排序、转 map
* `struct` 的分组与合并（业务建模）
* 结合字符串操作（如年月 key）与泛型技巧

---

📌 **下一步建议：**
如你希望更贴近项目实战，我还可以出题围绕以下方向：

* JSON 数据扁平化 + 转结构体切片
* 数据分页、过滤、搜索关键词匹配
* 树形结构构建（如菜单/组织结构）
* 使用 `sync.Map` / 并发 map 操作

想尝试哪一类挑战？你可以告诉我领域或场景，我为你定制 💪

