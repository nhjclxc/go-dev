package calc

func main() {

	/*
		Go语言中的测试依赖go test命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。
		go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。

	*/
}

func Add(num1, num2 int) int {
	return num1 + num2
}

func Sub(num1, num2 int) int {
	return num1 - num2
}
