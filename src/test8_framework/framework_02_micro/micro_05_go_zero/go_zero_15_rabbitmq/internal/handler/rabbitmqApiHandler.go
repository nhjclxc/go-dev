package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_15_rabbitmq/internal/logic"
	"go_zero_15_rabbitmq/internal/svc"
	"go_zero_15_rabbitmq/internal/types"
)

// 获取用户信息
func RabbitmqApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if err := recover(); nil != err {
				fmt.Println("程序异常：", err)
			}
		}()

		l := logic.NewRabbitmqApiLogic(r.Context(), svcCtx)
		err := l.RabbitmqApi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "success")
		}
	}
}

func RabbitmqApiSimpleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if err := recover(); nil != err {
				fmt.Println("程序异常：", err)
			}
		}()

		l := logic.NewRabbitmqApiLogic(r.Context(), svcCtx)
		err := l.RabbitmqApiSimple(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "success")
		}
	}
}
func RabbitmqApiPublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if err := recover(); nil != err {
				fmt.Println("程序异常：", err)
			}
		}()

		l := logic.NewRabbitmqApiLogic(r.Context(), svcCtx)
		err := l.RabbitmqApiPublish(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "success")
		}
	}
}
func RabbitmqApiRouterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if err := recover(); nil != err {
				fmt.Println("程序异常：", err)
			}
		}()

		l := logic.NewRabbitmqApiLogic(r.Context(), svcCtx)
		err := l.RabbitmqApiRouter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "success")
		}
	}
}
func RabbitmqApiTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if err := recover(); nil != err {
				fmt.Println("程序异常：", err)
			}
		}()

		l := logic.NewRabbitmqApiLogic(r.Context(), svcCtx)
		err := l.RabbitmqApiTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "success")
		}
	}
}
func RabbitmqApiRPCHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		defer func() {
			if err := recover(); nil != err {
				fmt.Println("程序异常：", err)
			}
		}()

		l := logic.NewRabbitmqApiLogic(r.Context(), svcCtx)
		err := l.RabbitmqApiRPC(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "success")
		}
	}
}
