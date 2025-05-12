package middleware

import (
	"context"
	"fmt"
	"net/http"
)

type UserAgentMiddleware struct {
}

func NewUserAgentMiddleware() *UserAgentMiddleware {
	return &UserAgentMiddleware{}
}

func (m *UserAgentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation


		val := r.Header.Get("User-Agent")
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "User-Agent", val)
		newReq := r.WithContext(ctx)


		fmt.Println("UserAgentMiddleware 111")
		// Passthrough to next handler if need
		// 如果想在这里改变下一个处理的请求数据，在这里把新的请求对象传入即可
		next(w, newReq)

		// 下一个处理器返回之后，接着执行这里的代码
		// 中间件与 handler 之间的执行就像函数调用一样
		fmt.Println("UserAgentMiddleware 222")

	}
}