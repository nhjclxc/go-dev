package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"grpc_coupon/internal/logic"
	"grpc_coupon/internal/svc"
	"grpc_coupon/internal/types"
)

// 获取优惠劵信息
func GetCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCouponLogic(r.Context(), svcCtx)
		resp, err := l.GetCoupon(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
