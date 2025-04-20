package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 集合 Set
func main06() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})

	// 延迟关闭
	defer redisClient.Close()

	// 往一个集合里面添加元素
	redisClient.SAdd("set", "a", "b", "c")
	redisClient.SAdd("set2", "c", "d", "e")

	// 获取集合中的所有成员
	fmt.Println(redisClient.SMembers("set"))
	// 判断一个元素是否属于这个集合
	fmt.Println(redisClient.SIsMember("set", "a"))
	// 随机返回count个元素
	fmt.Println(redisClient.SRandMemberN("set", 1))
	// 获取一个集合的元素个数
	fmt.Println(redisClient.SCard("set"))

	// 返回给定集合的差集
	fmt.Println(redisClient.SDiff("set", "set2"))
	// 将给定集合的差集保存在结果集里，返回结果集的长度
	fmt.Println(redisClient.SDiffStore("store", "set", "se2"))
	// 返回给定集合的交集
	fmt.Println(redisClient.SInter("set", "set2"))
	// 将给定集合的交集保存在结果集里，返回结果集的长度
	fmt.Println(redisClient.SInterStore("store", "set", "set2"))
	// 返回给定集合的并集
	fmt.Println(redisClient.SUnion("set", "set2"))
	// 将给定集合的并集保存在结果集里，返回结果集的长度
	fmt.Println(redisClient.SUnionStore("store", "set", "store"))

	// 弹出并删除该元素
	fmt.Println(redisClient.SPop("set"))
	// 弹出并删除N给元素
	fmt.Println(redisClient.SPopN("set", 2))
	// 删除指定元素
	fmt.Println(redisClient.SRem("set", "a", "b"))

	// 从源集合移动指定元素刀目标集合
	redisClient.SMove("set", "set2", "a")

	// 遍历集合
	fmt.Println(redisClient.SScan("set", 0, "", 2))
}
