package main

import (
	"fmt"
	"os"
)

// 1、文件的打开与关闭操作

func main() {

	// 使用绝对路径和相对路径打开都可以

	// 1、打开文件
	// 通常file被叫为file对象、file指针或是file句柄
	file, err := os.Open("D:\\code\\go\\go-dev\\.gitignore")
	//f2, err2 := os.Open(".gitignore")

	// 2、判断文件是否打开成功
	if err == nil {
		panic("文件打开失败！！！" + err.Error())
	}

	fmt.Println(file)

	// 3、打开文件成功之后，建议  立马  搭配defer关键字来调用文件关闭方法，一面造成文件关闭之后的一些问题
	// func (f *File) Close() error
	// 基于以上Close方法的定义，调用Close方法的时候，会返回一个文件关闭错误的 error
	// 因此以下，在关闭文件的时候要使用一个闭包来处理关闭文件可能遇到的问题，
	// 注意：不能直接使用：defer f.Close()。这样的问题就是无法处理关闭失败时发送的错误
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

			// 在这里进行处理文件关闭异常！！！
			// ...

			panic("文件关闭异常，" + err.Error())
		}
	}(file)

	fmt.Println(file)

}
