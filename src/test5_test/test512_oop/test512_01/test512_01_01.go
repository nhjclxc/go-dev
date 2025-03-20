package main

import (
	"fmt"
	"unsafe"
)

// https://www.cnblogs.com/taotaozhuanyong/p/14713515.html
/*
1、Golang 也支持面向对象编程(OOP)，但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言。所以我们说 Golang 支持面向对象编程特性是比较准确的。

2、Golang 没有类(class)，Go 语言的结构体(struct)和其它编程语言的类(class)有同等的地位，可以理解 Golang 是基于 struct 来实现 OOP 特性的。

3、Golang 面向对象编程非常简洁，去掉了传统 OOP 语言的继承、方法重载、构造函数和析构函数、隐藏的 this 指针等等

4、Golang 仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它 OOP 语言不一样，比如继承 ：Golang 没有 extends 关键字，继承是通过匿名字段来实现。

5、Golang 面向对象(OOP)很优雅，OOP 本身就是语言类型系统(type system)的一部分，通过接口(interface)关联，耦合性低，也非常灵活。在 Golang 中面向接口编程是非常重要的特性。
*/
func main1() {

	p1 := Person{"zhangsan", 18, []int{1, 2, 3, 4, 5, 6}}
	fmt.Println(p1)
	p2 := p1 // 这一步可以被看做Java里面的属性拷贝，注意：基本数据类型复制一份新的数据出来，但是引用数据类型是
	fmt.Println(p2)
	println("---------------------")
	// p1 地址 is 0xc000008030,p2 地址 is 0xc000008060
	// 因此 通过p2 := p1创建赋值的对象是两个不同的对象
	fmt.Printf("p1 地址 is %p,p2 地址 is %p \n", &p1, &p2)

	// 以下修改可以看，出对于基本数据类型是值拷贝，
	p2.Name = "里斯"
	fmt.Printf("p1.Name is %v,p2.Name is %v \n", p1.Name, p2.Name)

	// 以下可以看出，对于引用数据类型指针、切片和映射是地址引用的赋值，底层并没有创建一个新数据
	p2.arr[2] = 666
	fmt.Printf("p1.arr is %v,p2.arr is %v \n", p1.arr, p2.arr)

	println("---------------------")

	var person1 Person
	person1.Age = 18
	person1.Name = "bingle1111"

	/*
		var person2 Person = person1
		可以理解为以下两行代码
		var person2 Person  // 先创建一个对象，并且这个时候是0值
		person2 = person1  // 在进行一次数据拷贝
	*/

	var person2 Person
	fmt.Printf("person2.Name is %v,person1.Name is %v \n", person2.Name, person1.Name)
	person2 = person1
	fmt.Println(person2.Age)
	person2.Name = "bingle2222"
	fmt.Printf("person2.Name is %v,person1.Name is %v  \n", person2.Name, person1.Name)

	println("----------以下展示指针的指针----------------------")

	var person3 *Person = &person1 // person3 := &person1
	// person1 地址 is 0xc00001e270, person3 地址 is 0xc000064068
	fmt.Printf("person1 地址 is %p, person3 地址 is %p \n", &person1, &person3)
	// person1.Name is bingle1111, person3.Name is bingle1111
	fmt.Printf("person1.Name is %v, person3.Name is %v  \n", person1.Name, person3.Name)
	person3.Name = "你是谁"
	// person1.Name is 你是谁, person3.Name is 你是谁
	fmt.Printf("person1.Name is %v, person3.Name is %v  \n", person1.Name, person3.Name)

	// 从上面的输出可以看出，
	// 使用 “person3 := &person1” 这一条语句进行创建 person3 的时候，实际上是把person1的地址给了 person3，这个 person3 是一个指针变量
	// 这个指针变量又存储了一个指针地址，存储的就是 person1 的指针地址，因此 person1 和 person3 同时指向同一个地址
	// 实际上这个 person3 变量就是 Person 数据类型的指针的指针变量，
	// 因此 person3 里面存的就是 person1这个的地址，其中 person3 自己也有一个地址
	// 指针的指针变量的内存示意图看[指针的指针变量的内存示意图.png]

	println("----------以下： 验证结构体的所有字段在内存中是连续的----------------------")
	var str1 string
	fmt.Println(unsafe.Sizeof(str1)) // 输出 16
	var int1 int
	fmt.Println(unsafe.Sizeof(int1)) // 输出 8
	var intArr1 []int
	fmt.Println(unsafe.Sizeof(intArr1)) // 输出 24

	fmt.Printf("地址 %p \n", &p1)      // 地址 0xc00008a180
	fmt.Printf("地址 %p \n", &p1.Name) // 地址 0xc00008a180
	fmt.Printf("地址 %p \n", &p1.Age)  // 地址 0xc00008a190
	fmt.Printf("地址 %p \n", p1.arr)   // 地址 0xc0000a6000
	// p1 的地址和第一个元素的地址 p1.Name 相同
	// p1.Name 与 p1.Age相差 16 字节 ，与 unsafe.Sizeof(str1)输出一致

}

type Person struct {
	Name string
	Age  int
	arr  []int
}
