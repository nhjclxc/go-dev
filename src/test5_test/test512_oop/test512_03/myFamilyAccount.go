package main

import (
	"fmt"
)

type record struct {
	// 当前操作的金额
	balance float64
	// 定义操作之前金额
	beforeOptBalance float64
	// 定义操作之后金额
	afterOptBalance float64
	// 定义这个支出或收入的说明
	remark string
}

// 家庭账户结构体
type MyFamilyAccount struct {
	name string
	age  int
	// 当前总计的金额
	balance    float64
	recordList []record
}

// 收支明细
func (this *MyFamilyAccount) showRecordList() {
	fmt.Println("-----------------收支明细-----------------")
	fmt.Printf("我是：%s，今年 %d 岁了，我的最余额是：%f，我的消费明细如下：\n", this.name, this.age, this.balance)
	if 0 != len(this.recordList) {
		for i, r := range this.recordList {
			fmt.Printf("我的 %d 笔收入指出明细是：%v\n", i+1, r)
		}
	} else {
		fmt.Println("嘿嘿，我目前还没有消费记录哦！")
	}
	fmt.Println("=================收支明细=================")

}

// 登记收入
func (this *MyFamilyAccount) income(r record) {
	this.recordList = append(this.recordList, r)
}

// 登记支出
func (this *MyFamilyAccount) outcome(r record) {
	this.recordList = append(this.recordList, r)
}

// 获取当前余额
func (this *MyFamilyAccount) getLastTimeBalance() float64 {
	length := len(this.recordList)
	if 0 == length {
		return 0.00
	}
	item := this.recordList[length-1]
	return item.afterOptBalance

	// 或者直接获取MyFamilyAccount.balance
}

func main() {

	account := &MyFamilyAccount{
		name:       "zhangsan",
		age:        18,
		recordList: []record{},
	}

	fmt.Printf("%v 初始化账户完成！！！\n", account)

	var exitFlag bool = false
	// 先定义一个死循环来不断输出
	for !exitFlag {
		// 1. 先输出这个主菜单
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("			1 收支明细")
		fmt.Println("			2 登记收入")
		fmt.Println("			3 登记支出")
		fmt.Println("			4 退出")
		fmt.Print("请选择(1-4):")

		var opt int
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			// 1 收支明细
			account.showRecordList()
		case 2:
			// 2 登记收入
			doIncome(account)
		case 3:
			// 3 登记支出
			doIncome(account)

		case 4:
			exitFlag = exit()
		default:
			fmt.Println("请输入正确的操作类型！！！")
		}
	}
}

func doIncome(account *MyFamilyAccount) {

	var item record = record{}
	fmt.Print("本次的金额：")
	//fmt.Scanf("%f\n", &item.balance)
	fmt.Scanln(&item.balance)
	fmt.Print("本次的一个备注：")
	//fmt.Scanf("%s", &item.remark)
	fmt.Scanln(&item.remark)
	lastTimeBalance := account.getLastTimeBalance()
	item.afterOptBalance = lastTimeBalance
	item.beforeOptBalance = item.balance + lastTimeBalance

	account.income(item)
}

func doOutcome(account *MyFamilyAccount) {

	var item record = record{}
	fmt.Print("本次的金额：")
	//fmt.Scanf("%f\n", &item.balance)
	fmt.Scanln(&item.balance)
	fmt.Print("本次的一个备注：")
	//fmt.Scanf("%s", &item.remark)
	fmt.Scanln(&item.remark)
	lastTimeBalance := account.getLastTimeBalance()
	item.afterOptBalance = lastTimeBalance
	item.beforeOptBalance = lastTimeBalance - item.balance

	account.income(item)
}

// 退出
func exit() bool {
	fmt.Print("确定要退出吗？(输入Y或y直接退出)：")
	var opt rune

	fmt.Scanf("%c", &opt)
	//if int(opt) == int('Y') || int(opt) == int('y') {
	//	return true
	//}
	//return false

	return int(opt) == int('Y') || int(opt) == int('y')
}
