package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"

	// don't forget this， 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

// TabUser -
type TabUser struct {
	Id           int       `orm:"column(user_id);auto"`
	Name         string    `orm:"column(name);size(255)"`
	Email        string    `orm:"column(email);size(255);null"`
	Age          int8      `orm:"column(age)"`
	Birthday     time.Time `orm:"column(birthday);type(datetime);null"`
	MemberNumber string    `orm:"column(member_number);size(255);null"`
	Remark       string    `orm:"column(remark);size(128);null"`
	ActivatedAt  time.Time `orm:"column(activated_at);type(datetime);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);auto_now_add"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(TabUser))

	// need to register default database
	err := orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return
	}
}

func main() {
	// 开启调试模式
	orm.Debug = true

	// 创建不存在的表结构，如果存在则跳过
	// 存在则输出：table tab_user already exists, skip
	// automatically build table
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		return
	}

	// create orm object
	ormObj := orm.NewOrm()

	// data
	user := new(TabUser)
	user.Name = "mike"

	// insert data
	// [ORM]2025/04/29 21:36:32  -[Queries/default] - [  OK /     db.Exec /    27.7ms] -
	//[INSERT INTO `tab_user` (`name`, `email`, `age`, `birthday`, `member_number`, `remark`, `activated_at`, `created_at`, `updated_at`)
	//VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)] - `mike`, ``, `0`, `<nil>`, ``, ``, `<nil>`, `2025-04-29 21:36:32.2431918 +0800 CST`, `2025-04-29 21:36:32.2431918 +0800 CST`
	insert, err := ormObj.Insert(user)
	if err != nil {
		return
	}

	fmt.Println("insert = ", insert)

	//总体来说，可以分成以下几步：
	//	定义模型，并且注册，参考模型定义
	//	注册 DB，参考数据库注册
	//	创建 Orm 实例
	//	执行查询，
}
