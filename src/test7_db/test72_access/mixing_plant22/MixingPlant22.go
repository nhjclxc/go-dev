package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/alexbrainman/odbc"
)

type Config struct {
	ClientId string  `json:"clientId"`
	Path     string  `json:"path"`
	Pwd      string  `json:"pwd"`
	Fixed    int     `json:"fixed"`
	Env      int     `json:"env"`
	Win      int     `json:"win"`
	Tables   []Table `json:"tables"`
}

type Table struct {
	Table  string `json:"table"`
	Stamp  string `json:"stamp"`
	Fields string `json:"fields"`
}

type PushData struct {
	ClientId string                 `json:"clientId"`
	DataMap  map[string]interface{} `json:"dataMap"`
}

func readConfig(filepath string) Config {
	dir, _ := os.Getwd()
	fmt.Println("当前工作目录:", dir)
	// 打开文件
	//file, err := os.Open(dir + "\\" + filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("打开配置文件失败: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("文件关闭失败: %v", err)
		}
	}(file)

	// 读取文件内容
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 解析 JSON
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatalf("解析 JSON 失败: %v", err)
	}

	// 打印配置
	//fmt.Printf("应用名: %#v\n", config)

	return config
}

func main() {
	// 可执行文件构建命令
	// rsrc -ico favicon.ico -o rsrc.syso
	// go build -ldflags="-H windowsgui" -o MixingPlant.exe
	// MixingPlant.exe -h
	// MixingPlant.exe -uuid=BSBHZ01 -path=./BCS7.2.mdb -pwd=BCS7.2_SDBS -env=0 -fixed=5

	// go build -o MixingPlant2.exe MixingPlant.go

	// go build -ldflags="-s -w" -o MixingPlant722.exe MixingPlant22.go

	// win7编译命令
	// D:\develop\go17\bin\go.exe build -ldflags="-s -w" -o MixingPlant722.exe MixingPlant22.go

	config := readConfig("config.json")

	// 打印读取结果
	fmt.Println("运行参数如下:")
	fmt.Println("ㅤㅤㅤ全局唯一ClientId:", config.ClientId)
	fmt.Println("ㅤㅤㅤ数据库路径:", config.Path)
	fmt.Println("ㅤㅤㅤ数据库密码:", config.Pwd)
	fmt.Println("ㅤㅤㅤ推送间隔:", config.Fixed)
	fmt.Println("ㅤㅤㅤ启动环境:", config.Env)

	// win7
	driver := "*.mdb"
	if config.Win == 11 || config.Win == 10 {
		// win11
		driver = "*.mdb, *.accdb"
	}

	//dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;PWD=%s;", config.Path, config.Pwd)
	dsn := fmt.Sprintf("Driver={Microsoft Access Driver (%s)};Dbq=%s;PWD=%s;", driver, config.Path, config.Pwd)

	fmt.Println("dsn: " + dsn)

	db, err := sql.Open("odbc", dsn)
	if err != nil {
		log.Fatal("无法打开数据库:", err)
		return
	}
	defer db.Close()

	// 定时每分钟执行
	//ticker := time.NewTicker(time.Second * 5)
	ticker := time.NewTicker(time.Second * time.Duration(config.Fixed))
	defer ticker.Stop()

	// 启动环境配置
	domain := ""
	if config.Env == 0 {
		domain = "https://127.0.0.1:8080"
	} else if config.Env == 1 {
		domain = "https://api-sc-dev.hkznkj.com"
	} else if config.Env == 2 {
		domain = "https://api-sc.hkznkj.com"
	} else {
		log.Fatalf("环境配置错误，仅支持0、1、2的选项！！！")
	}

	var pushData *PushData = &PushData{
		ClientId: config.ClientId,
		DataMap: make(map[string]interface{}, len(config.Tables)),
	}

	// 立即执行一次
	doTask(config.Fixed, db, domain, config.ClientId, config.Tables, pushData)

	// 按固定间隔执行任务
	for range ticker.C {
		doTask(config.Fixed, db, domain, config.ClientId, config.Tables, pushData)

	}
}

func doTask(fixed int, db *sql.DB, domain string, clientId string, tables []Table, pushData *PushData) {
	// 获取当前时间和一分钟前时间
	now := time.Now()
	//now = now.AddDate(0, 0, -9)
	oneMinuteAgo := now.Add(time.Second * time.Duration(fixed) * -1)


	dataCount := 0
	for _, table := range tables {
		data := query(db, now, oneMinuteAgo, table.Table, table.Stamp, table.Fields)
		if data == nil {
			continue
		}
		dataCount += len(data)

		// 转为 JSON
		jsonBytes, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(jsonBytes))

		itemData := strings.ReplaceAll(string(jsonBytes), "\t", "")
		itemData = strings.ReplaceAll(itemData, "\n", "")

		pushData.DataMap[table.Table] = itemData
	}

	//fmt.Printf("查询时间段：start: " + oneMinuteAgo.String() + ", end: " + now.String() + "\n")
	fmt.Printf("查询时间段：start: " + oneMinuteAgo.String() + ", end: " + now.String() + ", dataCount = " + strconv.Itoa(dataCount) + "\n")


	if dataCount > 0 {
		uploadToServer(domain+"/device/mining/plant/"+clientId, pushData, dataCount)
	}
}

func uploadToServer(url string, pushData *PushData, dataCount int) {
	fmt.Printf("url = %s, size = %d. \n", url, dataCount)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 跳过证书校验
	}
	client := &http.Client{Transport: tr}

	jsonBytes, err := json.MarshalIndent(pushData, "", "  ")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(jsonBytes))

	resp, err := client.Post(url, "application/json", strings.NewReader(string(jsonBytes)))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf(fmt.Sprintf("上传失败，HTTP 状态码: %d", resp.StatusCode))
		return
	}

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 将响应体转换为字符串（如果需要）
	respStr := string(respBody)
	fmt.Printf("url = %s, response = %s. \n", url, respStr)

}


func query(db *sql.DB, now time.Time, oneMinuteAgo time.Time, table string, stamp string, fields string) []map[string]interface{} {

	// 拼接表名、字段名、时间字段
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE %s BETWEEN ? AND ?`, fields, table, stamp)
	rows, err := db.Query(query, now, oneMinuteAgo)
	//query := fmt.Sprintf(`SELECT %s FROM %s`, fields, table)
	//rows, err := db.Query(query)

	if err != nil {
		log.Println("查询失败:", err)
		return nil
	}
	defer rows.Close()

	results, err := readRowsAsMap(rows)
	if err != nil {
		log.Println("字段读取出错:", err)
		return nil
	}

	return results

	//results = results[:5]
	//return results
}

func readRowsAsMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// 用于保存所有行
	results := []map[string]interface{}{}

	// 每行处理
	for rows.Next() {
		// 创建一个 interface{} 切片用于存每个字段
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// 扫描一行
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// 构造 map
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]

			// 处理 []byte => string（Access/ODBC 经常返回 []byte）
			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}

	return results, nil
}

// dosage.FinTim = fromOADate(dosage.FinTim)
var nowTime = time.Now()

func fromOADate(oa string) string {
	const OABase = "1899-12-30T00:00:00Z"
	baseTime, err := time.Parse(time.RFC3339, OABase)
	if err != nil {
		baseTime = nowTime
	}

	f, err := strconv.ParseFloat(oa, 64)
	if err != nil {
		baseTime = nowTime
	}

	seconds := f * 24 * 60 * 60
	t := baseTime.Add(time.Duration(seconds * float64(time.Second)))

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t = t.In(loc)

	if t.Before(time.Date(2000, 1, 1, 0, 0, 0, 0, loc)) {
		t = nowTime
	}

	return t.Format("2006-01-02 15:04:05")
}