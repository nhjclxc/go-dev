
etcd内部接口
```shell
http://127.0.0.1:2379/health
http://127.0.0.1:2379/version
http://127.0.0.1:2379/metrics
```


启动一个etcd容器
```shell
docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --name etcd \
  quay.io/coreos/etcd:v3.5.15 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir=/etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380
```

测试
```shell
etcdctl put foo bar
etcdctl get foo
```

```shell
lxc20250729@lxc20250729deMacBook-Pro ~ % docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --name etcd2379 \
  quay.io/coreos/etcd:v3.5.15 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir=/etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380
6ffbbed4286ae7c93ce7c2900256980118999e243a1db13574e76858c605ac1e
lxc20250729@lxc20250729deMacBook-Pro ~ % 
lxc20250729@lxc20250729deMacBook-Pro ~ % docker ps -a
CONTAINER ID   IMAGE                                                  COMMAND                   CREATED         STATUS                      PORTS                                                             NAMES
6ffbbed4286a   quay.io/coreos/etcd:v3.5.15                            "/usr/local/bin/etcd…"   9 seconds ago   Up 9 seconds                0.0.0.0:2379-2380->2379-2380/tcp, [::]:2379-2380->2379-2380/tcp   etcd2379
lxc20250729@lxc20250729deMacBook-Pro ~ % 
lxc20250729@lxc20250729deMacBook-Pro ~ % etcdctl put foo bar
OK
lxc20250729@lxc20250729deMacBook-Pro ~ % etcdctl get foo
foo
bar
lxc20250729@lxc20250729deMacBook-Pro ~ % 

```