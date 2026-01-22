package delay2

import (
	"bp14_task_queue/queue"
	"time"
)

type DelayedTask struct {
	task    *queue.QueueTask
	nextRun time.Time
	index   int // 堆中索引
}

// DelayHeap 延迟任务堆
type DelayHeap []*DelayedTask

func (h DelayHeap) Len() int            { return len(h) }
func (h DelayHeap) Less(i, j int) bool  { return h[i].nextRun.Before(h[j].nextRun) }
func (h DelayHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].index = i; h[j].index = j }
func (h *DelayHeap) Push(x interface{}) { *h = append(*h, x.(*DelayedTask)) }
func (h *DelayHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
