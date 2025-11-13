package pay

//package pay

// PayStatusType 支付状态类型
type PayStatusType string

const (
	PayStatusPending    PayStatusType = "Pending"    // 未支付，用户尚未完成支付
	PayStatusPaid       PayStatusType = "Paid"       // 已支付，支付成功
	PayStatusFailed     PayStatusType = "Failed"     // 支付失败，如余额不足或渠道错误
	PayStatusCanceled   PayStatusType = "Canceled"   // 已取消，用户主动取消或超时取消
	PayStatusRefunded   PayStatusType = "Refunded"   // 已退款
	PayStatusExpired    PayStatusType = "Expired"    // 支付超时未完成，订单过期
	PayStatusProcessing PayStatusType = "Processing" // 支付处理中（银行/渠道异步确认）
)
