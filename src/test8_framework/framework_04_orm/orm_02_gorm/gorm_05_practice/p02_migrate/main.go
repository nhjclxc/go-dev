package main

import (
	"fmt"
	"gorm_05_practice/config"
)

// 根据 结构体 对表进行实时更新
func main() {

	db := config.DB
	_ = db

	err := (&Menu{}).CreateTable(db)
	if err != nil {
		fmt.Println("create table error", err)
		return
	}
	fmt.Println("create table success")

}
