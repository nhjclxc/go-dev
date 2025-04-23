package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义 NewsController 结构体来保存这个实例的相关数据
type NewsController struct {
	// newsService NewsService
}

// 定义 NewsController 对应的接口方法
func (this *NewsController) Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("我是 新闻页面首页。"),
	})
}
func (this *NewsController) GetNewsById(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": fmt.Sprintf("我是 id = %v 的文章内容。", context.Query("id")),
	})
}
