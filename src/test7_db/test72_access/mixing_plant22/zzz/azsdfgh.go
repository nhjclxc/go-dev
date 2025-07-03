package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "uuids.txt"

	// 写入示例
	newUUIDs := []string{
		"123e4567-e89b-12d3-a456-426614174200",
		"123e4567-e89b-12d3-a456-426614174201",
		"123e4567-e89b-12d3-a456-426614174202",
		"123e4567-e89b-12d3-a456-426614174203",
		"123e4567-e89b-12d3-a456-426614174204",
		"123e4567-e89b-12d3-a456-426614174205",
	}
	err := WriteUUIDsToFile(filename, newUUIDs)
	if err != nil {
		fmt.Println("写入出错:", err)
	}

	// 读取示例
	uuids, err := ReadUUIDsFromFile(filename)
	if err != nil {
		fmt.Println("读取出错:", err)
	}
	fmt.Println("读取到的 UUIDs:")
	for _, u := range uuids {
		fmt.Println(u)
	}

	fmt.Println("11111")

	now := []string{
		"123e4567-e89b-12d3-a456-426614174000",
		"123e4567-e89b-12d3-a456-426614174001",
		"123e4567-e89b-12d3-a456-426614174002",
	}

	unprocessedUUIDList2 := FilterNewUUIDs(now, uuids)

	fmt.Println("未处理的 UUID222：", unprocessedUUIDList2)

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