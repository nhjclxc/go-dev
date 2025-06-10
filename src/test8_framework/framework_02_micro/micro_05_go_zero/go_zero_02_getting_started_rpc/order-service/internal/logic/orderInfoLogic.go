package logic

import (
	"context"
	"errors"
	"fmt"

	"order-service/internal/svc"
	"order-service/internal/types"

	rpcUser "order-service/rpc/user"

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

	// 调用 rpc
	userReq := rpcUser.UserReq{UserId: "666"}
	userResp, err := l.svcCtx.UserService.GetUserByName(l.ctx, &userReq)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\norderInfoLogic.OrderInfo.GetUserByName.userResp = %v \n\n", userResp)

	resp = &types.OrderInfoResp{
		OrderId: req.OrderId,
		GoodsName: "飞机杯: " + fmt.Sprintf("orderInfoLogic.OrderInfo.GetUserByName.userResp = %v \n", userResp),
	}

	return
}
