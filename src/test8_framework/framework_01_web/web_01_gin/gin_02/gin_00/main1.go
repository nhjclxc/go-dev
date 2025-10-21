package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 将 log 输出重定向到文件
	logger := log.New(file, "", log.LstdFlags)

	logger.Println("✅ Go binary is running on Android!")

	for {
		time.Sleep(2 * time.Second)
		logger.Println("tick...")
	}
}
