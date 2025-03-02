package main

import "fmt"

/*
7.3 数组
数组是指一系列同一类型数据的集合。数组中包含的每个数据被称为数组元素（element），一个数组包含的元素个数被称为数组的长度。
*/
func main() {

	/*
		数组⻓度必须是常量，且是类型的组成部分。 [2]int 和 [3]int 是不同类型。
	*/
	//var n int = 10
	//var _ [n]int  //err, invalid array length n
	var _ [10]int //ok

	// 创建一个数组，并为其赋值
	var arr [10]int
	i := 0
	for ; i < 10; i++ {
		arr[i] = i * 10
	}

	// 使用for遍历数组
	for i = 0; i < 10; i++ {
		fmt.Printf("arr[%d] = %d \n", i, arr[i])
	}
	// 使用range遍历数组
	for index, val := range arr {
		println("index = ", index, ", val = ", val)
	}

	//println("printArray.start")
	//printArray(arr)
	//println("printArray.end")

	// 获取数组长度
	// 内置函数 len(长度) 和 cap(容量) 都返回数组⻓度 (元素数量)：
	// length  , capacity
	println(len(arr))
	println(cap(arr))

	// 创建数组时赋值,静态初始化,创建数组的时候就要赋值
	var arr1 [3]int = [3]int{1, 2} // 有一个数据没有初始化,那么没有得到初始化的值为0
	arr11 := [3]int{1, 2}
	printArray(arr1)
	printArray(arr11)
	println("---------------------------")

	// 已知数据,让go自行推断数据个数
	var arr2 []int = []int{1, 2, 3}
	arr21 := []int{1, 2, 3}
	for i, val := range arr2 {
		fmt.Printf("arr[%d] = %d \n", i, val)
	}
	for i, val := range arr21 {
		fmt.Printf("arr[%d] = %d \n", i, val)
	}
	println("---------------------------")

	// 更具索引赋值数组元素数据

	arr3 := []int{0: -1, 1: 111, 2: 222, 3: 666}
	for i, val := range arr3 {
		fmt.Printf("arr[%d] = %d \n", i, val)
	}

	// 多维数组
	arrD2 := [3][4]int{{1, 2, 3, 4}, {21, 22, 23, 24}, {31, 32, 33, 34}} // 3*4
	for i, arrInner := range arrD2 {
		for j, val := range arrInner {
			fmt.Printf("arr[%d][%d] = %d , ", i, j, val)
		}
		println()
	}
	q := [3][4]int{0: {1, 2, 3, 4}, 2: {21, 22, 23, 24}} // 3*4
	fmt.Println(q)

	// 比较两个数组是否完全相等【数组长度相等,数组类型相等,数组的每一个值相等】
	// 相同类型的数组之间可以使用 == 或 != 进行比较，但不可以使用 < 或 >，也可以相互赋值：

	arr2Test1 := [3][4]int{0: {1, 2, 3, 4}, 2: {21, 22, 23, 24}} // 3*4
	arr2Test2 := [3][4]int{0: {1, 2, 3, 4}, 2: {21, 22, 23, 24}} // 3*4

	println(arr2Test1 == arr2Test2)
	println(arr2Test1 != arr2Test2)

	arr2Test3 := [3][4]int{1: {1, 2, 3, 4}, 2: {21, 22, 23, 24}} // 数组元素位置不同
	println(arr2Test1 == arr2Test3)

	// 数组形状不同的直接不能比较
	//arr2Test4 := [3][3]int{1: {1, 2, 3}, 2: {21, 22, 23}} //
	//println(arr2Test1 == arr2Test4) // invalid operation: arr2Test1 == arr2Test4 (mismatched types [3][4]int and [3][3]int)

	//数组类型不同的不能比较
	//arr2Test4 := [3][4]float32{1: {1.0, 2.0, 3.0, 4.0}, 2: {21.0, 22.0, 23.0, 24.0}} //
	//println(arr2Test1 == arr2Test4) // invalid operation: arr2Test1 == arr2Test4 (mismatched types [3][4]int and [3][4]float32)

	/*
		【【【 go的函数参数传递是值传递的 】】】
		总是以值的方式传递的,传数组时会把数组的所有值传过去,数组越大,开销愈大
		   7.3.3 在函数间传递数组
		   	根据内存和性能来看，在函数间传递数组是一个开销很大的操作。
		   在函数之间传递变量时，总是以值的方式传递的。如果这个变量是一个数组，意味着整个数组，不管有多长，都会完整复制，并传递给函数。
	*/

	var arr222 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr222) // [1 2 3]
	testArr(arr222)
	// 以下输出表明直接通过数组方式传递数组,如何在其他函数里面修改是比起效果的
	fmt.Println(arr222) // [1 2 3]

	//要想在其他函数修改数组的元素值,那么必须使用指针来传递数组

	fmt.Println(arr222) // [1 2 3]
	testArr2(&arr222)
	// 以下输出表明直接通过数组方式传递数组,如何在其他函数里面修改是比起效果的
	fmt.Println(arr222) // [1 888 3]

}

// 通过指针数组来实现数组元素在自函数内部可以修改
func testArr2(arr *[3]int) {
	fmt.Println(*arr) // [1 2 3]
	arr[1] = 888
	fmt.Println(*arr) // [1 888 3]
}

func testArr(arr [3]int) {
	fmt.Println(arr) // [1 2 3]
	arr[1] = 666
	fmt.Println(arr) // [1 666 3]
}

func printArray(arr [3]int) {
	for i, val := range arr {
		fmt.Printf("arr[%d] = %d \n", i, val)
	}
}
