package main

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

/*
create table user1
(
    id1       bigint auto_increment primary key,
    username1 varchar(50) not null
);
INSERT INTO user1 (id1, username1) VALUES (11, 'u11');
INSERT INTO user1 (id1, username1) VALUES (12, 'u12');
INSERT INTO user1 (id1, username1) VALUES (13, 'u13');


create table user2
(
    id2       bigint auto_increment primary key,
    user1_id1       bigint ,
    username2 varchar(50) not null
);
INSERT INTO user2 (id2, user1_id1, username2) VALUES (21, 11, 'u21');
INSERT INTO user2 (id2, user1_id1, username2) VALUES (22, 12, 'u22');
INSERT INTO user2 (id2, user1_id1, username2) VALUES (23, 13, 'u23');


create table user3
(
    id3       bigint auto_increment primary key,
    user2_id2       bigint ,
    username3 varchar(50) not null
);
INSERT INTO user3 (id3, user2_id2, username3) VALUES (31, 21, 'u31');
INSERT INTO user3 (id3, user2_id2, username3) VALUES (32, 22, 'u32');
INSERT INTO user3 (id3, user2_id2, username3) VALUES (33, 23, 'u33');



create table user5
(
    id5       bigint auto_increment primary key,
    user1_id1       bigint ,
    username5 varchar(50) not null
);
INSERT INTO user5 (id5, user1_id1, username5) VALUES (51, 11, 'u51');
INSERT INTO user5 (id5, user1_id1, username5) VALUES (52, 11, 'u52');
INSERT INTO user5 (id5, user1_id1, username5) VALUES (53, 11, 'u53');
INSERT INTO user5 (id5, user1_id1, username5) VALUES (54, 12, 'u54');
INSERT INTO user5 (id5, user1_id1, username5) VALUES (55, 12, 'u55');
INSERT INTO user5 (id5, user1_id1, username5) VALUES (56, 12, 'u56');


create table user6
(
    id6       bigint auto_increment primary key,
    user5_id5       bigint ,
    username6 varchar(50) not null
);
INSERT INTO user6 (id6, user5_id5, username6) VALUES (61, 51, 'u61');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (62, 51, 'u62');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (63, 51, 'u63');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (64, 52, 'u64');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (65, 52, 'u65');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (66, 52, 'u66');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (67, 53, 'u67');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (68, 53, 'u68');
INSERT INTO user6 (id6, user5_id5, username6) VALUES (69, 53, 'u69');


user1 : user2 : user3 = 1:1:1
user1 : user5 = 1:n
user5 : user6 = n:m

*/

type User1 struct {
	Id1       int64  `gorm:"column:id1;type:bigint;primaryKey;autoIncrement;not null;" json:"id1" form:"id1"` // 用户ID
	Username1 string `gorm:"column:username1;type:varchar(50);not null;" json:"username1" form:"username1"`   // 用户名

	//foreignKey → 指的是 子表里的字段
	//references → 指的是 父表里的字段
	User2 User2 `gorm:"references:Id1;foreignKey:User1Id1"`

	User5s []User5 `gorm:"references:Id1;foreignKey:User1Id1"`
}

type User2 struct {
	Id2       int64  `gorm:"column:id2;type:bigint;primaryKey;autoIncrement;not null;" json:"id2" form:"id2"` // 用户ID
	User1Id1  int64  `gorm:"column:user1_id1;type:bigint;not null;" json:"user1_id1" form:"user1_id1"`        // 用户ID
	Username2 string `gorm:"column:username2;type:varchar(50);not null;" json:"username2" form:"username2"`   // 用户名

	User3 User3 `gorm:"references:Id2;foreignKey:User2Id2"`
}

type User3 struct {
	Id3       int64  `gorm:"column:id3;type:bigint;primaryKey;autoIncrement;not null;" json:"id3" form:"id3"` // 用户ID
	User2Id2  int64  `gorm:"column:user2_id2;type:bigint;not null;" json:"user2_id2" form:"user2_id2"`        // 用户ID
	Username3 string `gorm:"column:username3;type:varchar(50);not null;" json:"username3" form:"username3"`   // 用户名
}

type User5 struct {
	Id5       int64  `gorm:"column:id5;type:bigint;primaryKey;autoIncrement;not null;" json:"id5" form:"id5"` // 用户ID
	User1Id1  int64  `gorm:"column:user1_id1;type:bigint;not null;" json:"user1_id1" form:"user1_id1"`        // 用户ID
	Username5 string `gorm:"column:username5;type:varchar(50);not null;" json:"username5" form:"username5"`   // 用户名

	User6s []User6 `gorm:"references:Id5;foreignKey:User5Id5"`
}

type User6 struct {
	Id6       int64  `gorm:"column:id6;type:bigint;primaryKey;autoIncrement;not null;" json:"id6" form:"id6"` // 用户ID
	User5Id5  int64  `gorm:"column:user5_id5;type:bigint;not null;" json:"user5_id5" form:"user5_id5"`        // 用户ID
	Username6 string `gorm:"column:username6;type:varchar(50);not null;" json:"username6" form:"username6"`   // 用户名
}

//user1 err failed to parse field: User5, error: invalid field found for struct gorm_03/many_tables.User5's field User6: define a valid foreign key for relations or implement the Valuer/Scanner interface

func (*User1) TableName() string {
	return "user1"
}
func (*User2) TableName() string {
	return "user2"
}
func (*User3) TableName() string {
	return "user3"
}
func (*User5) TableName() string {
	return "user5"
}
func (*User6) TableName() string {
	return "user6"
}

func Test(t *testing.T) {

	DB := getDB("127.0.0.1", "test1")

	fmt.Println(DB)

	//user1 : user2 : user3 = 1:1:1
	//test111(DB)

	//user1 : user5 = 1:n
	//test222(DB)

	//user5 : user6 = n:m
	//test333(DB)

	// user1 : user5 : user6 = 1 : n:m
	test555(DB)

}

func test555(db *gorm.DB) {

	var user1s []User1
	tx1 := db.Model(User1{}).Preload("User2").Preload("User5s").Preload("User5s.User6s").Find(&user1s)
	fmt.Printf("user1s err %s \n", tx1.Error)
	fmt.Printf("user1s %v \n", user1s)
}

func test333(db *gorm.DB) {

	var user5s []User5
	tx := db.Model(User5{}).Preload("User6s").Find(&user5s)

	fmt.Printf("user5s err %s \n", tx.Error)
	fmt.Printf("user5s %v \n", user5s)
}

func test222(db *gorm.DB) {
	//user1 : user5 = 1:n
	var user1 User1
	tx := db.Model(User1{}).Preload("User5s").First(&user1)

	fmt.Printf("user1 err %s \n", tx.Error)
	fmt.Printf("user1 %v \n", user1)

	var user1s []User1
	tx2 := db.Model(User1{}).Preload("User5s").First(&user1)

	fmt.Printf("user1s err %s \n", tx2.Error)
	fmt.Printf("user1s %v \n", user1s)

}

func test111(db *gorm.DB) {
	//user1 : user2 : user3 = 1:1:1

	var user1 User1

	tx := db.Model(User1{}).
		Preload("User2").       // user1通过User2进行关联user2表
		Preload("User2.User3"). // user1通过User2进行里面的User3关联user3表
		First(&user1)

	fmt.Printf("user1 err %s \n", tx.Error)
	fmt.Printf("user1 %v \n", user1)

	var user1s []User1

	tx2 := db.Model(User1{}).
		Preload("User2").       // user1通过User2进行关联user2表
		Preload("User2.User3"). // user1通过User2进行里面的User3关联user3表
		Find(&user1s)

	fmt.Printf("user1s err %s \n", tx2.Error)
	fmt.Printf("user1s %v \n", user1s)
}

/*

1、Preload加条件

db.Preload("User2.User3", "status = ?", 1).Find(&list)

db.Preload("User2.User3", func(db *gorm.DB) *gorm.DB {
        return db.Where("status IN (?)", []int{1, 2})
    }).
    Find(&list)

db.Preload("User2.User3", func(db *gorm.DB) *gorm.DB {
        return db.Where("status = ?", 1).Order("id desc").Limit(5)
    }).
    Find(&list)





2、Preload取指定字段

db.Preload("User2.User3", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "status")
    }).
    Find(&list)


db.Preload("User2.User3", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "status", "device_id") // 取User3的这三个字段
    }).
    Preload("User2.User3.User666", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "title") // 只取User666的这两个字段
    }).
    Find(&list)





3、Preload取指定字段 + 条件
Preload指定字段 + 条件

db.Preload("User2.User3", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "status", "device_id").
                  Where("status IN (?)", []int{1, 2}).
                  Order("id DESC")
    }).
    Preload("User2.User3.User666", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "title").Where("deleted_at IS NULL")
    }).
    Find(&live)

*/
