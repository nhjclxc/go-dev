package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	/*
	import "regexp"

	func Compile(expr string) (*Regexp, error)       // 编译正则
	func MustCompile(expr string) *Regexp            // panic 版
	func (r *Regexp) MatchString(s string) bool      // 是否匹配
	func (r *Regexp) FindString(s string) string     // 第一个匹配字符串
	func (r *Regexp) FindAllString(s string, n int) []string // 所有匹配结果
	func (r *Regexp) ReplaceAllString(src, repl string) string // 字符串替换


	✅ 三、常见正则表达式语法
	| 表达式      | 含义              |          |
	| -------- | --------------- | -------- |
	| `.`      | 任意单字符（不包括换行）    |          |
	| `^abc`   | 以 abc 开头        |          |
	| `abc$`   | 以 abc 结尾        |          |
	| `[a-z]`  | 匹配小写字母 a 到 z    |          |
	| `[^a-z]` | 非小写字母           |          |
	| `\d`     | 数字（相当于 `[0-9]`） |          |
	| `\w`     | 单词字符（字母/数字/下划线） |          |
	| \`a      | b\`             | 匹配 a 或 b |
	| `(ab)+`  | 分组，匹配 1 次或多次 ab |          |

	*/

	{
		//1. 检查是否匹配
		r := regexp.MustCompile(`^a\d{3}$`) // 开头a+3位数字
		fmt.Println(r.MatchString("a123"))  // true
		fmt.Println(r.MatchString("b123"))  // false
		fmt.Println()
	}

	{

		// 2. 查找匹配（Find）
		r := regexp.MustCompile(`[a-z]+`)
		str := "123abc456def"
		fmt.Println(r.FindString(str))           // abc
		fmt.Println(r.FindAllString(str, -1))    // [abc def]
		fmt.Println()
	}



	//3. 查找位置（FindStringIndex）
	{
		r := regexp.MustCompile(`abc`)
		idx := r.FindStringIndex("xyzabcdef")
		fmt.Println(idx) // [3 6] 表示 abc 开始和结束位置
		fmt.Println()
	}

	{
		//4. 提取子组（FindStringSubmatch）
		r := regexp.MustCompile(`(\d+)-(\d+)`)
		match := r.FindStringSubmatch("订单号：123-456")
		fmt.Println(match)       // [123-456 123 456]
		fmt.Println()
	}

	{

		//5. 替换内容（ReplaceAllString）
		r := regexp.MustCompile(`\d`)
		str := "a1b2c3"
		newStr := r.ReplaceAllString(str, "*")
		fmt.Println(newStr) // a*b*c*
		fmt.Println()
	}

	{
		//6. 用函数替换（ReplaceAllStringFunc）
		r := regexp.MustCompile(`[a-z]+`)
		text := "Go i123s cool"

		result := r.ReplaceAllStringFunc(text, strings.ToUpper)
		fmt.Println(result) // GO I123S COOL
	}



}
