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

// redis集群的使用
func main() {

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

 */
