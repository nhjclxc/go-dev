package handler

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_09_http/internal/logic"
	"go_zero_09_http/internal/svc"
	"go_zero_09_http/internal/types"
	"net/http"
)


// 模拟sse相关功能
func sendSSEHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSSEReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("mian 出现错误，r = ", r)
			}
		}()

		l := logic.NewSendSSELogic(r.Context(), svcCtx)
		l.SendSSE(&req, w, r)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.Ok(w)
		//	httpx.OkJsonCtx(r.Context(), w, "操作成功")
		//}


		//xhttp "github.com/zeromicro/x/http"
		//code-data 响应格式
		//httpx.JsonBaseResponseCtx(r.Context(), w, err)
		// {"code":0,"msg":"ok","data":{"uid":1,"name":"go-zero"}}

	}
}
