Go 语言标准库中的 [`container`](https://pkg.go.dev/container) 包，提供了一些**常用数据结构的基础实现**，主要包括三个子包：

---

## 🧱 1. `container/list`：双向链表（Doubly Linked List）

### ✅ 特点

* 实现了双向链表（Doubly Linked List）。
* 节点结构为 `Element`，可以前后遍历。
* 插入、删除操作时间复杂度为 **O(1)**，但查找需要遍历。

### 📦 常用方法

```go
l := list.New()
e1 := l.PushBack("a")   // 尾部插入
e2 := l.PushFront("b")  // 头部插入
l.InsertAfter("c", e1)  // 在 e1 后插入
l.InsertBefore("d", e2) // 在 e2 前插入
l.Remove(e1)            // 删除 e1

for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)  // 遍历
}
```

---

## 📚 2. `container/heap`：堆（Heap）

### ✅ 特点

* 提供了**堆结构的接口**（最小堆或最大堆）。
* 实际上是用切片实现的，但你需要**实现 `heap.Interface` 接口**（包括 `Len`, `Less`, `Swap`, `Push`, `Pop`）。

### 🧱 示例（最小堆）：

```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
```

### 📦 使用方式

```go
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
fmt.Println(heap.Pop(h)) // 1
```

---

## 📚 3. `container/ring`：循环链表（Ring）

### ✅ 特点

* 实现了一个**固定长度的循环链表**。
* 每个节点指向前后两个节点，并且**首尾相连**。

### 📦 使用方式

```go
r := ring.New(3)
for i := 0; i < r.Len(); i++ {
    r.Value = i
    r = r.Next()
}

// 遍历
r.Do(func(p any) {
    fmt.Println(p) // 输出 0, 1, 2
})
```

---

## 🔍 三者对比总结：

| 包名     | 数据结构 | 插入/删除效率  | 查找效率    | 典型应用         |
| ------ | ---- | -------- | ------- | ------------ |
| `list` | 双向链表 | O(1)     | O(n)    | LRU缓存、队列/栈实现 |
| `heap` | 堆结构  | O(log n) | O(1) 最值 | 优先队列、任务调度    |
| `ring` | 循环链表 | O(1)     | O(n)    | 循环缓存、轮询任务等   |

---

## 🧠 总结

虽然这些 `container` 包提供了底层数据结构，但在实际项目中：

* **多数人会直接用 `slice/map` 实现逻辑**，例如栈、队列、优先队列。
* 这些包适合对**性能敏感或结构明确**的场景，比如实现 LRU 缓存、任务调度器。

---

如需我帮你实现某个具体结构（比如队列、LRU 缓存、线程安全优先队列等），也可以继续问我，我可以一步步用 `container` 包为你实现。
