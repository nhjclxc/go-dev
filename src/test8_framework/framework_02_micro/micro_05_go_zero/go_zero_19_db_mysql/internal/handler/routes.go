// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package handler

import (
	"net/http"

	"go_zero_19_db_mysql/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// http://127.0.0.1:8090/mysql/insert
				/*
				{
				    "name": "namenamez",
				    "email": "emailemail",
				    "memberNumber": "memberNumbermemberNumber",
				    "remark": "remarkremark"
				}
				 */
				Method:  http.MethodPost,
				Path:    "/mysql/insert",
				Handler: MysqlInsertHandler(serverCtx),
			},

			//	// http://127.0.0.1:8090/mysql/update
			//	/*
			//	{
			//
			//	    "name": "namenamez",
			//	    "email": "emailemail",
			//	    "memberNumber": "memberNumbermemberNumber",
			//	    "remark": "remarkremark"
			//	}
			//	 */
			//	Method:  http.MethodPut,
			//	Path:    "/mysql/update",
			//	Handler: MysqlUpdateHandler(serverCtx),
			//},



			{
				// 获取用户信息
				Method:  http.MethodGet,
				Path:    "/mysql/get",
				Handler: MysqlApiHandler(serverCtx),
			},
		},
	)
}
