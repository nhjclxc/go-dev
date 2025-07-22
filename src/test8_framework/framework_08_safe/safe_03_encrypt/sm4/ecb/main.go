package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

func main() {
	// 明文
	originStr := "helloSM4world123"
	fmt.Println("明文:", originStr)
	plaintext := []byte(originStr) // 必须是 16 字节的倍数
	// 密钥（16字节 = 128位）
	key := []byte("1234567890abcdef")

	// 加密
	ciphertext, err := sm4.Sm4Ecb(key, plaintext, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文（hex）:", hex.EncodeToString(ciphertext))

	// 解密
	decrypted, err := sm4.Sm4Ecb(key, ciphertext, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后:", string(decrypted))
}
