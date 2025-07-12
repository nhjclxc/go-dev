package main


/*
实现一个结构体 `StringSet`，内部使用 `map[string]struct{}` 实现集合，支持以下操作：

* `Add(s string)`
* `Remove(s string)`
* `Contains(s string) bool`
* `Values() []string`
* `Len() int`

> ✅ 要求：集合中不能有重复值。

 */

type StringSet struct {
	// 集合中只关心元素是否存在，不需要占用额外的空间存储 true/false，使用 struct{} 占用内存为 0，更节省资源。
	data map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{
		data: make(map[string]struct{}),
	}
}

func (ss *StringSet) Add(s string)  {
	ss.data[s] = struct{}{}
}

func (ss *StringSet) Remove(s string)  {
	delete(ss.data, s)
}

func (ss *StringSet) Contains(s string) bool {
	_, exists := ss.data[s]
	return exists
}

func (ss *StringSet) Values() (res []string) {
	res = make([]string, 0, len(ss.data))
	for key, _ := range ss.data {
		res = append(res, key)
	}
	return
}

func (ss *StringSet) Len() int {
	return len(ss.data)
}








