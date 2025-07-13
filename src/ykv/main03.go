package main

import (
	"bytes"
	"fmt"
	"os"
)

// [ftyp][moov][mdat]... [ftyp][moov][mdat]... [ftyp][moov][mdat]...
func main03() {
	inputFile := "D:\\Program Files (x86)\\YOUKU\\9.2.59.1003\\???.ykv"

	data, _ := os.ReadFile(inputFile)

	keyword := []byte("mdat")
	offset := 0
	for {
		i := bytes.Index(data[offset:], keyword)
		if i == -1 {
			break
		}
		fmt.Printf("发现 mdat 起始于 offset: %d\n", offset+i-4) // 回溯4字节查看 box size
		offset += i + 4
	}
}
