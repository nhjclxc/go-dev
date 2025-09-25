package model

import "time"

type Menu struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Path      string
	Type      uint // 1=菜单,2=按钮/API
	CreatedAt time.Time
	UpdatedAt time.Time

	Roles []Role `gorm:"many2many:role_menu;"`
}

func (*Menu) TableName() string {
	return "menu"
}
