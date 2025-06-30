package main

import "fmt"

// 使用slice结构实现栈stack结构的push和pop
type Stack struct {
	data []int
}

func (receiver *Stack) push(item int) {
	//if nil == receiver.data {
	//	receiver.data = make([]int, 1)
	//}
	receiver.data = append(receiver.data, item)
}

func (receiver *Stack) pop() int {
	if nil == receiver.data || len(receiver.data) == 0 {
		panic("索引越界！！！")
	}

	// 获取栈顶元素
	top := receiver.data[len(receiver.data) - 1]

	// 移除栈顶数据
	receiver.data = receiver.data[:len(receiver.data) - 1]

	return top
}


func main01() {

	// 初始化一个栈
	//s := &Stack{data: make([]int,1)}
	s := &Stack{}

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)

	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())

}
