package main

import (
	"fmt"
	"time"
)

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
	customTime, err := BuildTimeWithCurrentClock("2025-06-03")
	if err != nil {
		fmt.Println("错误:", err)
		return
	}

	fmt.Println("构造后的时间:", customTime.Format("2006-01-02 15:04:05"))
}
