package task

import (
	"context"
	"fmt"
	"time"
)

// CloudAPITask 从云端API获取拉流转推的任务
type CloudAPITask struct {
	interval time.Duration
}

// NewTask01 创建一个新的请求任务
func NewTask01(interval int) *CloudAPITask {
	return &CloudAPITask{
		interval: time.Duration(interval) * time.Second,
	}
}

// Name 返回任务名称
func (t *CloudAPITask) Name() string {
	return "task01"
}

func (t *CloudAPITask) Enable() bool {
	return true
}

// Execute 执行任务
func (t *CloudAPITask) Execute(ctx context.Context) error {

	fmt.Printf("定时任务01被执行 %s, %#v", t.Name(), time.Now().String())

	return nil
}

func (t *CloudAPITask) GetInterval() time.Duration {
	return t.interval
}
