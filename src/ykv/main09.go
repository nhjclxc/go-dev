package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

// 使用golang和ffmpeg实现优酷ykv视频解码为mp4
func main() {

	// 环境准备：
	// 1、golang环境：go1.24.0，https://golang.google.cn/dl/，里面找到go1.24.*.windows-amd64.zip，只要是go1.24版本即可
	// 2、ffmpeg环境：ffmpeg-2025-07-10-git-82aeee3c19-essentials_build，下载地址：https://www.gyan.dev/ffmpeg/builds/或https://github.com/GyanD/codexffmpeg/releases/tag/2025-07-10-git-82aeee3c19找到ffmpeg-2025-07-10-git-82aeee3c19-essentials_build.7z之后解压到指定路径


	// 代码实现步骤：
	// 1、读取原ykv文件
	// 2、解码原ykv文件获取所有视频分片二级制数据
	// 3、

	inputFile := "D:\\code\\go\\go-dev\\src\\ykv\\video2.ykv"
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	offsets := findFtypOffsets(data)
	if len(offsets) == 0 {
		fmt.Println("未找到任何 ftyp 片段")
		return
	}
	fmt.Printf("共发现 %d 个 MP4 分片\n", len(offsets))

	// 补充最后一段终点为文件尾
	offsets = append(offsets, len(data))
	sort.Ints(offsets)

	// 创建 filelist.txt
	listFile, err := os.Create("filelist-2.txt")
	if err != nil {
		panic(err)
	}
	defer listFile.Close()

	for i := 0; i < len(offsets)-1; i++ {
		start := offsets[i]
		end := offsets[i+1]
		filename := fmt.Sprintf("part2-%d.mp4", i+1)

		err := os.WriteFile(filename, data[start:end], 0644)
		if err != nil {
			panic(err)
		}

		fmt.Printf("✅ 提取 %s 成功，大小：%d 字节\n", filename, end-start)

		// 写入到 filelist.txt
		_, err = listFile.WriteString(fmt.Sprintf("file '%s'\n", filename))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("\n🚀 所有分片已保存并生成 filelist-2.txt，你现在可以运行：")
	fmt.Println("ffmpeg -f concat -safe 0 -i filelist-2.txt -c copy full_output-2.mp4")



}

func findFtypOffsets(data []byte) []int {
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


/*
共发现 7 个 MP4 分片
✅ 提取 part1.mp4 成功，大小：51125482 字节
✅ 提取 part2.mp4 成功，大小：50638181 字节
✅ 提取 part3.mp4 成功，大小：54370259 字节
✅ 提取 part4.mp4 成功，大小：52553230 字节
✅ 提取 part5.mp4 成功，大小：55635110 字节
✅ 提取 part6.mp4 成功，大小：48883289 字节
✅ 提取 part7.mp4 成功，大小：43059133 字节

🚀 所有分片已保存并生成 filelist.txt，你现在可以运行：
ffmpeg -f concat -safe 0 -i filelist.txt -c copy full_output.mp4

*/