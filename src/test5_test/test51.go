package main

import "fmt"

/*
来源：/doc/学习Go语言(Golang).pdf
*/
func main() {

	//test51_1()

	//test51_2()

	//test51_3()

	//test51_4()

	//test51_5()

	//test51_6()

	//test51_7()

	//test51_8()

	//test51_9()

	//test51_10()

	//test51_11()

	test51_12()

}

func test51_12() {
	/*
		编写一个Go程序可以逆转字符串，例如“foobar”被打印成“raboof”。
		提示：不幸的是你需要知道一些关于转换的内容，参阅“转换”第59页的内容。
	*/

}

func test51_11() {
	/*
		现有字符串：asSASA ddd dsjkdsjs dk
		要求：替换位置4开始的三个字符为“abc”。
	*/
	str := "asSASA ddd dsjkdsjs dk"
	subStr := "abc"

	temp := ""
	j := 0
	for i := 0x0; i < len(str); i++ {
		if i >= 3 && i < 3+len(subStr) {
			temp += string(subStr[j])
			j++
		} else {
			temp += string(str[i])
		}

	}
	println(str)
	println(temp)

	// 使用切片方法
	index := 4
	temp2 := str[:index] + subStr + str[index+len(subStr):]
	println(temp2)

}

func test51_10() {
	/*
	   2. 建立一个程序统计字符串里的字符数量：
	   asSASA ddd dsjkdsjs dk
	   同时输出这个字符串的字节数。提示：看看unicode/utf8包。
	*/

	str := "asSASA ddd dsjkdsjs dk"
	//println(str[0])
	//println(str[1])
	//println(str[2])
	//println(int('a'))
	//println(int('z'))
	//println(int('A'))
	//println(int('Z'))
	count := 0
	for i := 0; i < len(str); i++ {
		if (int('a') <= int(str[i]) && int(str[i]) <= int('z')) || (int('A') <= int(str[i]) && int(str[i]) <= int('Z')) {
			count++
		}
	}
	println(count)

}

func test51_9() {
	/*
	   1. 建立一个Go程序打印下面的内容（到100个字符）：
	   A
	   AA
	   AAA
	   AAAA
	   AAAAA
	   AAAAAA
	   AAAAAAA
	*/

	count := 0
	for i := 1; i < 100; i++ {
		count += i
		if count > 100 {
			i = count - 100
		}

		for j := 0; j < i; j++ {
			print("A")
		}
		println()
		if count >= 100 {
			break
		}
	}

}

func test51_8() {
	//1. 解决这个叫做Fizz-Buzz[http://imranontech.com/2007/01/24/using-fizzbuzz-to-find-developers-who-grok-coding/] 的问题：
	//编写一个程序，打印从1到100的数字。
	//		当是三个倍数就打印“Fizz”代替数字，
	//		当是的五的倍数就打印“Buzz”。
	//		当数字同时是三和五的倍数时，打印“FizzBuzz”。

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else {
			fmt.Println(i)
		}
	}

}

func test51_7() {
	//1. 创建一个基于for的简单的循环。使其循环10次，并且使用fmt包打印出计数器的值。

	println("---------------111-----------------")
	count := 0
	for i := 0; i < 10; i++ {
		count++
	}
	println("count = ", count)

	//2. 用goto改写1的循环。关键字for不可使用。
	println("----------------222----------------")
	count2 := 0
MyFor:
	if count2 < 10 {
		count2++
		println("count2 = ", count2)
		goto MyFor
	}
	println("count2 = ", count2)

	//3. 再次改写这个循环，使其遍历一个array，并将这个array打印到屏幕上。
	println("----------------333----------------")
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
}

/*
*
map相关
map可以认为是一个用字符串做索引的数组（在其最简单的形式下）
一般定义map的方法是：map[<from type>]<to type>
*/
func test51_6() {

	//一个key是int类型，value是string类型的字典的定义：
	var dict1 map[int]string
	var dict2 map[int]string = map[int]string{
		1: "张三",
		2: "里斯",
		3: "王五", // 注意：最后一个逗号是必须的
	}
	dict3 := map[int]string{
		1: "张三111",
		2: "里斯333",
		3: "王五666", // 注意：最后一个逗号是必须的
	}
	var dict4 map[int]string = make(map[int]string)
	fmt.Println(dict1)
	fmt.Println(dict2)
	fmt.Println(dict3)
	fmt.Println(dict4)
	fmt.Println(dict2[1])
	dict2[1] = "你好"
	dict2[5] = "你是谁"
	fmt.Println(dict2)
	val, ok := dict2[5]
	fmt.Println(ok)  // 有对应key的数据这个ok返回true
	fmt.Println(val) // 同时返回其对应的数据
	val666, ok666 := dict2[666]
	fmt.Println(ok666)  // 没有对应key的数据，这个ok返回false
	fmt.Println(val666) // 同时value返回空字符串

	println("---------------")
	fmt.Println(dict2)
	delete(dict2, 5) // 删除dict里面key为5的数据
	fmt.Println(dict2)
	fmt.Println(dict2)
	delete(dict2, 5) // 删除dict里面key为5的数据
	fmt.Println(dict2)

}

/*
*
slice相关
slice 与 array 接近，但是在新的元素加入的时候可以增加长度。slice总是指向底层的一个array。
slice 是一个指向 array 的指针，这是其与array不同的地方；slice是引用类型，这意味着当赋值某个slice到另外一个变量，两个引用会指向同一个array。
例如，如果一个函数需要一个slice参数，在其内对slice元素的修改也会体现在函数调用者中，这和传递底层的array指针类似。
*/
func test51_5() {

	// 以下先讨论数组的两种创建方式
	// 注意：一个数组的长度是不能改变的，即数组在创建之初长度是多少就是多少？
	var arr1 [5]int               // 这个数组在声明的时候就是定义为长度是5了，后面不能改变
	arr2 := [5]int{1, 2, 3, 4, 5} //这个数组在声明的时候就是创建的时候就已经默认赋值5个数据了，被go编译器推断为长度5，后面不能改变
	fmt.Println(arr1)
	fmt.Println(arr2)
	//arr1[5] = 666 //无效的 数组 索引 '5' (5 元素的数组超出界限)
	//arr2[5] = 666 //无效的 数组 索引 '5' (5 元素的数组超出界限)
	//append(arr1, 5) // 无法将 'arr1' (类型 [5]int) 用作类型 []Type

	// 注意：数组array和切片slice的数据格式都相同，如[10]int，不同的是创建方式的不同
	// 讨论slice的创建方式
	var sli = make([]int, 3)
	fmt.Println(sli)
	sli[0] = 1
	sli[1] = 11
	sli[2] = 111
	fmt.Println(sli)
	sli = append(sli, 666)
	fmt.Println(sli)

}

/*
*
数组array
*/
func test51_4() {

	//数组的大小也是判断数组是否相等的一个条件

	var arr1 [5]int
	var arr2 [5]int
	//var arr3 [10]int
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr1 == arr2) // 长度相等，且每个元素都相等(默认元素都是0)。此时两个数组相等
	//fmt.Println(arr1 == arr3) // 无效运算: arr1 == arr3(类型 [5]int 和 [10]int 不匹配)
	arr2[1] = 111
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr1 == arr2) // 长度相等，但是有元素不相等，所以两个数组不相等

	// 数组同样是值类型的：将一个数组赋值给另一个数组，会复制所有的元素。
	//尤其是当向函数内传递一个数组的时候，它会获得一个数组的副本，而不是数组的指针。

	var arr21 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr22 [5]int = arr21
	fmt.Println(arr21)
	fmt.Println(arr22)
	arr22[2] = 666
	fmt.Println(arr21)
	fmt.Println(arr22) // 尤其输出可知，赋值符号是深拷贝

	// 在不使用指针传递数据的时候，go默认是会深拷贝一份数据出来作为参数传递给子函数
	fmt.Println(arr22) // [1 2 666 4 5]
	arrTest1(arr22)
	fmt.Println(arr22) // [1 2 666 4 5]

	// 将数组地址传入函数进行操作，原数据会被修改
	fmt.Println(arr22) // [1 2 666 4 5]
	arrTest2(&arr22)
	fmt.Println(arr22) // [1 999 666 4 5]

	// 二维数组
	a1 := [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	a2 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(a1)
	fmt.Println(a2)
}

func arrTest1(arr [5]int) {
	arr[1] = 369
}
func arrTest2(arr *[5]int) {
	arr[1] = 999
}

/*
*
golang内建函数
Table 1.3. Go 中的预定义函数

	print
	println
	new
	make
	close
	delete
	recover
	len
	cap
	append
	copy
	real
	complex
	panic
	imag
*/
func test51_3() {

}

/*
*
switch的高级用法
*/
func test51_2() {

	var i int

	// 第一个返回值是行数，第二个返回值是输入是遇到的错误，如果没有错误则返回nil
	scanln, err := fmt.Scanln(&i)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(scanln)
	println(i)

	println("---------------------")
	switch {
	case i > 0 && i <= 10:
		println("i > 0 && i <= 10")
	case i > 10 && i <= 100:
		println("i > 10 && i <= 100")
	case i > 100 && i <= 1000:
		println("i > 100 && i <= 1000")
	default:
		println("default")
	}

	// fallthrough 关键字
	/*
		📌	 fallthrough 的作用
				默认情况下，switch 语句 不会自动执行下一个 case，匹配到的 case 会执行，并在遇到 break（默认行为）后退出。
				fallthrough 会继续执行 紧接着的 case 语句（即使它的条件不匹配）。
				只能用于 case 语句的最后一行。
				不能跳过 case，只能向下一个 case 继续执行。
	*/
	// 当case已经匹配，且这时还想让其执行下一个操作时，在已经匹配的case的最后一行写fallthrough，那么此时还会执行下面的一个case语句
	// 此时无论下一个case的条件是什么，下一个case都会被执行
	// fallthrough可以被连续使用，此时效果即和被拆开的效果是一样的
	// 注意：❌ fallthrough 不能用于 default
	println("---------------------")
	switch {
	case i > 0 && i <= 10:
		println("211  i > 0 && i <= 10")
		// 如果这里匹配上了，而接下来的逻辑又和下一个匹配相同，则把这个匹配操作往下扔
		fallthrough
		//println("2  i > 0 && i <= 10")  // 注意：后面不能再加语句了
	case i > 10 && i <= 100:
		println("2  i > 10 && i <= 100")
		fallthrough
	case i > 100 && i <= 1000:
		println("2  i > 100 && i <= 1000")
	default:
		println("2  default")
		//fallthrough // 不能在 'switch' 语句的 final case 中使用 'fallthrough'
	}

	/*
		📌 什么时候使用 fallthrough？
				合并多个 case 的逻辑，但又想保持清晰的 case 结构。
				有意让多个 case 执行，但不想写多个 case 重复代码。
				避免 case 代码重复，但又要执行多个 case 逻辑。
	*/
}

/*
break 可以指定结束哪一个循环，该循环必须使用标签来指定
*/
func test51_1() {

	// 指定标签，这个标签的指定和goto的指定一样，用大写开头
Ifor:
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			fmt.Printf("i = %d, ", i)
			if i > 2 {
				break Ifor
			}
		}
		fmt.Println()
	}
	fmt.Println("\n----------------------------\n")
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			fmt.Printf("i = %d, ", i)
			if i > 2 {
				break
			}
		}
		fmt.Println()
	}

}
