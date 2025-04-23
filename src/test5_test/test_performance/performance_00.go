package main

import (
	"flag"
	"fmt"
)

// https://denganliang.github.io/the-way-to-go_ZH_CN/18.1.html
// 使用flag包解析命令行参数
func main00() {
	// go run performance_02.go -name=张三 -age=25 -v other1 other2
	// 定义命令行参数
	name := flag.String("name", "默认名称", "用户的名字")
	age := flag.Int("age", 18, "用户的年龄")
	verbose := flag.Bool("v", false, "是否输出详细信息")

	// 解析命令行参数
	flag.Parse()

	// 使用解析后的参数
	fmt.Printf("Name: %s\n", *name)
	fmt.Printf("Age: %d\n", *age)
	fmt.Printf("Verbose: %v\n", *verbose)

	// 还可以获取非标志参数（也就是剩下的参数）
	fmt.Println("Other args:", flag.Args())
}
