package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"time"
)

// eego 的 ORM 模块要求在使用之前要先注册好模型，并且 Beego 会执行一定的校验，用于辅助检查模型和模型之间的约束。

// 注册模型
//注册模型有三个方法：
//RegisterModel(models ...interface{})
//RegisterModelWithPrefix(prefix string, models ...interface{})：该方法会为表名加上前缀，例如RegisterModelWithPrefix("tab_", &User{})，那么表名是tab_user；
//RegisterModelWithSuffix(suffix string, models ...interface{})：该方法会为表名加上后缀，例如RegisterModelWithSuffix("_tab", &User{})，那么表名是user_tab

type TabUser struct {

	// 为字段设置 DB 列的名称
	// orm 表示的是 beego 里面的 orm
	// column(xxx) 表示对于的字段在数据库表中的字段名称
	// auto 表示 这个字段是自增主键 如果不想使用自增主键，那么可以使用 pk 设置为主键。这时建议使用UUID
	// default 默认值是一个扩展功能，必须要显示注册默认值的Filter，而后在模型定义里面加上default的设置。
	UserId       int       `orm:"column(user_id);auto"`
	Name         string    `orm:"column(name);size(255)"`
	Email        string    `orm:"column(email);size(255);null"`
	Age          int8      `orm:"column(age);default:12"`
	Birthday     time.Time `orm:"column(birthday);type(datetime);null"`
	MemberNumber string    `orm:"column(member_number);size(255);null"`
	Remark       string    `orm:"column(remark);size(128);null"`

	// 自动更新时间 ，对于批量的 update 此设置是不生效
	//auto_now 每次 model 保存时都会对时间自动更新
	//auto_now_add 第一次保存时才设置时间
	ActivatedAt time.Time `orm:"column(activated_at);type(datetime);null"`
	CreatedAt   time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(datetime);auto_now"`

	// 设置 - 即可忽略模型中的字段，即表示这个实体属性不是数据库里面的一个字段
	AnyField string `orm:"-"`



	// 设置表关联

	// 一对一设置
	//RelOneToOne（正向关系，reverse(one)）
	TabUserCard *TabUserCard `orm:"null;reverse(one);on_delete(set_null)"`

	// 一对一关系，判断方法
	//🧠 快速判断技巧：
	//看数据库中哪张表有外键字段（例如 user_id）→ 那张表就是正向关系（rel）的一方。
	//被外键引用的一方就是反向（reverse）关系的一方。



}

// 自定义 TabUser 实体类的表明为 tab_user
func (u *TabUser) TableName() string {
	return "tab_user"
}

type TabUserCard struct {
	Id         int       `orm:"column(user_card_id);auto"`
	UserId     int64     `orm:"column(user_id)"`
	CardTypeId int64     `orm:"column(card_type_id)"`
	Remark     string    `orm:"column(remark);size(255);null"`
	Status     string    `orm:"column(status);size(1);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);auto_now"`

	// 当前结构体持有外键字段（即数据库中它有指向对方的字段）
	//用于在主结构体中定义外键字段，即本结构体拥有另一个结构体的引用。也就是本表为主表
	// on_delete
	TabUser *TabUser `orm:"rel(one)"`

//
}

func (t *TabUserCard) TableName() string {
	return "tab_user_card"
}

func init() {

	orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8")

}

func main() {

}
