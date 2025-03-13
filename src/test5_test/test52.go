package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
*
函数练习
*/
func main() {

	//test52_1()

	//test52_2()

	//test52_3()

	//test52_4()

	//test52_5()

	//test52_6()

	//test52_7()

	//test52_8()

	//test52_9()

	//test52_10()

	//test52_11()

	//test52_12()

	//test52_13()

	/*
	   	在 Go 语言中，if 语句支持 短变量声明（Short Variable Declaration），也就是 := 语法，并且可以在 if 语句的条件部分声明局部变量。这种写法的具体语法如下：

	      go
	      复制
	      编辑
	      if 初始化语句; 条件判断 {
	          // 代码块
	      }
	      其中，初始化语句 和 条件判断 之间用 ; 进行分隔。这个 ; 允许在 if 语句内执行某些操作并声明变量，同时在 if 语句的条件判断部分继续使用这些变量。

	*/
	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Printf("%T, %v", num, num)
	}

}

/*
*
函数返回一个函数

 1. 编写一个函数返回另一个函数，返回的函数的作用是对一个整数+2。函数的名称叫做plusTwo。然后可以像下面这样使用：
    p := plusTwo()
    fmt.Printf("%v\n", p(2))
    应该打印4。
 2. 使1中的函数更加通用化，创建一个plusX(x)函数，返回一个函数用于对整数加上x。
*/
func test52_13() {

	p := plusX(2)
	fmt.Println(p(2))
	fmt.Println(plusX(5)(2))
	fmt.Println(plusX(6)(2))
	fmt.Println(plusX(9)(2))

}

func plusX(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
}

/*
*
map()函数是一个接受一个函数和一个列表作为参数的函数。
函数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回。因此：
map(f(), (a1,a2,a3,...,an)) = (f(a1),f(a2),f(a3),...,f(an))
 1. 编写Go中的简单的map()函数。它能工作于操作整数的函数就可以了。
 2. 扩展代码使其工作于字符串列表。
*/
func test52_12() {

	// 倍数
	multiple := func(num int) int {
		return num * 2
	}
	// 平方
	square := func(num int) int {
		return num * num
	}

	var arr []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)
	fmt.Println(myMap(multiple, arr))
	fmt.Println(myMap(square, arr))

}

func myMap(myFunc func(num int) int, arr []int) []int {
	sli := make([]int, 0, len(arr)) // 预分配容量，避免多次扩容
	for _, val := range arr {
		sli = append(sli, myFunc(val))
	}
	return sli
}

/*
*
Q10. (1) 斐波那契
 1. 斐波那契数列以：11235813开始。或者用数学形式表达：1=1; 2=1; n = (n−1) + (n−2)。
    编写一个接受int值的函数，并给出这个值得到的斐波那契数列。
*/
func test52_11() {
	fmt.Println(fib(6))
}

func fib(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

/*
实现栈的压栈与弹栈
 1. 创建一个固定大小保存整数的栈。它无须超出限制的增长。栈应当是后进先出（LIFO）的。
    定义push函数——将数据放入栈，和pop函数——从栈中取得内容。
 2. 更进一步。编写一个String方法将栈转化为字符串形式的表达。
    可以这样的方式打印整个栈：fmt.Printf("My stack %v\n", stack)栈可以被输出成这样的形式：[0:m] [1:l] [2:k]
*/
func test52_10() {

	var stack []int = []int{1, 2, 3, 4, 5}[0:5]
	fmt.Println(stack)

	push(&stack, 6)
	//fmt.Println(stack)
	printlnStack(stack)
	pop(&stack)
	//fmt.Println(stack)
	printlnStack(stack)
	push(&stack, 666)
	//fmt.Println(stack)
	printlnStack(stack)

}

// 输出栈的每一个元素
func printlnStack(stack []int) {
	for index, element := range stack {
		fmt.Printf("[%d : %d] \n", index, element)
	}
	fmt.Println("----------------------------")
}

// 将stack栈顶的一个元素弹出
func pop(stack *[]int) (flag bool, element int) {

	//fmt.Println(stack)
	//fmt.Println(*stack)
	//fmt.Println(&stack)
	//fmt.Println(len(*stack) - 1)
	//fmt.Println((*stack)[len(*stack)-1])
	if len(*stack) == 0 {
		return false, 0
	}
	top := len(*stack) - 1
	element = (*stack)[top]
	ints := (*stack)[0:top]
	*stack = ints
	return true, element
}

// 向栈stack里面压入一个元素element
// 参数1 stack：是栈对象
// 参数2 element：是要入栈的元素
// 返回参数1 flag：是否成功
// 返回参数2 err：出错是的消息
func push(stack *[]int, element int) (flag bool, err string) {

	ints := append(*stack, element)
	*stack = ints
	return true, ""
}

/*
*
Q6. (0) 整数顺序
 1. 编写函数，返回其（两个）参数正确的（自然）数字顺序：
    f(7,2) → 2,7
    f(2,7) → 2,7
*/
func test52_9() {

	println(sort(7, 2))
	println(sort(2, 7))
}

func sort(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

/*
*
Q5. (0) 平均值
1. 编写一个函数用于计算一个float64类型的slice的平均值。
*/
func test52_8() {
	var arr []float64 = []float64{1.1111111111111, 1.222222222222222222, 1.333333333333,
		1.5555555555, 1.666666666, 1.8888888}
	fmt.Println(arr)

	println("avg = ", calFloat64Avg(arr[0:0]))
	println("avg = ", calFloat64Avg(arr[0:1]))
	println("avg = ", calFloat64Avg(arr[0:6]))

}

func calFloat64Avg(sli []float64) float64 {
	if len(sli) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, value := range sli {
		sum += value
	}
	return sum / float64(len(sli))
}

/*
*
恐慌（Panic）和恢复（Recover）
Go 没有像Java那样的异常机制，例如你无法像在Java中那样抛出一个异常。
作为替代，它使用了恐慌和恢复（panic-and-recover）机制。
一定要记得，这应当作为最后的手段被使用，你的代码中应当没有，或者很少的令人恐慌的东西。
*/
func test52_7() {
	//Panic
	//是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。
	//当函数F调用panic，函数F的执行被中断，并且F中的延迟函数会正常执行，然后F返回到调用它的地方。
	//在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到程序崩溃时的所有goroutine返回。
	//恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。

	//Recover
	//是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。
	//recover仅在延迟函数中有效。
	//在正常的执行过程中，调用recover会返回nil并且没有其他任何效果。
	//如果前的goroutine 陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。

}

/*
*
函数变量做回调
*/
func test52_6() {

	myFunc := func(args ...interface{}) {
		for i, arg := range args {
			fmt.Printf("第 %d 个数据类型是 arg = %d \n", i, arg)
		}
	}

	test52_61(666, myFunc)

}

func test52_61(i int, myFunc func(args ...interface{})) {
	fmt.Printf(" i = ", i)
	// 以下假设对i进行数据处理
	i++
	i++
	i++

	// 执行函数的回调
	myFunc("执行回调", i, "azxcv", 1.23)
}

/*
*
函数变参
*/
func test52_5() {

	// 变参只能写在最后面
	test52_51()
	test52_51(1)
	test52_51(1, 2)
	test52_51(1, 2, 3)

	test52_52(123, 563, 789)
	test52_52(123, "aaaaa")
	test52_52(123, "aaaaa", true, false)
	test52_52(123, "aaaaa", true, 3.14)
	println(math.MinInt8)
	println(math.MaxInt8)
	println(math.MinInt16)
	println(math.MaxInt16)
	println(math.MinInt32)
	println(math.MaxInt32)
	println(math.MinInt64)
	println(math.MinInt64)

}

// 以下是适用多种数据类型
// 使用 interface{} 作为参数类型
func test52_52(args ...interface{}) {
	for i, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Printf("第 %d 个数据类型是 %s 类型，arg = %d \n", i, v, arg)
		case string:
			fmt.Printf("第 %d 个数据类型是 %s 类型，arg = %s \n", i, v, arg)
		case bool:
			fmt.Printf("第 %d 个数据类型是 %s 类型，arg = %t \n", i, v, arg)
		case float32:
			fmt.Printf("第 %d 个数据类型是 %s 类型，arg = %f \n", i, v, arg)
		default:
			fmt.Printf("未知的类型", v)
		}
	}
	println("-------------------------")
}

// 以下只能传递int数据类型的变参
func test52_51(args ...int) {
	for i, arg := range args {
		fmt.Printf("i = %d, arg = %d\n", i, arg)
	}
	println("--------")
}

/*
*
匿名函数
*/
func test52_4() {

	// 直接写一个匿名函数
	func(num int) {
		println("匿名函数，", num)
	}(666)

	// 在go中，函数也是一个变量，可以将一个函数直接赋值给一个变量
	// 在随后的作用域内可以使用这个函数变量来调用该函数
	func1 := func(num int, str string) {
		fmt.Printf("num = %d, str = %s \n", num, str)
	}

	func1(666, "who")
	func1(999, "nihao")

}

/*
*
延迟代码 defer

	在defer后指定的函数会在函数退出前调用。
*/
func test52_3() {

	println("a")
	defer test52_31(1)
	defer test52_32(2)
	defer test52_32(3)
	println("b")
	//a
	//b
	//e 3
	//e 2
	//c 1
	//d 1
	//e 1
	// 由以上输出结果可知，当go遇到defer语句时，会用一个栈来保存defer标识的语句。
	// 当本函数的所有语句都执行完毕之后，才会去执行这个栈里面保存的语句，使用栈的先进后出特性依此执行
}

func test52_31(num int) {

	println("c", num)
	defer test52_32(num)
	println("d", num)
}

func test52_32(num int) {
	println("e", num)
}

/*
*
函数多返回值测试
*/
func test52_2() {

	println(test52_21(8, "who"))

}

func test52_21(num int, name string) (name22 string, number int) {

	println("name22", name22)
	println("number", number)
	// 以下不使用 := 的理由是在函数返回列表里面已经定义了，go会初始化对应的返回值列表
	number = num
	name22 = name
	aaa := strconv.Itoa(number) + name
	println("aaa", aaa)

	//return name22, number
	//return aaa, number

	// 如果在函数体内定义的变量和函数定义的返回值变量一样的话，那么返回的时候可以省略return后面的变量名称
	return

}

var global int = 666

/*
函数定义
*/
func test52_1() {
	/*
		func (p mytype) funcName(q int, x string) (r,s int) { return 0,0 }

			func：用于定义一个函数
			(p mytype)：用于将函数绑定到特殊的类型上，这个叫做接收者，被称为method
			funcName：函数名称
			(q int, x string)：参数列表
			(r,s int)：返回值列表
			{ return 0,0 }：大括号内部定义代码体
	*/

	/**
	作用域：
		定义在函数外面的为全部变量，
		定义在大括号内部的为局部变量，从哪里开始定义，那么这个变量的作用域就从哪里开始，且到最近的一个大括号结束其定义域
	*/

	println("global = ", global)
	{
		//println("localInt1 = ", localInt1) // 未解析的引用 'localInt1'
		var localInt1 = 888
		println("global = ", global)
		println("localInt1 = ", localInt1)

		{
			println("localInt1 = ", localInt1)
			//println("localInt2 = ", localInt2) // 未解析的引用 'localInt2'
			var localInt1 = 999
			var localInt2 = 123
			println("global = ", global)
			println("localInt1 = ", localInt1)
			println("localInt2 = ", localInt2)

		}

	}

}
