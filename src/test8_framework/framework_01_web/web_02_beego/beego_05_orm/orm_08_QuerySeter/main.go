package main

import (
	"beego_05_orm/orm_08_QuerySeter/model"
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

// ORM 以 QuerySeter 来组织查询，每个返回 QuerySeter 的方法都会获得一个新的 QuerySeter 对象。
// https://beegodoc.com/zh/developing/orm/query_seter.html
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

	//ORM 以 QuerySeter 来组织查询，每个返回 QuerySeter 的方法都会获得一个新的 QuerySeter 对象。
	//基本使用方法:
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，anonymous_user 为表名
	qs := o.QueryTable("tab_user")
	//qs = o.QueryTable(anonymous_user.TableName()) // 返回 QuerySeter

	// 也可以直接使用 Model 结构体作为表名
	qs = o.QueryTable(&model.TabUser{})

	// 也可以直接使用对象作为表名
	user := new(model.TabUser)
	qs = o.QueryTable(user) // 返回 QuerySeter

	// 后面可以调用qs上的方法，执行复杂查询。
	fmt.Println(qs)

	// 1、查询表达式
	//test080101(qs)

	// 2、中间方法
	//test080102(qs)

}

func test080102(qs orm.QuerySeter) {
	// 2、中间方法

	// 2.1、Filter(string, ...interface{}) QuerySeter
	// 多次调用Filter方法，会使用AND将它们连起来。

	qs.Filter("profile__isnull", true).Filter("name", "slene")
	// WHERE profile_id IS NULL AND name = 'slene'

	// 2.2、FilterRaw，不推荐使用，会出现sql注入威胁
	// FilterRaw(string, string) QuerySeter
	// 该方法会直接把输入当做是一个查询条件，因此如果输入有错误，那么拼接得来的 SQL 则无法运行。Beego 本身并不会执行任何的检查。
	//  第一个参数是字段名，不加 __in 等修饰符。
	//第二个参数是你写好的 SQL 条件字符串（可以带括号、运算符等）。
	//注意：Beego ORM 不会做 SQL 注入保护，你需要确保传入的 SQL 条件是安全的。

	qs.FilterRaw("user_id", " IN (SELECT id FROM profile WHERE age>=18)")
	//sql-> WHERE user_id IN (SELECT id FROM profile WHERE age>=18)

	// 2.3、Exclude，排除什么条件
	//Exclude(string, ...interface{}) QuerySeter
	//准确来说，Exclude表达的是NOT的语义：

	qs.Filter("profile__age__in", 18, 20).Exclude("profile__age__in", 1000)
	// WHERE profile.age IN (18, 20) AND profile.age NOT IN (1000)

	// 2.4、SetCond
	//SetCond(*Condition) QuerySeter
	//设置查询条件：Condition中使用的表达式，可以参考查询表达式

	cond := orm.NewCondition()
	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	////sql-> WHERE T0.`profile_id` IS NOT NULL AND NOT T0.`Status` IN (?) OR T1.`age` >  2000
	num, err := qs.SetCond(cond1).Count()
	fmt.Println(num, err)

	// 2.5、

	// 2.6、

	// 2.7、

	// 2.8、

}

func test080101(qs orm.QuerySeter) {
	// 1、查询表达式
	qs.Filter("id", 1) // WHERE id = 1
	// 关联表
	// 其中tab_user_card表示哪个表的字段要进行限制，
	// age 表示具体的字段，
	// 18 表示具体的值
	// gt 表示大于号
	// __ 字段的分隔符号使用双下划线 __ ，可以认为表示一个空格
	qs.Filter("tab_user_card__age", 18)     // WHERE profile.age = 18
	qs.Filter("tab_user_card__Age", 18)     // 使用字段名和 Field 名都是允许的
	qs.Filter("tab_user_card__age__gt", 18) // WHERE profile.age > 18

	// 表达式的尾部可以增加操作符以执行对应的 sql 操作。
	// 当前支持的操作符号：
	//    exact / iexact 等于
	//    contains / icontains 包含
	//    gt / gte 大于 / 大于等于
	//    lt / lte 小于 / 小于等于
	//    startswith / istartswith 以...起始
	//    endswith / iendswith 以...结束
	//    in
	//    isnull
	//    后面以 i 开头的表示：大小写不敏感

	//1.1、exact
	//Filter / Exclude / Condition expr 的默认值
	qs.Filter("name", "slene")        // WHERE name = 'slene'
	qs.Filter("name__exact", "slene") // WHERE name = 'slene'
	// 使用 = 匹配，大小写是否敏感取决于数据表使用的 collation
	qs.Filter("profile_id", nil) // WHERE profile_id IS NULL

	//1.2、iexact
	qs.Filter("name__iexact", "slene")
	// WHERE name LIKE 'slene'
	// 大小写不敏感，匹配任意 'Slene' 'sLENE'

	//1.3、contains
	qs.Filter("name__contains", "slene")
	// WHERE name LIKE BINARY '%slene%'
	// 大小写敏感, 匹配包含 slene 的字符

	//1.4、icontains
	qs.Filter("name__icontains", "slene")
	// WHERE name LIKE '%slene%'
	// 大小写不敏感, 匹配任意 'im Slene', 'im sLENE'

	//1.5、in
	qs.Filter("age__in", 17, 18, 19, 20)
	// WHERE age IN (17, 18, 19, 20)

	ids := []int{17, 18, 19, 20}
	qs.Filter("age__in", ids)
	qs.Filter("age__in", ids)
	// WHERE age IN (17, 18, 19, 20)

	//1.6、gt / gte
	qs.Filter("profile__age__gt", 17)
	// WHERE profile.age > 17

	qs.Filter("profile__age__gte", 18)
	// WHERE profile.age >= 18

	//1.7、lt / lte
	qs.Filter("profile__age__lt", 17)
	// WHERE profile.age < 17

	qs.Filter("profile__age__lte", 18)
	// WHERE profile.age <= 18

	//1.8、startswith，istartswith，endswith，iendswith，isnull
	qs.Filter("name__startswith", "slene")
	// WHERE name LIKE BINARY 'slene%'
	// 大小写敏感, 匹配以 'slene' 起始的字符串

	qs.Filter("name__istartswith", "slene")
	// WHERE name LIKE 'slene%'
	// 大小写不敏感, 匹配任意以 'slene', 'Slene' 起始的字符串

	qs.Filter("name__endswith", "slene")
	// WHERE name LIKE BINARY '%slene'
	// 大小写敏感, 匹配以 'slene' 结束的字符串

	qs.Filter("name__iendswith", "slene")
	// WHERE name LIKE '%slene'
	// 大小写不敏感, 匹配任意以 'slene', 'Slene' 结束的字符串

	qs.Filter("profile__isnull", true)
	qs.Filter("profile_id__isnull", true)
	// WHERE profile_id IS NULL

	qs.Filter("profile__isnull", false)
	// WHERE profile_id IS NOT NULL
}
