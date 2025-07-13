package main

import (
	"bytes"
	"fmt"
	"os"
)

// 查找所有 ftyp 片段起始位置
func findFtypOffsets1(data []byte) []int {
	var offsets []int
	search := []byte("ftyp")
	i := 0
	for {
		index := bytes.Index(data[i:], search)
		if index == -1 {
			break
		}
		offset := i + index - 4 // 回溯4字节包含box length
		if offset >= 0 {
			offsets = append(offsets, offset)
		}
		i = i + index + 4
	}
	return offsets
}

func main05() {
	inputFile := "D:\\Program Files (x86)\\YOUKU\\9.2.59.1003\\???.ykv"

	outputFile := "full_output.mp4"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("len(data) = %d \n", len(data))

	offsets := findFtypOffsets1(data)
	if len(offsets) == 0 {
		fmt.Println("未找到任何 ftyp 片段，可能不是优酷 ykv 文件")
		return
	}

	fmt.Printf("共发现 %d 个 MP4 分片\n", len(offsets))

	// 最后一段需要以文件结尾为终点
	offsets = append(offsets, len(data))

	// 提取所有 MP4 分片并拼接到一起
	out, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	for i := 0; i < len(offsets)-1; i++ {
		start := offsets[i]
		end := offsets[i+1]

		fmt.Printf("提取第 %d 个片段：offset %d → %d（大小：%d 字节）\n", i+1, start, end, end-start)
		//_, err := out.Write(data[start:end])
		//if err != nil {
		//	panic(err)
		//}
	}

	fmt.Println("已完成所有片段的提取与拼接，输出为 full_output.mp4")
}
