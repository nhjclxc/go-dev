package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"grpc_user/internal/logic"
	"grpc_user/internal/svc"
	"grpc_user/internal/types"
)

// 新增用户
func InsertUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInsertUserInfoLogic(r.Context(), svcCtx)
		err := l.InsertUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
