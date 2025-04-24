package rsa

import (
	"fmt"
	"testing"
)

func TestParseRSAPrivateKeyFromPEM(t *testing.T) {

	// RSA
	pubPEM := []byte(`-----BEGIN PUBLIC KEY----------END PUBLIC KEY-----`)
	privPEM := []byte(`-----BEGIN RSA PRIVATE KEY----------END RSA PRIVATE KEY-----`)

	pubKey, _ := ParseRSAPublicKeyFromPEM(pubPEM)
	privKey, _ := ParseRSAPrivateKeyFromPEM(privPEM)

	cipher, _ := RsaEncrypt(pubKey, []byte("hello rsa"))
	plain, _ := RsaDecrypt(privKey, cipher)
	fmt.Println("RSA:", string(plain))
}
