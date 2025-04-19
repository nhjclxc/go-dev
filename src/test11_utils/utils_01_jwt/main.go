package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	SECRETKEY = "1234567890xqwertyuiop" //私钥
)

// golang之JWT实现
// https://www.bookstack.cn/read/golang_development_notes/zh-4.4.md
// go get github.com/golang-jwt/jwt/
// https://pkg.go.dev/github.com/golang-jwt/jwt/v5
// / https://jwt.io/
func main1() {

	//生成token
	tokenString, err := generateToken(60 * 60 * 24)
	fmt.Println(err)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEyMywiZXhwIjoxNzQ1MTUxNjg3fQ.sLl9JVzf7fJKF10Co5K3yoJjWOGUV-0sj7IXAUdTGac
	fmt.Printf("token: %v\n", tokenString)

	//解析token
	ret, err := ParseToken(tokenString)
	fmt.Println(err)
	fmt.Printf("userinfo: %v\n", ret)
}

// 自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

// generateToken 生成Token
func generateToken(expire int) (string, error) {
	customClaims := &CustomClaims{
		UserId: 123, // 保存用户信息
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(), // 过期时间，必须设置
		},
	}

	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		panic("Token 生成失败！！！" + err.Error())
	}
	return tokenString, err
}

// 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
