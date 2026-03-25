package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"testing"
	"time"
)

// **零拷贝（zero-copy）**指：尽量避免内存拷贝，把数据从源直接传给下一个处理环节。

/*
io.Reader / io.Writer 接口

	Go 的 I/O 基于接口：Read(p []byte) 和 Write(p []byte)
	数据可以直接从一个 Reader 传到 Writer，而不必先存到中间变量

bytes.Buffer

	内存缓冲区，但可以直接作为 Reader 或 Writer
	避免每次操作都创建新 slice

io.Copy

	内置函数 io.Copy(dst io.Writer, src io.Reader) 就是零拷贝的典型例子
	它直接从 src 读一块数据，写到 dst，无需你手动 copy

*/

/*
源数据 → tar.Writer → gzip.Writer → bytes.Buffer → 网络/文件
*/
func TestZeroCopy(t *testing.T) {

}

func TestZeroCopyTar(t *testing.T) {
	// 场景1：打包多个文件，压缩成 gzip，发送到客户端

	files := make(map[string][]byte)
	files["f1.txt"] = []byte("qwertyuiop")
	files["f2.txt"] = []byte("lkjhgfdsa")
	files["f3.txt"] = []byte("zxcvbnm")

	buf := new(bytes.Buffer)

	gzipWriter := gzip.NewWriter(buf)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for name, data := range files {
		hdr := &tar.Header{
			Name:    name,
			Mode:    0644,
			Size:    int64(len(data)),
			ModTime: time.Now(),
		}
		tarWriter.WriteHeader(hdr)
		tarWriter.Write(data)
	}

	reader := bytes.NewReader(buf.Bytes())

	// 直接输出到网络或文件
	//io.Copy(os.Stdout, reader)

	create, err := os.Create("zeorcopy.tar.gz")
	if err != nil {
		return
	}
	io.Copy(create, reader)

}
