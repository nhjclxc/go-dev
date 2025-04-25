package controller

import (
	"github.com/gin-gonic/gin"
	db "gorm_01/config"
	"gorm_01/model"
	"gorm_01/model/dto"
)

// 定义 GenTableColumn2Controller 结构体来保存这个实例的相关数据
type GenTableColumn2Controller struct {
}

func (this *GenTableColumn2Controller) GetById(context *gin.Context) {

	// http://localhost:8090/genTable2/getById?tableId=0

	// 获取前端的参数
	genTableColunm2DTO := dto.GenTableColunm2DTO{}
	context.BindQuery(&genTableColunm2DTO)

	// 添加条件
	if genTableColunm2DTO.ColumnId == 0 {
		panic("columnId 不能为空！！！")
	}

	// 创建查询对象
	dbQuery := db.DB.Model(&model.GenTableColumn2{})

	dbQuery = dbQuery.Where("column_id = ?", genTableColunm2DTO.ColumnId)

	// 执行查询
	res := []model.GenTableColumn2{}
	// 关联表 GenTable2
	dbQuery.Preload("GenTable2").Find(&res)

	// 响应数据
	context.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}
