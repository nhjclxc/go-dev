package main

import (
	"encoding/json"
	"fmt"
	"go-dev/src/test5_test/test512_oop/test512_06/model"
	"sort"
)

// CustomerService 定义结构体对象
type CustomerService struct {
	customerMap map[int]*model.Customer
	autoId      int
}

// 添加一个工厂方法，用于初始化 结构体数据
func NewCustomerService() (customerService *CustomerService) {
	return &CustomerService{make(map[int]*model.Customer, 0), 0}
}

// Insert 1、实现新增
func (this *CustomerService) Insert(name string, sex string, age int,
	phone string, email string, money int) int {
	if age < 18 {
		panic("小孩啊，我这个是18🈲哦")
	}
	this.autoId++
	customer := model.NewCustomer(this.autoId, name, sex, age, phone, email, money)
	this.customerMap[this.autoId] = &customer
	return 1
}

// Update 2、实现修改
func (this *CustomerService) Update(customer model.Customer) int {
	if customer.Id == 0 {
		panic("请输入id")
	}
	customerDB := this.customerMap[customer.Id]
	updateCustomerFields(customerDB, customer)

	return 1
}

// 执行更新客户信息，仅更新非零字段   （JSON 反序列化） 更高效。
func updateCustomerFields(target *model.Customer, updates model.Customer) {
	data, _ := json.Marshal(updates)
	// 把json格式的 data 字符串反序列化到 target 对象里面
	err := json.Unmarshal(data, target)
	if err != nil {
		return
	}
}

// FindById 3、通过 id 寻找客户信息（返回 slice 中的实际对象指针）
func (this *CustomerService) FindById(id int) *model.Customer {
	if id == 0 {
		panic("id不能为0")
	}
	return this.customerMap[id]
}

// Delete 4、根据 id 删除客户信息
func (this *CustomerService) Delete(id int) {
	if id == 0 {
		panic("id不能为0")
	}
	delete(this.customerMap, id)
}

// 5、打印所有数据
func (this *CustomerService) List() {
	fmt.Println("------------ Start -------------")
	for _, customer := range this.customerMap {
		fmt.Println(customer)
	}
	fmt.Println("------------  End  -------------")
}

// 实现有序的map遍历
func (this *CustomerService) SortListPrint() {

	fmt.Println(len(this.customerMap))
	// 1、获取所有的key
	keyList := make([]int, len(this.customerMap))

	// 遍历获取所有key，只接收一个传输的时候，默认返回key
	for key := range this.customerMap {
		keyList = append(keyList, key)
	}

	// 2、调用内部的排序sort方法，对key进行排序
	sort.Ints(keyList)

	// 3、按照keyList的顺序去除数据
	fmt.Println("------------ Start -------------")
	for _, key := range keyList {
		fmt.Println(this.customerMap[key])
	}
	fmt.Println("------------  End  -------------")

}
