package delay2

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

	delayTaskCh chan *DelayedTask // 通知延迟堆变化 用于唤醒延迟队列 goroutine
}

// NewDelayQueue 创建 DelayQueue
func NewDelayQueue(ctx context.Context, concurrentNum int64) *DelayQueue {

	ctx, cancel := context.WithCancel(context.Background())
	bq := &DelayQueue{
		taskCh:        make(chan *queue.QueueTask, 100),
		delayTaskCh:   make(chan *DelayedTask, 100),
		ctx:           ctx,
		cancel:        cancel,
		concurrentNum: concurrentNum,
	}
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

		task := &DelayedTask{
			task:    queueTask,
			nextRun: time.Now().Add(queueTask.Delay),
		}

		select {
		case b.delayTaskCh <- task:
		default: // 避免阻塞
			return fmt.Errorf("delay queue full")
		}
		log.Println("任务投递成功：", queueTask.TaskUuid)
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
		log.Println("<获取任务队列数据成功>", queueTask.TaskUuid)
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

func (b *DelayQueue) Close() error {
	b.cancel()
	b.wg.Wait()
	return nil
}

func (b *DelayQueue) AddWorker() error {
	if atomic.AddInt64(&b.addCount, 1) > b.concurrentNum {
		atomic.AddInt64(&b.addCount, -1)

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

		// 延迟堆 只在当前 goroutine 使用，保证并发安全
		delayHeap := DelayHeap{}
		//把一个“普通 slice”调整成“满足堆性质的结构”，让它能被当作堆来正确使用。
		heap.Init(&delayHeap)
		var timer *time.Timer
		// 在 delay 包的基础上，实现无锁设计
		for {

			var timerCh <-chan time.Time
			if len(delayHeap) > 0 {
				d := time.Until(delayHeap[0].nextRun)
				if d < 0 {
					d = 0
				}
				if timer == nil {
					timer = time.NewTimer(d)
				} else {
					// timer.Stop() 返回是否停止的标志
					// true		定时器 还没触发，被成功停止
					// false	定时器 已经触发 or 正在触发

					// 📌 保证：timer.C 是干净的
					if !timer.Stop() {
						select {
						//非阻塞地尝试从 timer.C 读一次
						//如果真有信号 → 吃掉
						//如果没有 → 立刻返回
						case <-timer.C: // 清掉旧的 tick
						default:
						}
					}
					// 将定时器的时间重置
					timer.Reset(d)
				}
				timerCh = timer.C
			}

			// 以下将等待三个通道
			// 1、<-b.ctx.Done()：是系统的ctx，当队列被关闭的时候触发
			// 2、task := <-b.delayTaskCh: 新任务通道，当有新任务被提交的时候触发该通道，重新进入for，即重新创建一个延时器
			// 3、<-timerCh：等待延时器到来，当第一个任务到了的时候，直接将该任务提交到执行队列
			select {
			case <-b.ctx.Done():
				fmt.Println("context canceled. [<UNK>]")
				return
			case task := <-b.delayTaskCh:
				// 新任务来了，直接加入堆
				heap.Push(&delayHeap, task)
			case <-timerCh:
				if len(delayHeap) > 0 {
					t := heap.Pop(&delayHeap).(*DelayedTask)
					b.taskCh <- t.task
				}
			}
		}
	}()
}
