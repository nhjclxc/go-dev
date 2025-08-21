package cron_task

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

// CronTaskManager 企业级定时任务管理器
type CronTaskManager struct {
	cron  *cron.Cron
	tasks map[string]cron.EntryID
	mu    sync.RWMutex
}

// NewCronTaskManager 构造函数
func NewCronTaskManager(withSeconds bool) *CronTaskManager {
	var c *cron.Cron
	if withSeconds {
		c = cron.New(cron.WithSeconds())
	} else {
		c = cron.New()
	}
	return &CronTaskManager{
		cron:  c,
		tasks: make(map[string]cron.EntryID),
	}
}

// Start 启动调度器
func (m *CronTaskManager) Start() {
	m.cron.Start()
	fmt.Println("CronTaskManager started")
}

// Stop 停止调度器
func (m *CronTaskManager) Stop() {
	ctx := m.cron.Stop() // 返回 context.Context，等待正在运行的任务完成
	fmt.Println("CronTaskManager stopped")
	<-ctx.Done()
}

// AddTask 添加新任务
func (m *CronTaskManager) AddTask(name string, spec string, job func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.tasks[name]; exists {
		return fmt.Errorf("task %s already exists", name)
	}

	id, err := m.cron.AddFunc(spec, func() {
		fmt.Printf("Task %s started \n", name)
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Task %s panic: %v \n", name, r)
			}
		}()
		job()
		fmt.Printf("Task %s finished \n", name)
	})
	if err != nil {
		return err
	}

	m.tasks[name] = id
	fmt.Printf("Task %s added with spec %s \n", name, spec)
	return nil
}

// RemoveTask 删除任务
func (m *CronTaskManager) RemoveTask(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, exists := m.tasks[name]
	if !exists {
		return fmt.Errorf("task %s not found", name)
	}

	m.cron.Remove(id)
	delete(m.tasks, name)
	fmt.Printf("Task %s removed \n", name)
	return nil
}

// ListTasks 列出所有任务
func (m *CronTaskManager) ListTasks() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.tasks))
	for name := range m.tasks {
		names = append(names, name)
	}
	return names
}

// GetTaskStatus 获取任务状态（Entry）
// 包含：ID、Next、Prev、Schedule
func (m *CronTaskManager) GetTaskStatus(name string) (*cron.Entry, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	id, exists := m.tasks[name]
	if !exists {
		return nil, fmt.Errorf("task %s not found", name)
	}

	entry := m.cron.Entry(id)
	return &entry, nil
}

//
//type CronTaskManager struct {
//	cron *cron.Cron
//	ctx  context.Context
//}
//
//func NewCronTaskManager() *CronTaskManager {
//	return &CronTaskManager{
//		//cron_task: cron_task.New(),
//		//cron_task: cron_task.New(cron_task.WithSeconds()), // 启用秒字段
//		//  用 经典的 7 字段 cron_task（秒 分 时 日 月 星期 年），可以使用 Quartz 解析器
//		cron: cron.New(
//			cron.WithParser(
//				cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor),
//			)),
//	}
//}
//
//// 启动定时任务
//func (u *CronTaskManager) Run(ctx context.Context) {
//	u.ctx = ctx
//	// 每10秒执行一次连接清理任务
//	u.cron.AddFunc("*/10 * * * * *", u.killLongConnections)
//	u.cron.Start()
//	<-u.ctx.Done()
//}
//
//// killLongConnections Orchestrates the disconnection of non-working and timed-out pairs.
//func (u *CronTaskManager) killLongConnections() {
//	slog.Info("CronTaskManager: 开始执行连接清理任务")
//
//	slog.Info("CronTaskManager: 连接清理任务执行完毕")
//}
