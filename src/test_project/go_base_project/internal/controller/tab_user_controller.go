package controller

import (
	"github.com/gin-gonic/gin"
	"go_base_project/internal/model"
	"go_base_project/internal/request"
	"go_base_project/internal/service"
	"net/http"
)

// TabUserController 用户  控制器层
type TabUserController struct {
	tabUserService *service.TabUserService
}

// NewTabUserController 创建 TabUser 用户  控制器层对象
func NewTabUserController(tabUserService *service.TabUserService) *TabUserController {
	return &TabUserController{
		tabUserService: tabUserService,
	}
}

// InsertTabUser 新增用户
// @Router /tab/user [post]
func (tuc *TabUserController) InsertTabUser(c *gin.Context) {
	var tabUser model.TabUser
	err := c.ShouldBindJSON(&tabUser)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := tuc.tabUserService.InsertTabUser(c.Request.Context(), tabUser)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "新增用户 失败："+err.Error())
		return
	}
	SuccessResponse(c, res)
}

// UpdateTabUser 修改用户
// @Router /tab/user [put]
func (tuc *TabUserController) UpdateTabUser(c *gin.Context) {
	var tabUser model.TabUser
	err := c.ShouldBindJSON(&tabUser)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := tuc.tabUserService.UpdateTabUser(c.Request.Context(), tabUser)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "修改用户 失败："+err.Error())
		return
	}
	SuccessResponse(c, res)
}

// DeleteTabUser 删除用户
// @Router /tab/user/:idList [delete]
func (tuc *TabUserController) DeleteTabUser(c *gin.Context) {
	idListStr := c.Param("idList") // 例如: "1,2,3"
	idList, err := ParseIds(idListStr)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "参数错误："+err.Error())
		return
	}

	res, err := tuc.tabUserService.DeleteTabUser(c.Request.Context(), idList)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "删除用户 失败："+err.Error())
		return
	}
	SuccessResponse(c, res)
}

// GetTabUserById 获取用户 详细信息
// @Router /tab/user/:id [get]
func (tuc *TabUserController) GetTabUserById(c *gin.Context) {
	idStr := c.Param("id") // 例如: "1"
	id, err := ParseId(idStr)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "参数错误："+err.Error())
		return
	}

	res, err := tuc.tabUserService.GetTabUserById(c.Request.Context(), id)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "查询用户 失败："+err.Error())
		return
	}
	SuccessResponse(c, res)
}

// GetTabUserList 查询用户 列表
// @Router /tab/user/list [get]
func (tuc *TabUserController) GetTabUserList(c *gin.Context) {
	var tabUserReq request.TabUserReq
	err := c.ShouldBindQuery(&tabUserReq)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := tuc.tabUserService.GetTabUserList(c.Request.Context(), tabUserReq)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "查询用户 列表失败："+err.Error())
		return
	}
	SuccessResponse(c, res)
}

// GetTabUserPageList 分页查询用户 列表
// @Router /tab/user/pageList [get]
func (tuc *TabUserController) GetTabUserPageList(c *gin.Context) {
	var tabUserReq request.TabUserReq
	err := c.ShouldBindQuery(&tabUserReq)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := tuc.tabUserService.GetTabUserPageList(c.Request.Context(), tabUserReq)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "查询用户 分页列表失败："+err.Error())
		return
	}
	SuccessResponse(c, res)
}

// ExportTabUser 导出用户 列表
// @Router /tab/user/export [get]
func (tuc *TabUserController) ExportTabUser(c *gin.Context) {
	var tabUserReq request.TabUserReq
	err := c.ShouldBindQuery(&tabUserReq)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	excelfile, err := tuc.tabUserService.ExportTabUser(c.Request.Context(), tabUserReq)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "导出用户 列表失败："+err.Error())
		return
	}

	// 将文件写入响应
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", `attachment; filename="report.xlsx"`)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	// 将内容直接写到响应体
	if err := excelfile.Write(c.Writer); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
}
