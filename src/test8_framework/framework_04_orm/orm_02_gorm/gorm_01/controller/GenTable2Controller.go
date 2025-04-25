package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	db "gorm_01/config"
	"gorm_01/model"
	"gorm_01/model/dto"
	"time"
)

// 定义 GenTable2Controller 结构体来保存这个实例的相关数据
type GenTable2Controller struct {
	// genTable2Service GenTable2Service

}

// 定义 GenTable2Controller 对应的接口方法
func (this *GenTable2Controller) Insert(context *gin.Context) {

	// 获取前端的参数
	genTable2DTO := dto.GenTable2DTO{}
	context.BindJSON(&genTable2DTO)

	// 在 Go 中使用 Gin + GORM 项目中，将前端传入的 DTO（数据传输对象）拷贝到数据库用的 Model（实体）结构体，是很常见的场景。
	//Go 没有像 Java 的 BeanUtils 或 MyBatis Plus 那样的自动属性拷贝内建功能，但可以通过以下几种方式实现：
	// github.com/jinzhu/copier

	genTable2 := model.GenTable2{}
	copier.Copy(&genTable2, &genTable2DTO) // 自动将字段拷贝进去

	genTable2.CreateTime = time.Now()
	genTable2.UpdateTime = model.JSONTime(time.Now())
	// 连接会话
	result := db.DB.Create(&genTable2)

	fmt.Println("新增ID：", genTable2.TableId)
	fmt.Println("受影响行数：", result.RowsAffected)

	// 响应数据
	context.JSON(200, gin.H{
		"code": 200,
		"data": genTable2.TableId,
	})

}

func (this *GenTable2Controller) Delete(context *gin.Context) {

	// 获取前端的参数
	genTable2DTO := dto.GenTable2DTO{}
	context.BindQuery(&genTable2DTO)

	// 创建查询对象
	dbQuery := db.DB.Model(&model.GenTable2{})

	// 添加条件
	if genTable2DTO.TableId == 0 {
		panic("id不能为空！！！")
	}
	// ?tableId=2
	dbQuery = dbQuery.Where("table_id = ?", genTable2DTO.TableId)

	// 执行查询
	result := dbQuery.Delete(&model.GenTable2{})

	// 响应数据
	context.JSON(200, gin.H{
		"code": 200,
		"data": result.RowsAffected,
	})
}

func (this *GenTable2Controller) Update(context *gin.Context) {

	// 获取前端的参数
	genTable2DTO := dto.GenTable2DTO{}
	context.BindJSON(&genTable2DTO)

	// dto 属性拷贝到 model 里面
	genTable2 := model.GenTable2{}
	copier.Copy(&genTable2, &genTable2DTO) // 自动将字段拷贝进去

	genTable2.CreateTime = time.Now()
	genTable2.UpdateTime = model.JSONTime(time.Now())
	// 连接会话
	result := db.DB.Create(&genTable2)

	fmt.Println("新增ID：", genTable2.TableId)
	fmt.Println("受影响行数：", result.RowsAffected)

	// 响应数据
	context.JSON(200, gin.H{
		"code": 200,
		"data": genTable2.TableId,
	})

	// ✅ 修改单字段：db.Model(&User{}).Where("id = ?", 1).Update("name", "NewName")

	// ✅ 修改多个字段：db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "Tom", Age: 32})

	// ✅ 使用 map 修改多个字段（推荐）：
	//db.Model(&User{}).Where("id = ?", 1).Updates(map[string]interface{}{
	//	"name": "Jerry",
	//	"age":  29,
	//})

}

func (this *GenTable2Controller) GetById(context *gin.Context) {

	// http://localhost:8090/genTable2/getById?tableId=0

	// 获取前端的参数
	genTable2DTO := dto.GenTable2DTO{}
	context.BindQuery(&genTable2DTO)

	// 创建查询对象
	dbQuery := db.DB.Model(&model.GenTable2{})

	// 添加条件
	if genTable2DTO.TableId == 0 {
		panic("id不能为空！！！")
	}
	// ?tableId=2
	dbQuery = dbQuery.Where("table_id = ?", genTable2DTO.TableId)

	// 执行查询
	res := []model.GenTable2{}
	dbQuery.Find(&res)

	// 响应数据
	context.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

func (this *GenTable2Controller) GetAll(context *gin.Context) {

	// 创建查询对象
	res := []model.GenTable2{}
	db.DB.Find(&res)

	context.JSON(200, gin.H{
		"data": res,
	})
}

func (this *GenTable2Controller) GetList(context *gin.Context) {

	// 获取前端的参数
	genTable2DTO := dto.GenTable2DTO{}
	context.BindQuery(&genTable2DTO)

	// 创建查询对象
	dbQuery := db.DB.Model(&model.GenTable2{})

	// 添加条件
	if genTable2DTO.Sort != 0 {
		// ? 是参数的占位符必填
		// ?sort=2
		dbQuery = dbQuery.Where("sort >= ?", genTable2DTO.Sort)
	}
	if genTable2DTO.TableId != 0 {
		// ?tableId=2
		dbQuery = dbQuery.Where("table_id = ?", genTable2DTO.TableId)
	}
	if genTable2DTO.TableName2 != "" {
		// ?tableName2=job
		dbQuery = dbQuery.Where("table_name2 like ?", "%"+genTable2DTO.TableName2+"%")
	}
	if genTable2DTO.TableComment != "" {
		// ?tableComment=定时
		dbQuery = dbQuery.Where("table_comment like ?", "%"+genTable2DTO.TableComment+"%")
	}

	// 执行查询
	res := []model.GenTable2{}
	dbQuery.Find(&res)

	// 响应数据
	context.JSON(200, gin.H{
		"data": res,
	})
}

func (this *GenTable2Controller) GetPageList(context *gin.Context) {

	// 获取前端的参数
	genTable2DTO := dto.GenTable2DTO{}
	context.BindQuery(&genTable2DTO)

	if genTable2DTO.PageNum < 1 {
		genTable2DTO.PageNum = 1
	}
	if genTable2DTO.PageSize < 1 {
		genTable2DTO.PageSize = 10
	}

	// 创建查询对象
	dbQuery := db.DB.Model(&model.GenTable2{})

	if genTable2DTO.TableName2 != "" {
		// ?tableName2=job
		dbQuery = dbQuery.Where("table_name2 like ?", "%"+genTable2DTO.TableName2+"%")
	}
	// http://localhost:8090/genTable2/getPageList?pageNum=2&pageSize=1

	// 查询总计
	var total int64 = 0
	dbQuery.Model(model.GenTable2{}).Count(&total)

	// 执行详细记录查询
	res := []model.GenTable2{}
	if total != 0 {
		// 先加 Limit Offset 条件
		// 再执行 Find 查询命令
		dbQuery.Limit(genTable2DTO.PageSize).Offset((genTable2DTO.PageNum - 1) * genTable2DTO.PageSize).Find(&res)
	}

	// 响应数据
	context.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"list":     res,
			"total":    total,
			"pageNum":  genTable2DTO.PageNum,
			"pageSize": genTable2DTO.PageSize,
		},
	})
}
