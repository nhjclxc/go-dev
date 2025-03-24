package service

import (
	"encoding/json"
	"fmt"
	"go-dev/src/test5_test/test512_oop/test512_05/model"
	"reflect"
)

// 必须定义一个 CustomerService 结构体
// 该 CustomerService 结构体就类似于 Java 的类一样，要把这个业务层的全局数据放到这个结构体里面，我们才能去管理数据
// 下面实现的方法再去绑定这个结构体，就实现了下面的方法对这个结构体进行操作了【自己好好体会体会】
type CustomerService struct {
	// 定义所有客户的切片字段
	allCustomerSlice []model.Customer
}

// 添加一个工厂方法，用于初始化 结构体数据
func NewCustomerService() *CustomerService {
	var temp *CustomerService = &CustomerService{}
	temp.allCustomerSlice = make([]model.Customer, 0) // 初始化切片，避免后续使用
	return temp
}

// Insert 1、实现新增
func (this *CustomerService) Insert(cus *model.Customer) int {
	if nil == this.allCustomerSlice {
		this.allCustomerSlice = make([]model.Customer, 0)
	}

	// 模拟做一些数据校验
	if (*cus).Age < 18 {
		panic("小孩，要成年人才可以成为我的客户哦！！！")
	}

	this.allCustomerSlice = append(this.allCustomerSlice, *cus)
	return 1
}

// Update 2、实现修改
func (this *CustomerService) Update(updates model.Customer) int {
	if nil == this.allCustomerSlice {
		this.allCustomerSlice = make([]model.Customer, 8)
	}

	// 模拟做一些数据校验
	if updates.Id == 0 {
		panic("缺少用户id")
	}

	customerDBP := this.FindById(updates.Id)

	if nil == customerDBP {
		panic("未找到该id对应的用户")
	}

	this.updateCustomerFields(customerDBP, updates)

	return 1
}

// 更新客户信息，
func (this *CustomerService) updateCustomerFields(customerDBP *model.Customer, updates model.Customer) {
	if "" != updates.Name {
		customerDBP.Name = updates.Name
	}
	if "" != updates.Sex {
		customerDBP.Sex = updates.Sex
	}
	if 0 != updates.Age {
		customerDBP.Age = updates.Age
	}
	if "" != updates.Phone {
		customerDBP.Phone = updates.Phone
	}
	if "" != updates.Email {
		customerDBP.Email = updates.Email
	}
	if 0 != updates.Money {
		customerDBP.Money = updates.Money
	}
}

// 更新客户信息，自动忽略零值  （反射） 更灵活。
func updateCustomerFieldsReflect(target *model.Customer, updates model.Customer) {
	targetVal := reflect.ValueOf(target).Elem()
	updatesVal := reflect.ValueOf(updates)

	for i := 0; i < updatesVal.NumField(); i++ {
		fieldValue := updatesVal.Field(i)
		if !fieldValue.IsZero() { // 只更新非零值字段
			targetVal.Field(i).Set(fieldValue)
		}
	}
}

// 仅更新非零字段   （JSON 反序列化） 更高效。
func updateCustomerFieldsJSON(target *model.Customer, updates model.Customer) {
	data, _ := json.Marshal(updates)
	// 把json格式的 data 字符串反序列化到 target 对象里面
	err := json.Unmarshal(data, target)
	if err != nil {
		return
	}
}

// FindById 3、通过 id 寻找客户信息（返回 slice 中的实际对象指针）
func (this *CustomerService) FindById(id int) *model.Customer {
	for i := range this.allCustomerSlice {
		if this.allCustomerSlice[i].Id == id {
			// 注意：如果找到了，这里必须返回该客户的地址，如果不返回地址，在其他地方修改了的话CustomerService也不会修改
			// 因此，返回地址可有调用者控制是否修改，更灵活
			return &this.allCustomerSlice[i] // ✅ 这里返回的是切片中的真实对象指针
		}
	}

	// for _, customer := range this.allCustomerSlice {
	//    if id == customer.id {
	//        return &customer  // ❌ 这里会返回局部变量的地址，可能导致错误
	//    }
	// }
	return nil
}

// Delete 4、根据 id 删除客户信息
func (this *CustomerService) Delete(id int) *model.Customer {

	for i, _ := range this.allCustomerSlice {
		if id == this.allCustomerSlice[i].Id {
			// 对于一些异常情况的判断
			// 当前切片里面有且仅有一个元素，且该元素被匹配上了，要怎么删除？？？
			// 被找到的元素处于第一个位置？？？
			// 被找到的元素处于最后一个位置？？？
			// 执行删除
			// 最后使用三个点...将this.allCustomerSlice[i+1:]数组展开为单个数据
			this.allCustomerSlice = append(this.allCustomerSlice[0:i], this.allCustomerSlice[i+1:]...)

			// 返回删除数据的地址
			return &this.allCustomerSlice[i]
		}
	}

	panic("删除失败！！！未找到该id对应的用户")
}

// 5、打印所有数据
func (this *CustomerService) List() []model.Customer {
	//if len(this.allCustomerSlice) <= 0 {
	//	fmt.Println("当前暂无数据！！！")
	//}
	//
	//fmt.Println("\t 编号 \t 姓名 \t 性别 \t 年龄 \t 手机号 \t 邮箱 \t  余额 \t  ")
	//for _, customer := range this.allCustomerSlice {
	//	fmt.Println(customer.String())
	//}
	return this.allCustomerSlice

}

func main2() {

	cs := &CustomerService{}

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
		Name:  "张22三",
		Sex:   "22男",
		Age:   220,
		Phone: "12211",
		Email: "22222",
		Money: 32000,
	})
	fmt.Println(cs)
	cs.Delete(666)
	fmt.Println(cs)
	cs.List()

}
