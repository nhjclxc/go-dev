package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_19_db_mysql/internal/logic"
	"go_zero_19_db_mysql/internal/svc"
	"go_zero_19_db_mysql/internal/types"
)

// 获取用户信息
func MysqlApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MysqlApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMysqlApiLogic(r.Context(), svcCtx)
		resp, err := l.MysqlApi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
