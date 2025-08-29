package main

import (
	"fmt"
	db "gorm_02/config"
	"gorm_02/model"
	"time"
)

func MyFunc() {

}

func main() {

	user := model.TabUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.DB.Create(&user) // 通过数据的指针来创建

	fmt.Println("anonymous_user.UserId：", user.UserId)       // 返回插入数据的主键
	fmt.Println("result.Error：", result.Error)               // 返回 error
	fmt.Println("result.RowsAffected：", result.RowsAffected) // 返回插入记录的条数

}
