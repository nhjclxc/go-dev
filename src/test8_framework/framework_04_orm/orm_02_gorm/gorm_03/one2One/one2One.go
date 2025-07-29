package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_03/config"
)

// User 用户结构体
type User struct {
	Id         uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`      // 用户id
	Username   string     `gorm:"column:username" json:"username"`                   // 用户名
	UserDetail UserDetail `gorm:"foreignKey:UserId;references:Id" json:"userDetail"` // 一对一关联
}

func (User) TableName() string {
	return "tab_user"
}

// UserDetail 用户详细结构体
type UserDetail struct {
	UserId  uint64 `gorm:"column:user_id;unique;not null" json:"userId"` // 与主表建立一对一关联
	Address string `gorm:"column:address" json:"address"`                // 用户地址
	Phone   string `gorm:"column:phone" json:"phone"`                    // 用户手机号
}

func (UserDetail) TableName() string {
	return "tab_user_detail"
}

// GetUserWithDetail 根据用户id获取用户信息和用户详细信息
func GetUserWithDetail(db *gorm.DB, userId uint64) (*User, error) {
	var user User
	err := db.Preload("UserDetail").First(&user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {
	detail, err := GetUserWithDetail(config.DB, 2)
	if err != nil {
		return
	}
	fmt.Println(detail)
}
