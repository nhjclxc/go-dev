package model

import (
    "gorm.io/gorm"
    "time"
)

// TabUser 用户 结构体
type TabUser struct {
	Id int64 `gorm:"column:id;type:bigint;primaryKey;autoIncrement;not null;comment:用户ID" json:"id" form:"id"` // 用户ID

	Username string `gorm:"column:username;type:varchar(50);unique;not null;comment:用户名" json:"username" form:"username"` // 用户名

	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码（加密存储）" json:"password" form:"password"` // 密码（加密存储）

	Email string `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email" form:"email"` // 邮箱

	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:创建时间" json:"createdAt" form:"createdAt"` // 创建时间

	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:更新时间" json:"updatedAt" form:"updatedAt"` // 更新时间

}

// TableName 返回当前实体类的表名
func (tu *TabUser) TableName() string {
	return "tab_user"
}

// CreateTable 根据结构体里面的gorm信息创建表结构
func (tu *TabUser) CreateTable(tx *gorm.DB) error {
	tableName := tu.TableName()
	if !tx.Migrator().HasTable(tableName) {
		return tx.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户 '").
			Table(tableName).Migrator().CreateTable(&TabUser{})
	}
	return nil
}

// 可用钩子函数包括：BeforeCreate / AfterCreate、BeforeUpdate / AfterUpdate、BeforeDelete / AfterDelete
// BeforeCreate 在插入数据之前执行的操作
func (tu *TabUser) BeforeCreate(tx *gorm.DB) (err error) {
	if err = tu.CreateTable(tx); err != nil {
		return err
	}

	tu.CreatedAt = time.Now()
	tu.UpdatedAt = time.Now()
	return
}

func (tu *TabUser) BeforeUpdate(tx *gorm.DB) (err error) {
	tu.UpdatedAt = time.Now()
	return
}
