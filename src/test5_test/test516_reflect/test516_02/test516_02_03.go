package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 给定一个变量 var v float64 = 1.2，请使用反射得到他的reflact.Value
	// 然后获取对应的Type、Kind和具体值，并将reflact.Value转化为infterface{}，再将infterface{}转为float64

	var v float64 = 1.2
	fmt.Println(v)

	var obj interface{} = v
	fmt.Println(obj)

	refVlaue := reflect.ValueOf(obj)
	fmt.Println(refVlaue)

	refType := refVlaue.Type()
	fmt.Println(refType)

	refKind := refVlaue.Kind()
	fmt.Println(refKind)

	var refObj interface{} = refVlaue.Interface()
	fmt.Println(refObj)

	num := refObj.(float64)
	fmt.Println(num)

	fmt.Println("---------------------------")

	str := "Hello"
	fmt.Println(str)

	var iStr interface{} = &str
	refValue := reflect.ValueOf(iStr)
	//refValue.SetString("Go!!!")
	refValue.Elem().SetString("World!!!")
	fmt.Println(str)

}
