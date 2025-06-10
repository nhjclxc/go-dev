package middleware

import (
	"fmt"
	"net/http"
)


// 使用自定义中间件的不在
// 1、定义中间件（customMiddleware.go），注意：中间件方法必须绑定到一个结构体上面
// 2、注册中间件（serviceContext.go），CustomMiddleware  *middleware.CustomMiddleware
// 3、注入中间件（serviceContext.go），CustomMiddleware: middleware.NewCustomMiddleware(),
// 4、使用中间件（routers.go），rest.WithMiddlewares( ...


type CustomMiddleware struct {
}

func NewCustomMiddleware() *CustomMiddleware {
	return &CustomMiddleware{}
}

// 自定义的中间件
func (m *CustomMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation


		fmt.Printf("执行到了自定义中间件 CustomMiddleware \n")


		fmt.Println("CustomMiddleware 111")
		// Passthrough to next handler if need
		// 如果想在这里改变下一个处理的请求数据，在这里把新的请求对象传入即可
		next(w, r)

		// 下一个处理器返回之后，接着执行这里的代码
		// 中间件与 handler 之间的执行就像函数调用一样
		fmt.Println("CustomMiddleware 222")

	}
}