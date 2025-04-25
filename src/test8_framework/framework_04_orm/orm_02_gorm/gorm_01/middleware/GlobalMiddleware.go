package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// GlobalPainc 全局异常处理
func GlobalPainc(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("全局异常处理 捕获到了异常：%v \n", err)
			context.JSON(200, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("操作失败：%v", err),
			})
		}
	}()

	log.Println("进入全局异常处理")

	// 执行剩下的 句柄
	context.Next()

}
