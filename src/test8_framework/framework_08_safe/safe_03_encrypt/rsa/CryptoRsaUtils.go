package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// ParseRSAPublicKeyFromPEM 解析 PEM 格式公钥
func ParseRSAPublicKeyFromPEM(pubPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

// ParseRSAPrivateKeyFromPEM 解析 PEM 格式私钥
func ParseRSAPrivateKeyFromPEM(privPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key format")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// RsaEncrypt 使用公钥加密数据
func RsaEncrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// RsaDecrypt 使用私钥解密数据
func RsaDecrypt(priv *rsa.PrivateKey, cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
