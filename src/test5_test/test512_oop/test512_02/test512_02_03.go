package main

import "fmt"

// 接口
// 什么是接口：
// 　　接口可以理解为对一组方法声明进行的统一命名,但是这些方法没有提供任何实现。
// 也就是说,把一组方法声明在一个接口中,然后继承于该接口的类都需要实现这些方法。
// 通过接口,你可以对方法进行统一管理,避免了在每一种类型中重复定义这些方法。
func main3() {

	p := Phone{
		size:    123,
		Product: Product{Name: "这是一个水果手机"},
	}
	fmt.Println(p)

	p.Connect()
	p.DisConnect()
	p.Send("这是一条消息哦哦哦哦哦！！！！！！")
	println("---------------------")

	c := Camera{
		remark:  "这是一个相机哦",
		Product: Product{Name: "这是一个索尼相机"},
	}
	fmt.Println(c)

	c.Connect()
	c.DisConnect()

	println("---------------------")

	println("_--------------模拟多态---------------")
	var usb USB = &p
	usb.Connect()
	usb.DisConnect()
	println("_--------------模拟多态---------------")
	usb = &c
	usb.Connect()
	usb.DisConnect()

	// 接口体现多态的两种形式
	//1、多态参数
	//　　在前面的 Usb 接口案例，Usb usb ，即可以接收手机变量，又可以接收相机变量，就体现了 Usb接口多态。
	//2、多态数组
	//　　给 Usb 数组中，存放Phone 结构体 和 Camera 结构体变量
	var useList []USB = []USB{&p, &c} // 接口要的必须是指针变量
	fmt.Println(useList)
	for i, usb := range useList {
		fmt.Printf("useList i = %d, usb = %v \n", i, usb)
	}

	println("---------------------")

	computer := Computer{
		cpu: 8,
		Product: Product{
			Name: "HP",
		},
	}
	computer.Working(&c)
	computer.Working(&p)

}

// interface 类型可以定义一组方法，但是这些不需要实现。并且 interface 不能包含任何变量。
// 到某个自定义类型(比如结构体 Phone)要使用的时候,在根据具体情况把这些方法写出来(实现)。
// 1、接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)。这点不难理解，和Java，C#是一样的概念

// 声明一个 USB 接口
type USB interface {
	// 1、接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低偶合的思想。
	Connect()
	DisConnect()
}

type Product struct {
	Name string
}

// Phone 定义一个手机对象，来使用接口
type Phone struct {
	size int
	Product
}

// 2、Golang 中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。
//因此，Golang 中没有 implement 或者C#中的 : 这样的关键字

// 实现 USB 的两个接口
func (this *Phone) Connect() {
	fmt.Printf("Phone.Connect.Name = %s, size = %d \n", this.Name, this.size)
}
func (this *Phone) DisConnect() {
	fmt.Printf("Phone.DisConnect.Name = %s, size = %d \n", this.Name, this.size)
}

//func (this *Phone) Call() {
//	fmt.Printf("Phone.Call.Name = %s, size = %d \n", this.Name, this.size)
//}

// Camera 定义一个相机来实现 USB 的两个接口
type Camera struct {
	remark string
	Product
}

// 实现 USB 的两个接口
func (this *Camera) Connect() {
	fmt.Printf("Camera.DisConnect.Name = %s, remark = %s \n", this.Name, this.remark)
}
func (this *Camera) DisConnect() {
	fmt.Printf("Camera.DisConnect.Name = %s, remark = %s \n", this.Name, this.remark)
}

// 将 USB 对象当成参数传入，类似接收者模式的
type Computer struct {
	cpu int
	Product
}

func (this *Computer) Working(usb USB) {
	//(*usb).Connect()
	usb.Connect()
	fmt.Printf("Computer.Working usb = %v \n", usb)
	usb.DisConnect()
}

// 声明一个 通信 接口
type Communicate interface {
	// 1、接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低偶合的思想。
	Send(msg string)

	// 接口继承接口
	// 如果加上了下面的 USB 接口到这里，那么如果一个结构体要想实现 Communicate 接口，
	// 那么其必须也实现 Communicate 接口继承的 USB 接口里面的方法，这个就是接口的继承
	USB
}

// 模拟多实现，Phone 实现多个接口
// Phone 实现 Communicate 接口
func (this *Phone) Send(msg string) {
	fmt.Printf("Phone.Send.msg = %s \n", msg)
}

//接口注意事项和细节
//1、接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)。这点不难理解，和Java，C#是一样的概念
//2、接口中所有的方法都没有方法体,即都是没有实现的方法。
//3、在 Golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口。（在C#中，一个类继承接口，需要实现这个接口的所有方法，而且VS编译器会提醒需要实现接口中的方法）
//4、一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
//5、只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
// 6、一个自定义类型可以实现多个接口（多实现，单继承）
//7、Golang 接口中不能有任何变量
//8、一个接口(比如 A 接口)可以继承多个别的接口(比如 B,C 接口)，这时如果要实现 A 接口，也必须将 B,C 接口的方法也全部实现。
//9、interface 类型默认是一个指针(引用类型)，如果没有对 interface 初始化就使用，那么会输出 nil
//10、空接口 interface{} 没有任何方法，所以所有类型都实现了空接口, 即我们可以把任何一个变量赋给空接口。

//小结：
//1、当 A 结构体继承了 B 结构体，那么 A 结构就自动的继承了 B 结构体的字段和方法，并且可以直接使用
//2、当 A 结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此我们可以认为：实现接口是对继承机制的补充
//3、实现接口可以看作是对 继承的一种补充
//4、接口和继承解决的解决的问题不同
//　　继承的价值主要在于：解决代码的复用性和可维护性。
//　　接口的价值主要在于：设计，设计好各种规范(方法)，让其它自定义类型去实现这些方法。
//5、接口比继承更加灵活
//　　接口比继承更加灵活，继承是满足 is - a 的关系，而接口只需满足 like - a 的关系。
//6、接口在一定程度上实现代码解耦
