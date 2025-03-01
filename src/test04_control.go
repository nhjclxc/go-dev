package main

import "fmt"

/*
	4. 流程控制
*/

func main() {

	//test04_control_01()
	//test04_control_02()
	//test04_control_03()
	test04_control_04()

}

/*
4.3 跳转语句
*/
func test04_control_04() {
	//在循环里面有两个关键操作break和continue，break操作是跳出当前循环，continue是跳过本次循环。

	var i int
	for i = 0; i < 10; i++ {
		if i%3 == 0 {
			print("是一个可被3整除的数，")
		}
		println(i)
	}

}

/*
4.2 循环语句
*/
func test04_control_03() {

	// for循环与Java的区别就在于少了一对()，但是循环变量必须提前定义号
	var i int
	for i = 0; i < 10; i++ {
		println(i)
	}

	// 关键字 range 会返回两个值，第一个返回值是元素的数组下标，第二个返回值是元素的值：

	str := "string"
	println("-----------------------")
	for i, s := range str { // 第一个参数是索引，第二个参数是对应的assic码值
		println(i, " - ", s)
	}
	println("--------111111--------")
	for _, s := range str {
		println(s, ", ", string(s))
	}
	println("--------222222---------")
	for i := range str { // 只接收一个参数的时候默认是索引
		println(i)
	}
	println("--------333333----------")

	// 死循环
	var num = 0
	for {
		num++
		println(num)

		//if num == 100 {
		//	break
		//}
	}

}

/*
switch
*/
func test04_control_02() {
	/*
	   4.1.2 switch语句
	   	Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码：
	   	go的switch里面自带break
	*/
	var myType int = 66

	switch myType {
	case 1:
		println("这是类型1")
	case 2:
		println("这是类型2")
	default:
		println("default类型")
	}

	// 可以使用任何类型或表达式来作为switch的条件判断

	var sroce int = 55
	switch {
	case sroce < 60:
		println("不及格")
		println("这是不及格")
	case sroce >= 60 && sroce < 70:
		println("及格")
		println("2及格")
	default:
		println("不知道你多少分")
	}

}

/*
if
*/
func test04_control_01() {
	var a int = 3
	if a == 3 { //条件表达式没有括号
		fmt.Println("a==3")
	}

	//支持一个初始化表达式, 初始化字句和条件表达式直接需要用分号分隔
	if b := 3; b == 3 {
		fmt.Println("b==3")
	}

	if c := 3; c == 5 {
		fmt.Println("c==3")
	} else { // 左大括号必须和条件语句或else在同一行
		fmt.Println("c != 3")
	}

	num := 666
	//num++
	num--
	if num > 666 {
		println("num > 666")
	} else if num < 666 {
		println("num < 666")
	} else {
		println("num == 666")
	}

}
