package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// Scan 用法
// Scan 是 Go 语言 Redis 客户端中常用的方法之一，尤其在使用 go-redis 时（例如 github.com/redis/go-redis/v9），
// 用于遍历 Redis 中的键（keys），非常适合替代 KEYS 命令来避免阻塞。
func main03() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	// 先初始化一些key进去，以便下面的使用
	for i := 0; i < 10; i++ {
		redisClient.Set("go-anonymous_user:"+strconv.Itoa(i), i, 60*time.Second)
	}

	var cursor uint64
	var keys []string
	var err error

	// 你可以自定义匹配模式和一次扫描的数量
	matchPattern := "go-anonymous_user:*"
	count := int64(20) // 表示扫描20个匹配到的key

	for {
		// 使用 Scan 遍历键
		var scannedKeys []string
		//🔍 参数说明
		//		cursor: 游标，初始值为 0，Redis 会返回新的游标直到返回 0 表示遍历结束。
		//		match: 匹配模式（支持通配符，比如 go-anonymous_user:*）。
		//		count: 每次扫描返回的最大数量（只是 hint，实际可能少于这个数）。
		scannedKeys, cursor, err = redisClient.Scan(cursor, matchPattern, count).Result()
		if err != nil {
			panic(err)
		}

		keys = append(keys, scannedKeys...)

		if cursor == 0 {
			break
		}
	}

	for _, key := range keys {
		fmt.Println("Found key:", key)
	}

}
