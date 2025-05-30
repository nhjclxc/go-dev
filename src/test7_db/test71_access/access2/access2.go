package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	// 注意双反斜杠
	dsn := `Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=D:\\code\\go\\go-dev\\src\\test7_db\\test71_access\\Database1.accdb;`

	db, err := sql.Open("odbc", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 测试查询
	//rows, err := db.Query("SELECT * FROM 表名")
	rows, err := db.Query("SELECT ID,字段1,字段2,字段3 FROM go_access")
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	for rows.Next() {
		values := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))
		for i := range values {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		for i, col := range cols {
			fmt.Printf("%s: %v\t", col, values[i])
		}
		fmt.Println()
	}
}
