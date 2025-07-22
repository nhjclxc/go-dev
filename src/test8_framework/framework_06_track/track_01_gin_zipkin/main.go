package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
	"net/http"
)

// Gin + OpenTelemetry + Zipkin 实现链路追踪
//go get go.opentelemetry.io/otel
//go get go.opentelemetry.io/otel/sdk
//go get go.opentelemetry.io/otel/exporters/zipkin
//go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin


// 启动一个zipkin实例
// docker pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/openzipkin/zipkin:latest
// docker run -d -p 9411:9411 swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/openzipkin/zipkin:latest
// http://39.106.59.225:9411/zipkin/

// go run main.go

// 发起一个请求，看zipkin UI的请求记录

// 继续增加zap日志配合
// go get go.uber.org/zap




func main() {
	// 初始化zipkin
	shutdown := InitTracer("track_02_gin_zipkin_prometheus", "http://39.106.59.225:9411/api/v2/spans")
	defer shutdown(context.Background())


	r := gin.Default()


	// 添加 OpenTelemetry 中间件（自动注入 traceId/spanId）
	r.Use(otelgin.Middleware("track_02_gin_zipkin_prometheus"))


	// 启用跨域支持
	r.Use(cors.Default())



	// 路由分组，所有需要鉴权的接口用 AuthMiddleware 包裹
	authGroup := r.Group("/api", AuthMiddleware())
	{
		userGroup := authGroup.Group("/user")
		{
			userGroup.POST("/logout", logout)
			userGroup.POST("/postUser", postUser)
			userGroup.PUT("/putUser", putUser)
			userGroup.DELETE("/deleteUser", deleteUser)
			userGroup.GET("/getUser", getUser)
		}
	}


	r.Run(":8080")
}

// @Tags 登录登出模块
// @Summary 登录-Summary
// @Description 登录-Description
// @Accept  json
// @Produce json
// @Param   username   path    string     true        "登录用户名"
// @Param   password   path    string     true        "登录密码"
// @Success 200 {string} string    "ok"
// @Router /login [post]
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, "Hello world "+username+"_"+password)
}

// @Tags 登录登出模块
// @Summary 退出-Summary
// @Description 退出-Description
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Param username query string true "登录用户名"
// @Success 200 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string}  "退出登录响应数据"
// @Failure 401 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string} "未授权"
// @Router /logout [post]
func logout(c *gin.Context) {

	authorization := c.GetHeader("Authorization")
	username := c.PostForm("username")

	fmt.Printf("getUser，authorization = %s, username = %s \n", authorization, username)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo"},
		},
	)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !validateToken(token) {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func validateToken(token string) bool {
	fmt.Printf("AuthMiddleware.validateToken = %s \n", token)
	return true
}


// @Tags 用户模块
// @Summary 获取用户详细-Summary
// @Description 获取用户详细-Description
// @Security BearerAuth
// @Param username query string true "登录用户名"
// @Success 200 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string}  "退出登录响应数据"
// @Failure 401 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string} "未授权"
// @Router /getUser [get]
func getUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	username := c.Query("username")

	fmt.Printf("getUser，authorization = %s, username = %s \n", authorization, username)

	msg := fmt.Sprintf("getUser，authorization = %s, username = %s", authorization, username)
	Logger.Info(msg)

	Logger.Info("getUser，authorization = %s, username = %s \n",
			zap.String("authorization", authorization),
			zap.String("username", username),
		)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo"},
		},
	)
}

// @Tags 用户模块
// @Summary 创建用户-Summary
// @Description 提交用户信息创建用户-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body UserDto true "用户信息"
// @Success 200 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string}  "退出登录响应数据"
// @Failure 401 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string} "未授权"
// @Router /postUser [post]
func postUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	var dto UserDto
	c.ShouldBindJSON(&dto)

	fmt.Printf("postUser，authorization = %s, dto = %v \n", authorization, dto)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo"},
		},
	)
}

// @Tags 用户模块
// @Summary 更新用户信息-Summary
// @Description 更新用户信息-Description
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string}  "更新用户信息-响应数据"
// @Failure 401 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string} "未授权"
// @Router /putUser/{userId} [put]
func putUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	var dto UserDto
	c.ShouldBindJSON(&dto)

	fmt.Printf("putUser，authorization = %s, dto = %v \n", authorization, dto)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo"},
		},
	)
}

// @Tags 用户模块
// @Summary 删除用户-Summary
// @Description 删除用户-Description
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Success 200 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string}  "更新用户信息-响应数据"
// @Failure 401 {object} JsonResponse{data=LogoutVo,msg=string,code=int,error=string} "未授权"
// @Router /deleteUser/{userId} [delete]
func deleteUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	userId := c.Param("userId") // 取出字符串形式

	fmt.Printf("deleteUser，authorization = %s, userId = %s \n", authorization, userId)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo"},
		},
	)
}

// 接口统一响应结构
type JsonResponse struct {
	Code    int    `json:"code"`    // 响应码
	Msg     string `json:"msg"`     // 失败消息
	Success bool   `json:"success"` // 是否操作成功，操作成功返回true，否则返回false
	Data    any    `json:"data"`    // 响应数据
}

type LogoutVo struct {
	Foo string `json:"foo"` // 测试的字段
}

type UserDto struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Foo      string `json:"foo"`      // 测试的字段
}
