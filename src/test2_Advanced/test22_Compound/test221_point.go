package main

/*
7.2 指针
指针是一个代表着某个内存地址的值。这个内存地址往往是在内存中存储的另一个变量的值的起始位置。
Go语言对指针的支持介于Java语言和C/C++语言之间，它既没有想Java语言那样取消了代码对指针的直接操作的能力，
也避免了C/C++语言中由于对指针的滥用而造成的安全和可靠性问题。
*/
func main() {

	//test01()

	//test02()

	test03()

}

func test03() {
	/*
		   7.2.3 指针做函数参数

			编写一个swap交换函数

	*/

	var a, b int = 10, 20
	println("a = ", a, " b = ", b)
	sawp(&a, &b)
	println("a = ", a, " b = ", b)
	println("----------------")
	println("a = ", a, " b = ", b)
	sawp2(&a, &b)
	println("a = ", a, " b = ", b)
}

func sawp2(aaa *int, bbb *int) {
	// 以下交换没有改变aaa和bbb所存储的地址，而是改变了所存储地址指向变量存的值，即修改了实参a和b的值（完成了变量值交换的功能）
	// *aaa和*bbb的操作是取出他们所存的值，在赋一个新的值上去，对指针变量指向的那个变量存的值进行了交换修改
	println("aaa = ", aaa) // 0xc00005ff38
	println("bbb = ", bbb) // 0xc00005ff30
	*aaa, *bbb = *bbb, *aaa
	println("aaa = ", aaa) // 0xc00005ff38
	println("bbb = ", bbb) // 0xc00005ff30
}

func sawp(aa *int, bb *int) {
	// 以下只是对aa和bb所存储的地址进行了交换，并没有改变实参a和b所存储的值
	println("aa = ", aa) // 0xc00005ff38
	println("bb = ", bb) // 0xc00005ff30
	aa, bb = bb, aa
	println("aa = ", aa) // 0xc00005ff30
	println("bb = ", bb) // 0xc00005ff38

}

func test02() {
	/*
	   7.2.2 new函数，【【【 new(T)为指定类型的指针变量开辟空间，返回的变量的数据类型是：*T类型 】】】
	   表达式new(T)将创建一个T类型的匿名变量，所做的是为T类型的新值分配并清零一块内存空间，然后将这块内存空间的地址作为结果返回，
	   	而这个结果就是指向这个新的T类型值的指针值，返回的指针类型为*T。

	   	【【【 使用new函数创建的指针变量go会自动管理，开发者无需进行内存管理 】】】
	   	我们只需使用new()函数，无需担心其内存的生命周期或怎样将其删除，因为Go语言的内存管理系统会帮我们打理一切。
	*/

	// 为int类型的指针变量开辟空间
	var p *int = new(int)
	*p = 666

	println("p = ", p)
	println("&p = ", &p)
	println("*p = ", *p)

	var a int = 666

	println("a = ", a)
	println("&a = ", &a)
}

func test01() {
	/*
	   Go语言虽然保留了指针，但与其它编程语言不同的是：
	   		默认值 nil，没有 NULL 常量
	   		操作符 "&" 取变量地址， "*" 通过指针访问目标对象
	   		不支持指针运算，不支持 "->" 运算符，直接⽤ "." 访问目标成员
	*/

	// 创建一个变量
	var a int = 10

	// 创建一个指针变量
	var p *int

	// 指针变量赋值
	p = &a

	println("a = ", a)   // 10，就是存的值10
	println("&a = ", &a) // 0xc00005ff28，表示取变量a的地址
	println("p = ", p)   // 0xc00005ff28，p是一个指针，直接打印指针那么输出的就是他保存的值，即指针指向的那个地址
	println("&p = ", &p) // 0xc00005ff38，表示取指针变量p的地址
	println("*p = ", *p) // 10，表示去除指针变量p所存地址的值，解引用

	println("---------------")
	a = 666
	println("a = ", a)   // 666
	println("p = ", p)   // 0xc00005ff28
	println("&p = ", &p) // 0xc00005ff38
	println("*p = ", *p) // 666

	println("---------------")
	*p = 666
	println("a = ", a)   // 666
	println("p = ", p)   // 0xc00005ff28
	println("&p = ", &p) // 0xc00005ff38
	println("*p = ", *p) // 666
}
