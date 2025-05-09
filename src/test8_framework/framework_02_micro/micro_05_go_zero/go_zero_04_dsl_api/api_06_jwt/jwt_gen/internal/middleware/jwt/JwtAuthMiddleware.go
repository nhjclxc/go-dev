package jwt

import (
	"context"
	"fmt"
	"jwt_gen/internal/config"
	"jwt_gen/utils/jwt"
	"net/http"
	"strings"
)


// JwtAuthMiddleware 自定义jwt认证中间件


// 自定义中间件步骤
// 1. 自定义中间件结构体 JwtAuthMiddleware
// 2. 实现中间件句柄，具体实现看 JwtAuthorization
// 3. 在 serviceContext.go 文件中，注入自定义的中间件 JwtAuthMiddleware
// 4. 在 routes.go 或 handler 中使用中间件，

// 注意，要在 config.go 中注册中间件
//JwtAuthMiddleware  *jwt.JwtAuthMiddleware


type JwtAuthMiddleware struct{
	// 通过依赖注入的方式，将配置文件加进来，以为解析token的时候要用到tokan密钥
	cfg config.Config
}

// NewAuthMiddleware 创建实例
func NewAuthMiddleware(cfg config.Config) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		cfg: cfg,
	}
}

// JwtAuthorization 自定义 token 校验
func (this *JwtAuthMiddleware) JwtAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("\n\n\n 自定义中间件被触发JwtAuthorization \n\n\n")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		tokenMapClaims, err := jwt.ParseToken(this.cfg.Auth.AccessSecret, tokenString)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		err = tokenMapClaims.Valid()
		if err != nil{
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 自定义字段写入 context
		id := tokenMapClaims["id"]

		// 注入 context
		ctx := context.WithValue(r.Context(), "id", id)
		r = r.WithContext(ctx)

		// 调用下一个处理器，即请求放行
		next(w, r)
	}
}
