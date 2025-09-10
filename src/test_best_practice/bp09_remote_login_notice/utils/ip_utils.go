package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

// GetUserAgent 获取 UserAgent
func GetUserAgent(c *gin.Context) string {
	return c.GetHeader("User-Agent") // 或 c.Request.UserAgent()
}

// GetClientIP 获取真实客户端 IP（推荐）
func GetClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

// GetProxyChain 获取代理 IP 链（转发链）
func GetProxyChain(c *gin.Context) []string {
	// X-Forwarded-For: client, proxy1, proxy2 ...
	forwardedFor := c.GetHeader("X-Forwarded-For")
	if forwardedFor == "" {
		// 没有转发链，直接用 RemoteAddr
		ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		return []string{ip}
	}

	// 按逗号分隔，并去掉空格
	ips := strings.Split(forwardedFor, ",")
	for i := range ips {
		ips[i] = strings.TrimSpace(ips[i])
	}
	return ips
}

// GetUserFingerprint 基于请求生成一个简单指纹
func GetUserFingerprint(c *gin.Context) string {
	parts := []string{
		c.GetHeader("User-Agent"),
		c.GetHeader("Accept"),
		c.GetHeader("Accept-Language"),
		c.GetHeader("Accept-Encoding"),
		c.GetHeader("Connection"),
		c.GetHeader("Referer"),
		c.ClientIP(), // 真实客户端 IP
	}

	// 合并为一个字符串
	raw := strings.Join(parts, "|")

	// 用 SHA256 生成指纹
	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])
}
