package queue

import (
	"context"
	"time"
)

// QueueTask 每一个任务
type QueueTask struct {
	Id       int            `json:"id" form:"id"`
	Name     string         `json:"name" form:"name"`
	BizType  string         `json:"bizType" form:"bizType"`
	DataMap  map[string]any `json:"dataMap" form:"dataMap"`
	Delay    time.Duration
	DelayTmp int `json:"delayTmp" form:"delayTmp"`
	Priority int `json:"priority" form:"priority"`
	JoinTime time.Time

	// 内部使用的字段
	RetryCount int
	TaskUuid   string
}

// Queue 任务队列接口
type Queue interface {
	// Name 队列名称（用于监控、日志）
	Name() string

	// Enqueue 投递任务（生产）
	Enqueue(ctx context.Context, queueTask *QueueTask) error

	// Dequeue 获取任务（消费）
	Dequeue(ctx context.Context) (*QueueTask, error)

	// Ack 确认任务成功（可选，分布式必须）
	Ack(ctx context.Context, queueTask *QueueTask) error

	// Nack 失败处理（重试 / 丢弃 / 延迟）
	Nack(ctx context.Context, queueTask *QueueTask, delayDuration time.Duration, reason error) error

	// Len 当前队列长度（可选实现）
	Len() int

	// Close 关闭队列
	Close() error
}

// base_queue.go 基础任务队列
// delay_queue.go 延迟任务队列
// priority_queue.go 优先级延迟任务队列
