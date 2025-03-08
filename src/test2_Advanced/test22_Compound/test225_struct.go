package main

import (
	"fmt"
)

/*
*
7.6 结构体
结构体是一种聚合的数据类型，它是由一系列具有相同类型或不同类型的数据构成的数据集合。每个数据称为结构体的成员。

结构体定义：

	type StructName struct {
		var1 type
		var2 type
		var3 type
		...
	}
*/
func main() {

	//test2251()

	test2252()

}

// 定义学生结构体
type Student struct {
	name  string
	age   int
	addr  string
	score float32
}

/*
*
结构体对象作为函数参数
*/
func test2252() {
	println("------111---------")
	var stu1 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu1) // {zhangsan 18 北京 88.92}
	test2252_1(stu1)
	fmt.Println(stu1) // {zhangsan 18 北京 88.92}
	// 由上可知，结构体传参是传数据，拷贝一份数据的传送

	println("-------222--------")
	var stu2 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu2) // {zhangsan 18 北京 88.92}
	test2252_2(&stu2)
	fmt.Println(stu2) // {不知道是谁test2252_2 18 北京 88.92}
	// 以上将结构体对象的地址传入函数，那么原始数据可以被改变

	println("-------333--------")
	var stu3 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu3) // {zhangsan 18 北京 88.92}
	test2252_3(&stu3)
	fmt.Println(stu3) // {不知道是谁test2252_3 18 北京 88.92}
	// 以上将结构体对象的地址传入函数

	println("-------444--------")
	var stu4 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu4) // {zhangsan 18 北京 88.92}
	test2252_3(&stu4)
	fmt.Println(stu4) // {不知道是谁test2252_4 18 北京 88.92}
	// 以上将结构体对象的地址传入函数

	// 测试浅拷贝
	println("-------555--------")
	var stu5 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu5) // {zhangsan 18 北京 88.92}
	stu6 := stu5
	fmt.Println(stu6) // {zhangsan 18 北京 88.92}
	stu6.name = "stu6stu6stu6"
	fmt.Println(stu5) // {zhangsan 18 北京 88.92}
	fmt.Println(stu6) // {stu6stu6stu6 18 北京 88.92}

	println("-------666--------")
	var stu7 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu7) // {zhangsan 18 北京 88.92}
	var stu8 *Student = &stu7
	fmt.Println(stu8) // {zhangsan 18 北京 88.92}
	stu8.name = "stu8stu8stu8"
	fmt.Println(stu7)  // {stu8stu8stu8 18 北京 88.92}
	fmt.Println(*stu8) // {stu8stu8stu8 18 北京 88.92}

}

func test2252_4(pstu *Student) Student {
	fmt.Println(pstu) // {zhangsan 18 北京 88.92}
	stu := *pstu
	stu.name = "不知道是谁test2252_4"
	fmt.Println(stu) // {不知道是谁test2252_4 18 北京 88.92}
	return stu
}

func test2252_3(pstu *Student) {
	fmt.Println(pstu) // {zhangsan 18 北京 88.92}
	(*pstu).name = "不知道是谁test2252_3"
	fmt.Println(pstu) // {不知道是谁test2252_3 18 北京 88.92}
}

func test2252_2(pstu *Student) {
	fmt.Println(pstu) // {zhangsan 18 北京 88.92}
	pstu.name = "不知道是谁test2252_2"
	fmt.Println(pstu) // {不知道是谁test2252_2 18 北京 88.92}
}

func test2252_1(stu Student) {
	fmt.Println(stu) // {zhangsan 18 北京 88.92}
	stu.name = "不知道是谁test2252_1"
	fmt.Println(stu) // {不知道是谁test2252_1 18 北京 88.92}
}

/*
*
初识结构体
*/
func test2251() {

	// 创建结构体数据对象
	// 指定key创建
	var stu1 = Student{name: "zhangsan", age: 18, addr: "北京", score: 88.92}
	fmt.Println(stu1)

	// 顺序创建
	var stu2 = Student{"lisi", 20, "上海", 85.26}
	fmt.Println(stu2)

	println("-------------------")
	// 结构体对象赋值指针变量
	var pstu1 *Student = &stu1
	fmt.Println(pstu1)
	fmt.Println(*pstu1)
	fmt.Println(&pstu1)
	// 指针变量修改原始数据
	pstu1.name = "你是谁"
	fmt.Println(*pstu1)
	fmt.Println(stu1)

	var pstu2 *Student = &Student{name: "你好，世界"}
	fmt.Println(pstu2)

	// 直接创建指针变量
	var pstu3 = new(Student) // 一创建一个对象，那么这个对象的每一个变量都将被赋值一个默认值
	fmt.Println(pstu3)
	pstu3.name = "张一三" // 直接使用指针变量来操作
	fmt.Println(pstu3)
	(*pstu3).name = "张一三111" // 取出指针存储的地址之后在操作数据
	fmt.Println(pstu3)

	// 结构体的比较，类似于Java里面的equals方法
	var stu5 Student = Student{name: "张一三111"}
	fmt.Println(*pstu3 == stu5)
	var stu6 Student = Student{name: "张一三111"}
	fmt.Println(stu6 == stu5)
	stu6.age = 18
	fmt.Println(stu6 == stu5) // 比较的是结构体里面的每一个字段的值，所有都为true，那么结构体对象比较的值也是true

}
