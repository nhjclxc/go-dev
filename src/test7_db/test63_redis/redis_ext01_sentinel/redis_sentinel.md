# redis搭建哨兵模式以及golang连接


你提供的 Go 代码使用了 **Redis Sentinel 模式**（`NewFailoverClient` + `MasterName: "mymaster"` + `SentinelAddrs`），所以你的 Redis 服务器要 **同时配置 Redis 主节点和 Sentinel**，而不是普通单机 Redis 就行。下面我帮你梳理如何配置 `redis.conf` 和 `sentinel.conf` 来满足这个客户端连接。

---

## 1️⃣ 主节点 Redis 配置（`redis.conf`）

假设主节点 IP 是 `127.0.0.1`，端口默认 6379：

```conf
# redis.conf 主节点配置
bind 127.0.0.1
port 6379
daemonize yes           # 后台运行，可选
protected-mode yes      # 默认 yes
requirepass ""          # 如果客户端没密码可以空
dir ./                  # 数据目录
dbfilename dump.rdb     # RDB 文件名
```

> 注意：如果你只有一个 Redis 节点且不需要密码，可以空着 `requirepass`。

---

## 2️⃣ Sentinel 配置（`sentinel.conf`）

Redis Sentinel 用来监控主节点、发现主从，并在主节点宕机时做自动切换。你需要创建一个或多个 Sentinel 配置文件，例如 `sentinel-6379.conf`：

```conf
# sentinel-6379.conf
port 26379
daemonize yes
protected-mode yes

# 哨兵监控的 master 名称 + 主节点 IP + 主节点端口 + quorum
# quorum: 至少多少个 sentinel 同意主节点挂掉才 failover
sentinel monitor mymaster 127.0.0.1 6379 1

# 哨兵认证（可选，如果主节点设置了密码）
# sentinel auth-pass mymaster <redis-password>

# 自动故障转移超时时间
sentinel down-after-milliseconds mymaster 5000
sentinel failover-timeout mymaster 10000
sentinel parallel-syncs mymaster 1
```

> 这里 `MasterName` 要和 Go 客户端里的一致：`"mymaster"`
> `SentinelAddrs` 对应的是 Sentinel 的端口（这里是 26379），而不是 Redis 主节点 6379

---

## 3️⃣ 启动流程

1. 启动主节点 Redis：

```bash
redis-server /path/to/redis.conf
```

2. 启动 Sentinel：

```bash
redis-sentinel /path/to/sentinel-6379.conf
```

3. Go 客户端配置：

```go
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "mymaster",
    SentinelAddrs: []string{"127.0.0.1:26379"}, // 指哨兵端口
    Password:      "", // 如果主节点有密码就填
    DB:            0,
})
```

---

### ✅ 注意事项

* **端口**：Sentinel 默认 26379，Redis 主节点默认 6379。
* **MasterName**：客户端和 sentinel.conf 必须一致。
* **多哨兵**：生产环境建议至少配置 3 个 Sentinel 节点，提高可用性。
* **DB**：Sentinel 模式下支持数据库号，可以用 0 或其他。

---

如果你需要，我可以帮你画一张 **Redis + Sentinel 架构图**，让你直观理解客户端如何通过 `FailoverClient` 自动发现主节点。

你希望我画吗？
