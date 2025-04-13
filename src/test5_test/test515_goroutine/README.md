


# GPT：go中的并发、并行和协程

Go 语言以其强大的并发（Concurrency）支持而闻名，它通过 **goroutine** 和 **channel** 提供了一种简单、高效的并发编程方式。在 Go 中，并发（Concurrency）和并行（Parallelism）是两个不同的概念，同时，Go 通过 **goroutine** 实现了轻量级的并发能力。下面我们详细介绍这三者的区别和 Go 语言中的具体实现。

---

## 1. 并发（Concurrency）
**并发（Concurrency）** 指的是程序能够同时处理多个任务，但不一定是同时运行多个任务，而是任务之间交替执行（在单核 CPU 上通过任务切换实现）。Go 通过 **goroutine** 实现了高效的并发。

### Go 如何实现并发？
- 通过 **goroutine** 启动并发任务
- 通过 **channel** 进行 goroutine 之间的通信
- 通过 **sync 包** 提供的 **WaitGroup**、**Mutex** 等工具进行同步控制

#### 代码示例：
```go
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, "执行第", i, "次")
		time.Sleep(time.Millisecond * 500) // 模拟任务执行
	}
}

func main() {
	go task("任务A") // 并发执行任务A
	go task("任务B") // 并发执行任务B

	time.Sleep(time.Second * 3) // 防止主 goroutine 结束
	fmt.Println("主程序退出")
}
```
**输出示例（可能不同，因为 goroutine 执行是非确定性的）：**
```
任务A 执行第 0 次
任务B 执行第 0 次
任务A 执行第 1 次
任务B 执行第 1 次
...
```
可以看到，任务 A 和任务 B 在“交替”执行，而不是严格的串行执行，这就是 **并发** 的体现。

---

## 2. 并行（Parallelism）
**并行（Parallelism）** 指的是多个任务真正“同时”运行，通常需要 **多核 CPU** 支持。例如，在一个四核 CPU 上，最多可以同时运行四个 goroutine（如果 GOMAXPROCS 允许的话）。

### Go 如何实现并行？
- Go 运行时默认使用 **GOMAXPROCS** 变量来决定可以使用多少个 CPU 核心。
- 默认情况下，Go 运行时会使用 **所有可用的 CPU** 来执行 goroutine。

#### 代码示例：
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d 执行第 %d 次\n", id, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // 设置最多使用 2 个 CPU 核心

	for i := 0; i < 4; i++ {
		go task(i)
	}

	time.Sleep(time.Second * 3) // 等待 goroutine 执行完毕
	fmt.Println("主程序退出")
}
```
**输出示例（并行执行，多个 goroutine 可能同时运行）：**
```
Goroutine 0 执行第 0 次
Goroutine 1 执行第 0 次
Goroutine 2 执行第 0 次
Goroutine 3 执行第 0 次
Goroutine 0 执行第 1 次
Goroutine 1 执行第 1 次
...
```
在 **多核 CPU** 上，这些 goroutine 可能会**真正并行**运行。

---

## 3. 协程（Goroutine）
**Goroutine** 是 Go 语言中的一种轻量级线程，它由 Go 运行时管理，而不是由操作系统线程直接管理。相比于传统的线程，**goroutine 更加高效、占用更少的资源**。

### 特点：
1. **轻量级**：Goroutine 启动时只占用 **2KB 栈空间**，相比操作系统线程更节省资源。
2. **调度由 Go 运行时管理**：Goroutine 运行在 **Go 运行时的调度器** 之上，Go 会自动进行 **M:N 线程调度**（即多个 goroutine 共享多个 OS 线程）。
3. **自动扩展栈空间**：Goroutine 的栈大小可以 **动态增长**，而不像操作系统线程那样有固定的栈大小。

#### Goroutine 代码示例：
```go
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Hello")  // 创建 goroutine
	go printMessage("World")  // 创建 goroutine
	time.Sleep(time.Second * 2) // 等待 goroutine 运行
}
```
**输出示例（goroutine 并发执行）：**
```
Hello 0
World 0
Hello 1
World 1
Hello 2
World 2
```
这里 `printMessage("Hello")` 和 `printMessage("World")` 是 **并发** 运行的，而不是一个执行完才执行另一个。

---

## 4. 并发 vs 并行的区别
| 特性      | 并发（Concurrency） | 并行（Parallelism） |
|----------|------------------|----------------|
| 含义      | 任务交替执行，多个任务同时处理，但不一定同时运行 | 任务真正同时运行（需要多核 CPU） |
| 依赖 CPU | 不一定需要多核 CPU，在单核 CPU 也能实现 | 需要多核 CPU 来提高性能 |
| Go 支持 | 通过 goroutine + channel 实现 | 通过 `GOMAXPROCS` 设置多核执行 |
| 现实例子 | 在餐厅，服务员交替处理多个顾客的点单 | 在工厂里，多名工人同时生产多个产品 |

---

## 5. Goroutine 的调度原理
Go 运行时（Goroutine Scheduler）使用 **M:N 调度模型** 来管理 goroutine：
- **M（Threads）：** 操作系统线程
- **P（Processor）：** 逻辑处理器，控制 goroutine 在 CPU 上的调度
- **G（Goroutine）：** Go 的协程，运行在 M 线程上

调度过程：
1. **创建 Goroutine**：Go 代码调用 `go func()` 时，会创建一个新的 goroutine。
2. **调度 Goroutine**：Go 运行时的调度器会将 goroutine 分配给 **P** 处理器，再由 **P** 交给 **M**（OS 线程）执行。
3. **Goroutine 挂起 & 复用**：当一个 goroutine 发生 I/O 阻塞时，Go 运行时会将其挂起，并调度其他 goroutine 运行，以提高 CPU 利用率。

---

## 6. 何时使用并发和并行？
- **使用并发（Concurrency）：**
    - 任务间需要频繁切换（如 I/O 密集型任务）
    - 需要同时处理多个任务，如 Web 服务器并发处理 HTTP 请求

- **使用并行（Parallelism）：**
    - 计算密集型任务（如图像处理、大量数据计算）
    - 需要最大化 CPU 使用率，利用多核 CPU 提高计算效率

---

## 总结
1. **并发（Concurrency）** 是任务交替执行，**并行（Parallelism）** 是任务真正同时运行。
2. **Goroutine 是 Go 实现并发的核心机制**，它轻量级且高效。
3. **Go 通过 GOMAXPROCS 控制并行度**，默认情况下会利用所有 CPU 核心。
4. **使用 goroutine + channel 实现高效的并发控制**，避免传统线程带来的高开销。

如果你的应用是 I/O 密集型，建议使用 **goroutine + channel**；如果是计算密集型，建议优化 **GOMAXPROCS** 以充分利用多核 CPU。


# go的数据类型或一些常用包的方法中，哪些是线程安全的，哪些是线程不安全的

在 Go 语言中，某些数据类型和标准库包是线程安全的，而有些则是线程不安全的。以下是一些常见的线程安全和线程不安全的数据类型和方法：

---

### **1. 线程安全的**
#### **(1) 内置数据类型**
- `sync.Mutex` / `sync.RWMutex`  —— 互斥锁/读写锁，能确保数据在并发访问时安全。
- `sync.Once` —— 只执行一次，适用于单次初始化。
- `sync.WaitGroup` —— 等待一组 goroutine 结束。
- `sync.Cond` —— 用于 goroutine 之间的条件同步。
- `sync.Map` —— 并发安全的 map（但比普通 map 慢）。
- `sync.Pool` —— 对象池，多个 goroutine 可以安全地获取/放回对象。
- `atomic` 包的方法（如 `atomic.AddInt32`、`atomic.LoadUint64`）—— 提供原子操作，保证基本数据类型的安全读写。
- `chan`（管道） —— 在大多数情况下是线程安全的，但需确保正确的使用方式，避免死锁、非缓冲通道阻塞等问题。

#### **(2) 标准库中的线程安全方法**
- `time.After()` / `time.Tick()` —— 用于定时器，多个 goroutine 读写是安全的。
- `log.Logger`（使用 `log.Default()` 时是线程安全的，默认带有 `sync.Mutex` 保护）。
- `net/http` 中的 `http.Client`（但要注意 `Transport` 配置的共享问题）。
- `context.Context` —— 用于控制超时和取消，一般是线程安全的。
- `bytes.Buffer`（通过 `sync.Pool` 结合使用可以提高并发性能）。

---

### **2. 线程不安全的**
#### **(1) 内置数据类型**
- `map`（原生 `map` 不是线程安全的，多 goroutine 读写会导致 `fatal error: concurrent map writes`）。
- `slice`（多 goroutine 并发修改同一个 slice 可能会导致数据异常，如 `append` 导致容量扩展时发生的竞态条件）。
- `string`（虽然 `string` 是不可变的，但如果多个 goroutine 读写同一个 `string` 变量的指针，可能会导致竞态问题）。

#### **(2) 标准库中线程不安全的方法**
- `bytes.Buffer.Write()`（多个 goroutine 并发写入会导致数据错乱）。
- `strings.Builder.Write()`（多个 goroutine 并发写入会导致数据异常）。
- `encoding/json.Marshal()` / `encoding/json.Unmarshal()`（并发对同一个对象调用 `Marshal` 或 `Unmarshal` 可能会出现数据竞态）。
- `rand.New()` / `math/rand.Rand`（默认 `math/rand` 不是线程安全的，需要使用 `rand.New(rand.NewSource(seed))` 创建不同的 `Rand` 实例）。
- `net/http.ServeMux`（Go 1.20 之前默认不是线程安全的，需要加锁）。
- `os.File.Write()`（多个 goroutine 同时写入同一个文件，可能会发生数据错乱）。
- `fmt.Fprintf()`（多个 goroutine 对同一个 `io.Writer` 写入时，可能会出现并发问题）。

---

### **3. 线程安全 vs 线程不安全 对比表**

| **类别** | **线程安全** | **线程不安全** |
|----------|-------------|----------------|
| **内置数据类型** | `sync.Mutex`、`sync.RWMutex`、`sync.Once`、`sync.WaitGroup`、`sync.Map`、`sync.Pool`、`chan` | `map`、`slice`、`string` |
| **sync 包** | `sync.Mutex`、`sync.RWMutex`、`sync.Once`、`sync.WaitGroup`、`sync.Map`、`sync.Pool` | - |
| **atomic 包** | `atomic.AddInt32`、`atomic.LoadUint64` | - |
| **log 包** | `log.Default()` | `log.Logger`（如果未手动加锁） |
| **time 包** | `time.After()`、`time.Tick()` | - |
| **net/http 包** | `http.Client` | `http.ServeMux`（Go 1.20 之前） |
| **encoding/json 包** | - | `json.Marshal()`、`json.Unmarshal()` |
| **math/rand 包** | `rand.New(rand.NewSource(seed))` | `math/rand.Rand`（默认情况下） |
| **os 包** | - | `os.File.Write()` |
| **fmt 包** | - | `fmt.Fprintf()` |
| **bytes 包** | - | `bytes.Buffer.Write()` |
| **strings 包** | - | `strings.Builder.Write()` |

这样，你可以在开发时避免在多线程环境下使用不安全的数据类型或方法，或者使用锁、`sync`、`atomic` 等手段来确保线程安全。 🚀



# 14.0 协程 (goroutine) 与通道 (channel)

https://denganliang.github.io/the-way-to-go_ZH_CN/14.0.html


作为一门 21 世纪的语言，Go 原生支持应用之间的通信（网络，客户端和服务端，分布式计算，参见第 15 章和程序的并发。程序可以在不同的处理器和计算机上同时执行不同的代码段。Go 语言为构建并发程序的基本代码块是协程 (goroutine) 与通道 (channel)。他们需要语言，编译器，和 runtime 的支持。Go 语言提供的垃圾回收器对并发编程至关重要。

不要通过共享内存来通信，而通过通信来共享内存。

通信强制协作。

**在Go中：应用程序 > 进场 > 线程 > 协程，协程比线程轻很多，协程直接由go管理，不需要交于操作系统**

## 14.1.1 什么是协程
一个应用程序是运行在机器上的一个进程；进程是一个运行在自己内存地址空间里的独立执行体。一个进程由一个或多个操作系统线程组成，这些线程其实是共享同一个内存地址空间的一起工作的执行体。几乎所有’正式’的程序都是多线程的，以便让用户或计算机不必等待，或者能够同时服务多个请求（如 Web 服务器），或增加性能和吞吐量（例如，通过对不同的数据集并行执行代码）。一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。

并行是一种通过使用多处理器以提高速度的能力。所以并发程序可以是并行的，也可以不是。

公认的，使用多线程的应用难以做到准确，最主要的问题是内存中的数据共享，它们会被多线程以无法预知的方式进行操作，导致一些无法重现或者随机的结果（称作竞态）。

不要使用全局变量或者共享内存，它们会给你的代码在并发运算的时候带来危险。

解决之道在于同步不同的线程，对数据加锁，这样同时就只有一个线程可以变更数据。在 Go 的标准库 sync 中有一些工具用来在低级别的代码中实现加锁；我们在第 9.3 节中讨论过这个问题。不过过去的软件开发经验告诉我们这会带来更高的复杂度，更容易使代码出错以及更低的性能，所以这个经典的方法明显不再适合现代多核/多处理器编程：thread-per-connection 模型不够有效。

Go 更倾向于其他的方式，在诸多比较合适的范式中，有个被称作 Communicating Sequential Processes（顺序通信处理）（CSP, C. Hoare 发明的）还有一个叫做 message passing-model（消息传递）（已经运用在了其他语言中，比如 Erlang）。

在 Go 中，应用程序并发处理的部分被称作 goroutines（协程），它可以进行更有效的并发运算。在协程和操作系统线程之间并无一对一的关系：协程是根据一个或多个线程的可用性，映射（多路复用，执行于）在他们之上的；协程调度器在 Go 运行时很好的完成了这个工作。

协程工作在相同的地址空间中，所以共享内存的方式一定是同步的；这个可以使用 sync 包来实现（参见第 9.3 节），不过我们很不鼓励这样做：Go 使用 channels 来同步协程（可以参见第 14.2 节等章节）

当系统调用（比如等待 I/O）阻塞协程时，其他协程会继续在其他线程上工作。协程的设计隐藏了许多线程创建和管理方面的复杂工作。

协程是轻量的，比线程更轻。它们痕迹非常不明显（使用少量的内存和资源）：使用 4K 的栈内存就可以在堆中创建它们。因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。

协程可以运行在多个操作系统线程之间，也可以运行在线程之内，让你可以很小的内存占用就可以处理大量的任务。由于操作系统线程上的协程时间片，你可以使用少量的操作系统线程就能拥有任意多个提供服务的协程，而且 Go 运行时可以聪明的意识到哪些协程被阻塞了，暂时搁置它们并处理其他协程。

存在两种并发方式：确定性的（明确定义排序）和非确定性的（加锁/互斥从而未定义排序）。Go 的协程和通道理所当然的支持确定性的并发方式（例如通道具有一个 sender 和一个 receiver）。我们会在第 14.7 节中使用一个常见的算法问题（工人问题）来对比两种处理方式。

协程是通过使用关键字 go 调用（执行）一个函数或者方法来实现的（也可以是匿名或者 lambda 函数）。这样会在当前的计算过程中开始一个同时进行的函数，在相同的地址空间中并且分配了独立的栈，比如：go sum(bigArray)，在后台计算总和。

协程的栈会根据需要进行伸缩，不出现栈溢出；开发者不需要关心栈的大小。当协程结束的时候，它会静默退出：用来启动这个协程的函数不会得到任何的返回值。

任何 Go 程序都必须有的 main() 函数也可以看做是一个协程，尽管它并没有通过 go 来启动。协程可以在程序初始化的过程中运行（在 init() 函数中）。

在一个协程中，比如它需要进行非常密集的运算，你可以在运算循环中周期的使用 runtime.Gosched()：这会让出处理器，允许运行其他协程；它并不会使当前协程挂起，所以它会自动恢复执行。使用 Gosched() 可以使计算均匀分布，使通信不至于迟迟得不到响应。








