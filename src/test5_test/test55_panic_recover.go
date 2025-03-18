package main

import (
	"errors"
	"fmt"
)

/*
*
panic recover
*/
func main() {

	//test55_01()

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("mian 出现错误，r = ", r)
	//	}
	//}()
	//test55_02()
	//fmt.Println("我被执行了吗 4")

	test55_03()

	/*
		代码解析：
		defer 关键字
			defer 语句用于延迟执行代码，通常用于资源释放（如关闭文件、解锁互斥锁等）。
			在 panic 发生后，defer 仍然会执行，这样可以利用 recover 捕获 panic。

		panic 关键字
			panic("发生了一个错误") 表示程序进入不可恢复的错误状态，并终止当前 Goroutine 的正常执行。
			panic 之后，程序会向上回溯 defer 语句，直到 recover 捕获或程序崩溃。

		recover 关键字
			recover() 仅在 defer 函数中使用，作用是捕获 panic 并防止程序崩溃。
			recover() 返回 panic 传递的错误信息，如果没有 panic，则返回 nil。
	*/
	//关键点总结：
	//panic 会导致程序异常中止，除非在 defer 语句中使用 recover 进行捕获。
	//recover 只能在 defer 内部调用，否则无法生效。
	//recover 可以恢复 panic 让程序继续执行，避免整个进程崩溃。
	//这样，我们就可以使用 panic 处理严重错误，并通过 recover 进行适当的恢复，以确保程序的健壮性。 🚀
}

// 自定义错误
func test55_03() {
	// 在go中使用errors.New和panic内置函数来实现自定义错误
	// 1、errors.New(错误说明)，这个函数会返回一个 error 类型的值，错误一个错误
	// 2、panic内置函数接收一个 interface{} 类型的值作为一个参数，可以接收 error 类型的变量输出的错误信息，并退出程序

	file := "init.config111"

	fmt.Println("test55_03 1")

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("读取配置文件异常！！！")
	//	}
	//}()

	fmt.Println("test55_03 2")
	readConfig(file)
	fmt.Println("test55_03 6")

}

func readConfig(file string) bool {
	if "init.config" != file {
		// 不是指定的文件，那么就是说明此文件名不是预定义的配置文件，那么此时应该手动抛出一个异常，
		//使用 errors.New 来进行一个 error 异常创建
		err := errors.New(`not a default config file !!!`)
		// 使用panic来把这个异常向外抛出
		fmt.Println("test55_03 3")
		panic(err)
		fmt.Println("test55_03 4")
	}
	fmt.Println("test55_03 5")
	// 假设进行读取操作...
	return true

}

func test55_02() {
	num1 := 10
	num2 := 0

	fmt.Println("我被执行了吗 1")
	// 在有可能发生异常的语句前面，写一个defer的异常处理，defer会在函数的最后被执行，且这个闭包函数一定会被执行
	// 在闭包里面捕获recover的异常，已对其进行处理
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("出现错误，r = ", r)
		}
	}()

	fmt.Println("我被执行了吗 2")
	num3 := num1 / num2 // panic: runtime error: integer divide by zero
	// 在本函数内，发生panic的地方往下的代码不会被执行，此时这个函数就会推出了，
	// 但是如果写了defer的话就会转入defer语句标识的闭包执行，
	// 如果改闭包内调用了recover()函数，说明该painc被recover()函数捕获了，这个painc就不会往上传递了
	// 如果不使用recover()函数来捕获painc那么这个painc就会一层一层的往上传递，直至被recover()捕获或者程序奔溃导致退出
	fmt.Println("num3 = %d", num3)
	fmt.Println("我被执行了吗 3")

}

func test55_01() {

	fmt.Println("程序开始")

	// 使用 defer 机制调用捕获 panic 的函数
	defer func() {
		// 使用 recover 来捕获一个异常，类似一个catch
		// 如果异常不等于nil，即发生了异常
		// 那么执行相应的数据处理
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
		}
	}()

	fmt.Println("即将发生 panic...")
	panic("发生了一个错误") // 类似于Java抛出一个异常,

	// 这行代码不会被执行，因为 panic 之后程序会终止，除非被 recover
	fmt.Println("程序结束")
}
