package role

import (
	"context"

	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_group/api_gen/internal/svc"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_group/api_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleAddLogic {
	return &UserRoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRoleAddLogic) UserRoleAdd(req *types.UserRoleAddReq) (resp *types.UserRoleAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
