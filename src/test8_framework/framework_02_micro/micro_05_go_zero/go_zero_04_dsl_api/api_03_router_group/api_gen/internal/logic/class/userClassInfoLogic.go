package class

import (
	"context"

	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_group/api_gen/internal/svc"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_group/api_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserClassInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassInfoLogic {
	return &UserClassInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserClassInfoLogic) UserClassInfo(req *types.UserClassReq) (resp *types.UserClassResp, err error) {
	// todo: add your logic here and delete this line

	return
}
