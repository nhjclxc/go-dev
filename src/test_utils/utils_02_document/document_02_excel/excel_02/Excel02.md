


github.com/xuri/excelize


[官方教程](https://xuri.me/excelize/zh-hans/)


在 Go 中实现 Excel 列表数据的批量导出，常用的方式是使用 [github.com/xuri/excelize](https://github.com/xuri/excelize) 这个库，它功能强大，支持 XLSX 的读写。

下面是一个**完整示例**，演示如何批量导出一个结构体列表为 Excel 文件：

---

### ✅ 第一步：安装依赖

```bash
go get github.com/xuri/excelize/v2
```

```
go get -u github.com/xuri/excelize/v2
```

---

### ✅ 第二步：定义结构体和数据

```go
type User struct {
	ID    int
	Name  string
	Email string
}
```

---

### ✅ 第三步：实现 Excel 导出函数

```go
package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func exportUsersToExcel(users []User, filePath string) error {
	f := excelize.NewFile()
	sheet := "Sheet1"

	// 写入表头
	headers := []string{"ID", "姓名", "邮箱"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// 写入内容
	for rowIdx, user := range users {
		row := rowIdx + 2 // 从第2行开始写数据
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), user.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), user.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), user.Email)
	}

	// 保存文件
	if err := f.SaveAs(filePath); err != nil {
		return err
	}
	return nil
}

func main() {
	users := []User{
		{1, "张三", "zhangsan@example.com"},
		{2, "李四", "lisi@example.com"},
		{3, "王五", "wangwu@example.com"},
	}

	if err := exportUsersToExcel(users, "users.xlsx"); err != nil {
		fmt.Println("导出失败:", err)
		os.Exit(1)
	}
	fmt.Println("导出成功: users.xlsx")
}
```

---

### ✅ 导出效果（Excel 文件内容）

| ID | 姓名 | 邮箱                                                  |
| -- | -- | --------------------------------------------------- |
| 1  | 张三 | [zhangsan@example.com](mailto:zhangsan@example.com) |
| 2  | 李四 | [lisi@example.com](mailto:lisi@example.com)         |
| 3  | 王五 | [wangwu@example.com](mailto:wangwu@example.com)     |

---

### ✅ 补充功能（可选）

* 设置样式（字体加粗、单元格边框等）
* 设置列宽 `f.SetColWidth(sheet, "A", "C", 20)`
* 设置文件为流式输出（HTTP 下载）
* 支持多个 Sheet

---

如果你希望通过 Web API 导出 Excel 文件，我也可以提供 `gin` 版本的代码让用户直接下载。是否需要？
