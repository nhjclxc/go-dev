package main

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

// 补码（PKCS#7）
func pkcs7Padding(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// 去补码
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("数据为空，无法去补码")
	}
	unpadLen := int(data[length-1])
	if unpadLen > length {
		return nil, fmt.Errorf("补码长度无效")
	}
	return data[:(length - unpadLen)], nil
}

func sm4EncryptCBC(plaintext, key, iv []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcs7Padding(plaintext, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func sm4DecryptCBC(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("密文长度不是块大小的倍数")
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(plaintext, ciphertext)
	return pkcs7UnPadding(plaintext)
}

func main() {
	key := []byte("1234567890abcdef")       // 16字节密钥
	iv := []byte("abcdef1234567890")        // 16字节IV
	plaintext := []byte("golang-sm4测试数据")

	fmt.Println("原文:", string(plaintext))

	// 加密
	ciphertext, err := sm4EncryptCBC(plaintext, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文（hex）:", hex.EncodeToString(ciphertext))

	// 解密
	decrypted, err := sm4DecryptCBC(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后:", string(decrypted))
}
