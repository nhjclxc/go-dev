package model

import "time"

type TabUserCard struct {
	Id         int       `orm:"column(user_card_id);auto"`
	UserId     int64     `orm:"column(user_id)"`
	CardTypeId int64     `orm:"column(card_type_id)"`
	Remark     string    `orm:"column(remark);size(255);null"`
	Status     string    `orm:"column(status);size(1);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);auto_now"`

	// 当前结构体持有外键字段（即数据库中它有指向对方的字段）
	//用于在主结构体中定义外键字段，即本结构体拥有另一个结构体的引用。也就是本表为主表
	// on_delete
	TabUser *TabUser `orm:"rel(one)"`

}