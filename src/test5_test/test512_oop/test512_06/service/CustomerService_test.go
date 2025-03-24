package main

import (
	"go-dev/src/test5_test/test512_oop/test512_06/model"
	"testing"
)

var customerService *CustomerService

func init() {
	customerService = NewCustomerService()
}

func TestCustomerService_Insert(t *testing.T) {
	customerService.List()
	customerService.Insert("张三", "男", 19, "12345678", "qwerty", 123)
	customerService.Insert("张三22", "男22", 22, "22212345678", "22q2werty", 22123)
	customerService.Insert("张三33", "男33", 23, "33312345678", "33qwerty", 33123)
	customerService.List()
}

func TestCustomerService_Update(t *testing.T) {
	// 初始化插入操作
	TestCustomerService_Insert(nil)
	customer := model.NewCustomer(2, "里斯", "女", 28, "qqqqq", "azsxdc", 666)

	customerService.SortListPrint()
	customerService.Update(customer)
	customerService.SortListPrint()
}
