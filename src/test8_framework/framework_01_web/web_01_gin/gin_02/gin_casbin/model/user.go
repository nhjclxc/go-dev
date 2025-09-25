package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time

	Roles []Role `gorm:"many2many:user_role;"`
}

func (*User) TableName() string {
	return "user"
}
