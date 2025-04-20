package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 有序集合 ZSet
func main07() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	// 往有序集合中加入元素
	redisClient.ZAdd("zset", redis.Z{
		Score:  1,
		Member: "a",
	}, redis.Z{
		Score:  2,
		Member: "b",
	})

	// 返回有序集合中该元素的排名，从低到高排列
	fmt.Println(redisClient.ZRank("zset", "a").Val())
	fmt.Println(redisClient.ZRank("zset", "b").Val())
	fmt.Println(redisClient.ZRank("zset", "aaa").Val())
	// 返回有序集合中该元素的排名，从高到低排列
	fmt.Println(redisClient.ZRevRank("zset", "1"))

	// 返回介于min和max之间的成员数量
	fmt.Println(redisClient.ZCount("zset", "1", "2"))

	// 返回对元素的权值
	fmt.Println(redisClient.ZScore("zset", "a"))

	// 返回指定区间的元素
	fmt.Println(redisClient.ZRange("zset", 1, 2))

	// 返回介于min和max之间的所有成员列表
	redisClient.ZRangeByScore("zset", redis.ZRangeBy{
		Min:    "1",
		Max:    "2",
		Offset: 0,
		Count:  1,
	})

	// 给一个对应的元素增加相应的权值
	redisClient.ZIncr("ss", redis.Z{
		Score:  2,
		Member: "b",
	})

	// 删除指定元素
	redisClient.ZRem("ss", "a")
	// 删除指定排名区间的元素
	redisClient.ZRemRangeByRank("ss", 1, 2)
	// 删除权值在min和max区间的元素
	redisClient.ZRemRangeByScore("ss", "1", "2")

}
