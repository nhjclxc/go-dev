package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"grpc_coupon/internal/logic"
	"grpc_coupon/internal/svc"
	"grpc_coupon/internal/types"
)

// 新增优惠劵
func InsertCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInsertCouponLogic(r.Context(), svcCtx)
		err := l.InsertCoupon(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
