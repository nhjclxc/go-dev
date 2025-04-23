package middleware

import (
	"fmt"
	"gin_02_02/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

// 中间件三、全局中间件

// Authentication 接口鉴权
func Authentication(context *gin.Context) {

	log.Println("context.FullPath() = ", context.FullPath())

	// 放掉登录接口
	if "/user/login" == context.FullPath() {
		log.Println("登录接口不鉴权", context.FullPath())
		return
	}

	// 获取请求头
	var token string = context.GetHeader("Authorization") // Token

	log.Println("token: ", token)

	if token == "" {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "请登录",
		})

		// 鉴权不通过，终止该请求
		context.Abort()
	}

	token = strings.ReplaceAll(token, "Bearer ", "")
	uid, err := jwt.ParseToken(token)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 500,
			"msg":  "登录过期，请重新登录",
		})
		// 鉴权不通过，终止该请求
		context.Abort()
	}

	// 鉴权通过了，将该用户的详细放入本次请求的上下文中，类似于Java中的 ThreadLocal 存储一样
	context.Set("uid", uid)

}

// 请求参数日志
func RequestParamLog(context *gin.Context) {
	log.Println("1-记录接口日志，", context.FullPath())

	// 记录程序的执行时间
	start := time.Now().Nanosecond()

	// 调用该请求的剩余处理程序
	context.Next()

	log.Println("2-记录接口日志，", context.FullPath())

	end := time.Now().Nanosecond()

	log.Println("接口执行时间：", (end - start))
}

// 响应参数日志
func ResponseDataLog(context *gin.Context) {
	log.Println("记录接口响应数据，", context.FullPath())
}

// GlobalPainc 全局异常处理
func GlobalPainc(context *gin.Context) {
	defer func() {
		log.Println("全局异常处理 执行到了")
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
