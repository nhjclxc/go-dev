package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 中间件二、路由分组中间件

// 记录用户接口的所有接口日志
func UserMiddleware(context *gin.Context) {

	log.Println("UserMiddleware: ", context.FullPath())

}
