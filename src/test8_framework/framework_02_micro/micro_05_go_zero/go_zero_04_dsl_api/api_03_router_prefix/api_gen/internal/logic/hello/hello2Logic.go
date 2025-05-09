package hello

import (
	"context"

	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_prefix/api_gen/internal/svc"
	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_04_dsl_api/api_03_router_prefix/api_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Hello2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHello2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Hello2Logic {
	return &Hello2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Hello2Logic) Hello2(req *types.HelloReq) (resp *types.HelloResp, err error) {
	// todo: add your logic here and delete this line

	return
}
