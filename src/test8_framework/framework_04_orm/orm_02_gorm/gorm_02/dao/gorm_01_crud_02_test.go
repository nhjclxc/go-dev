package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	db "gorm_02/config"
	"gorm_02/model"
	"testing"
	"time"
)

// 数据库 新增 Query 查询
// https://gorm.io/zh_CN/docs/query.html

// 1.2.1、检索单个对象
func TestMyFunc0201(t *testing.T) {
	// GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误

	tabUser := model.TabUser{}

	// 获取一条记录，没有指定排序字段
	db.DB.Take(&tabUser)
	// SELECT `tab_user`.`user_id`,`tab_user`.`name`,`tab_user`.`email`,`tab_user`.`age`,`tab_user`.`birthday`,`tab_user`.`member_number`,`tab_user`.`remark`,`tab_user`.`activated_at`,`tab_user`.`created_at`,`tab_user`.`updated_at` FROM `tab_user` LIMIT 1
	fmt.Println("db.DB.Take(&tabUser): ", tabUser)

	// 获取最后一条记录（主键降序）
	tabUser = model.TabUser{}
	db.DB.Last(&tabUser)
	// SELECT `tab_user`.`user_id`,`tab_user`.`name`,`tab_user`.`email`,`tab_user`.`age`,`tab_user`.`birthday`,`tab_user`.`member_number`,`tab_user`.`remark`,`tab_user`.`activated_at`,`tab_user`.`created_at`,`tab_user`.`updated_at` FROM `tab_user` ORDER BY `tab_user`.`user_id` DESC LIMIT 1
	fmt.Println("db.DB.Last(&tabUser): ", tabUser)

	// 获取第一条记录（主键升序）
	tabUser = model.TabUser{}
	result := db.DB.First(&tabUser)
	// SELECT `tab_user`.`user_id`,`tab_user`.`name`,`tab_user`.`email`,`tab_user`.`age`,`tab_user`.`birthday`,`tab_user`.`member_number`,`tab_user`.`remark`,`tab_user`.`activated_at`,`tab_user`.`created_at`,`tab_user`.`updated_at` FROM `tab_user` ORDER BY `tab_user`.`user_id` LIMIT 1
	fmt.Println("db.DB.First(&tabUser): ", tabUser)

	// 检查 ErrRecordNotFound 错误
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound)) // false
	// 如果你想避免ErrRecordNotFound错误，你可以使用Find，比如db.Limit(1).Find(&anonymous_user)，Find方法可以接受struct和slice的数据。

}

// 1.2.2、检索多个对象
func TestMyFunc0202(t *testing.T) {
	// First and Last 方法会按主键排序找到第一条记录和最后一条记录 (分别)。 只有在目标 struct 是指针或者通过 db.Model() 指定 model 时该方法才有效。
	//此外，如果相关 model 没有定义主键，那么将按 model 的第一个字段进行排序。 例如：

	var user model.TabUser
	//var users []model.TabUser

	// works because destination struct is passed in
	db.DB.First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// works because model is specified using `db.Model()`
	result1 := map[string]interface{}{}
	db.DB.Model(&model.TabUser{}).First(&result1)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	fmt.Println("result1 = ", result1)

	// doesn't work
	result2 := map[string]interface{}{}
	db.DB.Table(user.TableName()).First(&result2)

	// works with Take
	result3 := map[string]interface{}{}
	db.DB.Table(user.TableName()).Take(&result3)

	// no primary key defined, results will be ordered by first field (i.e., `Code`)
	type Language struct {
		Code string
		Name string
	}
	db.DB.First(&Language{})
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1

}

// 1.2.3、根据主键检索
func TestMyFunc0203(t *testing.T) {
	// 如果主键是数字类型，您可以使用 内联条件 来检索对象。 当使用字符串时，需要额外的注意来避免SQL注入；查看 Security 部分来了解详情。\

	var user model.TabUser
	db.DB.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.DB.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	var users []model.TabUser
	db.DB.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);
	fmt.Println("users = ", users)

	// 如果主键是字符串(例如像uuid)，查询将被写成如下：
	db.DB.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	//// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	//当目标对象有一个主键值时，将使用主键构建查询条件，例如：
	var user2 = model.TabUser{UserId: 10}
	db.DB.First(&user2)
	// SELECT * FROM users WHERE id = 10;

	var result model.TabUser
	db.DB.Model(user2).First(&result)
	// SELECT * FROM users WHERE id = 10;

	fmt.Println("---------------------")

	var user11 = model.TabUser{UserId: 10}
	db.DB.First(&user11)
	// SELECT * WHERE `tab_user`.`user_id` = 10 ORDER BY `tab_user`.`user_id` LIMIT 1

	var user22 = []model.TabUser{model.TabUser{UserId: 10}}
	db.DB.Find(&user22)
	// SELECT * FROM `tab_user` WHERE `tab_user`.`user_id` = 10

	fmt.Println("user11", user11)
	fmt.Println("user22", user22)

	// 查一条记录
	var user111 model.TabUser
	db.DB.Where("user_id = ?", 10).First(&user111)
	// SELECT * WHERE `tab_user`.`user_id` = 10 ORDER BY `tab_user`.`user_id` LIMIT 1

	// 查多条记录
	var users222 []model.TabUser
	db.DB.Where("user_id = ?", 10).Find(&users222)
	// SELECT * FROM `tab_user` WHERE `tab_user`.`user_id` = 10

}

// 1.2.4、条件 - String 条件
func TestMyFunc0204(t *testing.T) {

	var user model.TabUser
	var users []model.TabUser

	// Get first matched record
	db.DB.Where("name = ?", "WangWu").First(&user)
	// SELECT * FROM `tab_user` WHERE name = 'WangWu' ORDER BY `tab_user`.`user_id` LIMIT 1
	fmt.Println("anonymous_user：", user)

	// Get all matched records
	db.DB.Where("name <> ?", "WangWu").Find(&users)
	// SELECT * FROM `tab_user` WHERE name <> 'WangWu'
	//fmt.Println("users：", users)

	// IN
	db.DB.Where("name IN ?", []string{"WangWu", "张三"}).Find(&users)
	// SELECT * FROM `tab_user` WHERE name IN ('WangWu','张三')

	// LIKE
	db.DB.Where("name LIKE ?", "%ang%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%ang%';

	// AND
	db.DB.Where("name = ? AND age >= ?", "WangWu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'WangWu' AND age >= 22;

	// Time
	today := time.Now()
	lastWeek := time.Now().Add(-1 * 7 * 24 * time.Hour)
	db.DB.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM `tab_user` WHERE updated_at > '2025-04-19 10:24:13.593'
	fmt.Println("today = ", today.String()) // 2025-04-26 10:24:13.5934478 +0800 CST m=+0.021203701

	// BETWEEN
	db.DB.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM `tab_user` WHERE created_at BETWEEN '2025-04-19 10:24:13.593' AND '2025-04-26 10:24:13.593'

	//如果对象设置了主键，条件查询将不会覆盖主键的值，而是用 And 连接条件。 例如：

	var user22 = model.TabUser{UserId: 10}
	result := db.DB.Where("id = ?", 20).First(&user22)
	// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
	//这个查询将会给出record not found错误 所以，在你想要使用例如 anonymous_user 这样的变量从数据库中获取新值前，需要将例如 id 这样的主键设置为nil。
	// 即：将 user22 变量声明为：var user22 = model.TabUser{}
	// 检查 ErrRecordNotFound 错误
	fmt.Println("ErrRecordNotFound: ", errors.Is(result.Error, gorm.ErrRecordNotFound))

}

// 1.2.5、条件 - Struct & Map 条件
func TestMyFunc0205(t *testing.T) {
	var user model.TabUser
	var users []model.TabUser

	// 参数为结构体时，直传非0字段
	// 参数为map时，map有什么条件就查什么条件

	// Struct
	// 注意：当使用结构体进行查询时，GORM 只会使用非零字段进行查询，这意味着如果你的字段值为 0、''、false 或其他零值，它将不会用于构建查询条件，例如：
	db.DB.Where(&model.TabUser{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM `tab_user` WHERE `tab_user`.`name` = 'jinzhu' AND `tab_user`.`age` = 20 ORDER BY `tab_user`.`user_id` LIMIT 1
	fmt.Println(user)

	// Map
	db.DB.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM `tab_user` WHERE `age` = 20 AND `name` = 'jinzhu'
	fmt.Println(users)

	// 要在查询条件中包含零值，可以使用映射，它将包含所有键值作为查询条件，例如：
	db.DB.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	users = []model.TabUser{}
	// Slice of primary keys
	db.DB.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM `tab_user` WHERE `tab_user`.`user_id` IN (20,21,22)
	fmt.Println(users)

	// 指定只查结构体的哪些字段
	//使用结构体进行搜索时，可以通过传入相关字段名称或数据库名称来指定在查询条件中使用结构体中的哪些特定值，Where()例如：
	db.DB.Where(&model.TabUser{Name: "jinzhu"}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	db.DB.Where(&model.TabUser{Name: "jinzhu"}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;

}

// 1.2.6、条件 - 内联条件
func TestMyFunc0206(t *testing.T) {

	user := model.TabUser{}
	users := []model.TabUser{}

	// 如果主键是非整数类型，则通过主键获取
	db.DB.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// 纯 SQL
	db.DB.Find(&user, "name = ? ", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.DB.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// 结构
	db.DB.Find(&users, model.TabUser{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// 映射
	db.DB.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;
}

// 1.2.7、Not 条件
func TestMyFunc0207(t *testing.T) {
	// 建立 NOT 条件，工作原理类似于Where
	// where not ...

	user := model.TabUser{}
	users := []model.TabUser{}

	db.DB.Not("name = ?", "jinzhu").First(&user)
	//SELECT * FROM users WHERE NOT name =“jinzhu”ORDER BY id LIMIT 1;

	// 不在
	db.DB.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.DB.Not(model.TabUser{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// 不在主键切片中
	db.DB.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

}

// 1.2.8、或条件
func TestMyFunc0208(t *testing.T) {
	// where xxx or zzz

	users := []model.TabUser{}

	db.DB.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// 结构
	db.DB.Where("name = 'jinzhu'").Or(model.TabUser{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// 映射
	db.DB.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

}

// 1.2.9、选择特定字段
func TestMyFunc0209(t *testing.T) {
	// 查询指定字段
	// Select允许您指定要从数据库检索的字段。否则，GORM 将默认选择所有字段

	user := model.TabUser{}
	users := []model.TabUser{}

	//从用户中选择姓名、年龄；
	db.DB.Select("name", "age").Find(&users)
	// SELECT `name`,`age` FROM `tab_user`

	//从用户中选择姓名、年龄；其中 姓名为 WangWu 的人
	db.DB.Select("name", "age").Where(map[string]any{"name": "WangWu"}).Find(&users)
	// SELECT `name`,`age` FROM `tab_user` WHERE `name` = 'WangWu'

	// 从用户中选择姓名、年龄；
	db.DB.Select([]string{"name", "age"}).Find(&users)
	// SELECT `name`,`age` FROM `tab_user`

	// 从用户中选择 COALESCE(age,'42')；
	db.DB.Table(user.TableName()).Select("COALESCE(age,?)", 42).Rows()
	// SELECT COALESCE(age,42) FROM `tab_user`

}

// 1.2.10、排序
func TestMyFunc02010(t *testing.T) {
	//指定从数据库检索记录时的顺序

	//anonymous_user := model.TabUser{}
	users := []model.TabUser{}

	db.DB.Order("age desc,name ").Find(&users)
	//SELECT * FROM users ORDER BY age desc, name;

	// 多个订单
	db.DB.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&model.TabUser{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

}

// 1.2.11、限制和偏移
func TestMyFunc02011(t *testing.T) {
	//Limit指定要检索的最大记录数，Offset指定在开始返回记录之前要跳过的记录数

	//anonymous_user := model.TabUser{}
	users := []model.TabUser{}
	users2 := []model.TabUser{}

	db.DB.Limit(3).Find(&users)
	// [10.077ms] [rows:3] SELECT * FROM `tab_user` LIMIT 3

	// 使用 -1 取消限制条件
	db.DB.Limit(10).Find(&users).Limit(-1).Find(&users2)
	// [0.532ms] [rows:10] SELECT * FROM `tab_user` LIMIT 10 (users1)
	// [0.536ms] [rows:14] SELECT * FROM `tab_user (users2)

	db.DB.Offset(3).Find(&users)
	// [0.553ms] [rows:0] SELECT * FROM `tab_user` OFFSET 3

	db.DB.Limit(10).Offset(5).Find(&users)
	// [0.000ms] [rows:9] SELECT * FROM `tab_user` LIMIT 10 OFFSET 5

	// 用 -1 取消偏移条件
	db.DB.Offset(10).Find(&users)
	// [0.532ms] [rows:0] SELECT * FROM `tab_user` OFFSET 10 (users1)

	db.DB.Offset(-1).Find(&users2)
	// [0.000ms] [rows:14] SELECT * FROM `tab_user`

}

// 1.2.11、限制和偏移
func TestMyFunc020111(t *testing.T) {

	// 测试分页效果
	query := db.DB.Where("age = ?", 18)

	var pageUserList []model.TabUser = Paginate(2, 5, query)

	for idx, user := range pageUserList {
		fmt.Printf("idx = %d, userId = %d \n", idx, user.UserId)
	}

}

// 测试分页
func Paginate(pageNum, pageSize int, tx *gorm.DB) []model.TabUser {

	/// https://gorm.io/zh_CN/docs/scopes.html#pagination

	if pageNum <= 0 {
		pageNum = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	users := []model.TabUser{}
	offset := (pageNum - 1) * pageSize
	tx.Offset(offset).Limit(pageSize).Find(&users)
	return users
}

// 1.2.12、Group By & Having
func TestMyFunc02012(t *testing.T) {
	user := model.TabUser{}
	users := []model.TabUser{}

	db.DB.Model(&model.TabUser{}).Select("name, sum(age) as total").Where("name LIKE ?", "zhang%").Group("name").Find(&users)
	// [1.629ms] [rows:1] SELECT name, sum(age) as total FROM `tab_user` WHERE name LIKE 'zhang%' GROUP BY `name`

	db.DB.Model(&model.TabUser{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "zhang").Find(&users)
	// [0.536ms] [rows:0] SELECT name, sum(age) as total FROM `tab_user` GROUP BY `name` HAVING name = 'zhang'

	// [0.543ms] [rows:-] SELECT date(created_at) as date, sum(age) as total FROM `tab_user` GROUP BY date(created_at)
	rows, _ := db.DB.Table(user.TableName()).Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Rows()
	defer rows.Close()
	for rows.Next() {

	}

	// SELECT date(created_at) as date, sum(age) as total FROM `orders` GROUP BY date(created_at) HAVING sum(amount) > 100
	rows2, _ := db.DB.Table(user.TableName()).Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Having("sum(age) > ?", 100).Rows()
	defer rows2.Close()
	for rows2.Next() {

	}

	db.DB.Table(user.TableName()).Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Having("sum(age) > ?", 100).Scan(&users)

}

// 1.2.13、distinct
func TestMyFunc02013(t *testing.T) {

	// 从模型中选择不同的值

	//anonymous_user := model.TabUser{}
	users := []model.TabUser{}

	db.DB.Distinct("name", "age").Order("name, age desc").Find(&users)
	// [2.134ms] [rows:6] SELECT DISTINCT `name`,`age` FROM `tab_user` ORDER BY name, age desc

}

// 1.2.14、Joins
func TestMyFunc02014(t *testing.T) {

	// 连表查询指定连接条件

	user := model.TabUser{}
	users := []model.TabUser{}

	db.DB.Model(&model.TabUser{}).Select("tab_user.user_id, tab_user.name, tab_user_card.user_card_id, tab_user_card.created_at").Joins("left join tab_user_card on tab_user_card.user_id = tab_user.user_id").Scan(&users)
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

	//fmt.Println(users)

	rows, _ := db.DB.Table(user.TableName()).Select("tab_user.name, tab_user.email").Joins("left join tab_user_card on tab_user_card.user_id = tab_user.user_id").Rows()
	for rows.Next() {

	}

	db.DB.Table(user.TableName()).Select("tab_user.name, tab_user.email").Joins("left join tab_user_card on tab_user_card.user_id = users.user_id").Scan(&users)

	// multiple joins with parameter
	//db.DB.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&anonymous_user)

}
