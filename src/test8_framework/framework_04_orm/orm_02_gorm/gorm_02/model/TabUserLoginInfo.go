package model

/*
create table tab_user_login_info
(

	login_info_id  bigint auto_increment comment '登录记录ID' primary key,
	user_id        bigint           null comment '用户ID',
	ipaddr         varchar(128)     null comment '登录IP地址',
	login_location varchar(255)     null comment '登录地点',
	browser        varchar(50)      null comment '浏览器类型',
	os             varchar(50)      null comment '操作系统',
	status         char default '1' null comment '登录状态（1成功 0失败）',
	msg            varchar(255)     null comment '提示消息',
	login_time     datetime         null comment '访问时间'

)

	comment '用户登录记录';
*/
type TabUserLoginInfo struct {

	// 登录记录ID
	LoginInfoId uint `gorm:"primaryKey;autoIncrement;column:login_info_id"`

	// 用户ID
	UserId uint `gorm:"column:user_id"`

	// 登录IP地址
	Ipaddr string `gorm:"column:ipaddr"`

	// 登录地点
	LoginLocation string `gorm:"column:login_location"`

	// 浏览器类型
	Browser string `gorm:"column:browser"`

	// 操作系统
	Os string `gorm:"column:os"`

	// 登录状态（1成功 0失败）
	// 使用 default:1 来赋默认值
	// 这些默认值会被当作结构体字段的零值插入到数据库中
	// 注意，当结构体的字段默认值是零值的时候比如 0, '', false，这些字段值将不会被保存到数据库中，你可以使用指针类型或者Scanner/Valuer来避免这种情况。
	Status rune `gorm:"column:status;default:1"`

	// 注意，当结构体的字段默认值是零值的时候比如 0, '', false，这些字段值将不会被保存到数据库中，你可以使用指针类型或者Scanner/Valuer来避免这种情况。
	// 由于要给数据库的默认值是0，而0在go里面是一个零值，此时go不会将这个零值传给数据库
	// 因此，我们就需要传地址（*rune）进去，不让go过滤掉零值
	//Status *rune `gorm:"column:status;default:0"`

	// 提示消息
	Msg string `gorm:"column:msg"`

	// 访问时间
	LoginTime string `gorm:"column:login_time"`
}
