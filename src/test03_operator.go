package main

/*
*3. 运算符
 */
func main() {
	/*
		   	++	后自增，没有前自增	a=0; a++	a=1
		       --	后自减，没有前自减	a=2; a--	a=1
		   	注意没有++a或--a，只有a++或a--

		在 Go 语言中，a++ 是一个独立的语句，而不是一个表达式。这意味着它不能与其他操作符（如赋值操作符 =）一起使用。因此，a = a++ 这种写法是非法的，会导致编译错误。

	*/
	var a int = 666
	println(a)
	//a = a++  // syntax error: unexpected ++ at end of statement，++或--必须单独占一行
	a++
	println(a)
	a--
	println(a)
	a += 1
	println(a)

	/*
		3.6 其他运算符
		运算符	术语	示例	说明
			&	取地址运算符	&a	变量a的地址
			*	取值运算符	*a	指针变量a所指向内存的值
	*/

	println(&a) // 0xc00005ff30
	//println(*a) // invalid operation: cannot indirect a (variable of type int)
	println(*&a)

}
