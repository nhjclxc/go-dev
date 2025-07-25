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

// SysJobApi 定时任务调度 api 层
type SysJobApi struct{
    sysJobService service.SysJobService
}

// InsertSysJob 新增定时任务调度
// @Tags 定时任务调度模块
// @Summary 新增定时任务调度-Summary
// @Description 新增定时任务调度-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysJob body model.SysJob true "修改定时任务调度实体类"
// @Success 200 {object} commonModel.JsonResult "新增定时任务调度响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job [post]
func (this *SysJobApi) InsertSysJob(c *gin.Context) {
    var sysJob model.SysJob
    c.ShouldBindJSON(&sysJob)

    res, err := this.sysJobService.InsertSysJob(&sysJob)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("新增定时任务调度失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// UpdateSysJob 修改定时任务调度
// @Tags 定时任务调度模块
// @Summary 修改定时任务调度-Summary
// @Description 修改定时任务调度-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sysJob body model.SysJob true "修改定时任务调度实体类"
// @Success 200 {object} commonModel.JsonResult "修改定时任务调度响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job [put]
func (this *SysJobApi) UpdateSysJob(c *gin.Context) {
    var sysJob model.SysJob
    c.ShouldBindJSON(&sysJob)

    res, err := this.sysJobService.UpdateSysJob(&sysJob)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("修改定时任务调度失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// DeleteSysJob 删除定时任务调度
// @Tags 定时任务调度模块
// @Summary 删除定时任务调度-Summary
// @Description 删除定时任务调度-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param jobIdList path string true "定时任务调度主键List"
// @Success 200 {object} commonModel.JsonResult "删除定时任务调度响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job/:jobIdList [delete]
func (this *SysJobApi) DeleteSysJob(c *gin.Context) {
    jobIdListStr := c.Param("jobIdList") // 例如: "1,2,3"
    jobIdList, err := commonUtils.ParseIds(jobIdListStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.sysJobService.DeleteSysJob(jobIdList)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("删除定时任务调度失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysJobByJobId 获取定时任务调度详细信息
// @Tags 定时任务调度模块
// @Summary 获取定时任务调度详细信息-Summary
// @Description 获取定时任务调度详细信息-Description
// @Security BearerAuth
// @Accept path
// @Produce path
// @Param jobId path int64 true "定时任务调度主键List"
// @Success 200 {object} commonModel.JsonResult "获取定时任务调度详细信息"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job/:jobId [get]
func (this *SysJobApi) GetSysJobByJobId(c *gin.Context) {
    jobIdStr := c.Param("jobId") // 例如: "1"
    jobId, err := commonUtils.ParseId(jobIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, commonModel.JsonResultError("参数错误：" + err.Error()))
        return
    }

    res, err := this.sysJobService.GetSysJobByJobId(jobId)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询定时任务调度失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysJobList 查询定时任务调度列表
// @Tags 定时任务调度模块
// @Summary 查询定时任务调度列表-Summary
// @Description 查询定时任务调度列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysJobDto body model.SysJobDto true "定时任务调度实体Dto"
// @Success 200 {object} commonModel.JsonResult "查询定时任务调度列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job/list [get]
func (this *SysJobApi) GetSysJobList(c *gin.Context) {
    var sysJobDto dto.SysJobDto
    c.ShouldBindQuery(&sysJobDto)

    res, err := this.sysJobService.GetSysJobList(&sysJobDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询定时任务调度列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// GetSysJobPageList 分页查询定时任务调度列表
// @Tags 定时任务调度模块
// @Summary 分页查询定时任务调度列表-Summary
// @Description 分页查询定时任务调度列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysJobDto body model.SysJobDto true "定时任务调度实体Dto"
// @Success 200 {object} commonModel.JsonResult "分页查询定时任务调度列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job/pageList [get]
func (this *SysJobApi) GetSysJobPageList(c *gin.Context) {
    var sysJobDto dto.SysJobDto
    c.ShouldBindQuery(&sysJobDto)

    res, err := this.sysJobService.GetSysJobPageList(&sysJobDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("查询定时任务调度列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}

// ExportSysJob 导出定时任务调度列表
// @Tags 定时任务调度模块
// @Summary 导出定时任务调度列表-Summary
// @Description 导出定时任务调度列表-Description
// @Security BearerAuth
// @Accept param
// @Produce param
// @Param sysJobDto body model.SysJobDto true "定时任务调度实体Dto"
// @Success 200 {object} commonModel.JsonResult "导出定时任务调度列表响应数据"
// @Failure 401 {object} commonModel.JsonResult "未授权"
// @Failure 500 {object} commonModel.JsonResult "服务器异常"
// @Router /sys/job/export [get]
func (this *SysJobApi) ExportSysJob(c *gin.Context) {
    var sysJobDto dto.SysJobDto
    c.ShouldBindQuery(&sysJobDto)

    res, err := this.sysJobService.ExportSysJob(&sysJobDto)

    if err != nil {
        c.JSON(http.StatusInternalServerError, commonModel.JsonResultError("导出定时任务调度列表失败：" + err.Error()))
        return
    }
    c.JSON(http.StatusOK, commonModel.JsonResultSuccess[any](res))
}




