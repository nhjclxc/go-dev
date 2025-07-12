package main

import (
	"fmt"
	"testing"
)

func TestIndexRange(t *testing.T) {


//输入: []int{3, 1, 2, 3, 1, 4}
//输出: map[int][2]int{1: [1 4], 2: [2 2], 3: [0 3], 4: [5 5]}

	fmt.Println(IndexRange([]int{3, 1, 2, 3, 1, 4}))

}