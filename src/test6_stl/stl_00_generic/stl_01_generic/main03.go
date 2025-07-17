package main

func main() {

}

/*
泛型接口定义：类型参数可以使用接口进行 约束（constraints），表示该类型参数必须实现某些行为或是某些类型的集合。
type Number interface {
	~int | ~float64 | ~int64
}
~int 表示可以是 int 或自定义的 int 类型别名（保持底层类型一致）
使用 | 表示多个类型的并集

然后你就可以像这样使用：
func Add[T Number](a, b T) T {
	return a + b
}

*/
