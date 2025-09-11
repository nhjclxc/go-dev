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

	redisConfig := RedisConfig{
		Host:     []string{"127.0.0.1:26379"}, // 注意这个端口是哨兵的端口，不是节点的端口
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
