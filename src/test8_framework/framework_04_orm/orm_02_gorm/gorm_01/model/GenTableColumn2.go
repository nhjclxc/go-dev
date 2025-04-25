package model

import "time"

type GenTableColumn2 struct {

	// 使用 gorm 将数据库里面的字段对应

	// 代码生成业务字段表主键
	ColumnId int64 `gorm:"column:column_id"`

	// 归属表编号
	TableId int64 `gorm:"column:table_id"`

	// 列名称
	ColumnName string `gorm:"column:column_name"`

	// 列描述
	ColumnComment string `gorm:"column:column_comment"`

	// 列类型
	ColumnType string `gorm:"column:column_type"`

	// 是否查询字段（1是）
	IsQuery rune `gorm:"column:is_query"`

	// 创建者
	CreateBy string `gorm:"column:create_by"`

	// 创建时间
	CreateTime time.Time `gorm:"column:create_time"`

	// 更新者
	UpdateBy string `gorm:"column:update_by"`

	// 更新时间
	UpdateTime JSONTime `gorm:"column:update_time"`
}

func (GenTableColumn2) TableName() string {
	return "gen_table_column2"
}
