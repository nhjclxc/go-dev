package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	// 修改为你的 .accdb 文件路径
	dbPath := `D:\code\go\go-dev\src\test7_db\test71_access\Database1.accdb`
	dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;", dbPath)

	db, err := sql.Open("odbc", dsn)
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close()

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal("无法连接到数据库:", err)
	}
	fmt.Println("连接成功！")

	// 查询 go_access 表
	rows, err := db.Query("SELECT * FROM go_access")
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("无法获取列名:", err)
	}
	fmt.Println("列名:", columns)

	// 构造一个 values slice 存储每行数据
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 逐行输出结果
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal("读取行数据失败:", err)
		}
		for i, col := range values {
			fmt.Printf("%s: %s\n", columns[i], string(col))
		}
		fmt.Println("----------")
	}
}
