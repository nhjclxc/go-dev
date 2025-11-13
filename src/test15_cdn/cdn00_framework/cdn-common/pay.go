package common

// cdn-common/pay.go

// PaymentProcessor 处理支付接口
type PaymentProcessor interface {

	// Charge 向用户收款
	//
	// Params:
	// 	- userID: 用户ID
	// 	- amount:支付金额
	// Returns:
	// 	- error: 错误信息
	Charge(userID string, amount int64) error

	//QueryStatus 查询支付状态
	QueryStatus(orderID string) (string, error)
}
