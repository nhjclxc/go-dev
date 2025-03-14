package main

import "fmt"

/*
*
panic recover
*/
func main() {

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
