package main

import (
	"bufio"
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
	Date     string  `json:"date"`
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

// BuildTimeWithCurrentClock 传入日期字符串（格式 "2006-01-02"），返回这个日期 + 当前时分秒 的时间对象
func BuildTimeWithCurrentClock(dateStr string) (time.Time, error) {
	// 使用当前时区解析日期
	loc := time.Now().Location()
	parsedDate, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("解析日期失败: %w", err)
	}

	now := time.Now()

	// 构造新时间：用日期字符串中的年月日 + 当前时分秒
	result := time.Date(
		parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		now.Hour(), now.Minute(), now.Second(), now.Nanosecond(),
		loc,
	)

	return result, nil
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
	// D:\develop\go17\bin\go.exe build -ldflags="-s -w" -o MixingPlant7_070205.exe MixingPlant7.go

	config := readConfig("config.json")

	// 打印读取结果
	fmt.Println("运行参数如下:")
	fmt.Println("ㅤㅤㅤ全局唯一ClientId:", config.ClientId)
	fmt.Println("ㅤㅤㅤ数据库路径:", config.Path)
	fmt.Println("ㅤㅤㅤ数据库密码:", config.Pwd)
	fmt.Println("ㅤㅤㅤ推送日期:", config.Date)
	fmt.Println("ㅤㅤㅤ推送间隔:", config.Fixed)
	fmt.Println("ㅤㅤㅤ启动环境:", config.Env)

	// win7
	driver := "*.mdb"
	if config.Win == 11 || config.Win == 10 {
		// win11
		driver = "*.mdb, *.accdb"
	}

	//dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;PWD=%s;", config.Path, config.Pwd)
	//dsn := fmt.Sprintf("Driver={Microsoft Access Driver (%s)};Dbq=%s;PWD=%s;", driver, config.Path, config.Pwd)
	dsn := fmt.Sprintf("Driver={Microsoft Access Driver (%s)};Dbq=%s;PWD=%s;ReadOnly=1;", driver, config.Path, config.Pwd)

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

	var now time.Time
	if config.Date == "" {
		now = time.Now()
	} else {
		now, err = BuildTimeWithCurrentClock(config.Date)
		if err != nil {
			fmt.Println("错误:", err)
			return
		}
	}

	now = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 0, 0, now.Location())

	// 获取今天的起始时间（即 00:00:00）
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 立即执行一次
	doTask(startOfDay, now, db, domain, config.ClientId, config.Tables, pushData)

	// 按固定间隔执行任务
	for range ticker.C {
		doTask(startOfDay, now, db, domain, config.ClientId, config.Tables, pushData)

	}
}

func doTask(startOfDay time.Time, now time.Time, db *sql.DB, domain string, clientId string, tables []Table, pushData *PushData) {
	// 获取当前时间和一分钟前时间
	//now := time.Now()
	//t := -14
	//startOfDay = startOfDay.AddDate(0, 0, t )
	//now = now.AddDate(0, 0, t )
	//oneMinuteAgo := now.Add(time.Second * time.Duration(fixed) * -1)


	dataCount := 0
	for _, table := range tables {
		data := query(db, startOfDay, now, table.Table, table.Stamp, table.Fields)
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
	fmt.Printf("查询时间段：start: " + startOfDay.String() + ", end: " + now.String() + ", dataCount = " + strconv.Itoa(dataCount) + "\n")


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

	rowsed, err := readRowsAsMap(rows)
	if err != nil {
		log.Println("字段读取出错:", err)
		return nil
	}

	// 1、先读取已推送的id
	// 2、基于现在读取到的数据来遍历
	// 3、看看当前的数据的id是不是在已推送的id里面
	// 4、将新的id写入文件


	filename := table + ".txt"

	// 1、先读取已推送的id
	pushedUUIDList, err := ReadUUIDsFromFile(filename)

	// 2、基于现在读取到的数据来遍历
	results := []map[string]interface{}{}

	var nowUUIDList []string
	for _, item := range rowsed {
		if rawId, ok := item["ID"]; ok {
			if idStr, ok := rawId.(string); ok {
				// 3、看看当前的数据的id是不是在已推送的id里面
				flag := containsString(pushedUUIDList, idStr)
				if !flag {
					results = append(results, item)
					nowUUIDList = append(nowUUIDList, idStr)
				}
			}
		}
	}

	// 4、将新的id写入文件
	_ = WriteUUIDsToFile(filename, nowUUIDList)


	return results

	//results = results[:5]
	//return results
}

// 返回true表示存在，false表示不在
func containsString(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
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

// pushedUUIDList: 已处理过的 UUID
// nowUUIDList: 当前读取到的 UUID
// 返回：未处理的新 UUID
func FilterNewUUIDs(nowUUIDList, pushedUUIDList []string) []string {

	// 将 pushedUUIDList 转成 map 用于快速查找
	pushedMap := make(map[string]int8)
	for _, uuid := range nowUUIDList {
		pushedMap[uuid] = 1
	}

	// 过滤出未处理的 uuid
	var unprocessedUUIDList []string
	for _, uuid := range pushedUUIDList {
		//value1, exists1 := pushedMap[uuid]
		//fmt.Println(value1)
		//fmt.Println(exists1)
		if value, exists := pushedMap[uuid]; !exists {
			fmt.Println(value)
			unprocessedUUIDList = append(unprocessedUUIDList, uuid)
		}
	}
	return unprocessedUUIDList
}


// 读取所有 UUID（每行一个）
func ReadUUIDsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var uuids []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			uuids = append(uuids, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return uuids, nil
}

// 批量写入 UUID 到文件（每个 UUID 占一行，追加写入）
func WriteUUIDsToFile(filename string, uuids []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, uuid := range uuids {
		if _, err := writer.WriteString(uuid + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}