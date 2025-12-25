package health

import (
	"context"
	"sync"
	"time"
)

// Status 健康状态
type Status string

const (
	StatusUp   Status = "up"
	StatusDown Status = "down"
)

// CheckResult 健康检查结果
type CheckResult struct {
	Status  Status                 `json:"status"`
	Details map[string]CheckDetail `json:"details,omitempty"`
}

// CheckDetail 检查详情
type CheckDetail struct {
	Status  Status `json:"status"`
	Message string `json:"message,omitempty"`
	Latency string `json:"latency,omitempty"`
}

// Checker 健康检查接口
type Checker interface {
	Name() string
	Check(ctx context.Context) CheckDetail
}

// Health 健康检查器
type Health struct {
	checkers []Checker
	mu       sync.RWMutex
}

// New 创建健康检查器
func New() *Health {
	return &Health{
		checkers: make([]Checker, 0),
	}
}

// Register 注册检查器
func (h *Health) Register(checker Checker) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.checkers = append(h.checkers, checker)
}

// Check 执行健康检查
func (h *Health) Check(ctx context.Context) CheckResult {
	h.mu.RLock()
	defer h.mu.RUnlock()

	result := CheckResult{
		Status:  StatusUp,
		Details: make(map[string]CheckDetail),
	}

	// 并发执行所有检查
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, checker := range h.checkers {
		wg.Add(1)
		go func(c Checker) {
			defer wg.Done()

			detail := c.Check(ctx)

			mu.Lock()
			result.Details[c.Name()] = detail
			if detail.Status == StatusDown {
				result.Status = StatusDown
			}
			mu.Unlock()
		}(checker)
	}

	wg.Wait()

	return result
}

// CheckWithTimeout 带超时的健康检查
func (h *Health) CheckWithTimeout(timeout time.Duration) CheckResult {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return h.Check(ctx)
}
