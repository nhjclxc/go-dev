package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"grpc_order/internal/logic"
	"grpc_order/internal/svc"
	"grpc_order/internal/types"
)

// 新增订单
func InsertOrderInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertOrderInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInsertOrderInfoLogic(r.Context(), svcCtx)
		err := l.InsertOrderInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
