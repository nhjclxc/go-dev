package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

//const (
//	SECRETKEY = "1234567890xqwertyuiop" //私钥
//)

func GenerateToken(id, secretkey string, expire int64) (string, error) {
	//或者用下面自定义claim
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Duration(expire) * time.Second).Unix(), // 过期时间，必须设置,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretkey))
	return tokenString, err

	//     claims := make(jwt.MapClaims)
	//    claims["exp"] = now + l.svcCtx.Config.Auth.AccessExpire // 过期时间
	//    claims["iat"] = now                                     // 签发时间
	//    claims["userId"] = anonymous_user.ID                              // 可以加入任意字段
	//    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//    accessToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}

// 解析token
func ParseToken(secretkey, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretkey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
