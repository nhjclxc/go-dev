package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// go-redis 自动使用连接池，你只需要设置这些参数即可，不需要手动管理连接池或连接释放。
	// 创建 Redis 客户端，内部自动使用连接池
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379", // Redis地址
		Password:     "",               // Redis无密码可留空
		DB:           0,                // 使用默认DB
		PoolSize:     20,               // 最大连接数（连接池大小）
		MinIdleConns: 5,                // 最小空闲连接
		PoolTimeout:  30 * time.Second, // 连接池等待超时时间
	})

	// 确保连接成功
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis连接成功！")

	// 简单 set/get 示例
	err = client.Set("mykey", "hello redis", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("mykey").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("获取到的值是：", val)

	// 关闭连接池
	defer client.Close()
}
