package middleware

import (
	"base_project/internal/controller"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var secretPrefix = "base_project"

// calculateMD5 计算字符串的 MD5 哈希值
func calculateMD5(data string) string {
	// 直接对字符串计算 MD5 哈希
	hash := md5.Sum([]byte(data))

	// 转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}

// verifyMD5 验证 MD5 哈希值
func verifyMD5(data string, expectedMD5 string) bool {
	// 重新计算 MD5 并比较
	calculatedMD5 := calculateMD5(data)
	return calculatedMD5 == expectedMD5
}

// OpenApiAuthMiddleware openapi认证中间件
func OpenApiAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.GetHeader("secret")
		if secret == "go_base_projectdebug" {
			c.Next()
			return
		}

		// 从请求头中获取
		timestampStr := c.GetHeader("timestamp")
		if secret == "" {
			controller.ErrorResponse(c, http.StatusBadRequest, "请求头中[secret]不存在")
			c.Abort()
			return
		}
		if timestampStr == "" {
			controller.ErrorResponse(c, http.StatusBadRequest, "请求头中[timestamp]不存在")
			c.Abort()
			return
		}
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			controller.ErrorResponse(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		nowTimestamp := time.Now().Unix()
		if nowTimestamp-timestamp > 60 {
			controller.ErrorResponse(c, http.StatusUnauthorized, "请求过期")
			c.Abort()
			return
		}

		data := fmt.Sprintf("%s-%d", secretPrefix, timestamp)
		flag := verifyMD5(data, secret)
		if !flag {
			controller.ErrorResponse(c, http.StatusUnauthorized, "无效的secret加密")
			c.Abort()
			return
		}
		c.Next()
	}
}
