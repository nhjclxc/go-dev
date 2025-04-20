package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 字符串 String
func main02() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	key := "go-redis-key"

	// 设置键值对，0就是永不过期
	//redisClient.Set("go-key", "Hello World!!", 0)
	// 过期时间
	redisClient.Set(key+"1111", "Hello World!!", 30*time.Second)

	// key 重命名
	fmt.Println(redisClient.Get(key + "1111").Val())
	redisClient.Rename(key+"1111", key)
	fmt.Println(redisClient.Get(key).Val())

	// 查询过期时间
	fmt.Println(redisClient.TTL(key))
	fmt.Println(redisClient.PTTL(key))
	fmt.Println("---------------------------")

	// 读取
	stringCmd := redisClient.Get(key)
	fmt.Println(stringCmd.Result())
	fmt.Println(stringCmd.String())
	fmt.Println(stringCmd.Name())
	fmt.Println(stringCmd.Args())
	fmt.Println(stringCmd.Val())
	fmt.Println(stringCmd.Err())
	fmt.Println("---------------------------")

	// 查询某个key的数据类型
	fmt.Println(redisClient.Type(key))

	// 扫描
	fmt.Println(redisClient.Scan(0, "go-", 4))

	// 批量存取
	redisClient.MSet("cookie", "12345", "token", "abcefg")
	fmt.Println(redisClient.MGet("cookie", "token").Val())

	// 数字增减，一般用于计数器的使用，如登录限制等等
	redisClient.Set("age", "1", 0)
	// 自增
	redisClient.Incr("age")
	redisClient.Incr("age")
	redisClient.Incr("age")
	redisClient.Incr("age")
	redisClient.Incr("age")
	fmt.Println(redisClient.Get("age").Val())
	// 自减
	redisClient.Decr("age")
	fmt.Println(redisClient.Get("age").Val())

	// 删除
	intCmd := redisClient.Del(key)
	fmt.Println(intCmd.Result())
	fmt.Println(intCmd.String())
	fmt.Println(intCmd.Name())
	fmt.Println(intCmd.Args())
	fmt.Println(intCmd.Val())
	fmt.Println(intCmd.Err())
	fmt.Println("---------------------------")

	// 再次读取
	fmt.Println(redisClient.Get(key).Result())

}
