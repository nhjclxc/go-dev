package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// OpenApiAuthMiddleware openapi认证中间件
func OpenApiAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Content-Type 校验
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodGet {
			if !strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
				Error(c, fmt.Errorf("invalid content-type"))
				c.Abort()
				return
			}
		}

		// 身份认证头部字段，
		//例如： HMAC-SHA256 Certificate=AccessKey/Timestamp/Service, Sign=c84d4c6e8854931d5609e58c666352062a70f3a5e591b84d3941374ed5e198ea
		//其中，
		//	- HMAC-SHA256：签名方法，目前固定取该值；
		//	- Certificate：签名凭证，AccessKey；Timestamp为当前时间戳；Server为请求服务名称；
		//	- Signature：签名摘要，计算过程详见 4
		authorization := c.GetHeader("X-WY-Authorization") // 身份认证头部字段，
		accessKey := c.GetHeader("X-WY-AccessKey")         // 用户的AccessKey
		timestampStr := c.GetHeader("X-WY-Timestamp")      // 当前时间戳
		_ = c.GetHeader("X-WY-Version")                    // 目前为1.0
		service := c.GetHeader("X-WY-Service")             // 请求服务名称，当前默认传CDN

		if authorization == "" || accessKey == "" || timestampStr == "" || service == "" {
			Error(c, fmt.Errorf("missing auth headers"))
			c.Abort()
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			Error(c, err)
			c.Abort()
			return
		}

		// 超过1分钟认为请求过期
		nowTimestamp := time.Now().Unix()
		if nowTimestamp-timestamp > 60 {
			Error(c, fmt.Errorf("request expired"))
			c.Abort()
			return
		}

		// 解析签名
		clientSign, err := parseAuthorization(authorization)
		if err != nil {
			Error(c, fmt.Errorf("invalid authorization format"))
			c.Abort()
			return
		}

		// todo 根据 ak 获取对应的 sk
		accessSecret := "sk"
		// 本地计算签名
		stringToSign := service + timestampStr + accessKey + accessSecret
		serverSign := hmacSHA256(stringToSign, accessSecret)

		// 对比签名
		if !hmac.Equal([]byte(clientSign), []byte(serverSign)) {
			Error(c, fmt.Errorf("signature mismatch"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// Error 错误响应
func Error(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  err.Error(),
	})
}

// hmacSHA256 计算签名
func hmacSHA256(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

// parseAuthorization 解析客户端签名数据
func parseAuthorization(auth string) (string, error) {
	// 提取Sign部分，前面的HMAC-SHA256 Certificate=xxx目前没有用
	parts := strings.Split(auth, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if strings.HasPrefix(p, "Sign=") {
			return strings.TrimPrefix(p, "Sign="), nil
		}
	}
	return "", fmt.Errorf("signature not found")
}

// ============= 以下可以用于生成用户的ak、sk ========================
func GenerateAccessKey() string {
	u := uuid.New().String() // 36 位
	u = strings.ReplaceAll(u, "-", "")
	return "GKCDNAK_" + u[:16]
}

func GenerateAccessSecret() (string, error) {
	const length = 32 // 32 字节 = 256 bit
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GenerateCredential() (string, string, error) {
	sk, err := GenerateAccessSecret()
	if err != nil {
		return "", "", err
	}

	return GenerateAccessKey(), sk, nil
}
