package hello

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_prefix/api_gen/internal/logic/hello"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_prefix/api_gen/internal/svc"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_prefix/api_gen/internal/types"
)

func Hello1Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hello.NewHello1Logic(r.Context(), svcCtx)
		resp, err := l.Hello1(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
