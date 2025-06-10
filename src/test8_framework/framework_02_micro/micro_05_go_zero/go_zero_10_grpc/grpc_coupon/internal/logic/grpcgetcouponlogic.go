package logic

import (
	"context"

	"grpc_coupon/grpc/coupon"
	"grpc_coupon/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrpcGetCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrpcGetCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrpcGetCouponLogic {
	return &GrpcGetCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrpcGetCouponLogic) GrpcGetCoupon(in *coupon.GrpcGetCouponReq) (*coupon.GrpcGetCouponResp, error) {
	// todo: add your logic here and delete this line

	return &coupon.GrpcGetCouponResp{}, nil
}
