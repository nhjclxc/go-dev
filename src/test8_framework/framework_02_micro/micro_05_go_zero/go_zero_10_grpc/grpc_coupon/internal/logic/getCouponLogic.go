package logic

import (
	"context"
	"fmt"
	"grpc_coupon/grpc/order"
	"grpc_coupon/grpc/user"
	"grpc_coupon/internal/svc"
	"grpc_coupon/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取优惠劵信息
func NewGetCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCouponLogic {
	return &GetCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCouponLogic) GetCoupon(req *types.GetCouponReq) (resp *types.GetCouponResp, err error) {
	// todo: add your logic here and delete this line


	fmt.Printf("\n\n GetCoupon.req: %#v \n\n", req)


	// 调用微服务 grpc
	gRpcOrderReq := order.GrpcGetOrderReq{OrderId: strconv.FormatInt(req.CouponId, 10)}
	gRpcOrderResp, err := l.svcCtx.GrpcOrderService.GrpcGetOrder(l.ctx, &gRpcOrderReq)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\n GRPC调用 order 返回成功：UserInfo.OrderGrpcService.GetGrpcOrderByName = %v \n\n", gRpcOrderResp)

	userReq := user.GrpcGetUserReq{UserId: strconv.FormatInt(req.CouponId, 10)}
	getUserByNameResp, err := l.svcCtx.GrpcUserService.GrpcGetUser(l.ctx, &userReq)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\n GRPC调用 user 返回成功：UserInfo.UserGRpcService.GetUserByName = %v \n\n", getUserByNameResp)

	return &types.GetCouponResp{
		CouponId:   req.CouponId,
		CouponName: "这是一个消费券",
		DeAmount:   123,
	}, nil
}
