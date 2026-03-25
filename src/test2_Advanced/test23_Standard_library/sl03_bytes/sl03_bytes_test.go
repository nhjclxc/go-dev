package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

// bytes 是 Go 标准库中一个非常核心的包，主要用于：高效处理 []byte（字节切片）数据

func TestName(t *testing.T) {
	bytess := []byte(" a - ")
	fmt.Println(bytess)
	fmt.Println(bytes.TrimSpace(bytess))
	fmt.Println(string(bytes.TrimSpace(bytess)))
}

// 操作 []byte（查找、替换、分割等）
func TestBaseOptIndex(t *testing.T) {
	byteArr := []byte("qwertytrewq")
	fmt.Printf("Index=%d\n", bytes.Index(byteArr, []byte("e")))
	fmt.Printf("IndexAny=%d\n", bytes.IndexAny(byteArr, "e"))
	fmt.Printf("IndexByte=%d \n", bytes.IndexByte(byteArr, byte('e')))
	fmt.Printf("IndexRune=%d \n", bytes.IndexRune(byteArr, 'e'))
	fmt.Printf("IndexFunc=%d \n", bytes.IndexFunc(byteArr, func(r rune) bool {
		return r == 'e'
	}))
	fmt.Printf("LastIndex=%d\n", bytes.LastIndex(byteArr, []byte("e")))

	fmt.Printf("Contains=%t\n", bytes.Contains(byteArr, []byte("e")))

}

func TestBaseOptReplace(t *testing.T) {
	byteArr := []byte("qwertytrewq")
	// n is replacing times, if n=-1 that replacing times no limit
	fmt.Printf("Replace 1 =%s\n", string(bytes.Replace(byteArr, []byte("e"), []byte("jk"), 1)))
	fmt.Printf("Replace -1 =%s\n", string(bytes.Replace(byteArr, []byte("e"), []byte("jk"), -1)))
	//bytes.ReplaceAll()
}

func TestBaseOptSplit(t *testing.T) {
	byteArr := []byte("qwertytrewq")
	fmt.Printf("Split=%s\n", bytes.Split(byteArr, []byte("e")))
	fmt.Printf("SplitN=%s\n", bytes.SplitN(byteArr, []byte("e"), 2))
	for part := range bytes.SplitSeq(byteArr, []byte("e")) {
		fmt.Printf("SplitSeq.part=%s\n", part)
	}
	//bytes.SplitAfter()

}

func TestBuffer(t *testing.T) {
	var buf bytes.Buffer

	buf.WriteString("hello ")
	buf.Write([]byte("world"))
	fmt.Println(buf.String())
	buf.WriteString("! ")
	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())

	buf.Reset()
	fmt.Println(buf.Len())
	fmt.Println(buf.String())

}

func TestBuffer2(t *testing.T) {
	// bytes.Buffer 做字符串拼接
	var buf bytes.Buffer
	for i := range 10 {
		buf.WriteByte(byte(i + 97))
	}
	fmt.Println(buf.String())

}

// bytes.Reader 把 []byte 转成 io.Reader
func TestReader(t *testing.T) {
	reader := bytes.NewReader([]byte("hello world"))
	bytess, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("io.ReadAll: %v", err)
		return
	}
	fmt.Println(string(bytess))
}
