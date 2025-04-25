package dto

/*
 */
type GenTable2DTO struct {

	// 使用 gorm 将数据库里面的字段对应

	// 归属表编号
	TableId int64 `json:"tableId" form:"tableId"`

	// 表名称
	TableName2 string `json:"tableName2" form:"tableName2"`

	// 表名称
	TableComment string `json:"tableComment" form:"tableComment"`

	// 排序
	Sort int `json:"sort" form:"sort"`

	// 状态（0=删除，1=在用）
	State int `json:"state" form:"state"`

	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`

	//// 创建者
	//CreateBy string `json:"name1" form:"username"`
	//
	//// 创建时间
	//CreateTime time.Time `gorm:"column:create_time"`
	//
	//// 更新者
	//UpdateBy string `gorm:"column:update_by"`
	//
	//// 更新时间
	//UpdateTime JSONTime `gorm:"column:update_time"`
}
