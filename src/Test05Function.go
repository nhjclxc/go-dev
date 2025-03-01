package main

/*
5. 函数
*/
func main() {

	//test05Function01()

	//test05Function02()

	//test05Function03()

	//test05Function04()

	//test05Function05()

	test05Function06()

}

// 5.8 作用域
func test05Function06() {
	//5.8.3 不同作用域同名变量
	//在不同作用域可以声明同名的变量，其访问原则为：在同一个作用域内，就近原则访问最近的变量，如果此作用域没有此变量声明，则访问全局变量，如果全局变量也没有，则报错。

}

// 5.6 延迟调用defer
func test05Function05() {
	/*
		   关键字 defer ⽤于延迟一个函数或者方法（或者当前所创建的匿名函数）的执行。注意，defer语句只能出现在函数或方法的内部。
		      defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的defer应该直接跟在请求资源的语句后。

		5.6.2 多个defer执行顺序
		如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执⾏。

	*/

	test05Function051()
}

func test05Function051() {
	defer println(1)
	defer println(2)
	println(3)
	test05Function052()
	defer println(4)
	defer println(5)
}

func test05Function052() {
	defer println("test05Function052 - 1")
	println("test05Function052 - 2")
	defer println("test05Function052 - 3")
}

// 5.5 匿名函数与闭包
func test05Function04() {

}

// 5.4 函数类型
func test05Function03() {
	/*
	   5.4 函数类型
	   	在Go语言中，函数也是一种数据类型，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型。

	   	也就是说函数其实是一个变量，函数可以传递。类似python的函数
	*/
	println(useFuncType(1, 2, add))
	println(useFuncType(1, 2, minus))

	var f MyFuncType = add

	println(f(66, 99))

}

/*
定义一个函数类型
函数类型定义格式：type 函数类型名称 func (只要参数类型的参数列表) (只要返回值类型的返回值列表)
不需要函数体，函数体由具体的函数实现
这个其实就是类似与c的宏定义
*/
type MyFuncType func(int, int) int

// 定义加法函数
func add(num1 int, num2 int) (result int) {
	return num1 + num2
}

// 定义减法函数
func minus(num1 int, num2 int) (result int) {
	return num1 - num2
}

// 定义使用函数类型的函数
func useFuncType(num1 int, num2 int, funcType MyFuncType) (res int) {
	return funcType(num1, num2)
}

// 以下为递归函数的示例， 5.3 递归函数
func test05Function02() {

	// 使用递归函数实现求：1+2+3...+100的值

	println(add22(1))

}

func add22(num int) (currentNum int) {
	if num == 100 {
		return 100
	}
	return num + add22(num+1)
}

// 以下为普通函数的示例
func test05Function01() {
	/*
		Go 语言函数定义格式如下：
		func FuncName( 参数列表 ) (返回类型) {
			//函数体
			return v1, v2 //返回多个值
		}

		函数定义说明：
			func：函数由关键字 func 开始声明
			FuncName：函数名称，根据约定，函数名首字母小写即为private，大写即为public
			参数列表：函数可以有0个或多个参数，参数格式为：变量名 类型，如果有多个参数通过逗号分隔，不支持默认参数
			返回类型：
			①　上面返回值声明了两个变量名o1和o2(命名返回参数)，这个不是必须，可以只有类型没有变量名
			②　如果只有一个返回值且不声明返回值变量，那么你可以省略，包括返回值的括号
			③　如果没有返回值，那么就直接省略最后的返回信息
			④　如果有返回值， 那么必须在函数的内部添加return语句

	*/

	func01Test01()
	func01Test02(666, "啊哈", "娘子！")
	//func01Test03()
	func01Test03(1, 2, 3)
	func01Test03(1, 2, 3, 4, 5)

	var num = func01Test041()
	println(num)
	var num2, str2 = func01Test042()
	println(num2)
	println(str2)
}

func func01Test041() int {
	// 5.2.3 无参有返回值，返回一个参数
	// 有返回值的函数，必须有明确的终止语句，否则会引发编译错误。
	return 666
}

func func01Test042() (num int, str string) {
	// 5.2.3 无参有返回值，返回多个参数
	// 有返回值的函数，必须有明确的终止语句，否则会引发编译错误。
	return 888, "abc"
}
func func01Test043() (num int, str string) {
	// 5.2.3 无参有返回值，返回多个参数
	// 有返回值的函数，必须有明确的终止语句，否则会引发编译错误。
	//var num = 999 // 此块中重新声明了 'num'
	num = 999
	return num, "abc"
}
func func01Test044() (num int) {
	// 5.2.3 无参有返回值，返回多个参数
	// 有返回值的函数，必须有明确的终止语句，否则会引发编译错误。
	num = 999 // 这样也可以返回999
	return
}

func func01Test03(args ...int) {
	// 5.2.2 有参无返回值，不定参数列表，即参数个数可变
	// 不定参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要将函数定义为接受不定参数类型：
	println("func01Test03 - start")
	for _, arg := range args {
		println(arg)
	}
	println("func01Test03 - end")

	func01Test031(args...)
	func01Test031(args[1:]...) // 从索引为1的参数开始传
}
func func01Test031(args ...int) {
	// 有定参数的传递
	println("func01Test031 - start")
	for _, arg := range args {
		println(arg)
	}
	println("func01Test031 - end")
}

func func01Test02(num int, s1 string, s2 string) {
	// 5.2.2 有参无返回值，普通参数列表
	println("5.2.2 有参无返回值")
	println(num, ",", s1, ",", s2)
}

func func01Test01() {
	// 5.2.1 无参无返回值
	println("5.2.1 无参无返回值")
}
