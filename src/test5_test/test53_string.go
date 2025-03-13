package main

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	strings "strings"
	"time"
)

/*
*
go的strings相关的操作
https://learnku.com/go/wikis/61175
*/
func main() {

	//test53_01()

	//test53_02()

	//test53_03()

	//test53_04()

	//test53_05()

	//test53_06()

	//test53_07()

	//test53_08()

	//test53_09()

	//test53_10()

	//test53_11()

	//test53_12()

	test53_13()

	//hexData1 := "0241422B30303630393033313403"
	//// 将十六进制字符串转换为字节数组
	//data, _ := hex.DecodeString(hexData1) // 解析重量数据（ASCII 数字）
	//fmt.Println(data)
	//weightStr := string(data[4:10]) // 6 个字符
	//fmt.Println(strconv.Atoi(weightStr))

	// 测量 exampleFunction 函数的执行时间
	// 测量 exampleFunction 函数的执行时间，传递多个参数
	//duration := measureExecutionTimeWithArgs(exampleFunction, 42, "Go", true)
	//fmt.Printf("The function executed in: %v\n", duration)
}

/*
*
13、Compare比较方法
*/
func test53_13() {

	str1 := "a"
	str2 := "b"

	fmt.Println(strings.Compare(str1, str2)) // str1小于str2返回-1
	fmt.Println(strings.Compare(str1, str1)) // str1等于str2返回0
	fmt.Println(strings.Compare(str2, str1)) // str1大于str2返回1

}

// 计算函数执行时间的工具函数，支持传递多个参数
func measureExecutionTimeWithArgs(f interface{}, args ...interface{}) time.Duration {
	// 获取函数的反射对象
	funcValue := reflect.ValueOf(f)

	// 将参数转换为反射对象
	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	start := time.Now() // 记录开始时间

	// 调用函数
	funcValue.Call(reflectArgs)

	elapsed := time.Since(start) // 计算执行时间
	return elapsed
}

// 示例函数，接收多个参数
func exampleFunction(a int, b string, c bool) {
	// 模拟一些耗时的操作
	time.Sleep(2 * time.Second)
	fmt.Printf("Received arguments: %d, %s, %v\n", a, b, c)
}

// 12、读取内容 Read, ReadByte, ReadRune
// 在 Go 语言中，Read、ReadByte 和 ReadRune 用于从 io.Reader 接口读取数据，常用于处理文件、网络流或标准输入等数据源。

func test53_12() {

	// 1. Read（按字节块读取）
	//Read 方法从 io.Reader 读取数据到 []byte 缓冲区中，返回读取的字节数和错误信息。
	reader := strings.NewReader("Hello, Go!") // 创建一个字符串 Reader
	buf := make([]byte, 5)                    // 读取 5 个字节的缓冲区

	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("读取失败:", err)
	} else {
		fmt.Printf("读取 %d 字节: %s\n", n, string(buf)) // 输出: 读取 5 字节: Hello
	}

	// 2. ReadByte（按字节读取）
	//ReadByte 方法从 io.ByteReader 读取 单个字节

	reader2 := strings.NewReader("GoLang")

	b, err := reader2.ReadByte() // 一个一个字节往下读取
	if err != nil {
		fmt.Println("读取失败:", err)
	} else {
		fmt.Printf("读取的字节: %c\n", b) // 输出: 读取的字节: G
	}
	b2, _ := reader2.ReadByte()
	fmt.Printf("读取的字节: %c\n", b2)
	b3, _ := reader2.ReadByte()
	fmt.Printf("读取的字节: %c\n", b3)

	// 3. ReadRune（按 Unicode 字符读取）
	//ReadRune 方法用于读取 单个 Unicode 字符，返回 rune（UTF-8 编码的 Unicode 字符）。

}

// 11、转换为其他类型
// 在 Go 语言中，字符串（string）可以转换为多种数据类型，如整数（int）、浮点数（float64）、布尔值（bool）、字节切片（[]byte）等。
// Go 提供了 strconv 包和 fmt 包中的相关方法来进行类型转换。
// strconv是字符串转化类
func test53_11() {

	// 1. string 转换为 int
	// Go 使用 strconv.Atoi 或 strconv.ParseInt 将字符串转换为整数。
	// 使用 Atoi（仅支持十进制）
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", num) // 输出: 123
	}

	// 使用 ParseInt（支持不同进制）
	str2 := "a"
	num2, err := strconv.ParseInt(str2, 16, 64) // 16 代表16进制，64 代表 int64
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", num2) // 输出: 26 (因为 1a 是十六进制)
	}

	fmt.Println(strconv.ParseInt("101", 2, 64))

	//2. string 转换为 float64
	//使用 strconv.ParseFloat。
	str3 := "3.1415"
	num3, err3 := strconv.ParseFloat(str3, 64) // 64 表示 float64
	if err3 != nil {
		fmt.Println("转换失败:", err3)
	} else {
		fmt.Println("转换成功:", num3) // 输出: 3.1415
	}

	// 3. string 转换为 bool
	//使用 strconv.ParseBool。
	// "1", "t", "T", "true", "TRUE", "True" → true
	// "0", "f", "F", "false", "FALSE", "False" → false
	str4 := "true"
	b, err := strconv.ParseBool(str4)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", b) // 输出: true
	}

	// 4. string 转换为 []byte
	//字符串可以直接转换为 []byte，适用于处理二进制数据。
	str5 := "Hello"
	byteSlice := []byte(str5) // string -> []byte
	fmt.Println(byteSlice)    // 输出: [72 101 108 108 111]
	//[]byte 转 string 也可以直接转换，如 string([]byte{72, 101, 108, 108, 111})。
	fmt.Println(string([]byte{72, 101, 108, 108, 111}))
	fmt.Println(string([]byte{30, 31, 32, 31, 38, 30}))
	fmt.Println(strconv.Atoi(string([]byte{30, 31, 32, 31, 38, 30})))

	hexData1 := "0241422B30303630393033313403"
	// 将十六进制字符串转换为字节数组
	data, err := hex.DecodeString(hexData1) // 解析重量数据（ASCII 数字）
	weightStr := string(data[4:10])         // 6 个字符
	fmt.Println(strconv.Atoi(weightStr))

	//5. string 转换为 rune 切片
	//适用于处理 Unicode 字符。
	str6 := "你好"
	runeSlice := []rune(str6) // string -> []rune
	fmt.Println(runeSlice)    // 输出: [20320 22909]

	// 6. string 转换为时间（time.Time）
	//使用 time.Parse 进行时间解析。
	str7 := "2025-03-13 14:00:00"
	layout := "2006-01-02 15:04:05" // Go 使用固定格式
	t, err := time.Parse(layout, str7)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", t) // 输出: 2025-03-13 14:00:00 +0000 UTC
		fmt.Println(t.Year())
	}

}

// 10、拼接 Join
func test53_10() {

}

// 9、分割 Fields 和 Split
// func Fields(s string) []string
// strings.Fields 将字符串按空白字符（空格、换行符、制表符等）分割成子串，并且会自动忽略连续的空白字符。
func test53_09() {

	str := "Go is  awesome  "

	// 按空白字符分割
	result := strings.Fields(str)
	fmt.Println(result)    // 输出: [Go is awesome]
	fmt.Println(result[0]) //
	fmt.Println(result[1]) //
	fmt.Println(result[2]) //

	// 2. strings.Split 用法
	//strings.Split
	//作用：将字符串按指定的分隔符分割成多个子串。
	result2 := strings.Split(str, "s")
	fmt.Println(result2)    //
	fmt.Println(result2[0]) //
	fmt.Println(result2[1]) //
	fmt.Println(result2[2]) //

}

// 8、修剪 TrimSpace
// func TrimSpace(s string) string
// TrimSpace 用于去除字符串 两端 的空白字符。这里的“空白字符”指的是：空格、制表符、换行符、回车符等。
func test53_08() {

	str := "   Hello, Go!   "
	fmt.Println(str)

	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Trim("Hello, Go!", "Hel"))

}

// 7、大小写 ToLower 和 ToUpper
// ToLower 和 ToUpper 是用于处理大小写转换的两个常用函数。
// func ToLower(s string) string
// func ToUpper(s string) string
func test53_07() {

	str := "GoLang"

	// 转为小写
	fmt.Println(strings.ToLower(str)) // 输出: "golang"

	// 转为大写
	fmt.Println(strings.ToUpper(str)) // 输出: "GOLANG"

	str1 := "golang"
	str2 := "Golang"

	// 不区分大小写的比较
	if strings.ToLower(str1) == strings.ToLower(str2) {
		fmt.Println("Strings are equal (case insensitive).")
	}

}

// 6、重复 Repeat
// strings.Repeat 用于重复字符串。
func test53_06() {

	str := "hello world"
	fmt.Println(str)

	// 1. strings.Repeat 用法
	//作用：将字符串 s 重复 count 次，并返回拼接后的新字符串。
	fmt.Println(strings.Repeat("", 5))
	fmt.Println(strings.Repeat(" ", 5))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Repeat("aa", 5))
	println("--------------------------------")

	s := "abc"
	n := 4
	result := strings.Join([]string{strings.Repeat(s, n)}, ",")
	fmt.Println(strings.Repeat(s, n))
	fmt.Println([]string{strings.Repeat(s, n)})
	fmt.Println(result) // 输出: "abcabcabcabc"

	start := time.Now()

	fmt.Println("strings.Repeat 耗时:", time.Since(start))
}

// 5、统计出现次数 Count
// 使用 strings.Count 来统计子串在字符串中出现的次数。
func test53_05() {

	str := "hello world"
	fmt.Println(str)

	// 1. strings.Count 用法
	// func Count(s, substr string) int
	//作用：计算 substr 在 s 中的非重叠出现次数。
	fmt.Println(strings.Count(str, "l"))
	fmt.Println(strings.Count(str, "ll"))
	fmt.Println(strings.Count(str, "o"))
	fmt.Println(strings.Count(str, ""))
	fmt.Println(strings.Count(str, "x"))
	fmt.Println(strings.Count("aaa", "aa")) // 1
	println("--------------------------------")

}

// 4、字符串替换 Replace
// strings.Replace - 指定替换次数的字符串替换
// strings.ReplaceAll - 全部替换
// strings.Map - 按字符映射替换
// strings.NewReplacer - 多个子串替换
func test53_04() {
	str := "hello world"
	fmt.Println(str)

	// 1. strings.Replace
	// func Replace(s, old, new string, n int) string
	//作用：将 old 替换为 new，最多替换 n 次。如果 n < 0，则替换所有匹配项。
	fmt.Println(strings.Replace(str, "l", "x", 0))
	fmt.Println(strings.Replace(str, "l", "x", 1))
	fmt.Println(strings.Replace(str, "l", "x", 2))
	fmt.Println(strings.Replace(str, "l", "x", -1))
	println("--------------------------------")

	// 2. strings.ReplaceAll
	// func ReplaceAll(s, old, new string) string
	//作用：等价于 strings.Replace(s, old, new, -1)，即替换所有匹配项。
	fmt.Println(strings.Replace(str, "l", "x", -1))
	fmt.Println(strings.ReplaceAll(str, "l", "x"))
	println("--------------------------------")

	// 3. strings.Map
	// func Map(mapping func(rune) rune, s string) string
	//作用：对每个字符应用一个映射函数 func(r rune) rune，返回新的字符串。
	myMapping := func(char rune) rune {
		char += 1
		return char
	}
	fmt.Println(strings.Map(myMapping, str))
	println("--------------------------------")

	//  适用于逐字符修改，如：
	//过滤特定字符（如删除数字、特殊符号）
	//字符转换（如大小写转换）

	// mapping实现小写转大写的功能
	fmt.Println(int('A') - int('a')) // -32
	lower2UpperMapping := func(char rune) rune {
		// 判断是小写字符的时候才转化，其他不变返回
		if int('a') <= int(char) && int('a') <= int(char) {
			return char - 32
		}
		return char
	}
	fmt.Println(strings.Map(lower2UpperMapping, str))
	println("--------------------------------")

	// 4. strings.NewReplacer
	// func NewReplacer(oldnew ...string) *Replacer
	//作用：一次替换多个不同的子字符串。
	str2 := "I like apples and bananas"

	// 创建 Replacer  old和new必须一对一对出现
	replacer := strings.NewReplacer("apples", "oranges", "bananas", "grapes")
	//replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")

	// 替换
	result := replacer.Replace(str2)
	fmt.Println(result)
	// 输出: "I like oranges and grapes"

}

// 3、索引和位置 Index
// strings.Index 提供了一些函数来查找子字符串或字符的索引位置。
func test53_03() {

	str := "hello world"

	//1.1 strings.Index
	//作用：查找 substr 在 s 中的第一个出现位置，找不到返回 -1。注意区分大小写
	fmt.Println(strings.Index(str, ""))   // 0
	fmt.Println(strings.Index(str, " "))  // 5
	fmt.Println(strings.Index(str, "ll")) // 2
	fmt.Println(strings.Index(str, "h"))  // 0
	fmt.Println(strings.Index(str, "x"))  // -1
	println("--------------------------------")

	// 1.2 strings.LastIndex
	//作用：查找 substr 在 s 最后一次出现的位置，找不到返回 -1。
	fmt.Println(strings.LastIndex(str, ""))
	fmt.Println(strings.LastIndex(str, "l"))
	fmt.Println(strings.LastIndex(str, "x"))
	fmt.Println(strings.LastIndex(str, "h"))
	println("--------------------------------")

	//2. 按字符匹配的索引查找
	//2.1 strings.IndexAny
	//作用：查找 chars 任意一个字符 在 s 中的第一个位置，找不到返回 -1。
	fmt.Println(strings.IndexAny(str, "h"))
	fmt.Println(strings.IndexAny(str, "l"))
	fmt.Println(strings.IndexAny(str, "h"))
	println("--------------------------------")

	//2.2 strings.LastIndexAny
	//作用：查找 chars 任意一个字符 在 s 最后出现的位置。
	fmt.Println(strings.LastIndexAny(str, "l"))
	fmt.Println(strings.LastIndexAny(str, "h"))
	println("--------------------------------")

	//3. 按条件匹配索引查找
	//3.1 strings.IndexFunc
	//作用：返回第一个满足 func(rune) bool 的字符索引，找不到返回 -1。

}

// 2、包含关系 Contains
// strings.Contains 用于判断一个字符串是否包含另一个子字符串。以下是详细的用法和注意事项。
func test53_02() {
	str := "Hello World"
	fmt.Println("str = ", str)

	// 1. strings.Contains 的基本用法
	// strings.Contains(s, substr string) bool
	// 如果 s 中包含 substr，返回 true，否则返回 false。

	fmt.Println(strings.Contains(str, ""))
	fmt.Println(strings.Contains(str, "llo"))
	fmt.Println(strings.Contains(str, "llo "))
	fmt.Println(strings.Contains(str, "llo a"))
	fmt.Println(strings.Contains(str, "llo w"))
	fmt.Println(strings.Contains(str, "llo W"))
	println("----------------")

	// 2. strings.ContainsAny（匹配任意字符）
	// strings.ContainsAny(s, chars string) bool
	//如果 s 包含 chars 中的任意一个字符，返回 true。
	fmt.Println(strings.ContainsAny(str, ""))
	fmt.Println(strings.ContainsAny(str, "llo"))
	fmt.Println(strings.ContainsAny(str, "llo "))
	fmt.Println(strings.ContainsAny(str, "llo a"))
	fmt.Println(strings.ContainsAny(str, "llo w"))
	fmt.Println(strings.ContainsAny(str, "llo W"))
	fmt.Println(strings.ContainsAny(str, "x"))
	println("----------------")

	// 3. strings.ContainsRune（匹配单个 Unicode 字符）
	//用于检查字符串是否包含某个 rune（字符）。
	str2 := "你好啊，世界 World"
	fmt.Println("str2 = ", str2)
	fmt.Println(strings.ContainsRune(str, '啊'))
	fmt.Println(strings.ContainsRune(str, 'W'))
	println("----------------")

	// 4. strings.ContainsFunc 判断字符串中是否包含满足某个条件的字符。
	// func ContainsFunc(s string, f func(rune) bool) bool
	// 参数 s string：需要检查的字符串
	// 参数 f func(rune) bool：一个函数，定义字符的匹配条件（返回 true 代表匹配）
	// 返回：true：如果 s 中至少有一个字符满足 f 的条件。false：如果 s 中没有符合条件的字符

	// 在 Go 语言中，strings.ContainsFunc(s string, f func(rune) bool) bool
	//遍历字符串 s 的每个字符（rune），并将其传递给回调函数 f 进行判断。
	//一旦 f 返回 true，ContainsFunc 立即返回 true，否则继续检查下一个字符。

	// 是否包含数字
	containDigit := func(char rune) bool {
		if int('0') <= int(char) && int(char) <= int('9') {
			return true
		}
		return false
	}

	// 是否包含字符
	containsLetter := func(char rune) bool {
		// go的条件判断不能换行
		if (int('a') <= int(char) && int(char) <= int('z')) || int('A') <= int(char) && int(char) <= int('Z') {
			return true
		}
		return false
	}

	str3 := "asd你是谁123"
	fmt.Println("containDigit", strings.ContainsFunc(str3, containDigit))
	fmt.Println("containsLetter", strings.ContainsFunc(str3, containsLetter))

	//// 判断字符串中是否包含特殊字符
	//hasSpecial := strings.ContainsFunc(str, func(r rune) bool {
	//	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	//})
	//fmt.Println(hasSpecial)
}

// 1、HasPrefix 和 HasSuffix
// strings.HasPrefix 和 strings.HasSuffix 用于判断字符串是否具有特定的前缀或后缀。
// 必须导入strings模块 ，import (”strings“)
func test53_01() {

	str := "Hello World"
	fmt.Println("str = ", str)

	// 判断str是否有某个前缀
	// strings.HasPrefix(str, prefix)
	// 第一个参数str：是要检测的字符串
	// 第二个参数prefix：是前缀字符串
	fmt.Println(strings.HasPrefix(str, "")) // true（空字符串作为前缀，总是返回 true）
	fmt.Println(strings.HasPrefix(str, "H"))
	fmt.Println(strings.HasPrefix(str, "He"))
	fmt.Println(strings.HasPrefix(str, "Hello"))
	fmt.Println(strings.HasPrefix(str, "Hezzz"))
	println("-----------------------")

	// 判断str是否以某个后缀结尾
	//strings.HasSuffix(str, suffix)
	// 第一个参数str：是要检测的字符串
	// 第二个参数suffix：是后缀字符串
	fmt.Println(strings.HasSuffix(str, "")) // true（空字符串作为后缀，总是返回 true）
	fmt.Println(strings.HasSuffix(str, "Wo"))
	fmt.Println(strings.HasSuffix(str, "d"))
	fmt.Println(strings.HasSuffix(str, "ld")) // HasSuffix要从后往前找
	fmt.Println(strings.HasSuffix(str, "World"))

	// 使用场景1，过滤指定文件
	files := []string{"config.yaml", "config.json", "readme.md", "config.toml"}
	// 过滤出配置文件
	for _, file := range files {
		if strings.HasPrefix(file, "config") {
			fmt.Println("配置文件：", file)
		}
	}

	// 使用场景2：检查URL是否具有特定前缀
	urls := []string{"example.com", "google.com", "example.org", "github.io"}
	// 过滤出 .com结尾的所有网站
	for _, url := range urls {
		if strings.HasSuffix(url, ".com") {
			fmt.Println(".com结尾的网站：", url)
		}
	}

}
