package logic

import (
	"context"
	"fmt"

	"go_zero_09_http/internal/svc"
	"go_zero_09_http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除订单
func NewDeleteOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderInfoLogic {
	return &DeleteOrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrderInfoLogic) DeleteOrderInfo(req *types.DeleteOrderInfoReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf(" DeleteOrderInfo = %v \n", req)

	return nil
}
