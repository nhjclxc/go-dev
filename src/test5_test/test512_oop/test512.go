package main

import "fmt"

// https://www.cnblogs.com/taotaozhuanyong/p/14713515.html
// https://www.cnblogs.com/taotaozhuanyong/p/14714206.html
// 以下属于 go 语言面向对象基础知识点
func main() {

	// Go 语言虽然不是纯粹的面向对象编程（OOP）语言，但它仍然支持面向对象的编程思想，主要体现在以下几个方面：

	person := Person{
		Name: "张三",
		age:  18,
	}
	fmt.Println(person)
	fmt.Println(person.GetAge())
	person.SayHello()
	person.GrowUp()
	fmt.Println(person.GetAge())

	// 创建一个学生，要使用到上面的 person
	student := Student{
		Person: person,
		score:  80,
	}
	fmt.Println(student)
	student.study()
	student.SayHello()

	// 测试 Work 接口的多态功能
	teacher := Teacher{
		wages: 8000,
		Person: Person{
			Name: "李老师",
			age:  32,
		},
	}
	fmt.Println(teacher)
	teacher.doWork()
	student.doWork()

}

// 1. 结构体（Struct）
// Go 没有类（Class）的概念，而是使用 结构体（struct） 来封装数据和行为。
type Person struct {
	// 名字字段Name的开头大写，表示这个字段在外部包是公开访问的  public
	Name string
	// 年龄字段age的第一个字母是小写的，表示这个字段在外部包是不可访问的 private
	age int
}

// 2. 方法（Method）
// Go 支持在结构体上定义方法（方法是带有接收者的函数）。
// (p Person) 表示的是值接收者：方法不会修改结构体本身。
func (self Person) SayHello() {
	fmt.Println("Hello, my name is", self.Name)
}

// (self *Person) 表示的是指针接收者：可以修改结构体内容，避免值拷贝，提高效率。
func (self *Person) GrowUp() {
	self.age++
}

// 3. 封装（Encapsulation）
// Go 没有 private、protected、public 关键字，而是通过大小写来控制可见性：
// 大写开头的变量或方法：可以被包外访问（public）。
// 小写开头的变量或方法：只能在当前包内访问（private）。
// 比如如上字段age就是不可直接访问的，但是可以提供一个方法给外部包使用
func (self Person) GetAge() int {
	return self.age
}

// 4. 继承与组合（Composition）
// Go 不支持传统的类继承（extends），而是通过 结构体嵌套（Composition） 来实现类似继承的效果。
// 以上的Person是一个比较通用的关于人的实体类，
// 以下定义一个 Student 实体类，其中Student实体类通过匿名字段嵌套Person实体类
type Student struct {
	// 直接吧 Person 实体类摆在这，就是匿名字段，属性和方法都可以用
	Person
	score int
}

func (self Student) study() {
	fmt.Printf("%s 正在学习... \n", self.Name)
}

// 5. 接口（Interface）
//Go 通过 接口（interface） 进行多态。

// 定义一个接口 工作的接口，学生工作就是学习，老师工作就是教书
type Worker interface {
	// 通过 doWork 方法进行多态表示
	doWork()
}

// 继续定义一个老师 Teacher 实体类来说明多态
type Teacher struct {
	Person
	wages int
}

// 接口是隐式实现：无需 implements 关键字。
//接口支持多态：一个变量可以存储不同实现。

// Duck Typing（鸭子类型）
//Go 采用 结构化类型系统，接口实现是隐式的，即只要类型实现了接口方法，它就自动被认为实现了该接口。
// 也就是下面的 doWork() 方法实现的时候为什么不需要使用像 Java 里面的 implements 关键字的原因
// go会自动去绑定 接口

// 为 Student 绑定实现一个 doWork() 工作的方法
func (self Student) doWork() {
	fmt.Printf("我是一名学生，我的名字是：%s，我目前的学习成绩是：%d 分。\n", self.Name, self.score)
}

// 为 Teacher 绑定实现一个 doWork() 工作的方法
func (self Teacher) doWork() {
	fmt.Printf("我是一名教师，我的名字是：%s，我目前的工资是：%d 美元/月。 \n", self.Name, self.wages)
}

// 6. 多态与空接口（Empty Interface）
// Go 允许使用 空接口 interface{} 代表任意类型，实现泛型效果。
func PrintAnything(val interface{}) {
	fmt.Println(val)
}

// 可以结合 类型断言 和 类型判断（type switch） 处理不同类型：
func PrintType(v interface{}) {
	// 这对某一个变量是未知类型的时候很有用，
	// 这个 interface{} 就是类似于Java里面 Object 类型的一个类型
	switch v.(type) {
	case string:
		fmt.Println("It's a string")
	case int:
		fmt.Println("It's an int")
	default:
		fmt.Println("Unknown type")
	}
}
