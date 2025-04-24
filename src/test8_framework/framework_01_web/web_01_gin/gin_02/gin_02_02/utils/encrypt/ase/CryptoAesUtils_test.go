package ase

import (
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {

	// AES
	key := []byte("1234567890123456") // 16字节
	data := "hello aes"
	enc, _ := AesEncryptBase64([]byte(data), key)
	fmt.Println("AES:", enc)
	dec, _ := AesDecryptBase64(enc, key)
	dstr := string(dec)
	fmt.Println("Decrypted:", dstr)
	fmt.Println("Decrypted:", dstr == data)
}
