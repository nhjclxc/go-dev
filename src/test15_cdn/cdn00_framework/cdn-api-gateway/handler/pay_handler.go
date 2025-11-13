package handler

// cdn-api-gateway/handler/pay_handler.go

import (
	"github.com/yourorg/cdn-common"
)

type PayHandler struct {
	processor common.PaymentProcessor
}

// 构造函数注入
func NewPayHandler(p common.PaymentProcessor) *PayHandler {
	return &PayHandler{
		processor: p,
	}
}

func (h *PayHandler) ChargeUser(userID string, amount int64) error {
	return h.processor.Charge(userID, amount)
}
