package tasks

import (
	"context"
	"go_base_project/pkg/logger"
)

// Task 定时任务接口
type Task interface {
	// Name 返回任务名称
	Name() string
	// Execute 执行任务
	Execute(ctx context.Context) error
}

// ExampleTask 示例任务
type ExampleTask struct{}

func NewExampleTask() Task {
	return &ExampleTask{}
}

func (t *ExampleTask) Name() string {
	return "example_task"
}

func (t *ExampleTask) Execute(ctx context.Context) error {
	logger.Info("执行示例任务", "task", t.Name())
	// 这里添加具体的任务逻辑
	return nil
}

// TaskRegistry 任务注册表
var TaskRegistry = map[string]func() Task{
	"example_task": NewExampleTask,
}

// GetTask 根据名称获取任务
func GetTask(name string) Task {
	if factory, ok := TaskRegistry[name]; ok {
		return factory()
	}
	return nil
}
