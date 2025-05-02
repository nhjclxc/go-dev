package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init() {
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8")

	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8",
		orm.MaxIdleConnections(maxIdle),
		orm.MaxOpenConnections(maxConn))

	//也可以在注册之后修改：
	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxOpenConns("default", 50)

	// 时区
	//ORM 默认使用 time.Local 本地时区
	//
	//作用于 ORM 自动创建的时间
	//从数据库中取回的时间转换成 ORM 本地时间
	//如果需要的话，你也可以进行更改
	//
	//// 设置为 UTC 时间
	//orm.DefaultTimeLoc = time.UTC
	//ORM 在进行 RegisterDataBase 的同时，会获取数据库使用的时区，然后在 time.Time 类型存取时做相应转换，以匹配时间系统，从而保证时间不会出错。



}

func main() {

}
