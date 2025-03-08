package main

import (
	"fmt"
)

/*
*
切片并不是数组或数组指针，它通过内部指针和相关属性引⽤数组⽚段，以实现变⻓⽅案。

slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
*/
func main() {

	println("slice")

	//test1()

	//test2()

	//test3()

	//test4()

	test5()

	/**
	  切片方法[::]是浅拷贝，copy方法是深拷贝
	*/
}

/*
*
golang的内置函数copy与slice
*/
func test5() {
	//函数 copy 在两个 slice 间复制数据，复制⻓度以 len 小的为准，两个 slice 可指向同⼀底层数组。
	// copy(dst []Type, src []Type)

	var slice []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	var sli1 []int
	fmt.Println(sli1) // []
	copy(sli1, slice)
	fmt.Println(sli1) // []

	var sli2 []int = []int{11, 22, 33}
	fmt.Println(sli2) // [11 22 33]
	copy(sli2, slice)
	fmt.Println(sli2) // [0 1 2]

	var sli3 []int = make([]int, 5)
	fmt.Println(sli3) // [0 0 0 0 0]
	copy(sli3, slice)
	fmt.Println(sli3) // [0 1 2 3 4]

}

/*
*
golang的内建函数append与sliece
*/
func test4() {

	var sli []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sli)

	// append(array/slice, args...)

	sli2 := append(sli, 1)
	fmt.Println(sli)
	fmt.Println(sli2)
	sli3 := append(sli2, 22, 333)
	fmt.Println(sli)
	fmt.Println(sli2)
	fmt.Println(sli3)
	// 由以上操作可知，append不改变元数据

	println("-------------------------")
	var slice1 []int // 生成一个[]的数据
	var slice12 = append(slice1, 1)
	fmt.Println(slice1)
	fmt.Println(slice12)

	println("-------------------------")

	var slice2 []int = make([]int, 5) // 生成一个[0 0 0 0 0]的数据
	var slice22 = append(slice2, 1)
	fmt.Println(slice2)
	fmt.Println(slice22)
	var slice222 = append(slice22, 22, 333)
	fmt.Println(slice22)
	fmt.Println(slice222)

	// 根据以上实验可知，append不会改变原数据，append的追加形式是在原数据的基础上在末尾追加数据的

	/// append函数会自动地底层数组的容量增长，一旦超过原底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的数据：

	println("-------------------------------------")
	var sli1 []int
	for i := 0; i < 10; i++ {
		var sli2 = append(sli1, i)
		fmt.Println(len(sli2), cap(sli2), sli2)

		sli1 = append(sli2, i)
	}
	/*
		1 1 [0]
		3 4 [0 0 1]
		5 8 [0 0 1 1 2]
		7 8 [0 0 1 1 2 2 3]
		9 16 [0 0 1 1 2 2 3 3 4]
		11 16 [0 0 1 1 2 2 3 3 4 4 5]
		13 16 [0 0 1 1 2 2 3 3 4 4 5 5 6]
		15 16 [0 0 1 1 2 2 3 3 4 4 5 5 6 6 7]
		17 32 [0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8]
		19 32 [0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9]
	*/

}

/*
*
切片slice与数组array之间的关系
*/
func test3() {

	// 创建一个数组
	var arr []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 从数组中获得一个切片
	sli := arr[2:5:6]
	fmt.Printf("Type of sli: %T\n", sli)

	fmt.Println(arr) // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(sli) // [2 3 4]
	sli[0] = 666     // 注意：操作的时候索引只能在slice的索引编号内操作
	fmt.Println(arr) // [0 1 666 3 4 5 6 7 8 9]
	fmt.Println(sli) // [666 3 4]
	arr[3] = 999
	fmt.Println(arr) // [0 1 666 999 4 5 6 7 8 9]
	fmt.Println(sli) // [666 999 4]
	// 由上述代码输出，可以得出结论：slice其实就是arr的引用数据类型，即将arr的某一块内存区域使用一个引用变量slice来提取出来，但是slice操作的数据还是原数据

}

/*
*
slice的切片操作
*/
func test2() {

	var slice1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	/*
	   操作				含义
	   s[n]				切片s中索引位置为n的项
	   s[:]				从切片s的索引位置0到len(s)-1处所获得的切片
	   s[low:]			从切片s的索引位置low到len(s)-1处所获得的切片
	   s[:high]			从切片s的索引位置0到high处所获得的切片，len=high
	   s[low:high]		从切片s的索引位置low到high处所获得的切片，len=high-low
	   s[low:high:max]	从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low
	   len(s)			切片s的长度，总是<=cap(s)
	   cap(s)			切片s的容量，总是>=len(s)
	*/
	println(slice1)

	// s[n]				切片s中索引位置为n的项
	println(slice1[0])
	println(slice1[1])
	println(slice1[2])

	//s[:]				从切片s的索引位置0到len(s)-1处所获得的切片
	slice11 := slice1[:]
	fmt.Printf("Type of x: %T\n", slice1) // Type of x: []int
	fmt.Printf("Type of x: %T\n", slice11)
	print(slice1)
	print(slice11)

	//s[low:]			从切片s的索引位置low到len(s)-1处所获得的切片
	print(slice1[2:])
	//s[:high]			从切片s的索引位置0到high处所获得的切片，len=high
	print(slice1[:3])
	//s[low:high]		从切片s的索引位置low到high处所获得的切片，len=high-low
	print(slice1[2:5])
	//s[low:high:max]	从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low
	print(slice1[2:5:6]) // 必须为低值 <= 高值 <= 最大值
	//len(s)			切片s的长度，总是<=cap(s)
	//print(cap(len(slice1)))
	//cap(s)			切片s的容量，总是>=len(s)
	//print(cap(slice1))

}

func print(sli []int) {
	for i, val := range sli {
		fmt.Printf("i = %d, val = %d \n", i, val)
	}
	println("---------------------------")
}

/*
*
slice的创建
*/
func test1() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//for i := 0; i < len(arr); i++ {
	//	println(arr[i])
	//}

	for i, val := range arr {
		fmt.Printf("i = %d, val = %d \n", i, val)
	}

	// 通过数组获得切片
	// 0表示从arr的第0个位置开始切割
	// 3表示切割长度为3
	// 5表示切割后的找个slice的最大容量是5
	slice1 := arr[0:3:5]

	println(slice1)
	//slice1[3] = 666
	//println(slice1)

	// 声明一个数组
	//var arr11 [6]int
	// 创建slice方式一：声明一个切片slice，声明切片和声明一个array是一样的操作
	var slice11 []int // 这种方式声明的slice默认值为nil
	println(slice11)
	// go里面实验nil来表示像Java一个的null
	if slice11 == nil {
		println("slice11 == nil")
	} else {
		println("slice11 != nil")
	}

	// 创建slice方式二：使用 make 关键字。注意：make只能创建slice、map和channel，并且返回一个有初始值(非零)
	// 使用示例：make(slice数据类型, length, capacity)，其中capacity部分可以省略
	var slice22 []int = make([]int, 6) // 这种方式声明的slice默认值为所有0，长度是lemgth的长度
	var slice221 []int = make([]int, 6, 66)
	println(slice22)
	println(slice221)

	// 创建slice的方式三：声明的时候初始化数据
	var slice33 = []int{1, 2, 3, 4, 5, 6} // 此时的length=capacity
	println(len(slice33))
	println(slice33)

}
