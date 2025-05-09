package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jwt_gen/internal/logic"
	"jwt_gen/internal/svc"
	"jwt_gen/internal/types"
)

func getUserByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		req.ID = r.FormValue("id")

		l := logic.NewGetUserInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfoById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
