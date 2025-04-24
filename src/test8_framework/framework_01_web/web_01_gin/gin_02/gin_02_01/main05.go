package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 框架的各种请求传值
func main05() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由绑定
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "首页")
	})

	// 获取 query 参数
	router.GET("/getQuery", func(context *gin.Context) {

		fmt.Println("getQuery.id", context.Query("id"))
		fmt.Println("getQuery.name", context.Query("name"))
		fmt.Println("getQuery.age", context.Query("age"))

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// http://localhost:8090/getPathVariable/666/zhangsan?age=18
	// 获取 动态路径参数
	// 在路径上要声明不同的参数
	router.GET("/getPathVariable/:id/:name", func(context *gin.Context) {

		fmt.Println("getPathVariable.id", context.Param("id"))
		fmt.Println("getPathVariable.name", context.Param("name"))
		fmt.Println("getQuery.age", context.Query("age"))

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// post form 传参
	router.POST("/postForm", func(context *gin.Context) {

		fmt.Println("postForm.id", context.PostForm("id"))
		fmt.Println("postForm.name", context.PostForm("name"))
		fmt.Println("postForm.age", context.PostForm("age"))

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// post json 传参
	router.POST("/postJson", func(context *gin.Context) {

		// 声明一个变量 jsonData 来接收 json 数据
		var jsonData map[string]any = make(map[string]any)

		// 读取 json 数据
		err := context.ShouldBindJSON(&jsonData)
		if err != nil {
			return
		}

		fmt.Println("postJson.id", jsonData["id"])
		fmt.Println("postJson.name", jsonData["name"])
		fmt.Println("postJson.age", jsonData["age"])

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// put 接收 json 数据
	router.PUT("/putJson", func(context *gin.Context) {

		// 声明一个变量 jsonData 来接收 json 数据
		var jsonData map[string]any = make(map[string]any)

		// 读取 json 数据
		err := context.ShouldBindJSON(&jsonData)
		if err != nil {
			return
		}

		fmt.Println("putJson.id", jsonData["id"])
		fmt.Println("putJson.name", jsonData["name"])
		fmt.Println("putJson.age", jsonData["age"])

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// delete 接收 数据
	router.DELETE("/delete/:id", func(context *gin.Context) {

		fmt.Println("delete.id", context.Param("id"))
		fmt.Println("delete.status", context.Query("status"))

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// post 接收文件
	router.POST("saveFile", func(context *gin.Context) {

		file, err := context.FormFile("myFile")
		if err != nil {
			return
		}
		fmt.Println("file.Size ", file.Size)
		fmt.Println("file.Filename ", file.Filename)
		fmt.Println("file.Header ", file.Header)
		open, err := file.Open()
		if err != nil {
			return
		}
		var buf = make([]byte, 512)
		read, err := open.Read(buf)
		if err != nil {
			return
		}
		fmt.Println("read ", read, "data: ", string(buf))

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// gin 框架的 get post 请求参数与结构体绑定
	/*
		为了能够更方便的获取请求相关参数，提高开发效率，
		我们可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体中。
		下面的示例代码演示了.ShouldBind()强大的功能，它能够基于请求自动提取JSON、form表单和QueryString 类型的数据，并把值绑定到指定的结构体对象。
	*/

	// 获取 query 参数
	router.GET("/getQueryStruct", func(context *gin.Context) {

		var user = User{}

		// 使用 ShouldBind 将数据绑定上去
		err := context.ShouldBind(&user) // 注意：要传地址进去
		if err != nil {
			return
		}

		fmt.Printf("getQuery.user = %v \n", user)

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	// post form 传参
	router.POST("/postJosnStruct", func(context *gin.Context) {

		var user = User{}
		err := context.ShouldBindJSON(&user)
		if err != nil {
			return
		}
		fmt.Printf("postJosn.user = %v \n", user)

		context.JSON(http.StatusOK, map[string]any{"success": true})
	})

	//启动端口监听
	// 默认是：0.0.0.0:8080
	//router.Run(":8090")
	router.Run("localhost:8090")
}

type User struct {

	// json:... 是为了接收json数据的映射
	// form:... 是为了接收get请求的请求参数的数据的映射， Gin 中 GET 请求读取 URL 查询参数，需要用 form 或 query tag
	// Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`

	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}
