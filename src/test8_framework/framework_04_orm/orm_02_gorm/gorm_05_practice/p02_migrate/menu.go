package main

import (
	"gorm.io/gorm"
	"time"
)

// Menu Menu结构体
type Menu struct {
	Id int64 `gorm:"column:id;type:bigint;primaryKey;autoIncrement;not null" json:"id" form:"id"` //

	ParentId int64 `gorm:"column:parent_id;type:bigint;default:0" json:"parentId" form:"parentId"` //

	Name string `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"` //

	Path string `gorm:"column:path;type:varchar(255);not null;comment:URL或权限标识" json:"path" form:"path"` // URL或权限标识

	Type int8 `gorm:"column:type;type:tinyint;default:1;comment:1=菜单,2=按钮/API" json:"type" form:"type"` // 1=菜单,2=按钮/API

	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"createdAt" form:"createdAt"` //

	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updatedAt" form:"updatedAt"` //

	Version string `gorm:"column:version;default:1;type:varchar(255)" json:"version"` // 乐观锁（版本控制）

	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deletedAt"` // 删除标记, 删除时间 GORM 默认启用了“软删除（Soft Delete）”只要存在这个字段，GORM 默认启用软删除。

}

// TableName 返回当前实体类的表名
func (m *Menu) TableName() string {
	return "menu1"
}

// CreateTable 根据结构体里面的gorm信息创建表结构
func (m *Menu) CreateTable(tx *gorm.DB) error {
	tableName := m.TableName()
	//if !tx.Migrator().HasTable(tableName) {
	//var err error
	//	//AutoMigrate 创建表
	//	err = tx.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='Menu'").
	//			Table(tableName).AutoMigrate(&Menu{})
	//
	//	//CreateTable 创建表
	//	err = tx.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='Menu'").
	//		Table(tableName).Migrator().CreateTable(&Menu{})
	//	return err
	//}
	//return nil

	return tx.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='Menu'").
		Table(tableName).AutoMigrate(&Menu{})
}

// BeforeCreate 在插入数据之前执行的操作
func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	if err = m.CreateTable(tx); err != nil {
		return err
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *Menu) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
