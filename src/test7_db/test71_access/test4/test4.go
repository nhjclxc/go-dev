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

func main() {
	// 修改为你的实际路径
	dbPath := `D:\code\go\go-dev\src\test7_db\test71_access\Database1.accdb`
	dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;", dbPath)

	db, err := sql.Open("odbc", dsn)
	if err != nil {
		log.Fatal("无法打开数据库:", err)
	}
	defer db.Close()

	// 定时每分钟执行
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	// 第一次立即执行一次
	//queryLastMinuteData(db)

	for range ticker.C {
		var results []Record
		results = queryLastMinuteData(db)

		if results != nil {
			uploadToServer(results)
		}

		fmt.Printf("执行上传22\n")
	}
}

func queryLastMinuteData(db *sql.DB) ([]Record) {
	// 获取当前时间和一分钟前时间
	//now := time.Now()
	//oneMinuteAgo := now.Add(-1 * time.Minute)

	// Access 中时间格式为 #yyyy-mm-dd hh:nn:ss#
	//fromStr := oneMinuteAgo.Format("2006-01-02 15:04:05")
	//toStr := now.Format("2006-01-02 15:04:05")

	//query := fmt.Sprintf(`
	//	SELECT id, name, value FROM go_access
	//	WHERE created_at BETWEEN #%s# AND #%s#
	//`, fromStr, toStr)

	query := fmt.Sprintf(`
		SELECT id, name, value FROM go_access
	`)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("查询失败:", err)
		return nil
	}
	defer rows.Close()

	var results []Record
	for rows.Next() {
		var r Record
		err := rows.Scan(&r.ID, &r.Name, &r.Value)
		if err != nil {
			log.Println("解析数据失败:", err)
			continue
		}
		results = append(results, r)
	}

	// 输出结果
	//fmt.Printf("---- [%s] 读取到 %d 条记录 ----\n", now.Format("15:04:05"), len(results))
	for _, rec := range results {
		fmt.Printf("%+v\n", rec)
	}
	return results
}



func uploadToServer(data []Record) error {
	if true {
		fmt.Printf("执行上传\n")
		return nil
	}
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
