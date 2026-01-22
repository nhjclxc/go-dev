package delay

import (
	"bp14_task_queue/queue"
	"container/heap"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// DelayQueue 实现 带有延迟队列的 任务队列，
// 无锁设计【Do not communicate by sharing memory; share memory by communicating.】
type DelayQueue struct {
	taskCh        chan *queue.QueueTask // 立马可以被执行的任务队列
	wg            sync.WaitGroup
	ctx           context.Context
	cancel        context.CancelFunc
	concurrentNum int64

	// 记录消费者数量
	addCount int64

	// 延迟有关的
	delayHeap DelayHeap     // 延迟堆
	delayCh   chan struct{} // 通知延迟堆变化 用于唤醒延迟队列 goroutine
	mu        sync.Mutex    // 保护 delayHeap
}

// NewDelayQueue 创建 DelayQueue
func NewDelayQueue(ctx context.Context, concurrentNum int64) *DelayQueue {

	ctx, cancel := context.WithCancel(context.Background())
	bq := &DelayQueue{
		taskCh:        make(chan *queue.QueueTask, 100),
		delayHeap:     DelayHeap{},
		ctx:           ctx,
		cancel:        cancel,
		concurrentNum: concurrentNum,
	}
	bq.delayCh = make(chan struct{})
	heap.Init(&bq.delayHeap)
	bq.startDelayWorker()
	return bq
}

func (b *DelayQueue) Name() string {
	return "DelayQueue"
}

// Enqueue 投递任务（生产）
func (b *DelayQueue) Enqueue(ctx context.Context, queueTask *queue.QueueTask) error {
	if queueTask.TaskUuid == "" {
		queueTask.TaskUuid = "uuid-" + strconv.FormatInt(int64(queueTask.Id), 10)
	}

	select {
	case <-b.ctx.Done():
		fmt.Println("context canceled. [任务队列被关闭]")
		return b.ctx.Err()
	case <-ctx.Done():
		fmt.Println("context canceled. [投递任务被取消]")
		return ctx.Err()
	default:

		b.mu.Lock()
		heap.Push(&b.delayHeap, &DelayedTask{
			task:    queueTask,
			nextRun: time.Now().Add(queueTask.Delay),
		})
		b.mu.Unlock()
		fmt.Println("任务投递成功：", queueTask.TaskUuid)

		// 通知延迟队列 goroutine
		select {
		case b.delayCh <- struct{}{}:
		default: // 避免阻塞
		}
	}
	return nil
}

func (b *DelayQueue) Dequeue(ctx context.Context) (*queue.QueueTask, error) {
	select {
	case <-b.ctx.Done():
		fmt.Println("context canceled. [任务队列被关闭]")
		return nil, b.ctx.Err()
	case <-ctx.Done():
		fmt.Println("context canceled. [获取队列数据失败，操作被取消]")
		return nil, ctx.Err()
	case queueTask := <-b.taskCh:
		fmt.Println("<获取任务队列数据成功>", queueTask.TaskUuid)
		return queueTask, nil
	}
}

func (b *DelayQueue) Ack(ctx context.Context, queueTask *queue.QueueTask) error {

	select {
	case <-b.ctx.Done():
		fmt.Println("context canceled. [任务队列被关闭]")
		return b.ctx.Err()
	case <-ctx.Done():
		fmt.Println("context canceled. [<Ack操作被取消>]")
		return ctx.Err()
	default:
		// 单机内存队列无需操作
		return nil
	}
}

// Nack 失败处理（重试 / 丢弃 / 延迟）
func (b *DelayQueue) Nack(ctx context.Context, queueTask *queue.QueueTask, delayDuration time.Duration, reason error) error {
	select {
	case <-b.ctx.Done():
		fmt.Println("context canceled. [任务队列被关闭]")
		return b.ctx.Err()
	case <-ctx.Done():
		fmt.Println("context canceled. [<Nack<操作被取消>>]")
		return ctx.Err()
	default:
		if queueTask.RetryCount > 3 {
			err := fmt.Errorf("[%s]被多次标记为失败，以丢弃该任务\n", queueTask.TaskUuid)
			fmt.Println(err.Error())
			return err
		}
		queueTask.RetryCount++

		// 实现延迟将任务再次放入队列
		queueTask.Delay = delayDuration
		err := b.Enqueue(ctx, queueTask)
		if err != nil {
			return err
		}

		return nil
	}

}

func (b *DelayQueue) Len() int {
	//len(taskCh)：获取当前 chan 中“已有多少任务”，对于创建的时候是有缓冲区的，那么是当前 buffer 里 已经排队的任务数量。创建的时候是无缓冲区的那么永远是0
	//cap(taskCh)：获取 channel 的“容量上限”

	return len(b.taskCh)
}

func (b *DelayQueue) LenDelay() int {

	return len(b.delayHeap)
}

func (b *DelayQueue) Close() error {
	b.cancel()
	b.wg.Wait()
	return nil
}

func (b *DelayQueue) AddWorker() error {
	if b.addCount+1 > b.concurrentNum {
		s := fmt.Sprintf("消费者数量达到上限，【max=%d】", b.concurrentNum)
		log.Println(s)
		return fmt.Errorf(s)
	}
	b.wg.Add(1)
	atomic.AddInt64(&b.addCount, 1)

	fmt.Println("当前已加入任务队列的消费者数: ", b.addCount)
	return nil
}
func (b *DelayQueue) DoneWorker() {
	b.wg.Done()
	atomic.AddInt64(&b.addCount, -1)
}

// startDelayWorker 开启延迟队列任务
func (b *DelayQueue) startDelayWorker() {
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()

		for {
			b.mu.Lock()
			if len(b.delayHeap) == 0 {
				b.mu.Unlock()
				continue
			}

			top := b.delayHeap[0]
			now := time.Now()
			duration := top.nextRun.Sub(now)
			if duration > 0 {
				b.mu.Unlock()
				continue
			}

			// 获取到了可以立即执行的任务，则投递到主队列
			heap.Pop(&b.delayHeap)
			b.mu.Unlock()
			b.taskCh <- top.task

		}
	}()
}
