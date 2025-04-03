/*
在 package model 上面定义这个model包的包文档

	用于解释说明该包的作用
*/
package model

import "fmt"

// service层使用映射 map[int]Customer 实现，其中int表示Id，Customer表示这个id对应的客户实例数据

// Customer 客户结构体
type Customer struct {
	Id    int
	Name  string
	Sex   string
	Age   int
	Phone string
	Email string
	Money int
}

// 在 Go 语言中，工厂方法通常返回指针，而不是普通变量
// 1. 避免值拷贝，提高性能
// 2. 允许结构体方法修改对象
// 3. 适用于 nil 判断
// 4. 与 sync.Pool、接口实现兼容
// 返回指针 (*T) → 避免拷贝，提高性能，可修改，可返回 nil

// Customer 工厂方法
func NewCustomer(id int, name string, sex string, age int,
	phone string, email string, money int) Customer {
	return Customer{
		Id:    id,
		Name:  name,
		Sex:   sex,
		Age:   age,
		Phone: phone,
		Email: email,
		Money: money,
	}
}

func (this *Customer) String() string {
	return fmt.Sprintf("\t %d \t %s \t %s \t %d \t %s \t %s \t  %d \t  ",
		this.Id, this.Name, this.Sex, this.Age, this.Phone, this.Email, this.Money)
}

func main1() {
	c := Customer{
		Id:    1,
		Name:  "aaa",
		Sex:   "zz",
		Age:   0,
		Phone: "aaaa",
		Email: "aaaaaaa",
		Money: 666,
	}
	fmt.Println(c.String())

	cus := NewCustomer(111, "张安", "buzd ", 16, "123", "aaa", 123)
	fmt.Println(cus)
}
