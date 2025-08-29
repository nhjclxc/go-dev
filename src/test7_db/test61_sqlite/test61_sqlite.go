package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("数据库连接成功！")

	// panic: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
	// github.com/mattn/go-sqlite3 是一个基于 C 的 SQLite 封装，它底层调用了 SQLite 的 C 库，因此需要启用 Cgo 才能编译成功。

	// 如果你希望项目能跨平台构建、或部署到禁用 Cgo 的环境（如 AWS Lambda、alpine 等），可以使用以下纯 Go 实现的 SQLite 库：
	//🔹 modernc.org/sqlite
	// go get modernc.org/sqlite
}

// 📐 2. 创建表
func createTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS anonymous_user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	fmt.Println("表创建成功！")
}

// ➕ 3. 插入数据
func insertUser(db *sql.DB, name string, age int) {
	stmt, err := db.Prepare("INSERT INTO anonymous_user(name, age) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, age)
	if err != nil {
		panic(err)
	}
	fmt.Println("插入成功：", name, age)
}

// 📋 4. 查询数据（全部）
func queryAllUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM anonymous_user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID=%d, Name=%s, Age=%d\n", id, name, age)
	}
}

// 🔍 5. 查询单条数据
func queryUserByID(db *sql.DB, id int) {
	var name string
	var age int
	err := db.QueryRow("SELECT name, age FROM anonymous_user WHERE id = ?", id).Scan(&name, &age)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("没有找到数据")
			return
		}
		panic(err)
	}
	fmt.Printf("查询到：Name=%s, Age=%d\n", name, age)
}

// ✏️ 6. 更新数据
func updateUserAge(db *sql.DB, id int, newAge int) {
	stmt, err := db.Prepare("UPDATE anonymous_user SET age = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(newAge, id)
	if err != nil {
		panic(err)
	}
	count, _ := res.RowsAffected()
	fmt.Printf("更新成功：受影响行数 %d\n", count)
}

// ❌ 7. 删除数据
func deleteUserByID(db *sql.DB, id int) {
	stmt, err := db.Prepare("DELETE FROM anonymous_user WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	count, _ := res.RowsAffected()
	fmt.Printf("删除成功：受影响行数 %d\n", count)
}
