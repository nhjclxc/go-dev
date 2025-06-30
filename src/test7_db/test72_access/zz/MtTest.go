package main

import "fmt"

const (
	StatusInit    = iota // 0
	StatusRunning        // 1
	StatusFailed         // 2
	StatusSuccess        // 3
)

const (
	_                 = iota      // 忽略 0
	ReadPermission    = 1 << iota // 1 << 1 = 2
	WritePermission               // 1 << 2 = 4
	ExecutePermission             // 1 << 3 = 8
)

const (
	_  = iota      // 忽略 0
	aa = 1 << iota // 1 << 1 = 2
	ss = 3         // 1 << 2 = 4
	dd = 5         // 1 << 3 = 8
)

func main() {

	//fmt.Println(aa)
	//fmt.Println(ss)
	//fmt.Println(dd)

	//for i := 0; i < 10; i++ {
	//	go func(temp int) {
	//		fmt.Println(temp)
	//	}(i)
	//}


	//defer func() {
	//	fmt.Println(1)
	//}()
	//
	//
	//
	//defer func() {
	//	fmt.Println(2)
	//}()
	//
	//
	//
	//defer func() {
	//	fmt.Println(3)
	//}()
	//
	//zz(4)
	//
	//defer func() {
	//	zz(5)
	//}()
	//
	//fmt.Println(6)
	//
	//
	//
	//defer func() {
	//	fmt.Println(7)
	//}()
	//

	// 4 6 7 5 3 2 1


	var ii interface{}= "1"

	val, ok := ii.(int)

	fmt.Println(val)
	fmt.Println(ok)

	val2, ok2 := ii.(string)

	fmt.Println(val2)
	fmt.Println(ok2)


	var m1 map[int]string

	print(m1[1])
	print(m1[2])

	m1 = make(map[int]string)
	m1[2] = "sdcv"
	print(m1[2])

	//
	//list := new([]int)
	//// 编译错误
	//// new([]int) 之后的 list 是一个未设置长度的 *[]int 类型的指针
	//// 不能对未设置长度的指针执行 append 操作。
	//list = append(list, 1)
	//fmt.Println(list)
	//
	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//// 编译错误，s2需要展开
	//s1 = append(s1, s2...)
	//fmt.Println(s1)

}

func zz(i int)  {
	fmt.Println(i)
}

func funcMui(x,y int)(sum int, z error){
	return x+y,nil
}
