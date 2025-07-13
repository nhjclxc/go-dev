package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func main02() {
	inputFile := "D:\\Program Files (x86)\\YOUKU\\9.2.59.1003\\???.ykv"

	outputFile := "output.mp4"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	// 搜索 MP4 的 box "ftyp"，往前回溯 4 字节尝试读取 box size
	ftypIndex := bytes.Index(data, []byte("ftyp"))
	if ftypIndex == -1 || ftypIndex < 4 {
		fmt.Println("未找到有效的 ftyp 标记")
		return
	}

	// 向前回溯 4 字节，获取 box size
	boxStart := ftypIndex - 4
	boxSize := binary.BigEndian.Uint32(data[boxStart:ftypIndex])

	fmt.Printf("找到 ftyp box，起始位置: %d，box size: %d\n", boxStart, boxSize)

	// 简单判断 box size 是否合理
	if boxSize < 8 || boxSize > uint32(len(data))-uint32(boxStart) {
		fmt.Println("box size 不合理，可能不是合法的 MP4")
		return
	}

	// 从 boxStart 开始写入新文件
	err = os.WriteFile(outputFile, data[boxStart:], 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("mp4 提取完成，尝试使用 VLC 或 ffmpeg 打开 output.mp4")
}
