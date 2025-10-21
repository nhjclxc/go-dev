
# mac 安装 mongodb

## 安装
MongoDB 社区版在 Homebrew 中使用 mongodb-community：
```shell
brew tap mongodb/brew
brew install mongodb-community@7.0
```
这里的 @7.0 可以替换为你想安装的版本，比如 @6.0。

## 配置文件

创建基本的路径
```shell
sudo mkdir -p /usr/local/var/log/mongodb
sudo chown -R $(whoami) /usr/local/var/log/mongodb
sudo mkdir -p /usr/local/var/mongodb
sudo chown -R $(whoami) /usr/local/var/mongodb
sudo mkdir -p /usr/local/etc/
```

### 常用的基础的配置文件

`sudo vim /usr/local/etc/mongod.conf`
```mongod.conf
# ============================================
# MongoDB 基础配置文件
# 路径: /usr/local/etc/mongod.conf
# ============================================

# 数据存储路径
storage:
  dbPath: /usr/local/var/mongodb

# 日志配置
systemLog:
  destination: file
  path: /usr/local/var/log/mongodb/mongo.log
  logAppend: true

# 网络配置
net:
  port: 27017
  bindIp: 127.0.0.1   # 仅允许本机访问 (推荐开发时保持安全)
  # bindIp: 0.0.0.0   # 如需远程访问，取消注释，并开启认证

# 安全配置（默认关闭认证，方便开发） enabled/disabled
security:
  authorization: disabled

```


### 完整配置文件
```mongod.conf
# ============================================
# MongoDB 配置文件 (mongod.conf)
# 适用于 macOS Homebrew 安装
# 路径：/usr/local/etc/mongod.conf
# ============================================

# -------------------------------
# 数据存储相关配置
# -------------------------------
storage:
  dbPath: /usr/local/var/mongodb      # 数据存储目录
  engine: wiredTiger                  # 存储引擎 (默认 wiredTiger)

# -------------------------------
# 日志相关配置
# -------------------------------
systemLog:
  destination: file                   # 日志输出目标: file | syslog | stdout
  path: /usr/local/var/log/mongodb/mongo.log   # 日志文件路径
  logAppend: true                     # 是否追加写入日志 (true=追加，false=覆盖)
  verbosity: 0                        # 日志详细级别 (0=默认，1-5更详细)

# -------------------------------
# 网络相关配置
# -------------------------------
net:
  port: 27017                         # MongoDB 服务端口
  bindIp: 127.0.0.1                   # 绑定 IP 地址 (127.0.0.1=仅本地访问)
  # bindIp: 0.0.0.0                   # 如需远程访问，可改为 0.0.0.0 (不安全，需配合认证)

# -------------------------------
# 安全配置
# -------------------------------
security:
  authorization: disabled             # 是否启用用户认证
  # authorization: enabled            # 开启后需创建用户才能访问数据库

# -------------------------------
# 进程管理
# -------------------------------
processManagement:
  fork: false                         # 是否以守护进程模式运行 (Linux 常用，macOS/Homebrew 不建议)
  pidFilePath: /usr/local/var/run/mongodb.pid  # PID 文件路径

# -------------------------------
# 副本集配置 (可选)
# -------------------------------
# replication:
#   replSetName: rs0                  # 副本集名称，集群部署时使用

# -------------------------------
# 分片配置 (可选)
# -------------------------------
# sharding:
#   clusterRole: shardsvr             # 配置节点角色: shardsvr | configsvr

# -------------------------------
# 限制配置 (可选)
# -------------------------------
# setParameter:
#   enableLocalhostAuthBypass: true   # 本地连接是否允许免认证 (仅限初始配置阶段)

```


## 启动

临时启动（当前终端关闭即停止）：
`mongod --config /usr/local/etc/mongod.conf`

后台启动（开机自启）：
`brew services start mongodb/brew/mongodb-community`

停止 MongoDB：
`brew services stop mongodb/brew/mongodb-community`

## 检查是否启动成功

检查mongosh的位置：`which mongosh`

查看mogodb版本：`mongosh --version`

连接mogodb：`mongosh` 或 `mongodb://127.0.0.1:27017`，或者连接其他实例：`mongosh "mongodb://username:password@localhost:27017"`
