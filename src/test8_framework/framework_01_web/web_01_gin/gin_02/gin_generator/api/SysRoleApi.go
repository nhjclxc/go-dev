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

// SysRoleApi 角色信息 api 层
type SysRoleApi struct{
    sysRoleService service.SysRoleService
}

// InsertSysRole 新增角色信息
// @Tags 角色信息模块
// @Summary 新增角色信息-Summary
// @Description 新增角色信息-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysRole body model.SysRole true "修改角色信息实体类"
// @Success 200 {object} commonModel.JsonResult "新增角色信息响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role [post]
func (this *SysRoleApi) InsertSysRole(c *gin.Context) {
    var sysRole model.SysRole
    c.ShouldBindJSON(&sysRole)

    res, err := this.sysRoleService.InsertSysRole(&sysRole)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("新增角色信息失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// UpdateSysRole 修改角色信息
// @Tags 角色信息模块
// @Summary 修改角色信息-Summary
// @Description 修改角色信息-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysRole body model.SysRole true "修改角色信息实体类"
// @Success 200 {object} commonModel.JsonResult "修改角色信息响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role [put]
func (this *SysRoleApi) UpdateSysRole(c *gin.Context) {
    var sysRole model.SysRole
    c.ShouldBindJSON(&sysRole)

    res, err := this.sysRoleService.UpdateSysRole(&sysRole)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("修改角色信息失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// DeleteSysRole 删除角色信息
// @Tags 角色信息模块
// @Summary 删除角色信息-Summary
// @Description 删除角色信息-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param roleIdList path string true "角色信息主键List"
// @Success 200 {object} commonModel.JsonResult "删除角色信息响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role/:roleIdList [delete]
func (this *SysRoleApi) DeleteSysRole(c *gin.Context) {
    roleIdListStr := c.Param("roleIdList") // 例如: "1,2,3"
    roleIdList, err := commonUtils.ParseIds(roleIdListStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.sysRoleService.DeleteSysRole(roleIdList)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("删除角色信息失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysRoleByRoleId 获取角色信息详细信息
// @Tags 角色信息模块
// @Summary 获取角色信息详细信息-Summary
// @Description 获取角色信息详细信息-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param roleId path int64 true "角色信息主键List"
// @Success 200 {object} commonModel.JsonResult "获取角色信息详细信息"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role/:roleId [get]
func (this *SysRoleApi) GetSysRoleByRoleId(c *gin.Context) {
    roleIdStr := c.Param("roleId") // 例如: "1"
    roleId, err := commonUtils.ParseId(roleIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.sysRoleService.GetSysRoleByRoleId(roleId)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询角色信息失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysRoleList 查询角色信息列表
// @Tags 角色信息模块
// @Summary 查询角色信息列表-Summary
// @Description 查询角色信息列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysRoleDto body model.SysRoleDto true "角色信息实体Dto"
// @Success 200 {object} commonModel.JsonResult "查询角色信息列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role/list [get]
func (this *SysRoleApi) GetSysRoleList(c *gin.Context) {
    var sysRoleDto dto.SysRoleDto
    c.ShouldBindQuery(&sysRoleDto)

    res, err := this.sysRoleService.GetSysRoleList(&sysRoleDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询角色信息列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysRolePageList 分页查询角色信息列表
// @Tags 角色信息模块
// @Summary 分页查询角色信息列表-Summary
// @Description 分页查询角色信息列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysRoleDto body model.SysRoleDto true "角色信息实体Dto"
// @Success 200 {object} commonModel.JsonResult "分页查询角色信息列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role/pageList [get]
func (this *SysRoleApi) GetSysRolePageList(c *gin.Context) {
    var sysRoleDto dto.SysRoleDto
    c.ShouldBindQuery(&sysRoleDto)

    res, err := this.sysRoleService.GetSysRolePageList(&sysRoleDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询角色信息列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// ExportSysRole 导出角色信息列表
// @Tags 角色信息模块
// @Summary 导出角色信息列表-Summary
// @Description 导出角色信息列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysRoleDto body model.SysRoleDto true "角色信息实体Dto"
// @Success 200 {object} commonModel.JsonResult "导出角色信息列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/role/export [get]
func (this *SysRoleApi) ExportSysRole(c *gin.Context) {
    var sysRoleDto dto.SysRoleDto
    c.ShouldBindQuery(&sysRoleDto)

    res, err := this.sysRoleService.ExportSysRole(&sysRoleDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("导出角色信息列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

