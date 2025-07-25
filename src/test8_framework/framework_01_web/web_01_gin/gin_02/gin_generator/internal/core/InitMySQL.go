package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// 定义公开的连接信息
var GLOBAL_DB *gorm.DB

func init() {
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "用户名:密码@tcp(IP:Port)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印所有日志
		Logger: logger.Default.LogMode(logger.Info),
		// 显示出查询的所有字段
		QueryFields: true,
		// 禁用默认事务，true 表示不开起事务；不写或写false表示开启事务
		//SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("数据库连接失败！！！")
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	GLOBAL_DB = db
	log.Println("数据库连接成功")
}