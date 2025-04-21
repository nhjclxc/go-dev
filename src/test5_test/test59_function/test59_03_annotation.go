package main

import "fmt"

// Package mymath provides basic math functions.
//
// # Features
//
// - Add and Subtract
// - Works with float64
//
// # Example
//
// ```go
// result := mymath.Add(2, 3)
//
// fmt.Println(result) // Output: 5
// ```
//
// [你好](https://www.baidu.com)
func main() {
	/*
	   	在 Go 语言中，函数的注释通常遵循 Go 文档规范，主要有以下原则：

	      1. 使用完整的句子
	      注释通常以函数名开头，并以句号结尾。
	      2. 简洁明了
	      说明函数的作用、输入参数和返回值，必要时提供示例。
	      3. 与代码风格一致
	      //使用 // 进行单行注释，而不是 / * ... * / 形式的块注释。
	   	4. 可被 godoc 解析
	   	godoc 是 Go 语言的文档工具，它会解析与函数、方法、结构体等相关的注释，生成文档。
	*/
}

// 1. 普通函数

// Add 计算两个整数的和，并返回结果。
func Add(a, b int) int {
	return a + b
}

//2. 具有多个返回值

// Divide 执行除法运算，并返回商和余数。
// 如果除数为 0，则返回错误。
func Divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("除数不能为零")
	}
	return a / b, a % b, nil
}

// 3. 结构体方法
type User struct{ Name string }

// (User) Greet 返回用户的问候语。
func (u User) Greet() string {
	return "Hello, " + u.Name
}

// 4. 包级注释
//包级别的注释应该放在 package 语句之前。
// Package mathutils 提供数学计算相关的工具函数。
//package mathutils

//如果参数较复杂，可以单独在注释中说明

// Compute 计算并返回 a 和 b 之间的数学运算结果。
//
// 参数：
// - a: 第一个整数
// - b: 第二个整数
// - op: 操作符，可选值为 "+", "-", "*", "/"
//
// 返回：
// - 计算结果
// - 如果 op 无效，则返回错误
func Compute(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("无效的操作符: %s", op)
	}
}
