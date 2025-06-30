package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age int
}


// 自定义结构体排序
func main() {

	s1 := Student{ Name: "zhangsan", Age:  18, }
	s2 := Student{ Name: "lisi", Age:  25, }
	s3 := Student{ Name: "wangwu", Age:  15, }

	studentList := []Student{s1, s2, s3}

	sortDesc(studentList)
	fmt.Println(studentList)

	sortAsc(studentList)
	fmt.Println(studentList)

	sort.Find(18, func(i int) int {
		fmt.Println(i)
		return 1
	})


}

func sortDesc(studentList []Student) {
	sort.Slice(studentList, func(i, j int) bool {
		// 降序
		return studentList[i].Age > studentList[j].Age
	})
}

func sortAsc(studentList []Student) {
	sort.Slice(studentList, func(i, j int) bool {
		// 升序
		return studentList[i].Age < studentList[j].Age
	})
}
