package main

import (
	"fmt"
	"testing"
)

// map的遍历是无须的
func Test21(t *testing.T) {

	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}

}

//  模仿Java的 LinkedHashMap 实现go中map的有序遍历

type LinkedHashMap struct {
	keys []string
	data map[string]int
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		keys: make([]string, 0),
		data: make(map[string]int),
	}
}

func (m *LinkedHashMap) Put(key string, val int) {
	if _, ok := m.data[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.data[key] = val
}

func (m *LinkedHashMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *LinkedHashMap) Delete(key string) (int, bool) {
	for i := 0; i < len(m.keys); i++ {
		if m.keys[i] == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			delete(m.data, key)
			return i, true
		}
	}
	return -1, false
}

func (m *LinkedHashMap) Iterate() {
	for _, key := range m.keys {
		fmt.Printf("%d, ", m.data[key])
	}
	fmt.Println()
}

func Test212(t *testing.T) {
	var hashMap *LinkedHashMap = NewLinkedHashMap()
	hashMap.Put("qq", 11)
	hashMap.Put("ww", 22)
	hashMap.Put("ee", 33)
	hashMap.Put("rr", 55)
	hashMap.Put("tt", 66)
	hashMap.Put("tt", 98)
	hashMap.Iterate()
	hashMap.Delete("tt")
	hashMap.Iterate()
	hashMap.Delete("qq")
	hashMap.Iterate()

}

//
//// LinkedHashMap 模拟
//type LinkedHashMap struct {
//	keys []string
//	data map[string]int
//}
//
//func NewLinkedHashMap() *LinkedHashMap {
//	return &LinkedHashMap{
//		keys: []string{},
//		data: make(map[string]int),
//	}
//}
//
//func (m *LinkedHashMap) Put(key string, value int) {
//	if _, exists := m.data[key]; !exists {
//		m.keys = append(m.keys, key) // 只在新 key 时记录顺序
//	}
//	m.data[key] = value
//}
//
//func (m *LinkedHashMap) Get(key string) (int, bool) {
//	v, ok := m.data[key]
//	return v, ok
//}
//
//func (m *LinkedHashMap) Iterate() {
//	for _, k := range m.keys {
//		fmt.Println(k, m.data[k])
//	}
//}
//
//func main() {
//	lhm := NewLinkedHashMap()
//	lhm.Put("one", 1)
//	lhm.Put("two", 2)
//	lhm.Put("three", 3)
//	lhm.Put("four", 4)
//	lhm.Put("two", 22) // 更新不会改变顺序
//
//	lhm.Iterate()
//}
