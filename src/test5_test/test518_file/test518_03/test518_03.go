package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 写文件操作
func main() {

	/*
		os.flag
			const (
			    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
			    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
			    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
			    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
			    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
			    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
			    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
			    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
			)
	*/

	// 打开一个文件，并且写入数据
	// O_RDWR | O_CREAT
	//test518_03_01()

	// 追加数据
	// O_RDWR | O_APPEND
	// test518_03_02()

	// 先读取后追加
	// O_RDWR | O_APPEND
	//test518_03_03()

	test518_03_04()

}

func test518_03_04() {
	// 统计一个文件种含有多少个英文字符（大小写字母）、数字、空格、回车符以及其他字符数量

	// 打开文件
	file, _ := os.OpenFile("src/test5_test/test518_file/temp/test518_03.txt", os.O_RDONLY, 0644)

	// 关闭文件
	defer file.Close()

	// 读取文件
	reader := bufio.NewReader(file)

	wordCount, numCount, spaceCount, enterCount, otherCount := 0, 0, 0, 0, 0
	// 统计字符
	for true {
		// 一行一行读取
		newLine, err := reader.ReadString('\n')
		// line := scanner.Text() // ReadString('\n') 包括 '\n'，但 scanner.Text() 不包括

		// 判断是否到文件末尾
		if err == io.EOF {
			break
		}

		// 统计
		for _, val := range newLine {
			//fmt.Println(index, val)
			if ('a' <= val && val <= 'z') || ('A' <= val && val <= 'Z') {
				wordCount++
			} else if '0' <= val && val <= '9' {
				numCount++
			} else if ' ' == val {
				spaceCount++
			} else if '\n' == val {
				enterCount++
			} else {
				otherCount++
			}
		}
	}

	fmt.Printf("英文字符：%d，数字：%d，空格：%d，回车符：%d，其他字符：%d。\n",
		wordCount, numCount, spaceCount, enterCount, otherCount)

}

func test518_03_03() {

	// 1、打开
	file, _ := os.OpenFile("src/test5_test/test518_file/temp/test518_03.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	// 2、关闭
	defer file.Close()

	// 3、读取
	reader := bufio.NewReader(file)

	for true {
		newLine, readErr := reader.ReadString('\n')
		if readErr == io.EOF {
			break
		}
		fmt.Println(newLine)
	}

	// 4、写入
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		_, err := writer.WriteString("你好，GoLang 世界 " + strconv.Itoa((i * 5)) + "\n")
		if err != nil {
			return
		}
	}

	// 5、刷新缓冲区
	err := writer.Flush()
	if err != nil {
		return
	}

}

// 向一个已有的文件中追加新的数据
func test518_03_02() {

	// 打开文件
	file, _ := os.OpenFile("src/test5_test/test518_file/temp/test518_03.txt", os.O_RDWR|os.O_APPEND, 0644)

	// 关闭文件
	defer file.Close()

	// 创建缓冲区 bufio
	writer := bufio.NewWriter(file)

	// 追加文件
	for i := 0; i < 6; i++ {
		_, _ = writer.WriteString("你好， GoLang\n")
	}

	// 刷新缓冲区
	err := writer.Flush()
	if err != nil {
		return
	}

}

// 创建一个新文件写入一些数据
func test518_03_01() {

	// 创建一个新文件
	// func OpenFile(name string, flag int, perm FileMode) (*File, error)
	// 注意：
	//		第一个传输是要打开文件的路径
	// 		第二个参数：os.O_RDWR表示写权限，os.O_CREATE表示该文件不存在时创建该文件
	//		第三个参数：是 nux 类型操作系统的参数，用于控制文件权限的，在win系统下无效
	file, err := os.OpenFile("src/test5_test/test518_file/temp/test518_03.txt", os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic("打开文件出错" + err.Error())
	}

	// 关闭文件
	defer file.Close()

	// 基于 file 创建一个 带缓冲区的 Writer
	// func NewWriter(w io.Writer) *Writer
	writer := bufio.NewWriter(file)

	// 写数据
	for i := 0; i < 5; i++ {
		// writer.WriteString仅仅是将数据写到缓冲区
		byteLength, err := writer.WriteString("Hello, Go!!!\n")
		if err != nil {
			panic("写入缓冲区错误" + err.Error())
		}
		fmt.Println(byteLength)
	}

	// 由于 writer.WriteString 仅仅是将数据写到缓冲区
	// 因此，在 writer.WriteString 后必须调用 writer.Flush() 将缓冲区里面的数据写入磁盘
	// 如果不执行 writer.Flush() 操作，那么写入到缓冲区的数据将无法到磁盘，进而丢失重要数据
	err3 := writer.Flush()
	if err3 != nil {
		panic("刷新缓冲区错误" + err3.Error())
	}

}
