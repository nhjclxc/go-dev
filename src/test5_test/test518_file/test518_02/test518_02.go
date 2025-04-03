package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 使用带有缓存的bufio读取文件内容
func main() {

	// 1、打开文件
	//file, err := os.Open("D:\\code\\go\\go-dev\\.gitignore")
	file, err := os.Open(".gitignore")

	// 2、判断文件是否存在
	if nil != err {
		panic("文件打开失败！！！" + err.Error())
	}

	// 3、文件打开成功之后，立马配合defer进行文件关闭处理
	defer func(f *os.File) {
		err1 := f.Close()
		if err1 != nil {
			// ...
			fmt.Println("err1 = ", err1)
			panic("文件关闭失败！！！" + err1.Error())
		}
	}(file)

	// 4、创建缓冲区
	// func NewReader(rd io.Reader) *Reader
	reader := bufio.NewReader(file)

	// 5、循环读取文件内容
	for {
		// func (b *Reader) ReadString(delim byte) (string, error)
		// 按照 delim 读取数据，即遇到 delim 了就表示一次读取接收
		// 注意 '\n' 也会被读取到 line 变量里面
		line, err2 := reader.ReadString('\n') // 这里必须使用''

		// 判断是否读取结束
		// 如果读取到文件末尾了，在error里面会返回一个 io.EOF，文件末尾【end of file】
		if err2 == io.EOF {
			break // 退出循环，表示文件读取结束
		}

		// 输出每一行数据
		fmt.Print(line)

	}

	// 上面的方法适用于大文件
	// 下面的方法一次性将数据全部读取到内存，适用于小文件，且不需要读取和关闭操作，已经右工具封装了
	// func ReadFile(filename string) ([]byte, error)
	bytes, _ := ioutil.ReadFile("go.mod")
	fmt.Println("读取到的内容 %v", bytes)
	fmt.Println("读取到的内容 %s", string(bytes))

}
