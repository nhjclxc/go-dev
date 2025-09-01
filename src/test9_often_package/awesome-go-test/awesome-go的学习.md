# awesome-go的学习

[awesome-go](https://github.com/avelino/awesome-go)



## Actor Model


### Actor Model 解释

**Actor Model（演员模型）** 是一种并发计算模型，主要思想是将系统中的“对象”抽象为 **Actor（演员）**，每个 Actor 拥有以下特性：

1. **独立性**：每个 Actor 拥有自己的状态，不共享内存。
2. **接收消息**：Actor 通过消息通信，而不是直接调用其他 Actor 的方法。
3. **处理逻辑**：接收到消息后，Actor 可以：

    * 修改自身状态
    * 发送消息给其他 Actor
    * 创建新的 Actor
4. **无锁并发**：由于不共享内存，减少了传统锁竞争问题，更容易写出安全的并发程序。

**核心理念**：把系统分解成一组独立的、通过消息通信的 Actor，从而实现高度并发、分布式和容错的系统。这也是 Erlang 和 Akka 的设计思想来源。

在 Golang 中，你可以用 goroutine + channel 实现简单的 Actor，也可以使用一些专门的 Actor 框架，比如下面你提供的那些库。

---

**“Libraries for building actor-based programs.”**
用于构建基于 Actor 模型程序的库。

**库列表：**

* **Ergo** - 一个基于 Actor 的框架，具有网络透明性，用于在 Golang 中创建事件驱动架构。灵感来自 Erlang。
* **Goakt** - 一个快速分布式的 Actor 框架，使用 Protocol Buffers 作为消息格式，适用于 Golang。
* **Hollywood** - 超快且轻量的 Actor 引擎，用 Golang 编写。
* **ProtoActor** - 超高速分布式 Actor 框架，支持 Go、C# 和 Java/Kotlin。

---



#