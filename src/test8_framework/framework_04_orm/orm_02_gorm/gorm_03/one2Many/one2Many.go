package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_03/config"
)

// User 用户结构体
type User struct {
	Id       uint64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`  // 用户id
	Username string      `gorm:"column:username" json:"username"`               // 用户名
	Orders   []UserOrder `gorm:"foreignKey:UserId;references:Id" json:"orders"` // 一对多关联
}

func (User) TableName() string {
	return "tab_user"
}

// UserOrder 用户订单结构体
type UserOrder struct {
	Id      uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserId  uint64 `gorm:"column:user_id;index" json:"user_id"`
	OrderNo string `gorm:"column:order_no" json:"order_no"`
	Amount  string `gorm:"column:amount" json:"amount"`
}

func (UserOrder) TableName() string {
	return "tab_user_order"
}

// GetUserWithOrders 根据用户id获取用户信息和用户所有订单信息
func GetUserWithOrders(db *gorm.DB, userId uint64) (*User, error) {
	var user User
	err := db.Preload("Orders").First(&user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func main() {
	detail, err := GetUserWithOrders(config.DB, 1)
	if err != nil {
		return
	}
	fmt.Println(detail)
}
