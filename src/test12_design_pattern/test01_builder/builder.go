package test01_builder

/*
	separate the construction of a complex object from representation
	设计思想：
		*Builder interface (包含1.biuld_XX method 返回的是biulder接口，2.get_XX 返回对象)
		*父struct
		*Director struct, 属性为Builder, 实现Construct()和SetBuilder()方法
		*不同的子struct组合，实现接口builder
*/
// https://github.com/lee501/go-patterns/blob/master/01-builder-patterns/builder.go
