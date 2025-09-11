你这段代码使用了 **Redis Cluster 模式** (`NewClusterClient`)。要让 Redis 能被这个客户端连接，Redis 必须 **开启集群模式**，而不是普通单机模式。下面我给你梳理如何修改 `redis.conf`。

---

## 1️⃣ Redis 集群节点配置（redis.conf）

假设你要启动单个节点进行测试（生产环境至少 3 个主节点 + 从节点）：

```conf
# 基本配置
port 6379
bind 0.0.0.0           # 或者本机 IP
daemonize yes
protected-mode no      # 测试环境可关闭，生产环境用防火墙
requirepass redis-password  # 如果客户端要密码

# 数据目录
dir ./data
dbfilename dump.rdb

# 集群配置
cluster-enabled yes                 # 启用集群模式
cluster-config-file nodes.conf      # 集群节点信息文件
cluster-node-timeout 5000           # 节点超时时间（ms）
```

> 注意：
>
> * `requirepass` 要和 `ClusterOptions.Password` 一致
> * `cluster-config-file` 会自动创建和维护集群节点信息

---

## 2️⃣ 初始化 Redis Cluster

Redis 集群不是单节点就能用的，需要至少 **3 个主节点**：

```bash
# 假设你有 3 个节点 6379/6380/6381
redis-server ./redis-6379.conf
redis-server ./redis-6380.conf
redis-server ./redis-6381.conf
```

然后用 `redis-cli` 创建集群：

```bash
redis-cli --cluster create 127.0.0.1:6379 127.0.0.1:6380 127.0.0.1:6381 --cluster-replicas 0
```

> ⚠️ 生产环境至少 3 主 3 从，`--cluster-replicas 1`

---

## 3️⃣ Go 客户端配置

```go
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs:    []string{"127.0.0.1:6379", "127.0.0.1:6380", "127.0.0.1:6381"},
    Password: "redis-password",
})

// 测试
err := rdb.Set(ctx, "foo", "bar", 0).Err()
if err != nil {
    panic(err)
}
```

---

### 🔹 关键点

1. **Redis 必须启用集群模式**：`cluster-enabled yes`
2. **至少 3 个主节点**（生产环境需从节点）
3. **密码一致**：客户端和 redis.conf 的 `requirepass` 对应
4. **客户端使用 `NewClusterClient`**，不能用单机 `NewClient`

---

如果你希望，我可以帮你画一张 **Redis Cluster + Go 客户端连接示意图**，让你更直观理解 hash slot、节点、客户端路由关系。

你想让我画吗？
