// cdn-pay-service/pkg/pay_service.go
package pay

import (
	"fmt"
	common "github.com/yourorg/cdn-common"
)

type WeChatPayService struct{}

// 实现 PaymentProcessor 接口
func (p *WeChatPayService) Charge(userID string, amount int64) error {
	// 支付逻辑：调用支付网关、生成订单
	fmt.Println("WeChatPayService.Charge 支付逻辑：调用支付网关、生成订单")
	return nil
}

func (p *WeChatPayService) QueryStatus(orderID string) (string, error) {
	// 查询支付状态
	fmt.Println("WeChatPayService.QueryStatus 查询支付状态")
	return "WeChatPayService.QueryStatus", nil
}

// 工厂函数，返回接口类型
func NewWeChatPayProcessor() common.PaymentProcessor {
	return &WeChatPayService{}
}
