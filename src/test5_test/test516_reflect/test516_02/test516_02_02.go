package main

import (
	"fmt"
	"reflect"
)

func main02() {

	// 反射修改变量的值
	test516_02_02_01()

	// 反射修改结构体的值
	test516_02_02_02()

}

func test516_02_02_02() {

	stu := Stu{
		Name: "zhangsan",
		Age:  18,
	}

	fmt.Println(stu)
	test516_02_02_02_reflect(&stu)
	fmt.Println(stu)

}

func test516_02_02_02_reflect(obj interface{}) {
	refValue := reflect.ValueOf(obj)
	refElemValue := refValue.Elem()
	nameValue := refElemValue.FieldByName("Name")
	nameValue.SetString("wangwu")
	ageValue := refElemValue.FieldByName("Age")
	ageValue.SetInt(26)

}

type Stu struct {
	Name string
	Age  int
}

func test516_02_02_01() {
	var num int = 520

	fmt.Println(num)
	test516_02_02_01_reflect(&num)
	fmt.Println(num)

}

func test516_02_02_01_reflect(obj interface{}) {

	fmt.Println(obj)
	reNumValue := reflect.ValueOf(obj)
	//fmt.Printf("reNumValue.Kind = %v", reNumValue.Kind()) // ptr66
	//reNumValue.SetInt(666)

	// https://studygolang.com/pkgdoc
	//func (v Value) Elem() Value
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；
	//如果v持有的值为nil，会返回Value零值。
	// Elem -->>> Element
	reNumValue.Elem().SetInt(666) // 表示取指针指向的值，类似于*p
	// 由于 reNumValue 指向的是一个指针，我们要改变的是指针指向的值，因此在反射里面要想改变指针指向的值，必须先从指针接口变量里面取出指针所指向值的接口变量
	// 即要从 reNumValue 中取出他里面保存的值，即reNumValue.Elem()，接着在使用SetInt去修改他的值

	//fmt.Println("12345678")
}
