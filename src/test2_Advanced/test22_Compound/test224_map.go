package main

import (
	"fmt"
)

/*
Go语言中的map(映射、字典)是一种内置的数据结构，它是一个无序的key—value对的集合，比如以身份证号作为唯一键来标识一个人的信息。

map格式为：

	map[keyType]valueType
*/
func main() {

	//test2241()

	test2242()
}

/*
*
map的常用操作
*/
func test2242() {

	var dict map[int]string = map[int]string{1: "xhangsan"}
	fmt.Println(dict)

	// 追加
	dict[2] = "lisi"
	fmt.Println(dict)
	dict[3] = "aaa"
	fmt.Println(dict)

	// 修改
	dict[2] = "lisi-aaaaa"
	fmt.Println(dict)

	// 遍历
	for key, value := range dict {
		fmt.Printf("key = %d, value = %s\n", key, value)
	}
	// 第一个返回值是key，第二个返回值是value（可省略）
	for key := range dict {
		fmt.Printf("key = %d, value = %s\n", key, dict[key])
	}
	//判断某个key所对应的value是否存在, 第一个值返回值是value(如果存在的话)，第二个值返回的是是否存在的bool
	value, ok := dict[1]
	fmt.Println("value = ", value, ", ok = ", ok) //value =  mike , ok =  true
	value11, ok11 := dict[11]
	fmt.Println("value = ", value11, ", ok = ", ok11) //value =   , ok =  true
	println(value11 == "")

	// 删除操作
	fmt.Println(dict)
	delete(dict, 3)
	fmt.Println(dict)

	println("---------------")
	// map作为函数传递参数
	fmt.Println(dict) // map[1:xhangsan 2:lisi-aaaaa]
	updateDict(dict)
	fmt.Println(dict) // map[1:xhangsan 2:666666 3:99999]
	// 由以上输出可知，golang中dict作为函数参数传递的是引用，即数据会被改变
}

func updateDict(dict222 map[int]string) {
	fmt.Println(dict222)
	for key, value := range dict222 {
		fmt.Printf("key = %d, value = %s \n", key, value)
	}
	dict222[2] = "666666"
	dict222[3] = "99999"
	fmt.Println(dict222)
}

/*
*
认识map
*/
func test2241() {
	// map格式为： map[keyType]valueType

	var m1 map[int]string
	fmt.Println(m1)        // map[]
	fmt.Println(m1 == nil) // true

	var m2 map[int]string = map[int]string{}
	var m3 = make(map[int]string)
	fmt.Println(m2)
	fmt.Println(len(m2))
	fmt.Println(m2 == nil)
	fmt.Println(m3)
	fmt.Println(len(m3))
	fmt.Println(m3 == nil)

	var map1 map[int]string = map[int]string{
		1:   "zhangsan",
		22:  "lisi",
		333: "wangwu"} // 大括号必须在最后一个元素的后面，同一行后面

	println(map1[1])
	println(map1[22])
	println(map1[333])

	/*
		在一个map里所有的键都是唯一的，而且必须是支持==和!=操作符的类型，切片、函数以及包含切片的结构类型这些类型由于具有引用语义，不能作为映射的键，使用这些类型会造成编译错误：
	*/

	// 创建一个key为int类型，value为[]int（数组活slice）类型的字典
	dict := map[int][]int{
		1: {1, 2, 3},
		2: {11, 22, 33},
		3: {111, 222, 333},
		//3: {111, 222, 333}, //映射文字中的键 3 重复,此复合文字会使用给定值分配新的 映射 实例。
	}

	fmt.Println(dict[1])
	fmt.Println(dict[2])
	fmt.Println(dict[3])

	// 注意：map是无序的，我们无法决定它的返回顺序，所以，每次打印结果的顺利有可能不同。

	for key, val := range dict {
		fmt.Println(dict[key], val)
	}

}
