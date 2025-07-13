package main

import (
	"fmt"
	"os"
)

func main01() {
	filepath := "D:\\Program Files (x86)\\YOUKU\\9.2.59.1003\\???.ykv"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := make([]byte, 512)
	n, err := file.Read(header)
	if err != nil {
		panic(err)
	}

	fmt.Printf("前 %d 字节数据：\n%x\n", n, header[:n])
}
