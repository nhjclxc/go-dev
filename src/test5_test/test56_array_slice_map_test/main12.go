package main

/*

## 🧠 题目十二：分页切片数据

### 🔹题目描述：

实现函数 `Paginate[T any](data []T, page, pageSize int) []T`，返回切片的第 `page` 页数据（支持泛型）。

---

这些题目难度从中等到较难，非常适合用于：

* ✅ Go 的数据结构和算法练习
* ✅ 面试题训练
* ✅ 项目中常用切片和 map 的综合操作
* ✅ 泛型、排序、高阶函数的运用

 */

func Paginate[T any](data []T, page, pageSize int) []T {
	if page <= 0 || pageSize <= 0 {
		return []T{}
	}

	start := (page - 1) * pageSize
	if start >= len(data) {
		return []T{}
	}

	end := start + pageSize
	if end > len(data) {
		end = len(data)
	}
	return data[start: end]
}




