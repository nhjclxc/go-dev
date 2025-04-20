package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 使用 ZSet（有序集合） 做排行榜（排名）
func main() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	// 1、先添加一点数据
	redisClient.ZAdd("go-rank", []redis.Z{
		{Score: 1500, Member: "user_1"},
		{Score: 1200, Member: "user_2"},
		{Score: 900, Member: "user_3"},
	}...)
	redisClient.ZAdd("go-rank",
		redis.Z{Score: 800, Member: "user_4"},
		redis.Z{Score: 500, Member: "user_5"},
	)

	// 2. 获取排行榜前 N 名（Score降序）
	// ZRevRangeWithScores：从高分到低分排序（降序）
	// ZRangeWithScores：升序
	topN := 10
	res, err := redisClient.ZRevRangeWithScores("go-rank", 0, int64(topN-1)).Result()
	if err != nil {
		panic(err)
	}
	for i, z := range res {
		fmt.Printf("Rank %d: %v (Score: %.0f)\n", i+1, z.Member, z.Score)
	}

	// 3. 获取某个用户的排名
	// 注意：索引从0开始
	rank, err := redisClient.ZRevRank("go-rank", "user_3").Result()
	if err == redis.Nil {
		fmt.Println("用户不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("user_3 的排名是：%d \n", rank)
	}

	// 4. 获取某个用户的分数
	score, err := redisClient.ZScore("go-rank", "user_3").Result()
	if err == redis.Nil {
		fmt.Println("用户不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("user_3 的分数是：%.0f\n", score)
	}

	// 5. 增加/减少某个用户的分数
	redisClient.ZIncrBy("go-rank", 500, "user_3")
	fmt.Println(redisClient.ZRevRank("go-rank", "user_3").Val())

	//清除排行榜（调试时用）：
	redisClient.Del("go-rank")

	// 限制排行榜只保留前100名：
	redisClient.ZRemRangeByRank("go-rank", 100, -1)
}
