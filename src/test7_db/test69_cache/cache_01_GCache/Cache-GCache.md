好的，下面我将详细讲解 Go 中常用的缓存库 **GCache**，包括其 **基本概念、核心功能、用法、源码结构以及应用场景与注意事项**，帮助你从入门到深入理解 GCache 的机制和使用。

---

## 一、GCache 简介

[GCache](https://github.com/bluele/gcache) 是一个功能强大、灵活、并发安全的 Go 本地缓存库，提供 **多种缓存淘汰策略（LRU/LFU/ARC）**，并支持自动加载（类似 Guava Cache 的 `LoadingCache`）。

---

## 二、GCache 支持的缓存策略

GCache 支持如下三种经典淘汰策略：

| 策略      | 全称                         | 原理                   |
| ------- | -------------------------- | -------------------- |
| **LRU** | Least Recently Used        | 淘汰最近最少使用的元素（最近没被访问的） |
| **LFU** | Least Frequently Used      | 淘汰使用次数最少的元素          |
| **ARC** | Adaptive Replacement Cache | 结合 LRU 和 LFU 特性，自动调节 |

---

## 三、基本功能

### ✅ 核心特性：

* 指定最大缓存数量（自动淘汰）
* 支持缓存过期（Expiration）
* 支持自动加载（LoaderFunc）
* 并发安全（使用 `sync.RWMutex` 控制并发）
* 支持清除回调（evicted callback）

---

## 四、基本用法

### 1. 创建一个简单的 LRU 缓存

```go
import (
    "github.com/bluele/gcache"
    "time"
    "fmt"
)

func main() {
    cache := gcache.New(100). // 最大缓存项数
        LRU().                // 缓存策略
        Expiration(10 * time.Second). // 过期时间
        Build()

    _ = cache.Set("key", "value")
    
    val, err := cache.Get("key")
    if err == nil {
        fmt.Println(val) // 输出: value
    }
}
```

---

### 2. 使用自动加载功能（防缓存穿透）

```go
cache := gcache.New(20).
    LRU().
    LoaderFunc(func(key interface{}) (interface{}, error) {
        fmt.Println("load from slow source:", key)
        return "data-for-" + key.(string), nil
    }).
    Build()

value, _ := cache.Get("anonymous_user:123") // 第一次会触发 LoaderFunc
fmt.Println(value)

value2, _ := cache.Get("anonymous_user:123") // 第二次从缓存中取
fmt.Println(value2)
```

---

### 3. 使用淘汰回调函数

```go
cache := gcache.New(2).
    LRU().
    EvictedFunc(func(key, value interface{}) {
        fmt.Printf("Evicted key: %v, value: %v\n", key, value)
    }).
    Build()

cache.Set("a", 1)
cache.Set("b", 2)
cache.Set("c", 3) // a 会被淘汰
```

---

## 五、进阶功能与操作

### 判断是否存在：

```go
exists := cache.Has("key")
```

### 删除缓存项：

```go
cache.Remove("key")
```

### 获取所有键（slice）：

```go
keys := cache.Keys(true) // true 表示返回按顺序排序的 key
```

### 清空缓存：

```go
cache.Purge()
```

---

## 六、源码结构简要说明

GCache 核心结构如下：

```go
type Cache interface {
    Set(key, value interface{}) error
    Get(key interface{}) (interface{}, error)
    ...
}
```

构建器模式（Builder）：

```go
gcache.New(100).LRU().Expiration(10 * time.Second).Build()
```

它会根据调用的方法创建不同的内部实现（比如 `simpleCache`, `lruCache`, `lfuCache` 等），实际缓存项则通过 `entry` 结构包装存储：

```go
type entry struct {
    key        interface{}
    value      interface{}
    expiration *time.Time
}
```

淘汰策略是通过不同结构（如 `container/list` 或 `heap`）实现的，LRU 使用双向链表，LFU 使用最小堆。

---

## 七、应用场景

| 场景     | 说明                      |
| ------ | ----------------------- |
| 接口限频   | 缓存用户请求计数，10s 内只允许访问一次   |
| 数据预热   | 应用启动时缓存常用数据，减少后端压力      |
| 本地降级   | 后端宕机时返回缓存中的旧数据          |
| 缓存穿透防护 | 搭配 LoaderFunc，防止重复查询数据库 |

---

## 八、优缺点总结

| 优点            | 缺点               |
| ------------- | ---------------- |
| 支持多种淘汰策略，功能强大 | 不支持分布式部署（仅本地）    |
| 支持并发、安全性好     | 不支持持久化           |
| 自动加载和淘汰回调实用   | 最大容量控制基于数量而非内存大小 |

---

## 九、与其他缓存库对比

| 库名            | 特性                     | 适合场景         |
| ------------- | ---------------------- | ------------ |
| **GCache**    | 功能全面，支持 LRU/LFU/ARC    | 通用型          |
| **BigCache**  | 高并发性能优异，无锁设计           | 高频缓存（比如日志分析） |
| **freecache** | 内存分配优化，性能高             | 只需要简单 KV 缓存  |
| **ristretto** | 支持 cost-based eviction | 高级权重缓存场景     |

---

如果你需要实现一个**高并发本地缓存系统**，并希望控制**过期时间、自动加载、淘汰回调**等功能，**GCache 是非常合适的选择**。

---

是否需要我对 GCache 的 LoaderFunc 或 EvictedFunc 实现进行源码级讲解？
