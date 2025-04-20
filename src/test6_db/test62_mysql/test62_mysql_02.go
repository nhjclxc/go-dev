package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql包
	"github.com/google/uuid"
	"log"
	"time"
)

// 学习链接：https://segmentfault.com/a/1190000021693989
func main() {

	//_ "github.com/go-sql-driver/mysql" //导入mysql包
	// 必须导入mysql驱动才能执行下面的数据库连接，里面会调用init方法
	// 以下的 sql.Open("mysql" 需要这个包
	// 如果使用了：sql.Open("mysql"，就必须导入 github.com/go-sql-driver/mysql 包，因为这个包会在它的 init() 函数里注册名为 "mysql" 的驱动到 database/sql 里。

	// 1、打开数据库链接
	//mysql数据库的链接字符串组织：用户名:密码@tcp(数据库IP:端口)/数据库名?charset=utf8&parseTime=True，如果你的表里有应用到datetime字段，记得要加上parseTime=True，不然解析不了这个类型。
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalln("mysql数据库连接异常：", err)
	}
	// 连接成功

	// 延迟关闭
	defer db.Close()

	// 1、创建表结构 userinfo(uid, username, departname, created)

	// 2、插入一条数据
	username := uuid.New().String()
	departname := "测试部门"
	created := "2025-04-19"
	result, err := db.Exec("INSERT INTO userinfo (username, departname, created) values (?, ?, ?)", username, departname, created)
	if err != nil {
		log.Fatalln("数据插入失败：", err)
	}
	uid, err := result.LastInsertId()
	size, err := result.RowsAffected()
	fmt.Println("result.LastInsertId() = ", uid)
	fmt.Println("result.RowsAffected() = ", size)

	// 3、查询单条数据
	var ui UserInfo = UserInfo{}
	row := db.QueryRow("SELECT uid, username, departname, created FROM userinfo WHERE uid = ?", 1)
	row.Scan(&ui.uid, &ui.username, &ui.departname, &ui.created)
	fmt.Println("uid=1的数据为：", ui)

	fmt.Println("---------------------------")
	// 4、查询多条数据
	var uiList []UserInfo = make([]UserInfo, 8)
	rows, err := db.Query("SELECT uid, username, departname, created FROM userinfo")
	defer rows.Close()
	for rows.Next() {
		// 看看是否有数据，有数据就一直输出
		var ui0 UserInfo = UserInfo{}
		rows.Scan(&ui0.uid, &ui0.username, &ui0.departname, &ui0.created)
		uiList = append(uiList, ui0)
		fmt.Println("ui0 = ", ui0)
	}
	fmt.Println("uiList.cap = ", cap(uiList), ", len = ", len(uiList))

	// 5、修改数据
	res, err := db.Exec("UPDATE userinfo set username=? WHERE uid = ?", time.Now().String(), 1)
	affectedLen, err := res.RowsAffected()
	fmt.Println("UPDATE affectedLen = ", affectedLen)

	// 6、删除一条数据
	resd, err := db.Exec("DELETE FROM userinfo WHERE uid = ?", 7)
	affectedLend, err := resd.RowsAffected()
	fmt.Println("DELETE affectedLend = ", affectedLend)

	// 7、事务
	tx, _ := db.Begin()
	result5, _ := tx.Exec("update userinfo set username = ? where uid = ?", time.Now().String(), 5)
	result6, _ := tx.Exec("update userinfo set username = ? where uid = ?", time.Now().String(), 6)

	//影响行数，为0则失败
	i5, _ := result5.RowsAffected()
	i6, _ := result6.RowsAffected()
	if i5 > 0 && i6 > 0 {
		//2条数据都更新成功才提交事务
		err = tx.Commit()
		if err != nil {
			fmt.Println("事务提交失败", err)
			return
		}
		fmt.Println("事务提交成功")
	} else {
		//否则回退事务
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回退事务失败", err)
			return
		}
		fmt.Println("回退事务成功")
	}

	// Prepare和Exec的区别
	// 1. Exec 是直接执行 SQL： 这个时候，SQL 语句+参数一次性发给数据库执行。
	result22, err := db.Exec("INSERT INTO userinfo(username, created) VALUES (?, ?)", "Tom", "2025-05-19")
	fmt.Println(result22)
	fmt.Println(err)

	// 2. Prepare 是预编译 SQL（适合重复执行）：
	//等价于：第一步：发送 SQL 模板到数据库（只发送一次）；后续每次 Exec：只发参数，数据库重复使用预编译语句，提高效率
	stmt, err := db.Prepare("INSERT INTO userinfo(username, created) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec("Tom", "2025-06-19")
	stmt.Exec("Jerry", "2025-07-19")

	// Go 是如何防止 SQL 注入的？
	// 🔥 只要你使用 ? 占位符 + 参数绑定的方式（而不是字符串拼接），Go 的 database/sql 包天然就可以防止 SQL 注入。
	_ = db.QueryRow("SELECT * FROM userinfo WHERE username = ?", username)

}

type UserInfo struct {
	uid        string
	username   string
	departname string
	created    string
}
