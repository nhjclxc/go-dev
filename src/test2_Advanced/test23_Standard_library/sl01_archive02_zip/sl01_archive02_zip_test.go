package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

// archive/zip 是 Go 标准库中的一个包，用来 读取和创建 ZIP 压缩文件（.zip）。
// 和 archive/tar 不同，它自带压缩能力（通常是 deflate），不需要额外配合 gzip。
/*
| 特性   | zip  | tar   |
| ---- | ---- | ----- |
| 压缩   | ✅ 自带 | ❌ 不压缩 |
| 读取方式 | 随机访问 | 流式    |
| 结构   | 中央目录 | 顺序流   |
| 适合场景 | 文件分发 | 流式传输  |

zip = 打包 + 压缩 + 可随机读
tar = 纯打包（常配 gzip）
*/
func TestName(t *testing.T) {
	zip.NewWriter(nil)
}

func TestWriteZip(t *testing.T) {
	filename := "zip1.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read file: %v", err)
		return
	}

	testZip, err := os.Create("test.zip")
	if err != nil {
		log.Fatalf("create file: %v", err)
		return
	}
	defer testZip.Close()

	writer := zip.NewWriter(testZip)
	defer writer.Close()

	header, err := writer.CreateHeader(&zip.FileHeader{Name: filename, Method: zip.Deflate, Modified: time.Now()})
	if err != nil {
		log.Fatalf("create header: %v", err)
		return
	}
	write, err := header.Write(file)
	if err != nil {
		log.Fatalf("write header: %v", err)
		return
	}
	if write != len(file) {
		log.Fatalf("write wrong. wrote %d bytes, expected %d", write, len(file))
	}
	fmt.Println(".zip written !")

}

func TestReadZip(t *testing.T) {

	filename := "test.zip"
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatalf("Read zip file: %v", err)
		return
	}
	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			log.Fatalf("Open file: %v", err)
			return
		}

		data_bytes, err := io.ReadAll(rc)
		if err != nil {
			log.Fatalf("Read zip file: %v", err)
			return
		}
		fmt.Printf("data=%+v\n", string(data_bytes))

		//file.FileHeader.Name = file.Name
		fmt.Printf("file.Name=%+v\n", file.Name)
		fmt.Printf("file.Method=%+v\n", file.Method)
		fmt.Printf("file.Modified=%+v\n", file.Modified)

	}

}
