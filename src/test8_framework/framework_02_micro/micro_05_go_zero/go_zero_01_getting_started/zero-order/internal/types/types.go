// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type OrderInfoReq struct {
	OrderId int64 `json:"orderId"`
}

type OrderInfoResp struct {
	OrderId   int64  `json:"orderId"`   //订单id
	GoodsName string `json:"goodsName"` //商品名称
}
