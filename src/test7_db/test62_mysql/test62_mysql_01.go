package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"time"
)

// 学习地址：https://help.aliyun.com/zh/rds/apsaradb-rds-for-mysql/use-the-go-driver-package-go-mysql-driver-to-connect-to-a-database
func main1() {

	// 创建一个数据库连接
	var cfg *mysql.Config = mysql.NewConfig()
	cfg.User = "root"           //您需手动替换为实际数据库用户名
	cfg.Passwd = "root123"      //您需手动替换为实际数据库密码
	cfg.Net = "tcp"             //连接类型为TCP，保持默认，无需您手动修改
	cfg.Addr = "localhost:3306" //您需手动替换为实际MySQL数据库连接地址和端口号
	cfg.DBName = "test"         //您需手动替换为实际数据库名称

	// 连接数据库
	conn, err := mysql.NewConnector(cfg)
	if err != nil {
		panic(err.Error())
	}

	// 打开数据库
	db := sql.OpenDB(conn)

	// 延迟关闭数据库连接
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Error closing database connection: %v\n", err)
		}
	}(db)

	// 设置连接池参数
	db.SetMaxOpenConns(20)                  // 设置连接池中最大打开的连接数，您可根据实际需要手动调整
	db.SetMaxIdleConns(2)                   // 设置连接池中最大空闲的连接数，您可根据实际需要手动调整
	db.SetConnMaxIdleTime(10 * time.Second) // 连接在连接池中的最长的空闲时间，您可根据实际需要手动调整
	db.SetConnMaxLifetime(80 * time.Second) // 设置连接的最大生命周期，您可根据实际需要手动调整

	fmt.Println("数据库连接成功！！！")
	fmt.Println(cfg)
	fmt.Println(conn)
	fmt.Println(db)

	// 操作 1、创建一张表 userinfo
	//_, err = db.Exec("create table  if not exists userinfo( uid int auto_increment, username varchar(64) not null default '', departname varchar(20) not null default '', created varchar(10) not null default '', primary key(uid) );")
	//if err != nil {
	//	fmt.Println("create table error ", err)
	//	return
	//}

	// 操作 2、写入数据 userinfo
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	res, err := stmt.Exec("James"+uuid.New().String(), "Research", "2025-04-19")
	if err != nil {
		log.Fatalln("stmt.Exec：", err) // Error 1406 (22001): Data too long for column 'username' at row 1
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据插入成功：id = ", id)

	// 操作 3、更新数据 userinfo
	stmtUpdate, err := db.Prepare("UPDATE userinfo SET departname=?, created=? WHERE username=?")
	if err != nil {
		fmt.Println("Prepare update statement error:", err)
		return
	}
	resUpdate, err := stmtUpdate.Exec("Sa强强强强les", "2024-09-23", "James")
	if err != nil {
		fmt.Println(err)
	}
	rowCnt, _ := resUpdate.RowsAffected()
	fmt.Println(rowCnt)

	// 操作 4、更新数据 userinfo
	//rows, err := db.Query("SELECT username,departname,created FROM userinfo WHERE username=?", "James")
	rows, err := db.Query("SELECT username,departname,created FROM userinfo")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var username, departname, created string
		if err := rows.Scan(&username, &departname, &created); err != nil {
			fmt.Println(err)
		}
		fmt.Println("username:", username, "departname:", departname, " created:", created)
	}

	// 操作 4、更新数据 userinfo 通过绑定参数的方式执行参数化查询，预编译的方式
	stmtQuery, err := db.Prepare("SELECT username,departname,created FROM userinfo WHERE username like ?")
	defer rows.Close()
	if err != nil {
		fmt.Println("prepare error", err)
		return
	}
	rowData, err := stmtQuery.Query("%James%")
	if err != nil {
		fmt.Println("query data error", err)
		return
	}
	for rowData.Next() {
		var username, departname, created string
		if err := rowData.Scan(&username, &departname, &created); err != nil {
			fmt.Println(err)
		}
		fmt.Println("username:", username, "departname:", departname, "created:", created)
	}

	// 操作 5、删除数据 userinfo
	delStmt, _ := db.Prepare("DELETE FROM userinfo WHERE username=?")
	resultDel, err := delStmt.Exec("James")
	if err != nil {
		panic(err)
	}
	rowAffect, _ := resultDel.RowsAffected()
	fmt.Println("Data deletion completed.", rowAffect)

}
