package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_14_mqtt/internal/logic"
	"go_zero_14_mqtt/internal/svc"
	"go_zero_14_mqtt/internal/types"
)

// 获取用户信息
func mqttApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MqttApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMqttApiLogic(r.Context(), svcCtx)
		err := l.MqttApi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, "200")
		}
	}
}
