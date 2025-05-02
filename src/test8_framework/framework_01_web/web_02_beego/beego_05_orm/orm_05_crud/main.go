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
	UserId       int       `orm:"column(user_id);auto"`
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
	var ormer orm.Ormer = orm.NewOrm()

	//1、新增操作
	// 注意：md是指某个实体对象的 地 址
	// Insert(md interface{}) (int64, error)
	// InsertWithCtx(ctx context.Context, md interface{}) (int64, error)

	// data
	user := new(TabUser)
	user.Name = "orm_05_crud"

	insert, err := ormer.Insert(&user)
	if err != nil {
		return
	}
	fmt.Println("插入数据条数：", insert)

	// 2、保存或修改，在 MySQL 数据库中，使用的是 ON DUPLICATE KEY 语句
	// InsertOrUpdate(md interface{}, colConflitAndArgs ...string) (int64, error)
	// InsertOrUpdateWithCtx(ctx context.Context, md interface{}, colConflitAndArgs ...string) (int64, error)

	update, err := ormer.InsertOrUpdate(&user)
	if err != nil {
		return
	}
	fmt.Println("InsertOrUpdate：", update)

	user222 := new(TabUser)
	user222.Name = "orm_05_crud222"
	update222, err := ormer.InsertOrUpdate(&user222)
	if err != nil {
		return
	}
	fmt.Println("InsertOrUpdate222：", update222)

	// 3、批量插入
	// InsertMulti(bulk int, mds interface{}) (int64, error)
	// InsertMultiWithCtx(ctx context.Context, bulk int, mds interface{}) (int64, error)
	// 参数bulk是每一次批量插入的时候插入的数量。例如bulk<=1代表每一批插入一条数据，而如果bulk=3代表每次插入三条数据。
	//你需要仔细选择批次大小，它对插入性能有很大影响。大多数情况下，你可以把bulk设置成数据量大小。
	// mds必须是一个数组，或者是一个切片。

	users := make([]*TabUser, 2)
	users = append(users, user)
	users = append(users, user222)

	multi, err := ormer.InsertMulti(2, users)
	if err != nil {
		return
	}
	fmt.Println("multi：", multi)

	//4、修改数据，使用主键更新数据
	// 也就是如果你使用这个方法，Beego 会尝试读取里面的主键值，而后将主键作为更新的条件。
	// 如果你没有指定 cols 参数，那么所有的列都会被更新。
	// 即，如果有指定cols，则更新指定的字段
	// Update(md interface{}, cols ...string) (int64, error)
	// UpdateWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error)

	user.Name = "是否有发生修改呀！！！"
	user.MemberNumber = "123456789"
	user.Remark = "是否有发生修改呀"
	cols := []string{"member_number", "remark"}
	ormer.Update(&user, cols...) // 使用三个点将数据展开成 "member_number", "remark"

	// 5、删除操作 ，使用主键来删除数据，
	// Delete(md interface{}, cols ...string) (int64, error)
	//DeleteWithCtx(ctx context.Context, md interface{}, cols ...string) (int64, error)

	// 6、读数据
	// Read(md interface{}, cols ...string) error
	// ReadWithCtx(ctx context.Context, md interface{}, cols ...string) error
	//该方法的特点是：
	//读取到的数据会被放到 md；
	//如果传入了 cols 参数，那么只会选取特定的列；

	// 去读表中的全部数据
	u61s := make([]TabUser, 10)
	err = ormer.Read(u61s)

	// 读取 UserId = UserId 的全部列
	u62 := &TabUser{UserId: 666}
	err = ormer.Read(u62)

	// 只读取 UserId = UserId 的用户名这一个列
	u63 := &TabUser{UserId: 666}
	err = ormer.Read(u63, "UserName")

}
