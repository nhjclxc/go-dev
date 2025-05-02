package main

import (
	"beego_05_orm/orm_07_QueryBuilder/model"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/beego/beego/v2/client/orm"
	"log"

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

	// 模拟mybatis的动态条件拼接
	// 在 Go 中使用 SQL 构造器（如 Squirrel）可以非常优雅地实现“动态拼接 SQL 条件”，效果类似 MyBatis 的 <if> 标签。你当前的写法就类似 Squirrel 语法。
	// go get github.com/Masterminds/squirrel

	// 模拟前端传进来的条件
	age := 18
	includeEmail := true

	// 参数列表
	var args []any = []any{}

	// 查询动态列
	columns := []string{"tu.user_id", "tu.name", "tuc.card_type_id"}

	// 构造sql
	qb := sq.Select(columns...).
		From("tab_user AS tu").
		LeftJoin("tab_user_card AS tuc ON tu.user_id = tuc.user_id").
		OrderBy("name DESC").
		Limit(10).
		Offset(0)

	// 动态条件拼接
	if age > 0 {
		qb = qb.Where("age > ?", age)
		// 拼接sql的同时拼接参数列表
		args = append(args, age)
	}
	if includeEmail {
		qb = qb.Where("email IS NOT NULL")
	}

	sqlStr, args, err := qb.ToSql()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQL:", sqlStr)
	fmt.Println("ARGS:", args)



	// 执行 SQL 语句
	o := orm.NewOrm()
	var users []model.TabUser = []model.TabUser{}
	// args 就是sql的参数
	rows, err := o.Raw(sqlStr, args).QueryRows(&users)

	fmt.Println("rows = ", rows)


	for _, user := range users {
		fmt.Println(user)
	}


}

type UserQueryParams struct {
	MinAge int
	NeedCardType bool
	OrderBy string
	Limit uint64
	Offset uint64
}

func buildUserQuery(p UserQueryParams) sq.SelectBuilder {
	q := sq.Select("tu.user_id", "tu.name", "tuc.card_type_id").
		From("tab_user AS tu").
		LeftJoin("tab_user_card AS tuc ON tu.user_id = tuc.user_id").
		Where("age > ?", p.MinAge).
		Where("email IS NOT NULL")

	if p.NeedCardType {
		q = q.Where("tuc.card_type_id IS NOT NULL")
	}

	if p.OrderBy != "" {
		q = q.OrderBy(p.OrderBy + " DESC")
	}

	q = q.Limit(p.Limit).Offset(p.Offset)
	return q
}