package main

// 17.4 运算符模式和接口
// https://denganliang.github.io/the-way-to-go_ZH_CN/17.4.html
func main() {
	// 运算符是一元或二元函数，它返回一个新对象而不修改其参数，类似 C++ 中的 + 和 *，特殊的中缀运算符（+，-，* 等）可以被重载以支持类似数学运算的语法。
	//但除了一些特殊情况，Go 语言并不支持运算符重载：为了克服该限制，运算符必须由函数来模拟。既然 Go 同时支持面向过程和面向对象编程，我们有两种选择：
	// 17.4.1 函数作为运算符
	// 17.4.2 方法作为运算符

	mode_01()
	mode_02()
	mode_03()

}

func mode_03() {
	// 17.4.3 使用接口
	//当在不同类型上执行相同的方法时，创建一个通用化的接口以实现多态的想法，就会自然产生。
	//
	//例如定义一个代数 Algebraic 接口：
	//
	//type Algebraic interface {
	//	Add(b Algebraic) Algebraic
	//	Min(b Algebraic) Algebraic
	//	Mult(b Algebraic) Algebraic
	//	…
	//	Elements()
	//}
	//然后为我们的 matrix 类型定义 Add()，Min()，Mult()，……等方法。
	//
	//每种实现上述 Algebraic 接口类型的方法都可以链式调用。每个方法实现都应基于参数类型，使用 switch 类型断言来提供优化过的实现。另外，应该为仅依赖于接口的方法，指定一个默认处理分支：
	//
	//func (a *denseMatrix) Add(b Algebraic) Algebraic {
	//	switch b.(type) {
	//	case sparseMatrix:
	//		return addDenseToSparse(a, b.(sparseMatrix))
	//	…
	//	default:
	//		for x in range b.Elements() …
	//	}
	//}
	//如果一个通用的功能无法仅使用接口方法来实现，你可能正在处理两个不怎么相似的类型，此时应该放弃这种运算符模式。例如，如果 a 是一个集合而 b 是一个矩阵，那么编写 a.Add(b) 没有意义。就集合和矩阵运算而言，很难实现一个通用的 a.Add(b) 方法。遇到这种情况，把包拆分成两个，然后提供单独的 AlgebraicSet 和 AlgebraicMatrix 接口。
}

func mode_02() {
	//	17.4.2 方法作为运算符
	//	根据接收者类型不同，可以区分不同的方法。因此我们可以为每种类型简单地定义 Add 方法，来代替使用多个函数名称：
	//
	//	func (a *sparseMatrix) Add(b Matrix) Matrix
	//	func (a *denseMatrix) Add(b Matrix) Matrix
	//	每个方法都返回一个新对象，成为下一个方法调用的接收者，因此我们可以使用链式调用表达式：
	//
	//	m := m1.Mult(m2).Add(m3)
	//	比上一节面向过程的形式更简洁。
}

func mode_01() {
	// 17.4.1 函数作为运算符
	// 运算符由包级别的函数实现，以操作一个或两个参数，并返回一个新对象。函数针对要操作的对象，在专门的包中实现。
	//例如，假设要在包 matrix 中实现矩阵操作，就会包含 Add() 用于矩阵相加，Mult() 用于矩阵相乘，他们都会返回一个矩阵。
	//这两个函数通过包名来调用，因此可以创造出如下形式的表达式：
	// m := matrix.Add(m1, matrix.Mult(m2, m3))

	// 如果我们想在这些运算中区分不同类型的矩阵（稀疏或稠密），由于没有函数重载，我们不得不给函数起不同的名称，例如：
	//func addSparseToDense (a *sparseMatrix, b *denseMatrix) *denseMatrix
	//func addDenseToDense (a *denseMatrix, b *denseMatrix) *denseMatrix
	//func addSparseToSparse (a *sparseMatrix, b *sparseMatrix) *sparseMatrix

}
