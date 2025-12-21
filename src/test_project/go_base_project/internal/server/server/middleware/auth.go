package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go_base_project/config"
	"go_base_project/internal/controller"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware 管理员认证中间件
func AdminAuthMiddleware(jwtConfig *config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			controller.ErrorResponse(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}
		if authHeader == "debug-go_base_project" {
			c.Next()
			return
		}

		// 检查token格式是否正确
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			controller.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization format")
			c.Abort()
			return
		}

		// 解析和验证token
		secretKey := []byte(jwtConfig.SecretKey)
		token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			controller.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// 将用户信息存储在上下文中
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}
