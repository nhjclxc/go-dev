package logic

import (
	"context"

	"grpc_coupon/internal/svc"
	"grpc_coupon/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增优惠劵
func NewInsertCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCouponLogic {
	return &InsertCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertCouponLogic) InsertCoupon(req *types.InsertCouponReq) error {
	// todo: add your logic here and delete this line

	return nil
}
