package dto

/*
 */
type GenTableColunm2DTO struct {

	// 使用 gorm 将数据库里面的字段对应

	// 代码生成业务字段表主键
	ColumnId int64 `json:"columnId" form:"columnId"`

	// 归属表编号
	TableId int64 `json:"tableId" form:"tableId"`

	// 列名称
	ColumnName string `json:"columnName" form:"columnName"`

	// 列描述
	ColumnComment string `json:"columnComment" form:"columnComment"`

	// 列类型
	ColumnType string `json:"columnType" form:"columnType"`

	// 是否查询字段（1是）
	IsQuery rune `json:"isQuery" form:"isQuery"`
}
