package logic

import (
	"context"
	"fmt"
	"strconv"

	"grpc_order/grpc/coupon"
	"grpc_order/grpc/user"
	"grpc_order/internal/svc"
	"grpc_order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取订单信息
func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	// todo: add your logic here and delete this line


	fmt.Printf("\n\nOrderInfo.req: %#v \n\n", req)

	// 调用微服务 grpc
	gRpcUserReq := user.GrpcGetUserReq{UserId: strconv.FormatInt(req.OrderId, 10)}
	gRpcUserResp, err := l.svcCtx.UserGRpcService.GrpcGetUser(l.ctx, &gRpcUserReq)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\n GRPC调用 user 返回成功：orderInfoLogic.OrderInfo.GetGRpcUser.gRpcUserResp = %v \n\n", gRpcUserResp)



	couponReq := coupon.GrpcGetCouponReq{CouponId: strconv.FormatInt(req.OrderId, 10)}
	grpcCouponResp, err := l.svcCtx.GrpcCouponService.GrpcGetCoupon(l.ctx, &couponReq)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\n GRPC调用 coupon 返回成功：orderInfoLogic.GrpcCouponService.GetUserByName.grpcCouponResp = %v \n\n", grpcCouponResp)


	return &types.OrderInfoResp{
		OrderId:   req.OrderId,
		GoodsName: "返回的商品信息",
	}, nil
}
