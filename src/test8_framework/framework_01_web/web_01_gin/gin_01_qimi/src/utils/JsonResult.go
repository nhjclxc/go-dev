package main

import "github.com/gin-gonic/gin"

//go gin 如何封装一个统一的响应结构体

//const (
//	Success      = 0
//	ParamError   = 1001
//	AuthError    = 1002
//	ServerError  = 5000
//	// ...更多
//)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 成功响应
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// 失败响应
func Fail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
