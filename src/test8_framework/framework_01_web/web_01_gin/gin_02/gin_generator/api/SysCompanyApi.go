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

// SysCompanyApi 公司 api 层
type SysCompanyApi struct{
    sysCompanyService service.SysCompanyService
}

// InsertSysCompany 新增公司
// @Tags 公司模块
// @Summary 新增公司-Summary
// @Description 新增公司-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysCompany body model.SysCompany true "修改公司实体类"
// @Success 200 {object} commonModel.JsonResult "新增公司响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company [post]
func (this *SysCompanyApi) InsertSysCompany(c *gin.Context) {
    var sysCompany model.SysCompany
    c.ShouldBindJSON(&sysCompany)

    res, err := this.sysCompanyService.InsertSysCompany(&sysCompany)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("新增公司失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// UpdateSysCompany 修改公司
// @Tags 公司模块
// @Summary 修改公司-Summary
// @Description 修改公司-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysCompany body model.SysCompany true "修改公司实体类"
// @Success 200 {object} commonModel.JsonResult "修改公司响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company [put]
func (this *SysCompanyApi) UpdateSysCompany(c *gin.Context) {
    var sysCompany model.SysCompany
    c.ShouldBindJSON(&sysCompany)

    res, err := this.sysCompanyService.UpdateSysCompany(&sysCompany)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("修改公司失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// DeleteSysCompany 删除公司
// @Tags 公司模块
// @Summary 删除公司-Summary
// @Description 删除公司-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param idList path string true "公司主键List"
// @Success 200 {object} commonModel.JsonResult "删除公司响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company/:idList [delete]
func (this *SysCompanyApi) DeleteSysCompany(c *gin.Context) {
    idListStr := c.Param("idList") // 例如: "1,2,3"
    idList, err := commonUtils.ParseIds(idListStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.sysCompanyService.DeleteSysCompany(idList)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("删除公司失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysCompanyById 获取公司详细信息
// @Tags 公司模块
// @Summary 获取公司详细信息-Summary
// @Description 获取公司详细信息-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param id path int64 true "公司主键List"
// @Success 200 {object} commonModel.JsonResult "获取公司详细信息"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company/:id [get]
func (this *SysCompanyApi) GetSysCompanyById(c *gin.Context) {
    idStr := c.Param("id") // 例如: "1"
    id, err := commonUtils.ParseId(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.sysCompanyService.GetSysCompanyById(id)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询公司失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysCompanyList 查询公司列表
// @Tags 公司模块
// @Summary 查询公司列表-Summary
// @Description 查询公司列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysCompanyDto body model.SysCompanyDto true "公司实体Dto"
// @Success 200 {object} commonModel.JsonResult "查询公司列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company/list [get]
func (this *SysCompanyApi) GetSysCompanyList(c *gin.Context) {
    var sysCompanyDto dto.SysCompanyDto
    c.ShouldBindQuery(&sysCompanyDto)

    res, err := this.sysCompanyService.GetSysCompanyList(&sysCompanyDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询公司列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysCompanyPageList 分页查询公司列表
// @Tags 公司模块
// @Summary 分页查询公司列表-Summary
// @Description 分页查询公司列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysCompanyDto body model.SysCompanyDto true "公司实体Dto"
// @Success 200 {object} commonModel.JsonResult "分页查询公司列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company/pageList [get]
func (this *SysCompanyApi) GetSysCompanyPageList(c *gin.Context) {
    var sysCompanyDto dto.SysCompanyDto
    c.ShouldBindQuery(&sysCompanyDto)

    res, err := this.sysCompanyService.GetSysCompanyPageList(&sysCompanyDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询公司列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// ExportSysCompany 导出公司列表
// @Tags 公司模块
// @Summary 导出公司列表-Summary
// @Description 导出公司列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysCompanyDto body model.SysCompanyDto true "公司实体Dto"
// @Success 200 {object} commonModel.JsonResult "导出公司列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/company/export [get]
func (this *SysCompanyApi) ExportSysCompany(c *gin.Context) {
    var sysCompanyDto dto.SysCompanyDto
    c.ShouldBindQuery(&sysCompanyDto)

    res, err := this.sysCompanyService.ExportSysCompany(&sysCompanyDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("导出公司列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

