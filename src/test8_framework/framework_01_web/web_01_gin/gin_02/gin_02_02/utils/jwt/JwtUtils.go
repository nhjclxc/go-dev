package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

const JWT_SECRET_KEY = "zxcvbnmlkjhgfdsa123456789" // 自定义密钥（非常重要）

func main() {
	//生成token

	tokenString, err := GenerateToken(uuid.New().String(), 60*60*24)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", tokenString)

	//解析token
	ret, err := ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("userinfo: %v\n", ret)
	fmt.Println(ret.Valid())
	fmt.Println(ret.VerifyExpiresAt(123, true))
}

func GenerateToken(uid string, expire int) (string, error) {
	//或者用下面自定义claim
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Duration(expire) * time.Second).Unix(), // 过期时间，必须设置,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	return tokenString, err
}

// 解析token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(JWT_SECRET_KEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
