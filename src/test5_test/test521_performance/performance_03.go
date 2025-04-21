package main

import "fmt"

// / 18.5 接口
// https://denganliang.github.io/the-way-to-go_ZH_CN/18.5.html
func main() {
	/*
	   	（1）如何检测一个值 v 是否实现了接口 Stringer：

	      if v, ok := v.(Stringer); ok {
	          fmt.Printf("implements String(): %s\n", v.String())
	      }
	      （2）如何使用接口实现一个类型分类函数：

	      func classifier(items ...interface{}) {
	          for i, x := range items {
	              switch x.(type) {
	              case bool:
	                  fmt.Printf("param #%d is a bool\n", i)
	              case float64:
	                  fmt.Printf("param #%d is a float64\n", i)
	              case int, int64:
	                  fmt.Printf("param #%d is an int\n", i)
	              case nil:
	                  fmt.Printf("param #%d is nil\n", i)
	              case string:
	                  fmt.Printf("param #%d is a string\n", i)
	              default:
	                  fmt.Printf("param #%d’s type is unknown\n", i)
	              }
	          }
	      }
	*/

	var c Usb22    // 定义一个接口类型
	c = Camera22{} // 使用多态初始化一个接口对象

	// 接口断言，与类型断言类似
	// 这个方法就是类似Java里面的instanceof，用于判断某个对象是不是某个接口类型，是的话调用特定方法
	if v, ok := c.(Usb22); ok {
		fmt.Printf("implements String(): ", 111)
		v.Start22()
	}

}

// 声明/定义一个接口
type Usb22 interface {
	//声明了两个没有实现的方法
	Start22()
}

type Camera22 struct {
	name string
}

func (c Camera22) Start22() {
	fmt.Println("相机开始工作。。。")
}
