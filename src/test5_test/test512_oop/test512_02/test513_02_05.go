package main

import "fmt"

// 11.3 类型断言：如何检测和转换接口变量的类型
// https://denganliang.github.io/the-way-to-go_ZH_CN/11.3.html
func main() {

	/*
		// 一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。
		// 在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。
		// 通常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值：
		// v := varI.(T)       // unchecked type assertion
		// 类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。
		//如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：
			if v, ok := varI.(T); ok {  // checked type assertion
				Process(v)   // 转化成功，进行数据处理
				return
			}
			// varI is not of type T
		// 如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false，也没有运行时错误发生。

		// 注意：以上的 varI 必须是接口类型变量

	*/

	// 测试以上理论
	var shaperI Shaper
	var s = Square{6}
	var c = Circle{6}

	shaperI = &s
	//val1, ok1 := shaperI.(Square) // 不可能的类型断言: 'Square' 未实现 'Shaper'
	val1, ok1 := shaperI.(*Square) // 不可能的类型断言: 'Square' 未实现 'Shaper'
	fmt.Println(val1, ok1)
	println(val1.Area())
	val2, ok2 := shaperI.(*Circle)
	fmt.Println(val2, ok2)

	shaperI = &c
	val3, ok3 := shaperI.(*Square)
	fmt.Println(val3, ok3)
	val4, ok4 := shaperI.(*Circle)
	fmt.Println(val4, ok4)
	println(val4.Area())

}

// 正方形
type Square struct {
	side float32
}

// 圆形
type Circle struct {
	radius float32
}

// 计算面积的接口
type Shaper interface {
	Area() float32
}

// 实现接口
func (this *Square) Area() float32 {
	return this.side * this.side
}
func (this *Circle) Area() float32 {
	return (3.14 * this.radius * this.radius) / 2
}
