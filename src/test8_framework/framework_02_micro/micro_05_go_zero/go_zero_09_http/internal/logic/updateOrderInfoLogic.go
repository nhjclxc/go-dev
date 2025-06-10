package logic

import (
	"context"
	"fmt"

	"go_zero_09_http/internal/svc"
	"go_zero_09_http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改订单
func NewUpdateOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderInfoLogic {
	return &UpdateOrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderInfoLogic) UpdateOrderInfo(req *types.UpdateOrderInfoReq) error {
	// todo: add your logic here and delete this line

	fmt.Printf(" UpdateOrderInfo = %v \n", req)

	return nil
}
