package role

import (
	"context"

	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_group/api_gen/internal/svc"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_group/api_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleInfoLogic {
	return &UserRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleInfoLogic) UserRoleInfo(req *types.UserRoleReq) (resp *types.UserRoleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
