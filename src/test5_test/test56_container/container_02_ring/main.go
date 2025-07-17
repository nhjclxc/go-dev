package main

import (
	"container/ring"
	"fmt"
)

func main() {

	// 🧠 应用场景一：轮询调度器（Round Robin Scheduler）
	// 📉 应用场景二：固定容量的环形缓冲区（Ring Buffer）
	// 🎯 应用场景三：多协程之间的令牌轮转（Token Ring）
	// 🎲 应用场景四：循环菜单 / 状态机
	// ⏱ 应用场景五：滑动窗口计数器（Sliding Window）
	// 🧪 应用场景六：有限重试任务处理器
	// 🧮 应用场景七：哈希槽轮转（类似 Redis Cluster）


	/*
	| 方法/字段       | 功能说明               | 使用场景示例        |
	| ----------- | ------------------ | ------------- |
	| `New(n)`    | 创建长度为 n 的环形链表      | `ring.New(5)` |
	| `Len()`     | 当前环的长度             | 获取元素个数        |
	| `Value`     | 当前节点存储的值           | 赋值或读取节点内容     |
	| `Do(f)`     | 遍历所有节点执行函数         | 打印所有值         |
	| `Next()`    | 当前节点的下一个节点         | `r.Next()`    |
	| `Prev()`    | 当前节点的上一个节点         | `r.Prev()`    |
	| `Move(n)`   | 向前/后移动 n 步         | `r.Move(-1)`  |
	| `Link(s)`   | 将环 s 插入当前节点后       | 合并两个环         |
	| `Unlink(n)` | 移除当前节点后 n 个节点，返回子环 | 删除一段子链        |

	*/

	// 创建环形链表
	r := ring.New(5)

	// 为环形链表赋值
	for i := 0; i < r.Len(); i++ {
		r.Value = i * 10

		// 当前节点赋值完成之后，指针要向后偏移
		r = r.Next()
	}

	// 遍历
	r.Do(func(a any) {
		fmt.Println(a)
	})

	next := r.Next()
	fmt.Println(next)
	fmt.Println(next.Prev())

	move := next.Move(2)
	fmt.Println(move)
	fmt.Println(move.Next())
	fmt.Println(move.Prev())


	// 创建环形链表
	r2 := ring.New(3)
	r2.Value = "a"
	r2 = r2.Next()
	r2.Value = "b"
	r2 = r2.Next()
	r2.Value = "c"

	link := next.Link(r2)

	// 遍历
	next.Do(func(a any) {
		fmt.Println(a)
	})
	fmt.Println()
	fmt.Println()
	fmt.Println()
	// 遍历
	link.Do(func(a any) {
		fmt.Println(a)
	})




}
