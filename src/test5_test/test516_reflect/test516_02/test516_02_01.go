package main

import (
	"fmt"
	"reflect"
)

func main01() {

	// 演示：基本数据类型、infetface{}和reflact.Value进行反射的基本操作

	num := 18
	var iNum interface{} = num
	var reNumType = reflect.TypeOf(iNum)
	reNumValue := reflect.ValueOf(iNum)
	fmt.Printf("num = %v, type = %T \n", num, num)
	fmt.Printf("iNum = %v, type = %T \n", iNum, iNum)
	fmt.Printf("reNumType = %v, type = %T \n", reNumType, reNumType)
	fmt.Printf("reNumValue = %v, type = %T \n", reNumValue, reNumValue)
	//println(5 + reNumValue) // 无效运算: 5 + reNumValue(类型 untyped int 和 Value 不匹配)
	println(5 + reNumValue.Int())
	fmt.Println(reNumType.Kind())
	fmt.Println(reNumValue.Kind())

	iNum2 := reNumValue.Interface()
	num2 := iNum2.(int)
	fmt.Printf("iNum2 = %v, type = %T \n", iNum2, iNum2)
	fmt.Printf("num2 = %v, type = %T \n", num2, num2)

	// 演示：结构体数据类型Student、infetface{}和reflact.Value进行反射的基本操作

	stu := Student{
		name: "张同学",
		age:  18,
	}
	var iStu interface{} = stu
	reStu := reflect.ValueOf(iStu)
	fmt.Printf("stu = %v, type = %T \n", stu, stu)
	fmt.Printf("iStu = %v, type = %T \n", iStu, iStu)
	fmt.Printf("reStu = %v, type = %T \n", reStu, reStu)

	iStu2 := reStu.Interface()
	stu2 := iStu2.(Student)
	fmt.Printf("iStu2 = %v, type = %T \n", iStu2, iStu2)
	fmt.Printf("stu2 = %v, type = %T \n", stu2, stu2)
	stu22, ok := iStu2.(Student)
	fmt.Printf("stu22 = %v, type = %T , ok = %v \n", stu22, stu22, ok)

}

type Student struct {
	name string
	age  int
}
