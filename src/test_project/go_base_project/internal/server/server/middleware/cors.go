package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS 处理跨域请求的中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")

		// 允许的请求头（包括自定义的 secret 和 timestamp）
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, secret, timestamp")

		// 允许携带认证信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 预检请求的有效期（单位：秒）
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// 允许暴露的响应头
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

		// 处理预检请求（OPTIONS）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
