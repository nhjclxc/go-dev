package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义 BaseController 结构体来保存这个实例的相关数据
type BaseController struct {
	// baseService BaseService
}

// 定义 BaseController 对应的 一些通用方法
func (this *BaseController) Success(context *gin.Context, data any) {
	context.JSON(200, gin.H{
		"code": 200,
		//"data": fmt.Sprintf("恭喜你，访问成功了。"),
		"data": data,
	})
}
func (this *BaseController) Error(context *gin.Context, msg string) {
	context.JSON(200, gin.H{
		"code": 500,
		"msg":  fmt.Sprintf("很遗憾你这次访问失败了！！！" + msg),
	})
}
