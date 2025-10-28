# 协程池

非常棒 👍，你问到了 **Go 并发核心技术之一：协程池（Goroutine Pool）**。
我来给你分成两部分讲清楚：

1️⃣ **业界主流的开源协程池库有哪些、优缺点对比**
2️⃣ **如何自己从零实现一个简单可用的 Goroutine Pool（含源码讲解）**

---

## 一、Go 常用线程池 / 协程池开源库

Go 中“线程池”其实是 **goroutine pool** —— 因为 goroutine 已经是用户级轻量线程。
主要是为了解决：**任务多、但不能无限开 goroutine** 的问题。

### 📦 1. [ants](https://github.com/panjf2000/ants)

> ✅ **最流行、最成熟、性能最好的 goroutine pool 库**

* 🚀 特点：

    * 极高性能（减少内存分配和 GC 压力）
    * 动态调整池大小
    * 可复用 worker，任务自动回收
    * 支持超时、panic 保护
* 💡 适用场景：

    * 高并发短任务
    * Web 服务、爬虫、日志处理等
* 📘 使用示例：

  ```go
  import (
      "fmt"
      "github.com/panjf2000/ants/v2"
      "time"
  )

  func main() {
      pool, _ := ants.NewPool(10) // 最大10个协程
      defer pool.Release()

      for i := 0; i < 100; i++ {
          _ = pool.Submit(func() {
              fmt.Println("run task")
              time.Sleep(100 * time.Millisecond)
          })
      }
  }
  ```

---

### 📦 2. [tunny](https://github.com/Jeffail/tunny)

> 🧩 比 ants 简单的 worker pool，核心理念是同步阻塞式的任务执行。

* 🚀 特点：

    * 实现轻量，接口简单。
    * worker 数量固定，适合 CPU 密集型任务。
* 📘 示例：

  ```go
  import (
      "github.com/Jeffail/tunny"
      "fmt"
  )

  func main() {
      pool := tunny.NewFunc(4, func(payload interface{}) interface{} {
          return fmt.Sprintf("Processed: %v", payload)
      })
      defer pool.Close()

      output := pool.Process("Hello")
      fmt.Println(output)
  }
  ```

---

### 📦 3. [goworkerpool](https://github.com/alitto/pond)

> 🦢 另一个现代协程池库，性能和易用性介于 ants 和 tunny 之间。

* 🚀 特点：
    * 可调度任务数。
    * 动态扩容。
    * 自动 panic 恢复。
* 📘 示例：

  ```go
  import (
      "fmt"
      "github.com/alitto/pond"
  )

  func main() {
      pool := pond.New(10, 1000) // 10个worker, 队列1000任务
      defer pool.StopAndWait()

      for i := 0; i < 100; i++ {
          pool.Submit(func() {
              fmt.Println("Task executed")
          })
      }
  }
  ```

---

## 二、自己实现一个协程池（从零讲解）

我们自己动手实现一个最小可用的 goroutine pool，理解内部原理。

---

### 🧠 核心思想

1️⃣ 创建一个任务通道 `jobs`
2️⃣ 启动固定数量的 worker goroutine 从 `jobs` 取任务执行
3️⃣ 外部通过 `Submit()` 投递任务
4️⃣ 等待任务完成后关闭池

---

### 💻 实现示例

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Task func()

type Pool struct {
    jobs    chan Task
    wg      sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
    p := &Pool{
        jobs: make(chan Task),
    }

    // 启动固定数量 worker
    for i := 0; i < workerCount; i++ {
        go func(id int) {
            for job := range p.jobs {
                job()
                p.wg.Done()
            }
        }(i)
    }
    return p
}

// 提交任务
func (p *Pool) Submit(task Task) {
    p.wg.Add(1)
    p.jobs <- task
}

// 等待所有任务结束
func (p *Pool) Wait() {
    p.wg.Wait()
}

// 关闭池
func (p *Pool) Close() {
    close(p.jobs)
}

func main() {
    pool := NewPool(5)
    for i := 0; i < 20; i++ {
        n := i
        pool.Submit(func() {
            fmt.Printf("Worker: processing task %d\n", n)
            time.Sleep(200 * time.Millisecond)
        })
    }

    pool.Wait()
    pool.Close()
    fmt.Println("All tasks done.")
}
```

**执行效果：**

```
Worker: processing task 0
Worker: processing task 1
Worker: processing task 2
...
All tasks done.
```

---

### 💡 特性扩展（在生产中常见）

你可以在上面的基础上加入：

| 功能             | 实现思路                 |
| -------------- | -------------------- |
| 动态调整 worker 数量 | 定时检查任务队列长度，扩缩容       |
| 超时 / 取消        | `context.Context` 控制 |
| panic 恢复       | `defer recover()` 捕获 |
| 限流             | 使用带缓冲的 channel       |
| 统计             | 记录执行时间、任务数、错误数       |

---

## 三、总结对比

| 库/方式    | 性能    | 特点        | 难度 |
| ------- | ----- | --------- | -- |
| `ants`  | ⭐⭐⭐⭐⭐ | 高性能、成熟稳定  | 中等 |
| `pond`  | ⭐⭐⭐⭐  | 简洁 + 动态扩容 | 简单 |
| `tunny` | ⭐⭐⭐   | 轻量同步模型    | 简单 |
| 自己实现    | ⭐⭐    | 可定制化、学习原理 | 较高 |

---

✅ **一句话总结：**

> 小项目：自己写或用 tunny。
> 中大型服务：推荐 [ants](https://github.com/panjf2000/ants)。
> 想深入理解 Go 并发原理：自己实现一遍。

---

如果你愿意，我可以帮你写一个 **生产级可取消、带超时、自动恢复的协程池模板（带 Context 支持）**，直接能用于实际项目。
要我帮你写这个进阶版本吗？


