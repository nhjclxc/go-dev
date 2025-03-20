package main

import "fmt"

// golang接口类型断言
func main() {
	/*
		在 Golang 中，接口类型断言（Type Assertion）用于将一个接口类型的值转换为其具体类型。类型断言的语法如下：
				value, ok := interfaceValue.(TargetType)
		其中：
			interfaceValue 是一个接口类型的变量。
			TargetType 是要转换的具体类型。
			value 是转换后的值。
			ok 是一个布尔值，表示类型断言是否成功。如果成功，ok 为 true，否则 ok 为 false，value 为零值。

	*/

	mo := MyObject{
		name: "啊啊啊啊啊",
	}
	fmt.Println(mo)

	mo.doSomeThing1()

	// 必须要通过下面这句话来对其进行转化
	var int1 MyInterface1 = &mo
	value, ok := int1.(MyInterface1) // 断言 int1 为 MyInterface1 类型
	if ok {
		fmt.Println("111类型断言成功:", value)
	} else {
		fmt.Println("111类型断言失败")
	}
	value2, ok2 := int1.(MyInterface2) // // 断言 int1 为 MyInterface2 类型
	if ok2 {
		fmt.Println("222类型断言成功:", value2)
	} else {
		fmt.Println("222类型断言失败")
	}
	println("---------------------")

	// 使用多态来声明一个Animal类型的Dog实例变量
	// 多态：使用接口类型的变量指向接口实现类的对象
	// 在 Golang 中，多态的实现依赖于接口（interface）。接口类型的变量可以指向任何实现该接口的结构体对象，从而实现多态。
	var a Animal = Dog{} // a 变量是 Animal 接口类型，实际存储的是 Dog 类型的值

	dog, ok := a.(Dog) // 类型断言为 Dog
	if ok {
		fmt.Println("类型断言成功，转换为 Dog 类型")
		dog.Speak()
	} else {
		fmt.Println("类型断言失败")
	}

	/*
		2. 多态与方法接收者
		如果接口方法是由指针接收者实现的，则必须用结构体的指针赋值给接口变量，否则编译报错。

		3. 多态与切片
		Golang 允许存储多个实现了接口的对象，可以用 []接口类型 来存储不同类型的对象。

		4. 多态与函数参数
		接口变量可以作为函数参数，接收所有实现了该接口的结构体实例。

		5. 空接口与完全动态多态
		如果希望接口可以存储任何类型的对象，可以使用空接口 (interface{})，因为它是所有类型的父类。
	*/
}

type MyInterface1 interface {
	doSomeThing1()
}
type MyInterface2 interface {
	doSomeThing2()
}

type MyObject struct {
	name string
}

func (this *MyObject) doSomeThing1() {
	fmt.Printf("MyObject.doSomeThing.name = %s \n", this.name)
}

// 另一个示例
type Animal interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

// 另一个示例
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("整数:", v)
	case string:
		fmt.Println("字符串:", v)
	case bool:
		fmt.Println("布尔值:", v)
	default:
		fmt.Println("未知类型")
	}
}
