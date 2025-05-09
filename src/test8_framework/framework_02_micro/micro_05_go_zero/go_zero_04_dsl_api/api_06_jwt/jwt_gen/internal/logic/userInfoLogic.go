package logic

import (
	"context"
	"fmt"

	"jwt_gen/internal/svc"
	"jwt_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (this *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	// 获取 jwt 载体信息  // JWT 中设置的字段
	id := this.ctx.Value("id")

	fmt.Printf("UserInfo.jwt:uuid = %v", id)

	/*
	http://localhost:8090/user/info

	{
	    "id": "666"
	}

	 */

	return
}
