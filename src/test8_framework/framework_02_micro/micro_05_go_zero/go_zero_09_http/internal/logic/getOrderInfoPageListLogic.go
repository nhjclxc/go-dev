package logic

import (
	"context"
	"fmt"
	"go_zero_09_http/internal/svc"
	"go_zero_09_http/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取订单分页信息
func NewGetOrderInfoPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoPageListLogic {
	return &GetOrderInfoPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderInfoPageListLogic) GetOrderInfoPageList(req *types.OrderInfoPageListReq) (resp *types.OrderInfoPageListResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf(" OrderInfo = %v \n", req)

	offset := (req.PageNum - 1) * req.PageSize
	_ = offset

	// 第一步：查询总数, 一般就是去数据库里面查
	//total, err := l.svcCtx.OrderModel.Count(req.GoodsName)
	//if err != nil {
	//	return nil, err
	//}
	total := 15

	// 第二步：分页查询数据
	//orders, err := l.svcCtx.OrderModel.FindList(req.Keyword, req.PageSize, offset)
	//if err != nil {
	//	return nil, err
	//}
	var orders []types.OrderInfo = []types.OrderInfo{}
	orders = append(orders, types.OrderInfo{
		OrderId:   1,
		GoodsName: "1",
		Price:     1,
		Status:    "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	})
	orders = append(orders, types.OrderInfo{
		OrderId:   12,
		GoodsName: "12",
		Price:     12,
		Status:    "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	})
	orders = append(orders, types.OrderInfo{
		OrderId:   123,
		GoodsName: "123",
		Price:     123,
		Status:    "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	})

	// 将数据库的实体对象，转化为返回的结构体
	var list []types.OrderInfoResp
	for _, order := range orders {
		list = append(list, types.OrderInfoResp{
			OrderId:   order.OrderId,
			GoodsName: order.GoodsName,
		})
	}

	return &types.OrderInfoPageListResp{
		OrderInfoList: list,
		Total:         total,
		PageNum:       req.PageNum,
		PageSize:      req.PageSize,
	}, nil
}
