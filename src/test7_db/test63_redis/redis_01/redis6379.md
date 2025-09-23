
```conf
# 监听端口
port 6379

# 是否在后台运行
daemonize yes

# 数据文件存放路径
dir /tmp

# 数据库文件名称
dbfilename dump.rdb

# 开启持久化（RDB）
save 900 1
save 300 10
save 60 10000

```


```shell
redis-server /opt/homebrew/etc/redis6379.conf
```

lxc20250729@lxc20250729deMacBook-Pro etc % ps -ef | grep redis           
501 24206  6793   0 11:07上午 ttys000    0:00.01 grep redis
lxc20250729@lxc20250729deMacBook-Pro etc % redis-server /opt/homebrew/etc/redis6379.conf
lxc20250729@lxc20250729deMacBook-Pro etc % ps -ef | grep redis                          
501 24208     1   0 11:07上午 ??         0:00.01 redis-server *:6379
501 24210  6793   0 11:07上午 ttys000    0:00.01 grep redis