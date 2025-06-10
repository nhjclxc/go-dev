package logic

import (
	"context"
	"fmt"
	"strconv"

	"go_zero_09_http/internal/svc"
	"go_zero_09_http/internal/types"

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

	fmt.Printf(" OrderInfo = %v \n", req)

	//a := 0
	//_ = 1 / a

	return &types.OrderInfoResp{
		OrderId:   req.OrderId,
		GoodsName: "你好 " + strconv.FormatInt(req.OrderId, 10),
	}, nil
}

