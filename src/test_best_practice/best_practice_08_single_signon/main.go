package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type LoginInfo struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LoginTime time.Time `json:"loginTime"`
	Uuid      string    `json:"uuid"`
}
type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

var signinUserMap = make(map[string]LoginInfo)

// 使用内存存储登陆信息，实现一个账号同时只能登录一次，登陆不过期。
// 仅仅模拟实现一个账号只能由一个客户端登陆的功能
func main() {

	r := gin.Default()

	// 登陆接口
	// http://127.0.0.1:8090/api/login
	r.POST("/api/login", func(ctx *gin.Context) {
		var req LoginReq
		ctx.ShouldBindJSON(&req)

		if req.Username == "" || req.Password == "" {
			ctx.JSON(200, gin.H{
				"msg":           "用户名或密码为空！！！",
				"signinUserMap": signinUserMap,
			})
			return
		}

		// 模拟账号密码验证
		if (req.Username + "pwd") != req.Password {
			ctx.JSON(200, gin.H{
				"msg":           "密码错误！！！",
				"signinUserMap": signinUserMap,
			})
			return
		}

		msg := "登陆"
		// 检查当前用户是否已经登陆了
		_, ok := signinUserMap[req.Username]
		if ok {
			// 已经登陆了，踢掉旧会话，允许新会话生效
			delete(signinUserMap, req.Username)
			msg = "已经登陆了，踢掉旧会话，允许新会话生效"
		}

		uuidToken := uuid.New().String()

		signinUserMap[req.Username] = LoginInfo{
			Username:  req.Username,
			Password:  req.Password,
			LoginTime: time.Now(),
			Uuid:      uuidToken,
		}

		ctx.JSON(200, gin.H{
			"data":          uuidToken,
			"msg":           msg,
			"signinUserMap": signinUserMap,
		})
	})

	// 检查登陆会话是否还有效
	// http://127.0.0.1:8090/api/ping?username=admin
	r.GET("/api/ping", func(ctx *gin.Context) {
		var req LoginReq
		ctx.ShouldBindQuery(&req)
		if req.Username == "" {
			ctx.JSON(200, gin.H{
				"msg":           "请输入用户名！！！",
				"signinUserMap": signinUserMap,
			})
			return
		}
		loginInfo, ok := signinUserMap[req.Username]
		if !ok {
			ctx.JSON(200, gin.H{
				"msg":           "当前用户未登陆！！！",
				"signinUserMap": signinUserMap,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"data":          loginInfo,
			"signinUserMap": signinUserMap,
		})
	})

	// 踢出某个用户
	// http://127.0.0.1:8090/api/logout?username=admin
	r.GET("/api/logout", func(ctx *gin.Context) {
		var req LoginReq
		ctx.ShouldBindQuery(&req)
		if req.Username == "" {
			ctx.JSON(200, gin.H{
				"msg":           "请输入用户名！！！",
				"signinUserMap": signinUserMap,
			})
			return
		}
		_, ok := signinUserMap[req.Username]
		if !ok {
			ctx.JSON(200, gin.H{
				"msg":           "当前用户未登陆！！！",
				"signinUserMap": signinUserMap,
			})
			return
		}

		delete(signinUserMap, req.Username)

		ctx.JSON(200, gin.H{
			"data":          req.Username + " 已被踢出！！！",
			"signinUserMap": signinUserMap,
		})

	})

	r.Run(":8090")
}
