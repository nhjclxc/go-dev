package main

func main() {

}

/*
泛型方法定义：

	func (b Box[T]) Get() T {
		return b.Value
	}

注意：如果要定义泛型方法，则该方法绑定的结构体也必须是泛型
*/
type Box[T any] struct {
	Value T
}

func (b Box[T]) Get() T {
	return b.Value
}

type Box22 struct {
	Value string
}

// 该方法为泛型，因为接收器的基类型为泛型类型。方法规范必须声明与接收器基本类型的类型形参相对应的类型形参。
//方法类型实参必须使用接收器基类型中使用的类型约束。有关泛型的详细信息，
//func (box22 *Box22) Get[T any]() string {
//	return "nihao"
//}
