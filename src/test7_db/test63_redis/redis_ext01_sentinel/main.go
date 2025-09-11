package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisConfig struct {
	Host     []string `yaml:"host"`
	Password string   `yaml:"password"`
	DB       int      `yaml:"db"`
}

// redis哨兵模式的使用
func main() {

	ctx := context.Background()

	redisConfig := RedisConfig{
		Host:     []string{"127.0.0.1:26379"}, // 注意这个端口是哨兵的端口，不是节点的端口
		Password: "",
		DB:       0,
	}
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: redisConfig.Host,
		Password:      redisConfig.Password,
		DB:            redisConfig.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Println("redis client ping error", "error", err)
	} else {
		fmt.Println("redis client init success")
	}

	rdb.Set(ctx, "testkey", "12345678", 3*time.Second)
	fmt.Println("get ", rdb.Get(ctx, "testkey"))

	time.Sleep(5 * time.Second)
	fmt.Println("get ", rdb.Get(ctx, "testkey"))

}

/*

redis安装：https://blog.csdn.net/realize_dream/article/details/106227622

// 1、安装redis
> brew install redis
```
To restart redis after an upgrade:
  brew services restart redis
Or, if you don't want/need a background service you can just run:
  /opt/homebrew/opt/redis/bin/redis-server /opt/homebrew/etc/redis.conf
```

// 2、修改配置
redis.conf
redis-sentinel.conf
【修改的详细配置解析可以参考redis_sentinel.md】

// 3、启动redis
// 启动主节点 Redis
redis-server /opt/homebrew/etc/redis.conf

// 启动 Sentinel
redis-sentinel /opt/homebrew/etc/redis-sentinel.conf

// 关闭 redis
// ps -ef | grep redis
// kill -9 78668
// kill -9 78671

*/
