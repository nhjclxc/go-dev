package service

import (
	"fmt"
	"go-dev/src/test5_test/test512_oop/test512_05/model"
	"testing"
)

var cs *CustomerService

func init() {
	cs = &CustomerService{}
}

func TestCustomerService_Insert(t *testing.T) {
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
}

func TestCustomerService_Update(t *testing.T) {
	fmt.Println(cs)
	cs.Update(model.Customer{
		Id:    1,
		Name:  "张22三",
		Sex:   "22男",
		Age:   220,
		Phone: "12211",
		Email: "22222",
		Money: 32000,
	})
	fmt.Println(cs)
}

func TestCustomerService_FindById(t *testing.T) {
	fmt.Println(cs.FindById(666))
}

func TestCustomerService_Delete(t *testing.T) {
	//cs.Delete(1)
}

func TestCustomerService_List(t *testing.T) {
	cs.List()
}
