package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 使用gin实现接口响应数据的包装

/*
ChatGPT 说：
你说的场景是：
🧨 在 API 层已经调用了 c.JSON() 返回数据，但还想在中间件里 对响应内容（比如加密、重新包装等）做“修改”后再输出。

❗问题根源
一旦调用 c.JSON()，Gin 就会通过底层 http.ResponseWriter.Write() 直接把数据写入响应体，无法再修改，除非我们：

✅ 拦截并缓存写入的数据（即替换 ResponseWriter），让 handler 的 c.JSON() 写入内存而非直接返回。


✅ 正确做法（思路是：先“截流”，再“包装”）
你需要 用一个自定义 ResponseWriter 替代 Gin 默认的 writer，将 c.JSON() 输出的数据拦截下来，缓存到内存中，然后在中间件中统一处理并重新输出：

 */

type wrapperResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w wrapperResponseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b) // 不写出，先缓存
}

func wrapperResponseWriterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 替换 ResponseWriter
		writer := &wrapperResponseWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
		c.Writer = writer


		// 执行业务逻辑
		c.Next()


		// 获取原始响应体
		originBody := writer.body.Bytes()

		// 防止写两次
		c.Writer = writer.ResponseWriter

		// 解析原始 JSON
		var originalMap map[string]interface{}
		if err := json.Unmarshal(originBody, &originalMap); err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.Writer.Write([]byte(`{"code":500,"msg":"响应解析失败"}`))
			return
		}

		// 重新包装
		wrapped := gin.H{
			"code": originalMap["code"],
			"msg":  "统一包装成功",
			"data": originalMap["data"],
		}

		// 返回新的响应
		c.JSON(http.StatusOK, wrapped)

	}
}

func main() {

	r := gin.Default()

	// 启用跨域支持
	r.Use(cors.Default())

	wrapperResponseWriterGroup := r.Group("/api", wrapperResponseWriterMiddleware())
	wrapperResponseWriterGroup.GET("/getUser", getUser)

	r.Run(":8080")
}

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
			"data":  "Foo getUser " + time.Now().String(),
		},
	)
}
