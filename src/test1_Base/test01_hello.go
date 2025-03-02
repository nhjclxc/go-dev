package main

/*
里最基本的分发单位，也是工程管理中依赖关系的体现。

要生成Go可执行程序，必须建立一个名字为main的包，并且在该包中包含一个叫main()的函数（该函数是Go可执行程序的执行起点）。
Go语言的main()函数不能带参数，也不能定义返回值。

func 函数名(参数列表) {
    // 函数体
}
func 函数名(参数列表)(返回值列表) {
    // 函数体
}

*/

/*
命令行运行go代码
	【go run】直接运行，不编译，适合开发、调试
		单个文件：go run main.go
		多个文件：go run main.go other.go
	【go build】运行编译，适合生产、部署
		go build -o myapp.exe main.go，其中myapp就是编译后的输出文件
	【go run .】运行 Go 项目，此时必须要有module
*/

//导入单个包
//import "fmt"

// 同时导入多个包
import (
	"fmt"
	"html"
)

/*
标准输入输出包：fmt的了解
*/
func main() {
	println(html.EscapeString("123456"))

	fmt.Println("aaa")
	/*
		format
		该包主要用于格式化 I/O（输入/输出），提供类似 printf、print 和 scan 的功能。
		fmt 包的主要功能：
		输出（打印）
			fmt.Print()
			fmt.Println()
			fmt.Printf()

		输入（扫描）
			fmt.Scan()
			fmt.Scanln()
			fmt.Scanf()

		格式化字符串
			fmt.Sprintf() 生成格式化字符串
			fmt.Errorf() 创建格式化错误信息
	*/

	/*
		fmt 包的主要功能：
			输出（打印）
			fmt.Print()
			fmt.Println()
			fmt.Printf()
	*/

	fmt.Print("fmt.Print")
	fmt.Println("fmt.Println")
	fmt.Printf("fmt.Printf")
	fmt.Println("----------------------------")
	fmt.Println("----------------------------")

	/**
	变量声明：
		var 关键字用于声明变量，可以指定类型
			  var name string = "Alice"  // 显式声明变量，指定类型
			var age int = 25
			var height = 1.75           // 省略类型，由编译器推导为 float64
			var isStudent bool
		使用 := 进行简短声明
			name := "Alice"  // 变量名 := 值
			age := 25        // Go 自动推导 age 为 int 类型
			height := 1.75   // 推导为 float64

	*/
	var str = ""
	str2 := ""
	num := "0"

	/*
		输入（扫描）
			fmt.Scan()
			fmt.Scanln()
			fmt.Scanf()
	*/
	fmt.Scan(&str)
	fmt.Scanln(&str2)
	fmt.Scanf(num)
	fmt.Println("str = ", str)
	fmt.Println("str2 = ", str2)
	fmt.Println("num = ", num)

}

/*
Go语言优势
	可直接编译成机器码，不依赖其他库，glibc的版本有一定要求，部署就是扔一个文件上去就完成了。
	静态类型语言，但是有动态语言的感觉，静态类型的语言就是可以在编译的时候检查出来隐藏的大多数问题，动态语言的感觉就是有很多的包可以使用，写起来的效率很高。
	语言层面支持并发，这个就是Go最大的特色，天生的支持并发。Go就是基因里面支持的并发，可以充分的利用多核，很容易的使用并发。
	内置runtime，支持垃圾回收，这属于动态语言的特性之一吧，虽然目前来说GC(内存垃圾回收机制)不算完美，但是足以应付我们所能遇到的大多数情况，特别是Go1.1之后的GC。
	简单易学，Go语言的作者都有C的基因，那么Go自然而然就有了C的基因，那么Go关键字是25个，但是表达能力很强大，几乎支持大多数你在其他语言见过的特性：继承、重载、对象等。
	丰富的标准库，Go目前已经内置了大量的库，特别是网络库非常强大。
	内置强大的工具，Go语言里面内置了很多工具链，最好的应该是gofmt工具，自动化格式化代码，能够让团队review变得如此的简单，代码格式一模一样，想不一样都很困难。
	跨平台编译，如果你写的Go代码不包含cgo，那么就可以做到window系统编译linux的应用，如何做到的呢？Go引用了plan9的代码，这就是不依赖系统的信息。
	内嵌C支持，Go里面也可以直接包含C代码，利用现有的丰富的C库

*/

/*
1.2.2 标准命令概述
Go语言中包含了大量用于处理Go语言代码的命令和工具。其中，go命令就是最常用的一个，它有许多子命令。这些子命令都拥有不同的功能，如下所示。
	build：用于编译给定的代码包或Go语言源码文件及其依赖包。
	clean：用于清除执行其他go命令后遗留的目录和文件。
	doc：用于执行godoc命令以打印指定代码包。
	env：用于打印Go语言环境信息。
	fix：用于执行go tool fix命令以修正给定代码包的源码文件中包含的过时语法和代码调用。
	fmt：用于执行gofmt命令以格式化给定代码包中的源码文件。
	get：用于下载和安装给定代码包及其依赖包(提前安装git或hg)。
	list：用于显示给定代码包的信息。
	run：用于编译并运行给定的命令源码文件。
	install：编译包文件并编译整个程序。
	test：用于测试给定的代码包。
	tool：用于运行Go语言的特殊工具。
	version：用于显示当前安装的Go语言的版本信息。

*/
