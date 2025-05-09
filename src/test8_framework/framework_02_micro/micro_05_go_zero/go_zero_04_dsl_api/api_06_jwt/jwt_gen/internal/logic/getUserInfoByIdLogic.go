package logic

import (
	"context"
	"fmt"

	"jwt_gen/internal/svc"
	"jwt_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (this *GetUserInfoByIdLogic) GetUserInfoById(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	// 获取 jwt 载体信息  // JWT 中设置的字段
	uuid := this.ctx.Value("id")

	fmt.Printf("GetUserInfoByIdLogic.GetUserInfoById:uuid = %v", uuid)

	/*
	http://localhost:8090/user/getUserById?id=666

	 */

	return
}
