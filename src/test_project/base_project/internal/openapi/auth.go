package openapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const secretPrefix = "base_project"

// GenerateSecret 根据时间戳生成 openapi secret。
// 加密方式：对字符串 "go_base_project-{timestamp}" 直接计算 MD5 哈希值
func GenerateSecret(timestamp int64) (string, error) {
	// 构造原始数据：go_base_project-时间戳
	payload := fmt.Sprintf("%s-%d", secretPrefix, timestamp)

	// 直接对字符串计算 MD5（不进行 JSON 序列化）
	sum := md5.Sum([]byte(payload))
	return hex.EncodeToString(sum[:]), nil
}

// NewAuthHeaders 生成 openapi 请求所需的 secret 和 timestamp 头部值。
func NewAuthHeaders() (secret string, timestamp string, err error) {
	ts := time.Now().Unix()
	secret, err = GenerateSecret(ts)
	if err != nil {
		return "", "", err
	}
	return secret, strconv.FormatInt(ts, 10), nil
}
