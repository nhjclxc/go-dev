package ase

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// 使用 crypto/aes 实现对称加密
// AES 是对称加密算法，加密和解密使用同一个密钥。

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}

// AesEncryptBase64 加密并返回 Base64 编码字符串（带随机 IV）
func AesEncryptBase64(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = pkcs7Pad(plaintext, blockSize)

	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plaintext))
	mode.CryptBlocks(cipherText, plaintext)

	final := append(iv, cipherText...)
	return base64.StdEncoding.EncodeToString(final), nil
}

// AesDecryptBase64 解密 Base64 编码字符串
func AesDecryptBase64(cipherBase64 string, key []byte) ([]byte, error) {
	cipherData, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(cipherData) < blockSize {
		return nil, errors.New("cipher text too short")
	}

	iv := cipherData[:blockSize]
	cipherText := cipherData[blockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	return pkcs7Unpad(plainText)
}
