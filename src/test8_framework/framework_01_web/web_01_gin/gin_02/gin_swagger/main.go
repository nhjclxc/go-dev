package main

import (
	"fmt"
	_ "gin_swagger/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// gin生成swagger文档
// gin swagger

// go get github.com/swaggo/swag/cmd/swag
// go install github.com/swaggo/swag/cmd/swag
// go get github.com/swaggo/gin-swagger
// go get github.com/swaggo/files

// https://github.com/swaggo/gin-swagger
// https://swaggo.github.io/swaggo.io/declarative_comments_format/

/*
	下面将介绍如何给接口加上请求投认证信息

// @host用于描述接口像哪个baseUrl请求
*/

// @title 示例 API
// @version 1.0
// @description 使用 Swagger 演示 token 请求头设置
// @BasePath /
// @host localhost:8080
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// 启用跨域支持
	r.Use(cors.Default())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//// 以下配置可以开关的swagger，当前环境为dev时允许访问swagger，而当环境为prod时关闭swagger的访问权限
	//if config.Env == "dev" {
	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}

	r.POST("/login", login)

	// 路由分组，所有需要鉴权的接口用 AuthMiddleware 包裹
	authGroup := r.Group("/api", AuthMiddleware())
	{
		userGroup := authGroup.Group("/anonymous_user")
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

/*
	下面将介绍如何给接口加上请求投认证信息
    @Security BearerAuth：表示这个接口要加Bearer类型的认证信息
	@Failure 401：表示接口返回401时，是因为未授权的原因

✅ 小贴士
@securityDefinitions.apikey 必须放在 main.go 里，或者是生成文档的入口文件中；
@Security BearerAuth 每个需要认证的接口都要单独加；
可以封装一个统一的响应结构体返回。

*/

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
// @Param anonymous_user body UserDto true "用户信息"
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

/*
启动操作步骤：
	1、在执行完'go install github.com/swaggo/swag/cmd/swag'命令之后，在bin目录下会生成一个 'swag.exe'
	2、将上述 'swag.exe' 移动到和目前这个main.go的同级目录下
	3、在main.go的cmd里面执行 ‘swag init’ 命令，在这个目录下会生成一个 'doc'文件夹，里面包含docs.go、swagger.json、swagger.yaml
		swag init命令输出如下：[img.png](./img.png)
		```
			cmd输入：swag init
			2025/04/19 21:00:56 Generate swagger docs....
			2025/04/19 21:00:56 Generate general API Info, search dir:./
			2025/04/19 21:01:13 create docs.go at docs/docs.go
			2025/04/19 21:01:13 create swagger.json at docs/swagger.json
			2025/04/19 21:01:13 create swagger.yaml at docs/swagger.yaml
		```
	4、go run main.go 启动，浏览器访问：http://127.0.0.1:8282/swagger/index.html
		发现页面显示：[img_1.png](./img_1.png)
			```
				Failed to load API definition.
				Fetch error
				Internal Server Error doc.json
			```
		这是因为本go文件没有导入第3步生成的docs文件夹里面的内容，将 _ "gin_swagger/docs" 添加到imports下面即可【要将docs文件夹导入到注册"/swagger/*any"接口，这样去注册这个接口的时候才能找到】
	5、重新启动 go run main.go，浏览器访问：http://127.0.0.1:8282/swagger/index.html，效果如[img_2.png](./img_2.png)
	注意：如果swagger注解内容发送了变化，那么必须重新执行第3步，重新生成swagger文档在重启才能生效
	6、新增接口时，如何更新swagger接口文档，新写一个logout接口，执行`swag init`，重新启动项目，访问：http://127.0.0.1:8282/swagger/index.html


*/

/*

Failed to fetch.
Possible Reasons:
	CORS
	Network Failure
	URL scheme must be "http" or "https" for CORS request.

如果出现上述原因是：
*/
