package base

import (
	"bp14_task_queue/queue"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// BaseQueue 实现基础的 任务队列，
// 无锁设计【Do not communicate by sharing memory; share memory by communicating.】
type BaseQueue struct {
	taskCh        chan *queue.QueueTask // 任务队列
	wg            sync.WaitGroup
	ctx           context.Context
	cancel        context.CancelFunc
	concurrentNum int64

	// 记录消费者数量
	addCount int64
}

// NewBaseQueue 创建 BaseQueue
func NewBaseQueue(ctx context.Context, concurrentNum int64) *BaseQueue {
	ctx, cancel := context.WithCancel(ctx)

	return &BaseQueue{
		// 创建一个带有缓冲区的队列
		taskCh:        make(chan *queue.QueueTask, 100),
		ctx:           ctx,
		cancel:        cancel,
		concurrentNum: concurrentNum,
	}
}

// 以下用于快速实现一个接口的所有方法
//var _ Queue = &BaseQueue{}

func (b *BaseQueue) Name() string {
	return "BaseQueue"
}

// Enqueue 投递任务（生产）
func (b *BaseQueue) Enqueue(ctx context.Context, queueTask *queue.QueueTask) error {
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
	case b.taskCh <- queueTask:
		fmt.Println("任务投递成功：", queueTask.TaskUuid)
		return nil
	}
}

func (b *BaseQueue) Dequeue(ctx context.Context) (*queue.QueueTask, error) {
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

func (b *BaseQueue) Ack(ctx context.Context, queueTask *queue.QueueTask) error {

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
func (b *BaseQueue) Nack(ctx context.Context, queueTask *queue.QueueTask, delayDuration time.Duration, reason error) error {
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
		go func() {
			// 延迟 delayDuration 重试
			select {
			case <-time.After(delayDuration):
				_ = b.Enqueue(ctx, queueTask)
			case <-ctx.Done():
				return
			case <-b.ctx.Done():
				return
			}
		}()
		return nil
	}

}

func (b *BaseQueue) Len() int {
	//len(taskCh)：获取当前 chan 中“已有多少任务”，对于创建的时候是有缓冲区的，那么是当前 buffer 里 已经排队的任务数量。创建的时候是无缓冲区的那么永远是0
	//cap(taskCh)：获取 channel 的“容量上限”

	return len(b.taskCh)
}

func (b *BaseQueue) Close() error {
	b.cancel()
	b.wg.Wait()
	return nil
}

func (b *BaseQueue) AddWorker() error {
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
func (b *BaseQueue) DoneWorker() {
	b.wg.Done()
	atomic.AddInt64(&b.addCount, -1)
}
