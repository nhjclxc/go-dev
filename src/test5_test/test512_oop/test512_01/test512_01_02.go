package main

import "fmt"

// go 结构体对象属性拷贝
func main2() {
	//3、Golang 中的方法作用在指定的数据类型上的(即：和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅仅是 struct，
	//比如 int , float32 等都可以有方法（有点类似于 C# 中的扩展方法）

	var i integer = 666

	i.print()

}

// 为int实例变量自定义绑定一个方法，但是不能直接使用int，int是go的类型，因此我们要变换一下
type integer int

func (this *integer) print() {
	fmt.Println("i.print: ", this)
	fmt.Println("i.print: ", *this)
}
