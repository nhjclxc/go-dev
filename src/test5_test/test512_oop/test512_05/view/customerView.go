package main

import (
	"fmt"
	"go-dev/src/test5_test/test512_oop/test512_05/model"
	"go-dev/src/test5_test/test512_oop/test512_05/service"
)

// 定义视图层的结构体，用于保存 CustomerService 的实例对象
type CustomerView struct {
	// 定义 CustomerService 的实例变量，类似于Java里面在Controller引入一个Service层的实例对象
	customerService service.CustomerService
}

// 客户信息管理系统
// customer
func main() {

	fmt.Println(111)

	cs := &service.CustomerService{}

	defer func() {
		if err := recover(); nil != err {
			fmt.Println("程序异常：", err)
		}
	}()

	fmt.Println(cs)
	cs.Insert(&model.Customer{
		Id:    1,
		Name:  "张三",
		Sex:   "男",
		Age:   21,
		Phone: "111",
		Email: "222",
		Money: 3000,
	})
	fmt.Println(cs)
	cs.Update(model.Customer{
		Id:    1,
		Name:  "张222三",
		Sex:   "22男",
		Age:   22,
		Phone: "12211",
		Email: "222222",
		Money: 32000,
	})
	fmt.Println(cs)
	//cs.Delete(666)
	fmt.Println(cs)
	cs.List()

}
