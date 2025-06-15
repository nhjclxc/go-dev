package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_18_delay_queue_redis/internal/logic"
	"go_zero_18_delay_queue_redis/internal/svc"
	"go_zero_18_delay_queue_redis/internal/types"
)

// 获取用户信息
func DelayQueueRedisHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelayQueueRedisReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelayQueueRedisLogic(r.Context(), svcCtx)
		err := l.DelayQueueRedis(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
