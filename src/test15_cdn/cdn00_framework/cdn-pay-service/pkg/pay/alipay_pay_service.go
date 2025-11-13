// cdn-pay-service/pkg/pay_service.go
package pay

import (
	"fmt"
	common "github.com/yourorg/cdn-common"
)

type AlipayPayService struct{}

// 实现 PaymentProcessor 接口
func (p *AlipayPayService) Charge(userID string, amount int64) error {
	// 支付逻辑：调用支付网关、生成订单
	fmt.Println("AlipayPayService.Charge 支付逻辑：调用支付网关、生成订单")
	return nil
}

func (p *AlipayPayService) QueryStatus(orderID string) (string, error) {
	// 查询支付状态
	fmt.Println("AlipayPayService.QueryStatus 查询支付状态")
	return "AlipayPayService.QueryStatus", nil
}

// 工厂函数，返回接口类型
func NewAlipayPayProcessor() common.PaymentProcessor {
	return &AlipayPayService{}
}
