package openapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
	"time"
)

// calculateMD5 模拟中间件中的 MD5 计算逻辑
func calculateMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// verifyMD5 模拟中间件中的验证逻辑
func verifyMD5(data string, expectedMD5 string) bool {
	calculatedMD5 := calculateMD5(data)
	return calculatedMD5 == expectedMD5
}

// TestGenerateSecret 测试 secret 生成逻辑
func TestGenerateSecret(t *testing.T) {
	timestamp := time.Now().Unix()

	// 1. 使用客户端方法生成 secret
	secret, err := GenerateSecret(timestamp)
	if err != nil {
		t.Fatalf("生成 secret 失败: %v", err)
	}

	// 2. 使用服务端逻辑验证（模拟中间件验证过程）
	data := fmt.Sprintf("go_base_project-%d", timestamp)
	if !verifyMD5(data, secret) {
		t.Errorf("验证失败！客户端生成的 secret 与服务端验证不匹配")
		t.Logf("时间戳: %d", timestamp)
		t.Logf("原始数据: %s", data)
		t.Logf("客户端生成的 secret: %s", secret)
		t.Logf("服务端计算的 MD5: %s", calculateMD5(data))
	} else {
		t.Logf("✅ 验证成功！")
		t.Logf("时间戳: %d", timestamp)
		t.Logf("原始数据: %s", data)
		t.Logf("secret: %s", secret)
	}
}

// TestNewAuthHeaders 测试生成认证头部
func TestNewAuthHeaders(t *testing.T) {
	secret, timestamp, err := NewAuthHeaders()
	if err != nil {
		t.Fatalf("生成认证头部失败: %v", err)
	}

	t.Logf("生成的认证头部:")
	t.Logf("  secret: %s", secret)
	t.Logf("  timestamp: %s", timestamp)

	// 验证格式
	if secret == "" {
		t.Error("secret 不能为空")
	}
	if timestamp == "" {
		t.Error("timestamp 不能为空")
	}
	if len(secret) != 32 {
		t.Errorf("MD5 哈希应该是 32 个字符，实际: %d", len(secret))
	}
}

// TestDebugSecret 测试调试密钥
func TestDebugSecret(t *testing.T) {
	debugSecret := "go_base_projectdebug"

	t.Logf("调试模式密钥: %s", debugSecret)
	t.Logf("提示: 在开发环境可以使用此密钥绕过 MD5 验证")
}
