package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
)

func main() {
	// 1. 读取当前目录所有 mp4 文件（你可以改成指定目录）
	mp4Files, err := filepath.Glob("part2-*.mp4")
	if err != nil {
		log.Fatalf("读取 mp4 文件失败: %v", err)
	}
	if len(mp4Files) == 0 {
		log.Fatal("当前目录未找到任何 part-2*.mp4 文件")
	}

	// 2. 按文件名排序，确保顺序正确
	sort.Strings(mp4Files)

	// 3. 生成 filelist.txt
	filelistPath := "filelist-2.txt"
	filelist, err := os.Create(filelistPath)
	if err != nil {
		log.Fatalf("创建 filelist-2.txt 失败: %v", err)
	}
	defer filelist.Close()

	for _, f := range mp4Files {
		_, err := filelist.WriteString(fmt.Sprintf("file '%s'\n", f))
		if err != nil {
			log.Fatalf("写入 filelist-2.txt 失败: %v", err)
		}
	}

	fmt.Printf("生成 %s 完成，包含 %d 个文件\n", filelistPath, len(mp4Files))

	// 4. 调用 FFmpeg 执行合并
	outputFile := "merged_output-2.mp4"
	cmd := exec.Command("D:\\develop\\ffmpeg-2025-07-10-git-82aeee3c19-essentials_build\\ffmpeg-2025-07-10-git-82aeee3c19-essentials_build\\bin\\ffmpeg", "-f", "concat", "-safe", "0", "-i", filelistPath, "-c", "copy", outputFile)

	// 输出 ffmpeg 运行日志到控制台
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("开始执行 ffmpeg 合并视频，请稍候...")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("ffmpeg 执行失败: %v", err)
	}

	fmt.Printf("合并完成，输出文件：%s\n", outputFile)
}
