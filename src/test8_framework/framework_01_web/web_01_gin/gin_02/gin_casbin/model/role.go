package model

import "time"

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User `gorm:"many2many:user_role;"`
	Menus     []Menu `gorm:"many2many:role_menu;"`
}

func (*Role) TableName() string {
	return "role"
}
