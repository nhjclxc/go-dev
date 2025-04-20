package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 哈希表 HSet
func main04() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	// 单个设置
	redisClient.HSet("map", "name", "jack")
	// 批量设置
	redisClient.HMSet("map", map[string]interface{}{"a": "b", "c": "d", "e": "f"})
	// 单个访问
	fmt.Println(redisClient.HGet("map", "a").Val())
	// 批量访问
	fmt.Println(redisClient.HMGet("map", "a", "b").Val())
	// 获取整个map
	fmt.Println(redisClient.HGetAll("map").Val())

	// 删除map的一个字段
	fmt.Println(redisClient.HDel("map", "a"))

	// 判断字段是否存在
	fmt.Println(redisClient.HExists("map", "a"))

	// 获取所有的map的键
	fmt.Println(redisClient.HKeys("map"))

	// 获取map长度
	fmt.Println(redisClient.HLen("map"))

	// 遍历map中的键值对
	fmt.Println(redisClient.HScan("map", 0, "", 1))

}
