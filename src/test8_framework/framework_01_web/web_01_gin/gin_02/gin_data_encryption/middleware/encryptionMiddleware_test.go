package middleware

import (
	"fmt"
	"testing"
)

func TestEnDe(t *testing.T) {

	// 密钥（16字节 = 128位）
	key := []byte("ZAQ12WSXCDE34RFV")


	encStr := "username123"

	decrypted, _ := encryptSM4Hex(encStr, key)

	fmt.Printf("encrypted = %v \n", decrypted)


	decrypted111, _ := decryptSM4Hex(decrypted, key)

	fmt.Printf("decrypted111 = %v \n", decrypted111)


}
func TestEncrypt(t *testing.T) {

	// 密钥（16字节 = 128位）
	key := []byte("ZAQ12WSXCDE34RFV")


	decrypted, _ := encryptSM4Hex("username12345", key)
	fmt.Printf("encrypted = %v \n", decrypted)
	decrypted2, _ := encryptSM4Hex("password1234567876", key)
	fmt.Printf("decrypted2 = %v \n", decrypted2)
	decrypted222, _ := encryptSM4Hex("Proxy-Authorization", key)
	fmt.Printf("decrypted222 = %v \n", decrypted222)

	decrypted3, _ := encryptSM4Hex("{\"username\":\"username12345\",\"password\":\"password1234567876\"}", key)
	fmt.Printf("decrypted2 = %v \n", decrypted3)

}


func TestDecrypt(t *testing.T) {

	// 密钥（16字节 = 128位）
	key := []byte("ZAQ12WSXCDE34RFV")

	encStr := "22fd4876c6d60588d9af9e1511b07ab75467ff60e3a0dbb6f1902c3d0e5af0ff63339c3245900d8d17f63b0774e59151ce7d54107ddbee08ab298e0821db09c6651052c7b8adebf3ef2e045e1a21bc4d"

	decrypted, _ := decryptSM4Hex(encStr, key)

	fmt.Printf("decrypted = %v \n", decrypted)


}