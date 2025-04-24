package md5

import (
	"crypto/md5"
	"fmt"
)

// 编码
func Encrypt(str string) string {
	if str == "" {
		return ""
	}
	has := md5.Sum([]byte(str))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 匹配
func Match(str string, encryptedStr string) bool {
	// 注意：MD5 是不可逆的，无法从哈希值反推出原始内容。

	if str == "" || encryptedStr == "" {
		return false
	}
	has := md5.Sum([]byte(str))
	md5str := fmt.Sprintf("%x", has)
	return md5str == encryptedStr
}
