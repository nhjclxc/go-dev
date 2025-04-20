很棒的问题！在 Go 中使用 **Redis + SetNX 命令** 和 **Lua 脚本** 实现分布式锁是非常常见的方案。

我来一步步帮你整理两种方式：

---

## 🧩 一、使用 `SetNX` 实现分布式锁

### ✅ 实现思路
- 使用 `SETNX`（Go 中对应 `SetNX()`）来加锁；
- 加锁时设置一个**唯一标识符**（如 UUID）；
- 设置锁的**过期时间**（防止死锁）；
- 释放锁时先校验再删除（必须确保是自己加的锁）。

### 📦 依赖库

```bash
go get github.com/redis/go-redis/v9
go get github.com/google/uuid
```

### 🔧 实现代码（SetNX 版）

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/google/uuid"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// 加锁
func TryLock(key string, ttl time.Duration) (string, bool) {
    token := uuid.New().String()
    success, err := rdb.SetNX(ctx, key, token, ttl).Result()
    if err != nil {
        return "", false
    }
    if success {
        return token, true
    }
    return "", false
}

// 释放锁
func Unlock(key, token string) bool {
    val, err := rdb.Get(ctx, key).Result()
    if err != nil {
        return false
    }
    if val == token {
        _, err := rdb.Del(ctx, key).Result()
        return err == nil
    }
    return false
}

func main() {
    key := "my-lock"
    ttl := 5 * time.Second

    token, ok := TryLock(key, ttl)
    if !ok {
        fmt.Println("Failed to acquire lock")
        return
    }

    fmt.Println("Lock acquired:", token)

    // 做一些事情
    time.Sleep(2 * time.Second)

    if Unlock(key, token) {
        fmt.Println("Lock released")
    } else {
        fmt.Println("Failed to release lock")
    }
}
```

---

## 🧠 二、使用 Lua 脚本释放锁（避免非原子性）

因为 **Get + Del** 不是原子操作，可能导致误删别人的锁，因此使用 **Lua 脚本**来原子校验 + 删除。

### 🔐 Lua 脚本释放锁

```lua
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
```

### 💡 Go 实现释放锁（Lua 脚本版）

```go
var unlockScript = redis.NewScript(`
    if redis.call("GET", KEYS[1]) == ARGV[1] then
        return redis.call("DEL", KEYS[1])
    else
        return 0
    end
`)

func UnlockLua(key, token string) bool {
    result, err := unlockScript.Run(ctx, rdb, []string{key}, token).Result()
    if err != nil {
        return false
    }
    return result.(int64) == 1
}
```

在 `main()` 函数中替换为：

```go
if UnlockLua(key, token) {
    fmt.Println("Lock released (Lua)")
} else {
    fmt.Println("Failed to release lock (Lua)")
}
```

---

## ✅ 总结对比

| 方式 | 是否原子 | 安全性 | 推荐 |
|------|----------|--------|------|
| SetNX + Get + Del | 非原子 | 存在误删风险 | ❌ |
| SetNX + Lua 解锁 | 原子 | 更安全 | ✅ 推荐 |

---

## 🧰 进阶推荐
你也可以使用更高级的实现，比如：
- 使用 **RedLock 算法**（多节点 Redis）
- 用现成的库如 [`bsm/redislock`](https://github.com/bsm/redislock)

如果你想，我可以再帮你讲讲 RedLock 算法，或是封装一个更完整的分布式锁工具类，要不要来一发？