package middleware

import (
	"gin_02_02/model/user"
	"github.com/gin-gonic/gin"
	"log"
)

// 中间件一、路由中间件

// 单独为用户登录开的一个中间件处理方法
func UserLoginMiddleware(context *gin.Context) {

	// 记录有哪些用户登录过

	// 获取参数
	userLogin := user.UserLogin{}
	context.ShouldBind(&userLogin)

	log.Println("UserLoginMiddleware: ", userLogin.Username)


}
