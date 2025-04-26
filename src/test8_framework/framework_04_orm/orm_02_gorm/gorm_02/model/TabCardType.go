package model

import (
	"time"
)

/*
create table tab_card_type
(
    card_type_id   bigint auto_increment comment '卡类型ID' primary key,
    card_type_name varchar(128) null comment '卡类型名称',
    card_type_addr varchar(128) null comment '卡地址',
    status         char                  default '1' null comment '卡类型状态（1=可用 0=不可用）',
    `created_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
)
    comment '卡类型表';
*/

// 用户卡结构体
type TabCardType struct {

	// 卡类型ID
	CardTypeId uint `gorm:"primaryKey;autoIncrement;column:card_type_id"`

	// 卡类型名称
	CardTypeName string `gorm:"column:card_type_name"`

	// 卡地址
	Card_typeAddr string `gorm:"column:card_type_addr"`

	// 卡类型状态（1=可用 0=不可用）
	Status rune `gorm:"column:status"`

	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at"`

	// 修改时间
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
