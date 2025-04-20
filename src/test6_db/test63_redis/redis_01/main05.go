package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 列表 List
func main05() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	// 左边添加
	redisClient.LPush("list", "a", "b", "c", "d", "e")
	// 右边添加
	redisClient.RPush("list", "g", "i", "a")
	// 在参考值前面插入值
	redisClient.LInsertBefore("list", "a", "aa")
	// 在参考值后面插入值
	redisClient.LInsertAfter("list", "a", "gg")
	// 设置指定下标的元素的值
	redisClient.LSet("list", 0, "head")

	// 访问列表长度
	fmt.Println(redisClient.LLen("list"))

	// 左边弹出元素
	fmt.Println(redisClient.LPop("list"))
	// 右边弹出元素
	fmt.Println(redisClient.RPop("list"))
	// 访问指定下标的元素
	fmt.Println(redisClient.LIndex("list", 1))
	// 访问指定范围内的元素
	fmt.Println(redisClient.LRange("list", 0, 1))

	// 删除指定元素
	redisClient.LRem("list", 0, "a")
	// 删除指定范围的元素
	redisClient.LTrim("list", 0, 1)
	// 保留指定范围的元素
	redisClient.LTrim("list", 0, 1)

}
