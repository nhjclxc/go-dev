package main

import (
	"fmt"
	"time"
)

// 时间戳转换成日期字符串
func main() {

	unixTime := 1713152804
	timeObj := time.Unix(int64(unixTime), 0)
	var str = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str) //2024-04-15 11:46:44

}
