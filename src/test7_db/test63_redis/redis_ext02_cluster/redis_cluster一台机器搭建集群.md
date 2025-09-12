 # 自己搭建过程
 
[Mac下载redis](https://blog.csdn.net/realize_dream/article/details/106227622)

1. 创建每一个集群节点的文件夹
```shell
mkdir -p /opt/homebrew/etc/redis-cluster/7001 /opt/homebrew/etc/redis-cluster/7002 /opt/homebrew/etc/redis-cluster/7003 
mkdir -p /opt/homebrew/etc/redis-cluster/7004 /opt/homebrew/etc/redis-cluster/7005 /opt/homebrew/etc/redis-cluster/7006
```
Redis Cluster 必须覆盖所有 16384 个哈希槽。至少需要 3 个主节点 3 个从节点 才能构成一个“健壮的生产集群”

2. 每一个节点的配置，其他节点只需将下面的7001改为7002，7003，7004，7005，7006即可
```redis-7001.conf
port 7001
bind 127.0.0.1
daemonize yes
protected-mode no
dir /opt/homebrew/etc/redis-cluster/7001
dbfilename dump.rdb
cluster-enabled yes
cluster-config-file nodes-7001.conf
cluster-node-timeout 5000
appendonly yes
```

3. 启动每一个节点
```shell
redis-server /opt/homebrew/etc/redis-cluster/7001/redis-7001.conf
redis-server /opt/homebrew/etc/redis-cluster/7002/redis-7002.conf
redis-server /opt/homebrew/etc/redis-cluster/7003/redis-7003.conf
redis-server /opt/homebrew/etc/redis-cluster/7004/redis-7004.conf
redis-server /opt/homebrew/etc/redis-cluster/7005/redis-7005.conf
redis-server /opt/homebrew/etc/redis-cluster/7006/redis-7006.conf
```

查看是否启动
```shell
ps -ef | grep redis
```

4. 创建集群【注⚠️：上面的第三点只是启动了集群的所有节点，并没有把所有节点组成集群】
```shell
redis-cli --cluster create \
  127.0.0.1:7001 \
  127.0.0.1:7002 \
  127.0.0.1:7003 \
  127.0.0.1:7004 \
  127.0.0.1:7005 \
  127.0.0.1:7006 \
  --cluster-replicas 1
```
- `--cluster-replicas 1` 表示每个主节点分配一个从节点。
- 这样会得到一个 3 主 3 从的集群。
- 这一步会把 16384 个哈希槽 分配给 3 个主节点。

```shell
lxc20250729@lxc20250729deMacBook-Pro 7006 % redis-cli --cluster create \   
  127.0.0.1:7001 \
  127.0.0.1:7002 \
  127.0.0.1:7003 \
  127.0.0.1:7004 \
  127.0.0.1:7005 \
  127.0.0.1:7006 \
  --cluster-replicas 1

>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 127.0.0.1:7005 to 127.0.0.1:7001
Adding replica 127.0.0.1:7006 to 127.0.0.1:7002
Adding replica 127.0.0.1:7004 to 127.0.0.1:7003
>>> Trying to optimize slaves allocation for anti-affinity
[WARNING] Some slaves are in the same host as their master
M: 3ce5eed27387abfb5828942857c42406aac8b7e2 127.0.0.1:7001
   slots:[0-5460] (5461 slots) master
M: db99dbeb63f8dfb5ac6061e9b5bb082950af594c 127.0.0.1:7002
   slots:[5461-10922] (5462 slots) master
M: 2acac286f4859698a3526dabb1ce7f0d1fd40ea3 127.0.0.1:7003
   slots:[10923-16383] (5461 slots) master
S: d9b48fde87c8530ed307f4aba8b1ab2d1f96b74b 127.0.0.1:7004
   replicates 3ce5eed27387abfb5828942857c42406aac8b7e2
S: e775c2aea6662c9ff1636aa12b58c19ebfc6da3d 127.0.0.1:7005
   replicates db99dbeb63f8dfb5ac6061e9b5bb082950af594c
S: 04651b3657692a9383cd93c0ac61509254c391ae 127.0.0.1:7006
   replicates 2acac286f4859698a3526dabb1ce7f0d1fd40ea3
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
.
>>> Performing Cluster Check (using node 127.0.0.1:7001)
M: 3ce5eed27387abfb5828942857c42406aac8b7e2 127.0.0.1:7001
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: d9b48fde87c8530ed307f4aba8b1ab2d1f96b74b 127.0.0.1:7004
   slots: (0 slots) slave
   replicates 3ce5eed27387abfb5828942857c42406aac8b7e2
M: db99dbeb63f8dfb5ac6061e9b5bb082950af594c 127.0.0.1:7002
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: e775c2aea6662c9ff1636aa12b58c19ebfc6da3d 127.0.0.1:7005
   slots: (0 slots) slave
   replicates db99dbeb63f8dfb5ac6061e9b5bb082950af594c
S: 04651b3657692a9383cd93c0ac61509254c391ae 127.0.0.1:7006
   slots: (0 slots) slave
   replicates 2acac286f4859698a3526dabb1ce7f0d1fd40ea3
M: 2acac286f4859698a3526dabb1ce7f0d1fd40ea3 127.0.0.1:7003
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

- 这一步会把 16384 个哈希槽 分配给 3 个主节点。

5. 验证集群状态
```shell
redis-cli -p 7001 cluster info
# 输出中 cluster_state:ok 表示集群正常

redis-cli -p 7001 cluster nodes
# 输出每个节点信息，包括 master/slave、slot范围、复制关系
```

5. 使用go访问集群
```go
	ctx := context.Background()

	// 生产环境必须要有三个以上节点才能使用 NewClusterClient
	redisConfig := RedisConfig{
		Host:     []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}, // 写 7001、7002、7003 就够了，7004、7005、7006 会自动识别。
		Password: "",
		DB:       0,
	}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redisConfig.Host,
		Password: redisConfig.Password, // 如果集群开启密码
	})

```