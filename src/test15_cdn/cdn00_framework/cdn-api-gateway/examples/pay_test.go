package examples

import (
	"fmt"
	"github.com/yourorg/cdn-api-gateway/handler"
	common "github.com/yourorg/cdn-common"
	"github.com/yourorg/cdn-pay-service/pkg/pay"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(pay.PayStatusPaid)
	fmt.Println(pay.PayStatusProcessing)
	fmt.Println(pay.PayStatusCanceled)
}

func Test222(t *testing.T) {
	// 创建 PayService 实例，注入实现
	var processor common.PaymentProcessor = pay.NewAlipayPayProcessor()

	// 注入到 API Handler
	payHandler := handler.NewPayHandler(processor)

	// 使用
	err := payHandler.ChargeUser("user123", 100)
	if err != nil {
		panic(err)
	}
}

func Test333(t *testing.T) {
	// 创建 PayService 实例，注入实现
	var processor common.PaymentProcessor = pay.NewWeChatPayProcessor()

	// 注入到 API Handler
	payHandler := handler.NewPayHandler(processor)

	// 使用
	err := payHandler.ChargeUser("user123", 100)
	if err != nil {
		panic(err)
	}
}
