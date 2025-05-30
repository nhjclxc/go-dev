
package main

import (
"database/sql"
"encoding/json"
"fmt"
"log"
"net/http"
"strings"
"time"

_ "github.com/alexbrainman/odbc"
)
type Record struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func readAccessDB() ([]Record, error) {
	// 连接字符串，注意替换你的文件路径
	dsn := "Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=\\\\192.168.9.102\\C:\\Users\\maxin\\Documents\\Database2.accdb"
	db, err := sql.Open("odbc", dsn)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}
	defer db.Close()

	//
	//// 创建表
	//createTableSQL := `
    //    CREATE TABLE my_table (
    //        id AUTOINCREMENT PRIMARY KEY,
    //        name TEXT(100),
    //        value TEXT(100)
    //    )
    //`
	//_, err = db.Exec(createTableSQL)
	//if err != nil {
	//	panic(fmt.Sprintf("建表失败: %v", err))
	//}
	//
	//fmt.Println("表 my_table 创建成功！")


	//rows, err := db.Query("SELECT id, name, value FROM my_table")
	rows, err := db.Query("SELECT ID,字段1,字段2,字段3 FROM TABLE")
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	defer rows.Close()

	var results []Record
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Value); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}

func uploadToServer(data []Record) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://192.168.8.120:8080/aa", "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("上传失败，HTTP 状态码: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		log.Println("开始读取 Access 数据...")
		data, err := readAccessDB()
		if err != nil {
			log.Printf("读取失败: %v\n", err)
		} else {
			log.Printf("读取 %d 条记录，开始上传...\n", len(data))
			if err := uploadToServer(data); err != nil {
				log.Printf("上传失败: %v\n", err)
			} else {
				log.Println("上传成功。")
			}
		}

		<-ticker.C
	}
}
