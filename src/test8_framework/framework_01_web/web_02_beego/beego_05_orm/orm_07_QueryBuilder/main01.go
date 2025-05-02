package main

import (
	"beego_05_orm/orm_07_QueryBuilder/model"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	// don't forget this， 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)


func init() {
	// need to register models in init
	orm.RegisterModel(new(model.TabUser))
	orm.RegisterModel(new(model.TabUserCard))

	// need to register default database
	err := orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return
	}
}

// QueryBuilder 在功能上与 ORM 重合， 但是各有利弊。ORM 更适用于简单的 CRUD 操作，而 QueryBuilder 则更适用于复杂的查询，例如查询中包含子查询和多重联结。
func main01() {
	// 开启调试模式
	orm.Debug = true

	// 创建不存在的表结构，如果存在则跳过
	// 存在则输出：table tab_user already exists, skip
	// automatically build table
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		return
	}


	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")


	// 构建查询对象
	qb.Select("tu.user_id", "tu.name",
		"tuc.card_type_id").
		From("tab_user as tu").
		LeftJoin("tab_user_card as tuc").On("tu.user_id = tuc.user_id").
		Where("age > ?").
		And("email is not null").
		//And("tuc.card_type_id is not null").
		OrderBy("name").Desc().
		Limit(10).Offset(0)
	// [SELECT tu.user_id, tu.name, tuc.card_type_id FROM tab_user as tu LEFT JOIN tab_user_card as tuc ON tu.user_id = tuc.user_id WHERE age > ? ORDER BY name DESC LIMIT 10 OFFSET 0] - `2`


	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	var users []model.TabUser = []model.TabUser{}
	// args 就是sql的参数
	rows, err := o.Raw(sql, 2).QueryRows(&users)
	if err != nil {
		return
	}

	fmt.Println("rows = ", rows)

	for _, user := range users {
		fmt.Println(user)
	}


}

/*
type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	ForUpdate() QueryBuilder
	From(tables ...string) QueryBuilder
	InnerJoin(table string) QueryBuilder
	LeftJoin(table string) QueryBuilder
	RightJoin(table string) QueryBuilder
	On(cond string) QueryBuilder
	Where(cond string) QueryBuilder
	And(cond string) QueryBuilder
	Or(cond string) QueryBuilder
	In(vals ...string) QueryBuilder
	OrderBy(fields ...string) QueryBuilder
	Asc() QueryBuilder
	Desc() QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder
	GroupBy(fields ...string) QueryBuilder
	Having(cond string) QueryBuilder
	Update(tables ...string) QueryBuilder
	Set(kv ...string) QueryBuilder
	Delete(tables ...string) QueryBuilder
	InsertInto(table string, fields ...string) QueryBuilder
	Values(vals ...string) QueryBuilder
	Subquery(sub string, alias string) string
	String() string
}
 */