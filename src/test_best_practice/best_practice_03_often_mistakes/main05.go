package main

import "fmt"

func main05() {

	// 当切片作为参数传递时，切记不要解引用切片。
	// 我们已经知道，切片实际是一个指向潜在数组的指针。
	//我们常常需要把切片作为一个参数传递给函数是因为：实际就是传递一个指向变量的指针，在函数内可以改变这个变量，而不是传递数据的拷贝。

	sli := make([]int, 8)
	sli[0] = 11
	sli[1] = 22
	sli[2] = 33
	sli[3] = 44
	sli[4] = 55
	fmt.Println(sli, cap(sli), len(sli))
	dealSilce(sli)
	fmt.Println(sli, cap(sli), len(sli))
	fmt.Println(sli, cap(sli), len(sli))
	dealSilce2(&sli)
	fmt.Println(sli, cap(sli), len(sli))
	fmt.Println("----------------------")

	var arr []int = []int{1, 2, 3, 4, 5}
	fmt.Println(arr, cap(arr), len(arr))
	dealArray(arr)
	fmt.Println(arr, cap(arr), len(arr))
	fmt.Println(arr, cap(arr), len(arr))
	dealArray2(&arr)
	fmt.Println(arr, cap(arr), len(arr))

}

// 直接传切片
func dealSilce(sli []int) {
	sli[5] = 666
}

// 传切片的指针
func dealSilce2(sli *[]int) {
	(*sli)[6] = 888
}

// 直接传数组
func dealArray(arr []int) {
	arr[2] = 666
}

// 传数组的指针
func dealArray2(arr *[]int) {
	(*arr)[3] = 888
}
