package model

import "time"

// TabUser -
type TabUser struct {
	UserId       int       `orm:"column(user_id);auto"`
	Name         string    `orm:"column(name);size(255)"`
	Email        string    `orm:"column(email);size(255);null"`
	Age          int8      `orm:"column(age)"`
	Birthday     time.Time `orm:"column(birthday);type(datetime);null"`
	MemberNumber string    `orm:"column(member_number);size(255);null"`
	Remark       string    `orm:"column(remark);size(128);null"`
	ActivatedAt  time.Time `orm:"column(activated_at);type(datetime);null"`
	CreatedAt    time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(datetime);auto_now_add"`



	// 一对一设置
	//RelOneToOne（正向关系，reverse(one)）
	TabUserCard *TabUserCard `orm:"null;reverse(one);on_delete(set_null)"`

}

