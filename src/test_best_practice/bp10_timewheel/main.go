package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

/*
timewheel简介
	时间轮是一个环形队列，底层实现就是一个固定长度的数组，数组中的每个元素存储一个双向列表，这个列表存放着该时间内需要执行的所有任务

例子
	抽象点来说时钟表盘就是1秒为一个时间刻度，一共（一天）会有86400个刻度的时间轮，当指针走到那个刻度的时候就可以把对应的任务全部取出执行
	也就是说这里我们定义了一个可以延迟86400秒的时候轮，不过当我们定一个86401秒后执行的任务怎么办
		方案一 多级时间轮 （这里不展开描述）
		方案二 定义circle参数 （本文）, 一轮是86400秒的话，86401= 86400 + 1 也就是 1circle（一轮）加上一个刻度就可以取出该任务执行
*/

// 这里定义一个长度为60的时间轮，每个位置表示秒数

type TimeWheel struct {
	interval     time.Duration // 定时器轮训时间间隔
	slots        []list.List   // 这是每一个时间槽的双向连表，每一个索引位置存储了该时间槽的所有任务
	slotsNum     int
	currentSlots int
	ticker       *time.Ticker
	mu           sync.Mutex
	isRun        bool
	addTaskCh    chan *Task
	removeTaskCh chan string
	doneCh       chan struct{} // 确认退出完成
	closeCh      chan struct{} // 用来通知关闭
}

type Task struct {
	ID         string
	createTime time.Time
	delay      time.Duration
	slots      int
	circle     int // 多少圈
	job        string
	times      int //执行多少次 -1 一直执行
}

func (t *TimeWheel) do() {
	go func() {
		for {
			select {
			case ti := <-t.ticker.C:
				fmt.Println("time: ", ti.Format("2006-01-02 15:04:05"))
				//slot, _ := strconv.Atoi(fmt.Sprintf("%2d", ti.Second()))
				//slot /= 10
				slot := ti.Second()
				fmt.Println("slot = ", slot)

				t.mu.Lock()
				tasks := copyList(&t.slots[slot])
				t.mu.Unlock()

				for e := tasks.Front(); e != nil; e = e.Next() {
					tt, ok := e.Value.(*Task)
					if ok {
						go func(task *Task) {
							if task.times >= task.circle {
								t.removeTask(task.ID)
								return
							}

							time.Sleep(task.delay)
							fmt.Println("时间轮执行器：", task.ID)
						}(tt)
					}
				}
			}
		}

	}()
}

func copyList(src *list.List) *list.List {
	dst := list.New()
	for e := src.Front(); e != nil; e = e.Next() {
		dst.PushBack(e.Value) // 浅拷贝：只拷贝 Value 的引用
	}
	return dst
}

func (t *TimeWheel) Run() {
	if t.isRun {
		return
	}
	t.isRun = true
	t.do()

	// 开启时间轮
	go func() {
		for {
			select {
			case task, ok := <-t.addTaskCh:
				if ok {
					t.mu.Lock()
					t.slots[task.slots].PushBack(task)
					t.mu.Unlock()
				}
			case taskId, ok := <-t.removeTaskCh:
				if ok {

					t.removeTask(taskId)

				}
			case _, ok := <-t.doneCh:
				if ok {
					t.isRun = false
					t.closeCh <- struct{}{}
					return
				}
			}
		}

	}()
}

func (t *TimeWheel) removeTask(taskId string) {
	t.mu.Lock()
	for _, slot := range t.slots {
		for e := slot.Front(); e != nil; e = e.Next() {
			tt, ok := e.Value.(*Task)
			if ok && taskId == tt.ID {
				slot.Remove(e)
			}
		}
	}
	t.mu.Unlock()
}

func main() {

	slotsNum := 60
	interval := 1 * time.Second

	closeCh := make(chan struct{})
	w := TimeWheel{
		interval:     interval,
		slots:        make([]list.List, slotsNum),
		slotsNum:     slotsNum,
		currentSlots: 0,
		ticker:       time.NewTicker(interval),
		isRun:        false,
		addTaskCh:    make(chan *Task),
		removeTaskCh: make(chan string),
		doneCh:       make(chan struct{}),
		closeCh:      closeCh,
	}

	fmt.Println("时间轮开启中。")
	w.Run()
	fmt.Println("时间轮开启成功！")

	w.addTaskCh <- &Task{
		ID:         "11111",
		createTime: time.Now(),
		delay:      0,
		circle:     5,
		slots:      20,
		job:        "测试11111",
		times:      2,
	}
	w.addTaskCh <- &Task{
		ID:         "22222",
		createTime: time.Now(),
		delay:      0,
		circle:     5,
		slots:      20,
		job:        "测试11111",
		times:      2,
	}

	go func() {
		time.Sleep(111 * time.Second)
		w.doneCh <- struct{}{}
	}()

	<-closeCh
	fmt.Println("时间轮关闭！")

}
