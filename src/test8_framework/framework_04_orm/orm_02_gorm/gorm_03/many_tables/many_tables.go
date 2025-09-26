package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

/*
create table tab_user
(
    id       bigint auto_increment comment '用户ID' primary key,
    username varchar(50)  not null comment '用户名',
    password varchar(255) not null comment '密码',
    phone    varchar(100) null comment '手机号',
    email    varchar(100) null comment '邮箱',
    constraint username unique (username),
    index idx_username_email (username, email)
)
    comment '用户表';

create table tab_user_order
(
    id            bigint auto_increment comment '订单ID' primary key,
    user_id       bigint                             not null comment '用户ID',
    order_name    varchar(50)                        not null comment '订单名称',
    price         varchar(255)                       not null comment '订单价格',
    status        tinyint  default 1                 not null comment '订单状态：1=正常，0=删除',
    created_time  datetime default CURRENT_TIMESTAMP null comment '订单创建时间',
    finished_time datetime default CURRENT_TIMESTAMP null comment '订单完成时间',
    index idx_user_id (user_id)
)
    comment '用户订单表';

create table tab_user_order_detail
(
    id           bigint auto_increment comment '订单详细ID' primary key,
    order_id     bigint                             not null comment '用户ID',
    detail_name  varchar(50)                        not null comment '订单信息名称',
    status       tinyint  default 1                 not null comment '订单状态：1=未支付，2=已支付，3=商家接单，4=骑手接单，5=配送中，6=已送达，7=已完成',
    created_time datetime default CURRENT_TIMESTAMP null comment '创建时间',
    index idx_order_id_status (order_id, status)
)
    comment '用户订单详细表';

INSERT INTO tab_user (id, username, password, phone, email) VALUES (1, 'user1', 'user1-pwd', 'user1-phone', 'user1-email');
INSERT INTO tab_user (id, username, password, phone, email) VALUES (2, 'user2', 'user2-pwd', 'user2-phone', 'user2-email');
INSERT INTO tab_user (id, username, password, phone, email) VALUES (3, 'user3', 'user3-pwd', 'user3-phone', 'user3-email');


INSERT INTO tab_user_order (id, user_id, order_name, price, status, created_time, finished_time) VALUES (1, 1, '斐济杯', '666', 1, '2025-09-04 10:42:59', '2025-09-04 14:38:57');
INSERT INTO tab_user_order (id, user_id, order_name, price, status, created_time, finished_time) VALUES (2, 1, '猴油', '333', 1, '2025-09-04 10:43:38', null);
INSERT INTO tab_user_order (id, user_id, order_name, price, status, created_time, finished_time) VALUES (3, 1, '安全', '111', 2, '2025-09-04 10:43:58', null);
INSERT INTO tab_user_order (id, user_id, order_name, price, status, created_time, finished_time) VALUES (4, 2, '蛋糕', '1', 2, '2025-09-04 10:44:28', null);
INSERT INTO tab_user_order (id, user_id, order_name, price, status, created_time, finished_time) VALUES (5, 2, '唇膏', '123', 1, '2025-09-04 10:44:56', '2025-09-04 14:39:03');
INSERT INTO tab_user_order (id, user_id, order_name, price, status, created_time, finished_time) VALUES (6, 3, '玩具', '12', 1, '2025-09-04 10:44:45', null);


INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (1, 1, 'q QQ', 1, '2025-09-04 14:25:16');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (2, 1, 'q QQ', 2, '2025-09-04 14:25:20');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (3, 1, 'q QQ', 3, '2025-09-04 14:25:21');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (4, 1, 'q QQ', 4, '2025-09-04 14:25:23');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (5, 1, 'q QQ', 5, '2025-09-04 14:25:24');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (6, 1, 'q QQ', 6, '2025-09-04 14:25:26');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (7, 1, 'q QQ', 7, '2025-09-04 14:25:34');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (8, 5, 'aaa', 1, '2025-09-04 14:26:16');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (9, 5, 'aaa', 2, '2025-09-04 14:26:17');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (10, 5, 'aaa', 3, '2025-09-04 14:26:23');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (11, 5, 'aaa', 4, '2025-09-04 14:26:29');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (12, 5, 'aaa', 5, '2025-09-04 14:26:31');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (13, 5, 'aaa', 6, '2025-09-04 14:26:33');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (14, 5, 'aaa', 7, '2025-09-04 14:26:35');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (15, 2, 'www', 1, '2025-09-04 14:27:21');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (16, 2, 'www', 2, '2025-09-04 14:27:23');
INSERT INTO tab_user_order_detail (id, order_id, detail_name, status, created_time) VALUES (17, 2, 'www', 3, '2025-09-04 14:27:25');


*/

type TabUser struct {
	Id       int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement;not null;comment:用户ID" json:"id" form:"id"`                                         // 用户ID
	Username string `gorm:"column:username;type:varchar(50);unique;not null;index:idx_username_email,priority:1;comment:用户名" json:"username" form:"username"` // 用户名
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password" form:"password"`                                            // 密码
	Phone    string `gorm:"column:phone;type:varchar(100);comment:手机号" json:"phone" form:"phone"`                                                             // 手机号
	Email    string `gorm:"column:email;type:varchar(100);index:idx_username_email,priority:2;comment:邮箱" json:"email" form:"email"`                          // 邮箱
}

func (tu *TabUser) TableName() string {
	return "tab_user"
}

type TabUserOrder struct {
	Id           int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement;not null;comment:订单ID" json:"id" form:"id"`                                           // 订单ID
	UserId       int64     `gorm:"column:user_id;type:bigint;not null;index:idx_user_id,priority:1;comment:用户ID" json:"userId" form:"userId"`                          // 用户ID
	OrderName    string    `gorm:"column:order_name;type:varchar(50);not null;comment:订单名称" json:"orderName" form:"orderName"`                                         // 订单名称
	Price        string    `gorm:"column:price;type:varchar(255);not null;comment:订单价格" json:"price" form:"price"`                                                     // 订单价格
	Status       int8      `gorm:"column:status;type:tinyint;not null;default:1;comment:订单状态：1=正常，0=删除" json:"status" form:"status"`                                   // 订单状态：1=正常，0=删除
	CreatedTime  time.Time `gorm:"column:created_time;type:datetime;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:订单创建时间" json:"createdTime" form:"createdTime"`    // 订单创建时间
	FinishedTime time.Time `gorm:"column:finished_time;type:datetime;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:订单完成时间" json:"finishedTime" form:"finishedTime"` // 订单完成时间
}

func (tuo *TabUserOrder) TableName() string {
	return "tab_user_order"
}

type TabUserOrderDetail struct {
	Id          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement;not null;comment:订单详细ID" json:"id" form:"id"`                                                                              // 订单详细ID
	OrderId     int64     `gorm:"column:order_id;type:bigint;not null;index:idx_order_id_status,priority:1;comment:用户ID" json:"orderId" form:"orderId"`                                                    // 用户ID
	DetailName  string    `gorm:"column:detail_name;type:varchar(50);not null;comment:订单信息名称" json:"detailName" form:"detailName"`                                                                         // 订单信息名称
	Status      int8      `gorm:"column:status;type:tinyint;not null;default:1;index:idx_order_id_status,priority:2;comment:订单状态：1=未支付，2=已支付，3=商家接单，4=骑手接单，5=配送中，6=已送达，7=已完成" json:"status" form:"status"` // 订单状态：1=未支付，2=已支付，3=商家接单，4=骑手接单，5=配送中，6=已送达，7=已完成
	CreatedTime time.Time `gorm:"column:created_time;type:datetime;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:创建时间" json:"createdTime" form:"createdTime"`                                           // 创建时间
}

func (tuod *TabUserOrderDetail) TableName() string {
	return "tab_user_order_detail"
}

// 根据以上结构体实现 [{user, userOrder:[{order, orderDetail:[{},{},{}]}]}] 的返回结构
// 实现一对多查询常见有 4 种方式：
// 方法1:结构体 + GORM 关联标签 + Preload（最常用，ORM 风格，代码最简洁）。
// 方法2:分步 Association 查询（适合只要部分数据，避免大 preload）。
// 方法3:Joins + Scan（适合复杂条件和查询优化）。
// 方法4:Raw SQL + Scan（最灵活，适合性能要求高的复杂查询）
func main() {
	DB := getDB("", "test1")

	fmt.Println(DB)

	// tab_user -> tab_user_order
	// test1(DB)

	// tab_user -> tab_user_order -> tab_user_order_detail
	test2(DB)

}

func test2(db *gorm.DB) {
	userIds := []int64{1, 2, 3}

	// 方法1:结构体 + GORM 关联标签 + Preload（最常用，ORM 风格，代码最简洁）。
	//test21(db, userIds)

	// 方法2:分步 Association 查询（适合只要部分数据，避免大 preload）。
	// 方法3:Joins + Scan（适合复杂条件和查询优化）。
	test23(db, userIds)

	// 方法4:Raw SQL + Scan（最灵活，适合性能要求高的复杂查询）
}

func test23(db *gorm.DB, ids []int64) {
	// 方法3:Joins + Scan（适合复杂条件和查询优化）。
	// 先把所有数据都扁平化到一个结构体里面，之后在内存中构建返回结构体
	type FlatUserExt23 struct {
		TabUser
		TabUserOrder
		TabUserOrderDetail
	}

	var flatUserExt23List []FlatUserExt23
	err := db.Table("tab_user u").
		Select("u.*, uo.*, uod.*").
		Joins("left join tab_user_order uo on u.id = uo.user_id and uo.status = 1").
		Joins("left join tab_user_order_detail uod on uod.order_id = uo.id").
		Where("u.id in ?", ids).
		Find(&flatUserExt23List).Error
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// 组装
	type UserOrderExt23Result struct {
		TabUserOrder
		TabUserOrderDetail []*TabUserOrderDetail `gorm:"foreignKey:OrderId;references:Id"`
	}
	type UserExt23Result struct {
		TabUser
		UserOrderExt2 []*UserOrderExt23Result `gorm:"foreignKey:UserId;references:Id"`
	}
	var userMap map[int64]*UserExt23Result = make(map[int64]*UserExt23Result)                // 收集用户
	var userOrderMap map[int64]*UserOrderExt23Result = make(map[int64]*UserOrderExt23Result) // 收集订单
	for _, userExt23 := range flatUserExt23List {
		if _, ok := userMap[userExt23.TabUser.Id]; !ok {
			userMap[userExt23.TabUser.Id] = &UserExt23Result{
				TabUser:       userExt23.TabUser,
				UserOrderExt2: make([]*UserOrderExt23Result, 0),
			}
		}
		if _, ok := userOrderMap[userExt23.TabUserOrder.Id]; !ok {
			userOrderMap[userExt23.TabUserOrder.Id] = &UserOrderExt23Result{
				TabUserOrder:       userExt23.TabUserOrder,
				TabUserOrderDetail: make([]*TabUserOrderDetail, 0),
			}

			// 子第一次遍历到，把地址加入父级
			if userExt23.TabUserOrder.Id > 0 { // 空值的订单不加入
				userMap[userExt23.TabUser.Id].UserOrderExt2 = append(userMap[userExt23.TabUser.Id].UserOrderExt2, userOrderMap[userExt23.TabUserOrder.Id])
			}
		}
		// 孙第一次出现，加入子
		if userExt23.TabUserOrderDetail.Id > 0 { // 空值的订单详细不加入
			userOrderMap[userExt23.TabUserOrder.Id].TabUserOrderDetail = append(userOrderMap[userExt23.TabUserOrder.Id].TabUserOrderDetail, &userExt23.TabUserOrderDetail)
		}
	}

	// map2slice
	var res []UserExt23Result = make([]UserExt23Result, 0, len(userMap))
	for _, userExt23Result := range userMap {
		res = append(res, *userExt23Result)
	}

	for _, val := range res {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Printf("\n\n\n")
	for _, u := range res {
		fmt.Printf("User: %d %s\n", u.Id, u.Username)
		for _, order := range u.UserOrderExt2 {
			fmt.Printf("  Order: %d %s\n", order.Id, order.OrderName)
			for _, detail := range order.TabUserOrderDetail {
				fmt.Printf("    Detail: %d %s\n", detail.Id, detail.DetailName)
			}
		}
	}

	fmt.Println()

}

func test21(db *gorm.DB, ids []int64) {
	// 方法1:结构体 + GORM 关联标签 + Preload（最常用，ORM 风格，代码最简洁）。
	type UserOrderExt2 struct {
		TabUserOrder
		UserOrderDetail []TabUserOrderDetail `gorm:"foreignKey:OrderId;references:Id"`
	}
	type UserExt2 struct {
		TabUser
		UserOrderExt2 []UserOrderExt2 `gorm:"foreignKey:UserId;references:Id"`
	}

	var res []UserExt2
	err := db.Model(&TabUser{}).
		Select("id, username, phone").
		// Preload是独立的资查询，多个Preload直接没有直接的强关联关系，因此他们的顺序不重要
		//Preload("UserOrderExt2", "price >= ? and finished_time is not null", 123).
		//Preload("UserOrderExt2", map[string]interface{}{
		//	"finished_time <>": nil,
		//	"price >=":         123,
		//}).
		Preload("UserOrderExt2", func(db *gorm.DB) *gorm.DB {
			// 如果字表只要查询某些字段，则必须使用回调函数，
			// 注意：如果字表使用了Select()函数查询指定的字段，那么关联的字段必须要查出来，否则gorm无法进行数据组装。在这里就是tab_user_order.user_id 字段，
			return db.Select("id, user_id, order_name, price, status, created_time").
				Where("price >= ?", 123).Where("finished_time IS NOT NULL")
		}).
		Preload("UserOrderExt2.UserOrderDetail").
		Where("id in ? ", ids).
		Find(&res).Error
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	for _, val := range res {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()

}

func test33(db *gorm.DB, ids []int64) {

	// 方法3:Joins + Scan（适合复杂条件和查询优化）。
	type UserOrderDetailExt33 struct {
		TabUserOrder
		UserOrderDetail []TabUserOrderDetail `gorm:"foreignKey:OrderId;references:Id"`
	}
	type UserOrderExt33 struct {
		TabUser
		UserOrder []UserOrderDetailExt33 `gorm:"foreignKey:UserId;references:Id"`
	}
	var userOrderExt13 []UserOrderExt33
	err := db.Table("tab_user u").
		Select("u.*, uo.*, uod.*").
		Joins("left join tab_user_order uo on u.id = uo.user_id and uo.status = 1").
		Joins("left join tab_user_order_detail uod on uod.order_id = uo.id").
		Where("u.id in ?", []int64{1, 2}).
		Find(&userOrderExt13).Error
	// Scan 只能扫描 平铺的列到结构体字段，不能自动填充 slice 类型的关联（如 UserOrder、UserOrderDetail）。
	//Scan(&userOrderExt13).Error
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	for _, val := range userOrderExt13 {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()

}

func test1(db *gorm.DB) {
	// tab_user -> tab_user_order

	userIds := []int64{1, 2, 3}

	// 方法1:结构体 + GORM 关联标签 + Preload（最常用，ORM 风格，代码最简洁）。
	//test11(db, userIds)

	// 方法2:分步 Association 查询（适合只要部分数据，避免大 preload）。
	//test12(db, userIds)
	//test122(db, userIds) // 先查主表，再查子表，最后内存进行组装

	// 方法3:Joins + Scan（适合复杂条件和查询优化）。
	//test13(db, userIds)

	// 方法4:Raw SQL + Scan（最灵活，适合性能要求高的复杂查询）
	test14(db, userIds)
}

func test14(db *gorm.DB, userIds []int64) {
	// 方法4:Raw SQL + Scan（最灵活，适合性能要求高的复杂查询）
	// Raw SQL + Scan可以认为就是Joins + Scan的原生sql写法

	//SELECT u.*, uo.*, uod.*
	//	FROM tab_user u
	//left join tab_user_order uo on u.id = uo.user_id and uo.status = 1
	//left join tab_user_order_detail uod on uod.order_id = uo.id
	//WHERE u.id in (1,2)

	type UserOrderExt14 struct {
		TabUser
		TabUserOrder
	}
	type UserOrderExt14Result struct {
		TabUser
		UserOrders []TabUserOrder `gorm:"foreignKey:UserId;references:Id"`
	}

	var userOrderExt14 []UserOrderExt14
	err := db.Raw(`
		SELECT u.*, uo.*, uod.*
		FROM tab_user u
			left join tab_user_order uo on u.id = uo.user_id and uo.status = 1
			left join tab_user_order_detail uod on uod.order_id = uo.id
		WHERE u.id in ?`, userIds).
		Scan(&userOrderExt14).Error
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// 组装方式2
	userMap := make(map[int64]*UserOrderExt14Result)
	for _, userOrderExt13 := range userOrderExt14 {
		if _, ok := userMap[userOrderExt13.UserId]; !ok {
			// 没记录，标识第一次遍历到改userId，创建UserOrderExt14Result里面的slice对象
			userMap[userOrderExt13.UserId] = &UserOrderExt14Result{
				TabUser:    userOrderExt13.TabUser,
				UserOrders: make([]TabUserOrder, 0),
			}
		}
		// 追加数据
		userMap[userOrderExt13.UserId].UserOrders = append(userMap[userOrderExt13.UserId].UserOrders, userOrderExt13.TabUserOrder)

	}
	// map2slice
	var res []UserOrderExt14Result = make([]UserOrderExt14Result, 0, len(userMap))
	for _, userOrderExt13Result := range userMap {
		res = append(res, *userOrderExt13Result)
	}

	for _, val := range res {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()

}

func test13(db *gorm.DB, userIds []int64) {
	// 方法3:Joins + Scan（适合复杂条件和查询优化）。
	type UserOrderExt13 struct {
		TabUser
		TabUserOrder
	}
	// Scan/Find 不会自动填充嵌套 slice。 只能扫描 平铺的列到结构体字段，不能自动填充 slice 类型的关联（如 UserOrder、UserOrderDetail）。
	// 因此要想使用Scan + Join的方式实现一对多的查询，那么必须先将所有字段都扁平化到一个结构体里面，之后在内存中进行组装

	var userOrderExt13 []UserOrderExt13
	err := db.Table("tab_user u").
		Select("u.*, uo.*").
		Joins("left join tab_user_order uo on u.id = uo.user_id and uo.status = 1").
		Where("u.id in ?", userIds).
		//Find(&userOrderExt13).Error
		Scan(&userOrderExt13).Error
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	type UserOrderExt13Result struct {
		TabUser
		UserOrders []TabUserOrder `gorm:"foreignKey:UserId;references:Id"`
	}

	m := make(map[int64][]UserOrderExt13)
	for _, val := range userOrderExt13 {
		m[val.UserId] = append(m[val.UserId], val)
	}

	var res []UserOrderExt13Result = make([]UserOrderExt13Result, 0)
	for _, val := range m {
		var userOrders []TabUserOrder = make([]TabUserOrder, 0, len(val))
		for _, v := range val {
			userOrders = append(userOrders, v.TabUserOrder)
		}
		res = append(res, UserOrderExt13Result{
			TabUser:    val[0].TabUser,
			UserOrders: userOrders,
		})
	}

	for _, val := range res {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()

	// 组装方式2
	userMap := make(map[int64]*UserOrderExt13Result)
	for _, userOrderExt13 := range userOrderExt13 {
		if _, ok := userMap[userOrderExt13.UserId]; !ok {
			// 没记录，标识第一次遍历到改userId，创建UserOrderExt13Result里面的slice对象
			userMap[userOrderExt13.UserId] = &UserOrderExt13Result{
				TabUser:    userOrderExt13.TabUser,
				UserOrders: make([]TabUserOrder, 0),
			}
		}
		// 追加数据
		userMap[userOrderExt13.UserId].UserOrders = append(userMap[userOrderExt13.UserId].UserOrders, userOrderExt13.TabUserOrder)

	}
	// map2slice
	var res2 []UserOrderExt13Result = make([]UserOrderExt13Result, 0, len(userMap))
	for _, userOrderExt13Result := range userMap {
		res2 = append(res2, *userOrderExt13Result)
	}

	for _, val := range res2 {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()
}

func test122(db *gorm.DB, ids []int64) {
	// 先用一条sql去除所有user，在用一条sql取出所有user的所有订单，最后在内存中组装
	//这样： 只发 2 条 SQL（一次查用户，一次查订单）,避免 Association 的 N+1 查询和 panic 问题

	type UserOrderExt122 struct {
		TabUser
		UserOrders []TabUserOrder `gorm:"foreignKey:UserId;references:Id"`
	}

	// 1、查所有满足条件的user
	var users []TabUser
	err := db.Model(&TabUser{}).Where("id in ?", ids).Find(&users) // 可能要追加更多的where
	if err != nil {
		fmt.Println("err: ", err)
	}
	queryUserIds := make([]int64, 0)
	for _, user := range users {
		queryUserIds = append(queryUserIds, user.Id)
	}

	// 2、查user下面的 order
	var userOrders []TabUserOrder
	err = db.Model(&TabUserOrder{}).
		Where("user_id in ?", queryUserIds).
		Where("status = ?", 1).
		Find(&userOrders)
	if err != nil {
		fmt.Println("err2: ", err)
	}
	var userOrderMap map[int64][]TabUserOrder = make(map[int64][]TabUserOrder)
	for _, order := range userOrders {
		userOrderMap[order.UserId] = append(userOrderMap[order.UserId], order)
	}

	// 3、组装返回
	var userOrderExt122s []UserOrderExt122 = make([]UserOrderExt122, 0, len(users))
	for _, user := range users {
		userOrderExt122s = append(userOrderExt122s, UserOrderExt122{
			TabUser:    user,
			UserOrders: userOrderMap[user.Id],
		})
	}

	// 输出最后结果
	for _, val := range userOrderExt122s {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()

}

func test12(db *gorm.DB, userIds []int64) {
	// 方法2:分步 Association 查询（适合只要部分数据，避免大 preload）。
	//重要提示：Association 模式不支持链式条件（Where/Order/Limit 等）。
	//db.Model(&u).Association("Orders").Find(&orders) 只能按外键关系把所有订单取出来，想加过滤/排序，需改用普通查询

	type UserOrderExt12 struct {
		TabUser
		UserOrders []TabUserOrder `gorm:"foreignKey:UserId;references:Id"`
	}

	// 逐条关联加载（适合小结果集，按需取）实现方法：（要多次查询数据库，类似于先找主表的数据，之后在通过主表的关联字段查找字表的数据）
	//先拿到一批用户，再对每个用户用 Association 拉订单；
	//对每个订单再拉详情。优点是很直观、字段可控、避免一次性大 Preload；
	//缺点是N+1 问题（多次往返 DB）

	// 先读取一批用户详细
	var userOrderExt12List []UserOrderExt12
	err := db.Model(&TabUser{}).Where("id in ?", userIds).
		Find(&userOrderExt12List).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	// 根据userId去关联字表
	for i := range userOrderExt12List {
		var orders []TabUserOrder
		// 注意 Model 必须是单个 struct，不要传整个 slice
		err := db.Model(&userOrderExt12List[i]).Association("UserOrders").Find(&orders)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 字表的条件在这里进行内存过滤
		var res []TabUserOrder = make([]TabUserOrder, 0)
		for _, order := range orders {
			// 假设只要状态是 1 的
			if order.Status == 1 {
				res = append(res, order)
			}
		}
		userOrderExt12List[i].UserOrders = res

	}

	for _, val := range userOrderExt12List {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()

}

func test11(db *gorm.DB, userIds []int64) {
	// 方法1:结构体 + GORM 关联标签 + Preload（最常用，ORM 风格，代码最简洁）。

	type UserOrderExt11 struct {
		TabUser
		UserOrders []TabUserOrder `gorm:"foreignKey:UserId;references:Id"`
	}

	var res []UserOrderExt11
	err := db.Debug().Preload("UserOrders").Model(&TabUser{}).
		Where("id in ?", userIds).
		Find(&res).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("len(res) = ", len(res))

	for _, val := range res {
		fmt.Printf("val = %#v \n\n", val)
	}
	fmt.Println()
}

func getDB(ip, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("root:root123@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", ip, dbName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印所有日志
		Logger: logger.Default.LogMode(logger.Info),
		// 显示出查询的所有字段
		QueryFields: true,
		// 禁用默认事务，true 表示不开起事务；不写或写false表示开启事务
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("数据库连接失败！！！")
	}
	return DB
}
