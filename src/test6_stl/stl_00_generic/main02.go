package main

import "fmt"

// 使用slice结构实现栈stack结构的push和pop
// 注意：这里在main01的基础上硬入泛型
type StackG[T any] struct {
	data []T
}

func (receiver *StackG[T]) push(item T) {
	//if nil == receiver.data {
	//	receiver.data = make([]int, 1)
	//}
	receiver.data = append(receiver.data, item)
}

func (receiver *StackG[T]) pop() T {
	if nil == receiver.data || len(receiver.data) == 0 {
		panic("索引越界！！！")
	}

	// 获取栈顶元素
	top := receiver.data[len(receiver.data) - 1]

	// 移除栈顶数据
	receiver.data = receiver.data[:len(receiver.data) - 1]

	return top
}


func main() {

	// 初始化一个栈
	//s := &Stack{data: make([]int,1)}
	s := &StackG[int]{}

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



	s2 := &StackG[float32]{}

	s2.push(1.369)
	s2.push(2.369)
	s2.push(3.369)
	s2.push(4.369)
	s2.push(5.369)

	fmt.Println(s2.pop())
	fmt.Println(s2.pop())
	fmt.Println(s2.pop())
	fmt.Println(s2.pop())
	fmt.Println(s2.pop())

}
