package main

import (
	"fmt"
	"testing"
)

func TestDedupSlice(t *testing.T) {

	fmt.Println(DedupSlice([]string{"a", "b", "a", "c", "b"})) // [a b c]
	fmt.Println(DedupSlice1([]string{"a", "b", "a", "c", "b"})) // [a b c]

}

func TestCountFrequency(t *testing.T) {

//输入: []int{1, 2, 2, 3, 1, 1}
//输出: map[int]int{1: 3, 2: 2, 3: 1}

	fmt.Println(CountFrequency([]int{1, 2, 2, 3, 1, 1}))


}

func TestFilterInts(t *testing.T) {

	//FilterInts([]int{1, 2, 3, 4, 5}, func(i int) bool {
	//	return i%2 == 0
	//})
	//输出: []int{2, 4}

	fmt.Println(FilterInts([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i%2 == 0
	}))


}

func TestIntersect(t *testing.T) {


//输入: a = []int{1, 2, 3, 4}, b = []int{3, 4, 5, 6}
//输出: []int{3, 4}


	fmt.Println(Intersect([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
	fmt.Println(Union([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))

}