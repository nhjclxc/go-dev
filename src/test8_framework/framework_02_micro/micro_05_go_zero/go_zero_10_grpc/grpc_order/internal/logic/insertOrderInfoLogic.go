package logic

import (
	"context"

	"grpc_order/internal/svc"
	"grpc_order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertOrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增订单
func NewInsertOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertOrderInfoLogic {
	return &InsertOrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertOrderInfoLogic) InsertOrderInfo(req *types.InsertOrderInfoReq) error {
	// todo: add your logic here and delete this line

	return nil
}
