

## 1、 SQLite（入门级关系型数据库）
- 特点：轻量、零配置、嵌入式数据库，适合学习 SQL 基础语法、事务等。 
- 适合：本地存储、桌面/移动应用、小型服务。 
- Go推荐库：github.com/mattn/go-sqlite3 
- 🔸为什么先学：SQLite 易于安装和使用，可以先掌握 SQL 基本操作、数据库连接等概念。

[](https://github.com/mattn/go-sqlite3.git)
[](https://www.bilibili.com/video/BV1dZ4y1577v/)
### 1.1、
 安装 go-sqlite3 ：`go get github.com/mattn/go-sqlite3`
 


```go
package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func main() {
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS anonymous_user (id INTEGER PRIMARY KEY, name TEXT)")
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec("INSERT INTO anonymous_user(name) VALUES (?)", "Alice")
    if err != nil {
        log.Fatal(err)
    }
}
```