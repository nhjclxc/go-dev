package api

import (
    "gin_generator/model"
    commonModel "gin_generator/model/common"
    "gin_generator/model/dto"
    "gin_generator/service"
    "gin_generator/utils/commonUtils"
    "github.com/gin-gonic/gin"
    "net/http"
)

// GenTableApi 代码生成业务 api 层
type GenTableApi struct{
    genTableService service.GenTableService
}

// InsertGenTable 新增代码生成业务
// @Tags 代码生成业务模块
// @Summary 新增代码生成业务-Summary
// @Description 新增代码生成业务-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param genTable body model.GenTable true "修改代码生成业务实体类"
// @Success 200 {object} commonModel.JsonResult "新增代码生成业务响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table [post]
func (this *GenTableApi) InsertGenTable(c *gin.Context) {
    var genTable model.GenTable
    c.ShouldBindJSON(&genTable)

    res, err := this.genTableService.InsertGenTable(&genTable)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("新增代码生成业务失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// UpdateGenTable 修改代码生成业务
// @Tags 代码生成业务模块
// @Summary 修改代码生成业务-Summary
// @Description 修改代码生成业务-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param genTable body model.GenTable true "修改代码生成业务实体类"
// @Success 200 {object} commonModel.JsonResult "修改代码生成业务响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table [put]
func (this *GenTableApi) UpdateGenTable(c *gin.Context) {
    var genTable model.GenTable
    c.ShouldBindJSON(&genTable)

    res, err := this.genTableService.UpdateGenTable(&genTable)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("修改代码生成业务失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusInternalServerError, commonModel.JsonResultSuccess[any](res))
}

// DeleteGenTable 删除代码生成业务
// @Tags 代码生成业务模块
// @Summary 删除代码生成业务-Summary
// @Description 删除代码生成业务-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param tableIdList path string true "代码生成业务主键List"
// @Success 200 {object} commonModel.JsonResult "删除代码生成业务响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table/:tableIdList [delete]
func (this *GenTableApi) DeleteGenTable(c *gin.Context) {
    tableIdListStr := c.Param("tableIdList") // 例如: "1,2,3"
    tableIdList, err := commonUtils.ParseIds(tableIdListStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.genTableService.DeleteGenTable(tableIdList)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("删除代码生成业务失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetGenTableByTableId 获取代码生成业务详细信息
// @Tags 代码生成业务模块
// @Summary 获取代码生成业务详细信息-Summary
// @Description 获取代码生成业务详细信息-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param tableId path int64 true "代码生成业务主键List"
// @Success 200 {object} commonModel.JsonResult "获取代码生成业务详细信息"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table/:tableId [get]
func (this *GenTableApi) GetGenTableByTableId(c *gin.Context) {
    tableIdStr := c.Param("tableId") // 例如: "1"
    tableId, err := commonUtils.ParseId(tableIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.genTableService.GetGenTableByTableId(tableId)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询代码生成业务失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetGenTableList 查询代码生成业务列表
// @Tags 代码生成业务模块
// @Summary 查询代码生成业务列表-Summary
// @Description 查询代码生成业务列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param genTableDto body model.GenTableDto true "代码生成业务实体Dto"
// @Success 200 {object} commonModel.JsonResult "查询代码生成业务列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table/list [get]
func (this *GenTableApi) GetGenTableList(c *gin.Context) {
    var genTableDto dto.GenTableDto
    c.ShouldBindQuery(&genTableDto)

    res, err := this.genTableService.GetGenTableList(&genTableDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询代码生成业务列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetGenTablePageList 分页查询代码生成业务列表
// @Tags 代码生成业务模块
// @Summary 分页查询代码生成业务列表-Summary
// @Description 分页查询代码生成业务列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param genTableDto body model.GenTableDto true "代码生成业务实体Dto"
// @Success 200 {object} commonModel.JsonResult "分页查询代码生成业务列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table/pageList [get]
func (this *GenTableApi) GetGenTablePageList(c *gin.Context) {
    var genTableDto dto.GenTableDto
    c.ShouldBindQuery(&genTableDto)

    res, err := this.genTableService.GetGenTablePageList(&genTableDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询代码生成业务列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// ExportGenTable 导出代码生成业务列表
// @Tags 代码生成业务模块
// @Summary 导出代码生成业务列表-Summary
// @Description 导出代码生成业务列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param genTableDto body model.GenTableDto true "代码生成业务实体Dto"
// @Success 200 {object} commonModel.JsonResult "导出代码生成业务列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /gen/table/export [get]
func (this *GenTableApi) ExportGenTable(c *gin.Context) {
    var genTableDto dto.GenTableDto
    c.ShouldBindQuery(&genTableDto)

    res, err := this.genTableService.ExportGenTable(&genTableDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("导出代码生成业务列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

