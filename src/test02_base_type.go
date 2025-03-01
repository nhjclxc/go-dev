package main

import (
	"fmt"
	"reflect"
)

/*
*
go语言基础类型

Go语言中类似if和switch的关键字有25个(均为小写)。关键字不能用于自定义名字，只能在特定语法结构中使用。

	break       default        func         interface        select
	case        defer           go            map               struct
	chan        else            goto           package          switch
	const        fallthrough     if             range            type
	continue     for              import        return            var
*/
func main() {

	//test021_variable()

	test022_const()

}

/*
*
2.2 变量
*/
func test021_variable() {

	/*
			2.2.1 变量声明
				1. var 关键字声明
					特点：✅ 可以显式声明变量类型 ✅ 可以在函数外部（全局作用域）使用 ✅ 可以不立即赋值，默认为零值 ✅ 可以用于 const、func 等全局变量
					var v1 int
					var v2 int
					var vv = 666
				2. := 简短声明
					特点：✅ 必须在函数内部使用（不能用于全局变量） ✅ 不能显式指定变量类型（由编译器自动推导） ✅ 声明并初始化变量 ✅ 简洁、适合局部变量
					vvv := 10

					//一次定义多个变量
					var v3, v4 int

					var (
						v5 int
						v6 int
					)

		var vs := 何时使用？
		全局变量 → 必须使用 var
		局部变量（函数内部） → 推荐使用 :=，代码更简洁
		需要指定类型 → 使用 var
		可能需要后续初始化 → 使用 var
		同时声明多个变量 → var 和 := 都可以

	*/

	var vv int
	var vvv = 666
	fmt.Println(vv)
	fmt.Println(vvv)

	var v1 int = 111                               // 方式1，显示声明变量类型
	var v2 = 222                                   // 方式2，编译器自动推导出v2的类型
	v3 := 333                                      // 方式3，编译器自动推导出v3的类型
	fmt.Println("v3 type is ", reflect.TypeOf(v1)) //v3 type is  int
	fmt.Println("v3 type is ", reflect.TypeOf(v2)) //v3 type is  int
	fmt.Println("v3 type is ", reflect.TypeOf(v3)) //v3 type is  int

	//出现在 := 左侧的变量不应该是已经被声明过，:=定义时必须初始化
	//var v4 int
	//v4 := 2 //err

	// 变量赋值
	v1 = 666

	fmt.Println(v1, v2, v3)
	v1, v2, v3 = v2, v3, v1 // 多重赋值，直接实现两个变量的值交换
	fmt.Println(v1, v2, v3)

	/*
		2.2.4 匿名变量
		_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃：

		对于某个变量不想使用的适合就可以使用下划线_来接收，这样编译器就不会报错，如果不使用下划线的话，编译器会检查到这个变量没使用，这个时候就报错了

	*/
	_, str := test()
	fmt.Println(str)

	_ = 1

}
func test() (int, string) {
	return 250, "saaaab"
}

/*
*
2.2 变量
*/
func test022_const() {

	println("test022_const")

	/*
		2.3.1 字面常量(常量值)
		所谓字面常量（literal），是指程序中硬编码的常量，如：
			123
			3.1415  // 浮点类型的常量
			3.2+12i // 复数类型的常量
			true  // 布尔类型的常量
			"foo" // 字符串常量
	*/

	//var num int = 123
	//var pi float32 = 3.1415

	/*
		2.3.2 常量定义 【【【也是普通变量声明，直接把 var 变成 const 即可】】】
		    const Pi float64 = 3.14
		    const zero = 0.0 // 浮点常量, 自动推导类型

		    const ( // 常量组
		        size int64 = 1024
		        eof        = -1 // 整型常量, 自动推导类型
		    )
		    const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
		    const a, b, c = 3, 4, "foo"
		    // a = 3, b = 4, c = "foo"    //err, 常量不能修改

	*/

	const PI float32 = 3.1415926
	const PI2 float32 = 31415926
	println(PI)
	println(PI2)

	/*
			2.3.3 iota枚举
			常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。
		在 Go 语言中，iota 并不是某个单词的缩写，而是一个特殊的常量生成器。它用于在常量声明中生成一系列递增的值，通常用于枚举类型。

			在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。

			    const (
			        x = iota // x == 0
			        y = iota // y == 1
			        z = iota // z == 2
			        w  // 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
			    )

			    const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

			    const (
			        h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
			    )

			    const (
			        a       = iota //a=0
			        b       = "B"
			        c       = iota             //c=2
			        d, e, f = iota, iota, iota //d=3,e=3,f=3
			        g       = iota             //g = 4
			    )

			    const (
			        x1 = iota * 10 // x1 == 0
			        y1 = iota * 10 // y1 == 10
			        z1 = iota * 10 // z1 == 20
			    )

	*/

	/*
		Go语言内置以下这些基础类型：
		类型	名称	长度	零值	说明
		bool	布尔类型	1	false	其值不为真即为家，不可以用数字代表true或false
		byte	字节型	1	0	uint8别名
		rune	字符类型	4	0	专用于存储unicode编码，等价于uint32
		int, uint	整型	4或8	0	32位或64位
		int8, uint8	整型	1	0	-128 ~ 127, 0 ~ 255
		int16, uint16	整型	2	0	-32768 ~ 32767, 0 ~ 65535
		int32, uint32	整型	4	0	-21亿 ~ 21 亿, 0 ~ 42 亿
		int64, uint64	整型	8	0
		float32	浮点型	4	0.0	小数位精确到7位
		float64	浮点型	8	0.0	小数位精确到15位
		complex64	复数类型	8
		complex128	复数类型	16
		uintptr	整型	4或8		⾜以存储指针的uint32或uint64整数
		string	字符串		""	utf-8字符串

	*/

	/*
		布尔类型
			var v1 bool
			v1 = true
			v2 := (1 == 2) // v2也会被推导为bool类型

		//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
			var b bool
			b = 1 // err, 编译错误
			b = bool(1) // err, 编译错误

	*/
	bb := (1 == 1)
	bb2 := (2 == 1)
	println(bb)
	println(bb2)

	/*
		2.4.5 字符类型
			在Go语言中支持两个字符类型，一个是byte（实际上是uint8的别名），代表utf-8字符串的单个字节的值；另一个是 rune，代表单个unicode字符。
	*/
	var char rune = 'a' // rune 这个类型就是Java 里面的char类型
	//var char2 rune = "A" // cannot use "A" (untyped string constant) as rune value in variable declaration
	println(char) // 输出97的assic编码

	println("--------------------------")

	/*
		2.4.6 字符串
		在Go语言中，字符串也是一种基本类型：
		    var str string                                    // 声明一个字符串变量
		    str = "abc"                                       // 字符串赋值
		    ch := str[0]                                      // 取字符串的第一个字符
		    fmt.Printf("str = %s, len = %d\n", str, len(str)) //内置的函数len()来取字符串的长度
		    fmt.Printf("str[0] = %c, ch = %c\n", str[0], ch)

		    //`(反引号)括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
		    str2 := `hello
		    mike \n \r测试
		    `
		    fmt.Println("str2 = ", str2)
		        str2 =  hello
		        mike \n \r测试

	*/

	var str1 string = "string"
	println(str1[0])
	println(string(str1[1]))
	println(str1[2])
	println(len(str1))
	println(str1)

	var mStr string = `这是
多行
字符串
定义`
	println(mStr)

	println("--------------------------")

	/*
		2.4.7 复数类型
		复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）。
		    var v1 complex64 // 由2个float32构成的复数类型
		    v1 = 3.2 + 12i
		    v2 := 3.2 + 12i        // v2是complex128类型
		    v3 := complex(3.2, 12) // v3结果同v2

		    fmt.Println(v1, v2, v3)
		    //内置函数real(v1)获得该复数的实部
		    //通过imag(v1)获得该复数的虚部
		    fmt.Println(real(v1), imag(v1))

	*/
	println("--------------------------")

	/*
		2.6 类型转换
			Go语言中不允许隐式转换，所有类型转换必须显式声明，而且转换只能发生在两种相互兼容的类型之间。
				var ch byte = 97
				//var a int = ch //err, cannot use ch (type byte) as type int in assignment
				var a int = int(ch)

		2.7 类型别名
			    type bigint int64 //int64类型改名为bigint
			    var x bigint = 100
	*/

	var aa byte = 97
	var ch rune = rune(aa)
	var stra string = string(aa)
	println(aa)
	println(ch)
	println(stra)

	println("--------------------------")

	var com1 complex64 = 3 + 5i
	var int1 int = 666
	//var uint1 uint = -5 // cannot use -5 (untyped int constant) as uint value in variable declaration (overflows)
	var uint1 uint = 5 // cannot use -5 (untyped int constant) as uint value in variable declaration (overflows)
	var f1 float32 = 3.366
	var b bool = true
	var by byte = 1
	//var ru rune = 1
	var str11 string = "str11"
	//var err error = 1

	println(com1)
	println(int1)
	println(uint1)
	println(f1)
	println(b)
	println(by)
	println(str11)

}

/*
在goland中使用println()和	fmt.Println()的区别是什么？
1. println()
	来源：println() 是 Go 语言内置的函数，属于编译器提供的功能，而不是标准库的一部分。
	使用场景：通常用于调试或临时输出，不建议在生产代码中使用。

2. fmt.Println()
	来源：fmt.Println() 是 Go 标准库 fmt 包提供的函数。
	使用场景：适合在生产代码中使用，功能更强大且行为稳定。

总结：自己在学习go语言语法的时候可以使用println来实现快速输出内容；而在实际项目中禁止使用println，这时就要使用fmt.Println。

*/
