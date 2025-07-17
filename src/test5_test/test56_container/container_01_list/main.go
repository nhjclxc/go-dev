package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 🧱 1. container/list：双向链表（Doubly Linked List）

	l := list.New()
	_ = l.PushBack(9)
	back8 := l.PushBack(8)
	l.PushBack(7)
	l.PushFront(1)
	front2 := l.PushFront(2)
	l.PushFront(3)
	l.PushFront("a")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍历
	}
	fmt.Println()

	l.InsertAfter("aaa", back8)
	l.InsertBefore("qwerty", front2)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍历
	}
	fmt.Println()

	l.Remove(front2) // 删除 e1

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 遍历
	}
	fmt.Println()

}
