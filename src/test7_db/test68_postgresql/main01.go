package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)


// go get github.com/lib/pq
func main() {
	// 配置连接参数
	connStr := "host=127.0.0.1 port=5432 user=postgres password=pgsqldb123 dbname=pgsqldb sslmode=disable"

	// 连接数据库
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查连接
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect:", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// 创建表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age INT
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	_, err = db.Exec(`INSERT INTO users (name, age) VALUES ($1, $2)`, "Tom", 28)
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query(`SELECT id, name, age FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID=%d, Name=%s, Age=%d\n", id, name, age)
	}
}
