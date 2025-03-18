package main

import "fmt"

// 方法定义
// https://topgoer.com/方法/方法定义.html
func main() {

	/*
		Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)，这个就类似python的。
			• 只能为当前包内命名类型定义方法。这个类型是使用type Name struct定义的类型
			• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
			• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
			• 不支持方法重载，receiver 只是参数签名的组成部分。
			• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。
		一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
		所有给定类型的方法属于该类型的方法集。
	*/

	// 方法被用在对象上，对象来自于自定义的结构体对象，与Java中不带 static 的实例方法很像
	// 而函数是与实例对象无关的，与Java中带有 static 关键字的方法很像
	/*
		方法的定义：
			func (receiver Type) methodName(参数列表)(返回值列表){}
				* receiver Type：表示的就是这个方法绑定哪个对象类型的，receiver在此处被称为”接受者“
				* methodName：表示方法名称，注意：想要方法被外部包使用的话，方法名称必须使用大写打头
				* 参数列表：与普通的函数参数列表一样
				* 返回值列表：与普通的函数参数列表一样
	*/

	zhangsan := Person{
		id:   336699,
		name: "张三",
		age:  18,
	}

	fmt.Println(zhangsan)
	zhangsan.SetName("我不是药神")
	fmt.Println(zhangsan)
	fmt.Println(zhangsan.GetName())

	zhangsanP := &zhangsan
	fmt.Println(zhangsanP)
	zhangsanP.SetName("我是谁啊！！！")
	fmt.Println(zhangsanP)
	fmt.Println(*zhangsanP)
	fmt.Println(zhangsan)

	zhangsan.Print()

	println("-----------------------")
	zhangsan.ModifyName("ModifyName") // 无法修改原数据
	zhangsan.Print()
	zhangsan.ModifyNamePoint("ModifyNamePoint") // 可以修改原数据

	zhangsan.Print()

}

// type Person struct 只定义对象的数据结构
// 要想定义对象的方法，只能通过”func (receiver Type) methodName(参数列表)(返回值列表){}“的形式来Type类型的对象绑定一个方法
type Person struct {
	id   int32
	name string
	age  int
}

// 在这个例子中当我们使用指针时，Go 调整和解引用指针使得调用可以被执行。
//注意，当接受者不是一个指针时，该方法操作对应接受者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。
// 上面的意思就是方法如果要修改原数据，就必须使用对象的指针 receiver *T
// 如果方法内部不需要或者不允许修改原对象的数据，那么此时必须使用普通变量 receiver T

// (Person) setName 设置 Person 对象的名称
func (person *Person) SetName(name string) {
	person.name = name
}

func (person *Person) GetName() (name string) {
	return person.name
}

func (person Person) ModifyName(name string) {
	person.Print()
	person.SetName(name)
	person.Print()
}
func (person *Person) ModifyNamePoint(name string) {
	person.Print()
	person.SetName(name)
	person.Print()
}

func (person *Person) Print() {
	fmt.Printf("Person = {id = %d, name = %s, age = %d}.\n", person.id, person.name, person.age)
}
