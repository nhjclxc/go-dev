package main

import (
	"fmt"
	"strings"
)

// https://denganliang.github.io/the-way-to-go_ZH_CN/18.1.html
// 字符串
func main01() {

	// 修改一个字符串里面的某个字符

	str := "Hello!!!"
	fmt.Println(str)

	// 转化为字节数组
	bytes := []byte(str)
	fmt.Println(bytes)

	// 修改
	bytes[1] = 'g'

	fmt.Println(bytes)
	fmt.Println(str)

	// 在go中字符串是不可变的（与Java类似）
	// 因此，要将修改后的字符串放回去原来那个变量里面
	str = string(bytes)
	fmt.Println(str)

	// 获取字符串的子串数据
	fmt.Println(str[:])
	fmt.Println(str[0:])
	fmt.Println(str[2:])
	fmt.Println(str[:3])
	fmt.Println(str[1:3])

	// 使用 fori 或 for ... range 遍历一个字符串
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]), ", ")
	}
	fmt.Println()

	for i, ch := range str {
		fmt.Println(i, ", ", string(ch))
	}
	fmt.Println()

	//
	fmt.Println(len(str))         /// 字节数
	fmt.Println(len([]rune(str))) // 字符数

	// 字符串连接
	str2 := "World"
	str += str2
	fmt.Println("str = ", str)

	fmt.Println(strings.Join([]string{"hello", "world", "go"}, "--"))

}
