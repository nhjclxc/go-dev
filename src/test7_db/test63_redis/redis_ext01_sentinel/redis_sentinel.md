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






好的，我们按你提供的实例和代码仔细分析：

---

### 1️⃣ 你的部署情况

* **Redis 实例**：

  ```
  node1:6379
  node2:6379
  node3:6379
  ```

  假设主从关系为：

  ```
  node1(master)  -->  node2(slave)
                 -->  node3(slave)
  ```

* **Sentinel**：

  ```
  node1:26379
  node2:26379
  node3:26379
  ```

  监控 master 为 `mymaster`。

* **连接代码**：

  ```go
  r.Redis = redis.NewFailoverClient(&redis.FailoverOptions{
      MasterName:    "mymaster",
      SentinelAddrs: config.Common.Redis.Host,
      Password:      config.Common.Redis.Password,
      DB:            config.Common.Redis.DB,
  })
  ```

---

### 2️⃣ `redis.NewFailoverClient` 的行为

`NewFailoverClient` 是 **Go-Redis 官方客户端提供的哨兵高可用连接方式**，作用如下：

1. **连接哨兵列表** (`SentinelAddrs`)

    * 客户端会向这些哨兵询问当前的 master 节点地址。
2. **自动发现 master**

    * 哨兵返回 `(masterIP, masterPort)`。
    * 客户端直接连接到 master 节点进行读写操作。
3. **自动故障转移处理**

    * 如果 master 挂了，哨兵会选举新的 master。
    * 客户端会自动更新 master 地址，继续读写。

---

### 3️⃣ 数据写入的位置

* **所有写操作**都会写到 **当前 master 节点**（比如 `node1:6379`）。
* **slave 节点**会异步复制 master 的数据。

  ```
  master(node1)
       |
       | async replication
       v
  slave(node2)
       |
       v
  slave(node3)
  ```
* 也就是说：

    * 写入总是去 master
    * 数据同步到 slave 是异步的，存在微小延迟
    * 读取默认也是从 master，如果你配置了 `ReadOnly` 可以从 slave 读取

---

### 4️⃣ 三个实例的数据是否同步

* **是的，但异步**

    * master 写入 → slave 自动复制
    * 如果 master 写入刚刚完成，从节点可能暂时还没有最新数据

---

### ✅ 总结

| 连接方式                     | 数据写入位置    | 数据同步情况                                   |
| ------------------------ | --------- | ---------------------------------------- |
| 直接连接某个实例（6379）           | 连接到的节点    | slave 如果有会复制 master 数据，但客户端不会自动发现 master |
| `NewFailoverClient` + 哨兵 | 当前 master | master 写入 → slave 异步复制                   |
| Sentinel (26379)         | 不存数据      | 只提供 master 地址，不存数据                       |

> 核心：使用 `NewFailoverClient` 时，所有写入都是写到 master 的，slave 自动同步数据，客户端不用关心具体 master 节点是哪一台。

---

如果你愿意，我可以画一个 **带哨兵的 Redis 主从写入和同步示意图**，让数据流一目了然。

你希望我画吗？




好的，我来帮你画一个 **带哨兵的 Redis 主从架构 + 写入同步示意图**，让数据流和访问逻辑直观可见。

---

```
                 ┌─────────────┐
                 │  Sentinel1  │ 26379
                 └─────┬───────┘
                       │ monitors
                 ┌─────┴───────┐
                 │  Sentinel2  │ 26379
                 └─────┬───────┘
                       │ monitors
                 ┌─────┴───────┐
                 │  Sentinel3  │ 26379
                 └─────┬───────┘
                       │ monitors
                       ▼
                  ┌─────────┐
                  │ Master  │ node1:6379
                  │ (write) │
                  └────┬────┘
                       │ async replication
        ┌──────────────┴───────────────┐
        ▼                              ▼
   ┌─────────┐                     ┌─────────┐
   │  Slave  │ node2:6379           │  Slave  │ node3:6379
   │  (read) │                     │  (read) │
   └─────────┘                     └─────────┘

```

### 数据流说明

1. **写入**：

    * 客户端（通过 `NewFailoverClient` 连接哨兵）
    * 自动查询当前 master
    * 写入直接发送到 **master(node1)**

2. **同步**：

    * master 数据 **异步复制**到 slave(node2) 和 slave(node3)
    * slave 上可以读取数据（延迟可能存在）

3. **哨兵作用**：

    * 监控 master 和 slave 状态
    * 如果 master 挂掉，选举一个新的 master
    * 客户端自动切换到新 master

4. **读取**：

    * 默认从 master 读取（保证强一致性）
    * 配置 `ReadOnly` 可从 slave 读取，减轻 master 压力

---

💡 总结：

* **写入总是 master**
* **slave 异步同步**
* **哨兵不存数据，只负责监控和提供 master 地址**

---

如果你愿意，我可以帮你画一个 **更直观的带箭头的动态数据流图**，标出写入和复制路径，以及客户端访问流程。

你希望我画吗？
