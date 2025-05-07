package logic

import (
	"context"
	"errors"
	"fmt"

	"zero-order/internal/svc"
	"zero-order/internal/types"

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

	fmt.Printf("orderInfoLogic.OrderInfo.req = %v \n", req)
	fmt.Println()
	fmt.Printf("orderInfoLogic.OrderInfo.req = %#v \n", req)

	if req.OrderId == 0 {
		err = errors.New("OrderId不能为空！！！")
		return nil, err
	}

	resp = &types.OrderInfoResp{
		OrderId: req.OrderId,
		GoodsName: "飞机杯",
	}

	return
}
