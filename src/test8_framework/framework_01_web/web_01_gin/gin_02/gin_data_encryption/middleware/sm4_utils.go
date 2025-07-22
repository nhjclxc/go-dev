// middleware/sm4_utils.go
package middleware

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"io"

	"github.com/tjfoc/gmsm/sm4"
)

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据为空")
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("补码无效")
	}
	return data[:(length - unpadding)], nil
}

func encryptSM4Hex(plainText string, key []byte) (string, error) {
	data := pkcs7Padding([]byte(plainText), 16)
	cipherText, err := sm4.Sm4Ecb(key, data, true)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func decryptSM4Hex(hexStr string, key []byte) (string, error) {
	cipherData, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	plainData, err := sm4.Sm4Ecb(key, cipherData, false)
	if err != nil {
		return "", err
	}
	unpad, err := pkcs7UnPadding(plainData)
	if err != nil {
		return "", err
	}
	return string(unpad), nil
}

func decryptReader(r io.Reader, key []byte) (io.Reader, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(r, iv); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	return &cipher.StreamReader{S: stream, R: r}, nil
}
