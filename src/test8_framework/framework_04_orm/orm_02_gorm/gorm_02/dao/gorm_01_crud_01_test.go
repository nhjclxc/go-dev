package main

import (
	"fmt"
	"gorm.io/gorm"
	db "gorm_02/config"
	"gorm_02/model"
	"testing"
	"time"
)

// 数据库 新增 Insert 练习
// https://gorm.io/zh_CN/docs/create.html

// 1.1.1、单条记录新增
func TestMyFunc0101(t *testing.T) {

	user := model.TabUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// 你无法向 ‘create’ 传递结构体，所以你应该传入数据的指针.
	result := db.DB.Create(&user) // 通过数据的指针来创建
	// INSERT INTO `tab_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Jinzhu','',18,'2025-04-25 21:41:33.473',NULL,'2025-04-25 21:41:33.474','2025-04-25 21:41:33.474','2025-04-25 21:41:33.474')

	fmt.Println("anonymous_user.UserId：", user.UserId)       // 返回插入数据的主键
	fmt.Println("result.Error：", result.Error)               // 返回 error
	fmt.Println("result.RowsAffected：", result.RowsAffected) // 返回插入记录的条数
}

// 1.1.2、插入指定的字段数据
func TestMyFunc0103(t *testing.T) {

	user := model.TabUser{Name: "插入指定的字段数据", Age: 18, Birthday: time.Now()}

	// 没有的字段，数据库默认传 null 给数据库
	// 注意，下面写的是实体类对应的属性，有gorm自动转化为数据库字段
	result := db.DB.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `tab_user` (`name`,`age`,`activated_at`,`created_at`,`updated_at`) VALUES ('插入指定的字段数据',18,'2025-04-25 21:44:29.737','2025-04-25 21:44:29.737','202         5-04-25 21:44:29.737')

	fmt.Println("result.Error：", result.Error)               // 返回 error
	fmt.Println("result.RowsAffected：", result.RowsAffected) // 返回插入记录的条数

}

// 1.1.3、批量数据新增
func TestMyFunc0102(t *testing.T) {

	// 要高效地插入大量记录，请将切片传递给Create方法。 GORM 将生成一条 SQL 来插入所有数据，以返回所有主键值，并触发 Hook 方法。
	//当这些记录可以被分割成多个批次时，GORM会开启一个事务</0>来处理它们。

	userList := []model.TabUser{
		model.TabUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		model.TabUser{Name: "ZhangSan", Age: 19, Birthday: time.Now()},
		model.TabUser{Name: "WangWu", Age: 20, Birthday: time.Now()},
	}
	result := db.DB.Create(userList) // 通过数据的指针来创建
	// INSERT INTO `tab_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Jinzhu','',18,'2025-04-25 21:38:57.113',NULL,'2025-04-25 21:38:57.114','2025-04-25 21:38:57.114','2025-04-25 21:38:57.114'),('ZhangSan','',19,'2025-04-25 21:38:57.113',NULL,'2025-04-25 21:38:57.114','2025-04-25 21:38:57.114','2025-04-25 21:38:57.114'),('WangWu','',20,'2025-04-25 21:38:57.113',NULL,'2025-04-25 21:38:57.114','2025-04-25 21:38:57.114','2025-04-25 21:38:57.114')

	// 使用 CreateInBatches 来限制批量插入大小， 超出超入大小的数据会被丢弃
	//db.DB.CreateInBatches(userList, 100)

	fmt.Println("result.Error：", result.Error)               // 返回 error
	fmt.Println("result.RowsAffected：", result.RowsAffected) // 返回插入记录的条数

	// 看看主键有没有返回到实体类里面
	for _, val := range userList {
		fmt.Println(val.UserId, val.Name)
	}

}

// 1.1.4、创建钩子函数
func TestMyFunc0104(t *testing.T) {
	// GORM允许用户通过实现这些接口 BeforeSave, BeforeCreate, AfterSave, AfterCreate 来自定义钩子。
	//这些钩子方法会在创建一条记录时被调用，关于钩子的生命周期请参阅[Hooks](https://gorm.io/zh_CN/docs/hooks.html)

	// 部分钩子函数的实现在 TabUser 结构体里面

	// 如果在某个sql下不想执行钩子函数，则使用 Session(&gorm.Session{SkipHooks: true}) 来跳过钩子函数
	db.DB.Session(&gorm.Session{SkipHooks: true}).Create(nil)

}

// 1.1.5、根据 Map 传参来 保存数据
func TestMyFunc0105(t *testing.T) {
	db.DB.Model(&model.TabUser{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.DB.Model(&model.TabUser{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})

	// 当使用map来创建时，因为没有使用 GORM 指定的实体类
	// 所以：钩子方法不会执行，关联不会被保存且不会回写主键。
}

// 1.1.6、使用 SQL 表达式、Context Valuer 创建记录
func TestMyFunc0106(t *testing.T) {
	// GORM允许使用SQL表达式来插入数据，有两种方法可以达成该目的，使用map[string]interface{} 或者 Customized Data Types， 示例如下：

}

// 1.1.7、高级选项 - 关联创建
func TestMyFunc0107(t *testing.T) {

	// 关联创建 TabUser 和 TabUserCard
	// 新增用户的时候，同时给他开一张卡

	/*
		语法规则：
			db.Create(&User{
			  Name: "jinzhu",
			  CreditCard: CreditCard{Number: "411111111111"},
			})
			// INSERT INTO `users` ...
			// INSERT INTO `credit_cards` ...
	*/

	// 准备 TabUserCard 结构体
	tabUserCard := model.TabUserCard{
		CardTypeId: 1,
		Remark:     "这张卡是默认给张三创建的",
	}

	// 准备 TabUser 结构体
	tabUser := model.TabUser{
		Name:        "张三",
		Email:       "zhangsan@zzz.com",
		Age:         18,
		TabUserCard: tabUserCard,
	}

	// 关联创建
	result := db.DB.Create(&tabUser)
	// 先插入子表，再插主表
	// [1.056ms] [rows:1] INSERT INTO `tab_user_card` (`user_id`,`card_type_id`,`remark`,`status`,`created_at`,`updated_at`) VALUES (16,1,'这张卡是默认给张三创建的',0,'2025-04-26 09:26:14.064','2025-04-26 09:26:14.064') ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)
	//[15.968ms] [rows:1] INSERT INTO `tab_user` (`name`,`email`,`age`,`birthday`,`member_number`,`remark`,`activated_at`,`created_at`,`updated_at`) VALUES ('张三','zhangsan@zzz.com',18,'202  5-04-26 09:26:14.055',NULL,'e5fb0c89-cef0-4c32-a7b3-ab46197bc50b','2025-04-26 09:26:14.055','2025-04-26 09:26:14.055','2025-04-26 09:26:14.055')

	fmt.Println("anonymous_user.UserId：", tabUser.UserId)          // 返回插入数据的主键
	fmt.Println("tabUserCard.UserCardId：", tabUserCard.UserCardId) // 返回插入数据的主键
	fmt.Println("result.Error：", result.Error)                     // 返回 error
	fmt.Println("result.RowsAffected：", result.RowsAffected)       // 返回插入记录的条数

	// 有的时候，不想关联新增
	//你可以通过Select, Omit方法来跳过关联更新，示例如下：
	//db.Omit("CreditCard").Create(&anonymous_user)
	// skip all associations
	//db.Omit(clause.Associations).Create(&anonymous_user)

}

// 1.1.、高级选项 - 关联创建
func TestMyFunc010(t *testing.T) {

}
