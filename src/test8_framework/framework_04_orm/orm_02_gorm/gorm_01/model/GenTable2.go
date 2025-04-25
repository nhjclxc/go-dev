package model

import "time"

type GenTable2 struct {

	// 使用 gorm 将数据库里面的字段对应

	// 归属表编号
	// 指定为 primaryKey， 为的是新增的时候返回新增的 primaryKey 到实体类参数上面
	TableId int64 `gorm:"primaryKey;column:table_id"`

	// 表名称
	TableName2 string `gorm:"column:table_name2"`

	// 表名称
	TableComment string `gorm:"column:table_comment"`

	// 排序
	Sort int `gorm:"column:sort"`

	// 状态（0=删除，1=在用）
	State int `gorm:"column:state"`

	// 创建者
	CreateBy string `gorm:"column:create_by"`

	// 创建时间
	CreateTime time.Time `gorm:"column:create_time"`

	// 更新者
	UpdateBy string `gorm:"column:update_by"`

	// 更新时间
	UpdateTime JSONTime `gorm:"column:update_time"`

	// // 二、一对多：在主表中写：
	// 下面的 "foreignKey:TableId" 里面的 TableId 是 GenTableColumn2 里面与主表关联的id
	GenTableColumn2 []GenTableColumn2 `gorm:"foreignKey:TableId"`
}

func (GenTable2) TableName() string {
	return "gen_table2"
}
