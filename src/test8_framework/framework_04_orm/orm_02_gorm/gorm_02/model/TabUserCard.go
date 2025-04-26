package model

import (
	"time"
)

/*
create table tab_user_card
(
    user_card_id     bigint auto_increment comment '用户卡ID' primary key,
    user_id          bigint       not null comment '用户ID',
    card_type_id     bigint       not null comment '卡类型ID',
    remark 			 varchar(255) null comment '备注',
    status           char                  default '1' null comment '卡状态（1=在用 2=停用）',
    `created_at`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
)
    comment '用户卡关联表';
*/

// 用户卡结构体
type TabUserCard struct {

	// 用户卡ID
	UserCardId uint `gorm:"primaryKey;autoIncrement;column:user_card_id"`

	// 用户ID
	UserId uint `gorm:"column:user_id"`

	// 卡类型ID
	CardTypeId uint `gorm:"column:card_type_id"`

	// 备注
	Remark string `gorm:"column:remark"`

	// 卡状态（1=在用 2=停用）
	Status rune `gorm:"column:status"`

	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at"`

	// 修改时间
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// 必须使用 TableName() 方法来声明这个结构体对应的表名
// 如果不显示声明表名，那么 gorm 会默认把结构体的驼峰（TabUserCard）转化为下划线并在最后加一个s（复数形式），认为是结构体对应的表名，即（tab_user_cards）
func (TabUserCard) TableName() string {
	return "tab_user_card"
}
