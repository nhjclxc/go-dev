package main

import "fmt"

// go的内存分配
// Go 同样也垃圾收集，也就是说无须担心内存分配和回收。
func main() {

	/*
		用new 分配内存
			内建函数new本质上说跟其他语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，一个*T类型的值。
			用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。记住这点非常重要。

			这意味着使用者可以用new 创建一个数据结构的实例并且可以直接工作。如bytes.Buffer 的文档所述 “Buffer 的零值是一个准备好了的空缓冲。”
			类似的，sync.Mutex 也没有明确的构造函数或 Init 方法。取而代之，sync.Mutex 的零值被定义为非锁定的互斥量。
	*/

	var ms *MyStudent = new(MyStudent)
	ms2 := new(MyStudent)

	fmt.Println(ms)
	fmt.Println(ms2)

	/*
		用make 分配内存
			回到内存分配。内建函数make(T, args)与new(T)有着不同的功能。它只能创建slice，map 和 channel，并且返回一个有初始值（非零）的T类型，而不是*T。
			本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
			例如，一个slice，是一个包含指向数据（内部array）的指针，长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。
			对于slice，map和channel，make初始化了内部的数据结构，填充适当的值。

			例如，make([]int, 10, 100) 分配了100个整数的数组，然后用长度10和容量100创建了slice 结构指向数组的前10个元素。
			区别是，new([]int)返回指向新分配的内存的指针，而零值填充的slice结构是指向nil的slice值。
	*/

	// 使用 myInt 来代替int类型数据
	type myInt int

	var int1 int = 666
	var myInt1 myInt = myInt(int1)

	fmt.Println(int1)
	fmt.Println(myInt1)

	var int1p *int = &int1
	fmt.Println(int1p)
	*int1p++
	fmt.Println(int1)
	fmt.Println(int1p)
	fmt.Println(*int1p)
	println("---------")
	(*int1p)++
	fmt.Println(int1)
	fmt.Println(int1p)
	fmt.Println(*int1p)

}

type MyStudent struct {
	id   int
	name string
}
