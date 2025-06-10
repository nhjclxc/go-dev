package logic

import (
	"context"
	"fmt"

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
	fmt.Printf("\n\n\n GrpcGetCoupon 远程调用成功！！！ %#v \n\n\n", in)

	return &coupon.GrpcGetCouponResp{
		CouponId:   in.CouponId,
		CouponName: "这是一个远程调用的优惠券",
		DeAmount:   999,
	}, nil
}
