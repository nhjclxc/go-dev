package main

import (
	"fmt"
	"sort"
)

// 接口的最佳实践
// https://www.bilibili.com/video/BV1ME411Y71o?p=220
// https://studygolang.com/pkgdoc
// 实现的对Hero结构体切片的排序：sort.Sort(data Interface)
/*
func Sort
func Sort(data Interface)
Sort排序data。它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap。本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）。

type Interface
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
一个满足sort.Interface接口的（集合）类型可以被本包的函数进行排序。方法要求集合中的元素可以被整数索引。
*/

// 1、定义 Hero 结构体
type Hero struct {
	HeroScore int
}

// 2、定义 Hero 结构体的切片类型
type HeroSlice []Hero

// 3、实现接口的三个方法
// Len方法返回集合中的元素个数
// Len() int
func (self HeroSlice) Len() int {
	return len(self)
}

// Less方法报告索引i的元素是否比索引j的元素小
// Less(i, j int) bool
func (self HeroSlice) Less(i, j int) bool {
	// self[i]拿出了一个Hero结构体对象
	//var h Hero = self[i]
	//fmt.Println(h)

	// 通过 Hero 对象的英雄值大小来比较
	//if self[i].HeroScore < self[j].HeroScore {
	//	return true
	//}
	//return false

	// 优化写法
	return self[i].HeroScore < self[j].HeroScore
}

// Swap方法交换索引i和j的两个元素
// Swap(i, j int)
func (self HeroSlice) Swap(i, j int) {
	temp := self[i]
	self[i] = self[j]
	self[j] = temp
}

func main5() {

	// 首先了解一下 sort.Sort(data Interface) 方法
	var arr []int = []int{10, 2, 5, 8, 6, 9}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	// 对 sort.Sort(data Interface) 的实现
	var arrHero HeroSlice = HeroSlice{
		{HeroScore: 10},
		{HeroScore: 2},
		{HeroScore: 5},
		{HeroScore: 8},
		{HeroScore: 6},
		{HeroScore: 9},
	}
	fmt.Println(arrHero)
	sort.Sort(arrHero) // 方法接收的 data 必须是一个切片类型
	fmt.Println(arrHero)

}
