package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)
import "archive/tar"

// archive/tar 是 Go 标准库中的一个包，用来 读取和写入 tar 归档文件（类似 .tar、.tar.gz 里的 tar 部分）。
// 对文件进行打包生成.tar， 对.tar文件进行解包

// Test_write 写tar
func Test_write(t *testing.T) {
	create, err := os.Create("test.tar")
	if err != nil {
		log.Println("create error:", err)
		return
	}
	defer create.Close()

	data := []byte("对文件进行打包生成.tar")
	writer := tar.NewWriter(create)
	defer writer.Close()

	writer.WriteHeader(&tar.Header{
		Name: "test.tar",
		Size: int64(len([]byte(data))),
	})

	writen, err := writer.Write(data)
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	if writen != len(data) {
		log.Println("writen != len(data):", writen)
	}
	fmt.Println("writen:", writen)

}

// Test_read 读取tar
func Test_read(t *testing.T) {
	test_tar, err := os.Open("test.tar")
	if err != nil {
		log.Println("open error:", err)
		return
	}
	defer test_tar.Close()

	reader := tar.NewReader(test_tar)

	for {
		tar_header, err := reader.Next()
		if err != nil {
			log.Println("tar header error:", err)
			return
		}

		fmt.Printf("tar_header.Name=%+v\n", tar_header.Name)
		fmt.Printf("tar_header.Size=%+v\n", tar_header.Size)
		bytes := make([]byte, tar_header.Size)
		read, err := reader.Read(bytes)
		if err != nil && err != io.EOF {
			log.Println("read error:", err)
			return
		}
		fmt.Printf("read=%+v\n", read)
		fmt.Printf("data=%+v\n", string(bytes))

		if err == io.EOF {
			break
		}
	}

}

// Test_gzip_write 对 .tar 进行压缩
func Test_gzip_write(t *testing.T) {

	data := "对 .tar 进行压缩"
	name := "test2.tar.gz"
	test2TarGz, err := os.Create(name)
	if err != nil {
		log.Println("create error:", err)
		return
	}
	defer test2TarGz.Close()

	gzipWriter := gzip.NewWriter(test2TarGz)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	tarHeader := tar.Header{Name: name, Size: int64(len(data)), ModTime: time.Now(), Mode: 0644}
	tarWriter.WriteHeader(&tarHeader)
	write, err := tarWriter.Write([]byte(data))
	if err != nil {
		log.Println("write error:", err)
		return
	}
	if write != len(data) {
		log.Println("write != len(data):", write)
	}
	fmt.Println(".tar.gz write successful, size:", write)

}

// Test_gzip_read 读取 .tar.gzip 压缩的数据
func Test_gzip_read(t *testing.T) {

	test2TarGz, err := os.Open("test2.tar.gz")
	if err != nil {
		log.Println("open error:", err)
		return
	}
	defer test2TarGz.Close()

	gzipReader, err := gzip.NewReader(test2TarGz)
	if err != nil {
		log.Println("read error:", err)
		return
	}
	defer gzipReader.Close()

	reader := tar.NewReader(gzipReader)
	for {
		next, err := reader.Next()
		if err != nil {
			log.Println("next error:", err)
			return
		}

		fmt.Printf("tar_header.Name=%+v\n", next.Name)
		fmt.Printf("tar_header.Size=%+v\n", next.Size)
		fmt.Printf("tar_header.ModTime=%+v\n", next.ModTime)
		fmt.Printf("tar_header.Mode=%+v\n", next.Mode)

		bytes := make([]byte, next.Size)
		read, err := reader.Read(bytes)
		if err != nil && err != io.EOF {
			log.Println("read error:", err)
			return
		}
		fmt.Printf("read=%+v\n", read)
		fmt.Printf("data=%+v\n", string(bytes))

		if err == io.EOF {
			break
		}
	}
}

// 将多个数据打包成一个tar文件
func TestWriteMultData(t *testing.T) {
	datas := []string{
		"TestWriteMultData.tar",
		"TestWriteMultData.tar.gz",
		"TestWriteMultData.tar.xz",
	}

	create, err := os.Create("test3.tar")
	if err != nil {
		log.Println("create error:", err)
		return
	}
	defer create.Close()

	// ✅ 正确写法：一个 tar.Writer 写所有文件
	// 👉 关键原则：一个 tar 文件 = 一个 tar.Writer
	// tar.NewWriter(create) 必须放到外面
	tarWriter := tar.NewWriter(create)
	defer tarWriter.Close()

	for _, data := range datas {
		hdr := &tar.Header{Name: data, Size: int64(len(data)), Mode: 0644, ModTime: time.Now()}

		// 写 header
		if err := tarWriter.WriteHeader(hdr); err != nil {
			log.Println("write header error:", err)
			return
		}

		// 写内容
		if _, err := tarWriter.Write([]byte(data)); err != nil {
			log.Println("write error:", err)
			return
		}
		log.Printf("[%s] write len(data): %d \n", data, len(data))
	}

}

// 读取一个tar文件里面的多个数据
func TestReadMultData(t *testing.T) {

	test3Tar, err := os.Open("test3.tar")
	if err != nil {
		log.Println("open error:", err)
		return
	}
	defer test3Tar.Close()
	reader := tar.NewReader(test3Tar)
	for {
		next, err := reader.Next()
		if err == io.EOF {
			log.Println("EOF:", next)
			break
		}
		if err != nil && err != io.EOF {
			log.Println("read error:", err)
			return
		}
		if next == nil {
			break
		}

		fmt.Printf("tar_header.Name=%+v\n", next.Name)
		fmt.Printf("tar_header.Size=%+v\n", next.Size)
		fmt.Printf("tar_header.ModTime=%+v\n", next.ModTime)
		fmt.Printf("tar_header.Mode=%+v\n", next.Mode)

		bytes := make([]byte, next.Size)
		read, err := reader.Read(bytes)
		if err != nil && err != io.EOF {
			log.Println("read error:", err)
			return
		}
		fmt.Printf("read=%+v\n", read)
		fmt.Printf("data=%+v\n", string(bytes))

		fmt.Println()
	}
}
