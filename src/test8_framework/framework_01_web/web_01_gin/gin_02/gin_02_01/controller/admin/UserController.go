package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义 UserController 结构体来保存这个实例的相关数据
type UserController struct {
	// userService UsersService

	// UserController 继承 BaseController，使用匿名字段的方式，那么UserController的实例就有BaseController对应的方法了
	BaseController
}

// 定义 UserController 对应的接口方法
func (this *UserController) insertUser(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("insertUser。"),
	})
}
func (this *UserController) deleteUser(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("deleteUser。"),
	})
}
func (this *UserController) updateUser(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("updateUser。"),
	})
}
func (this *UserController) GetUserById(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("GetUserById。"),
	})
}
func (this *UserController) GetUserList(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("GetUserList。"),
	})
}
func (this *UserController) GetUserPageList(context *gin.Context) {
	this.Success(context, fmt.Sprintf("恭喜你，访问成功了。 %s", context.FullPath()))
}
