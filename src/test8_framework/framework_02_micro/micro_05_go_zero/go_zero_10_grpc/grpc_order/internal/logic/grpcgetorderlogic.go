package logic

import (
	"context"
	"fmt"

	"grpc_order/grpc/order"
	"grpc_order/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrpcGetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrpcGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrpcGetOrderLogic {
	return &GrpcGetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrpcGetOrderLogic) GrpcGetOrder(in *order.GrpcGetOrderReq) (*order.GrpcGetOrderResp, error) {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n\n GrpcGetOrder 远程调用成功！！！ %#v \n\n\n", in)

	return &order.GrpcGetOrderResp{
		OrderId:   in.OrderId,
		OrderName: "这是一个远程调用的商品",
		Price:     12369,
	}, nil
}
