package main

import (
	"fmt"
	redisUtils "go-dev/src/test6_db/test63_redis/redis_utils/utils"
	"time"
)

// https://blog.csdn.net/qq_44237719/article/details/128920821
func main() {
	redisUtils.Set("test", "哈哈哈")
	redisUtils.GetSet("test", "嘿嘿嘿")
	redisUtils.Get("test")
	//
	redisUtils.Expire("test", 5*time.Second)
	//
	fmt.Println(redisUtils.LRem("list-1", 1, 3))

	redisUtils.HSet("stu-1", "username", "qqq")
	//redisUtils.HSet("stu-1", "pwd", "111")
	//redisUtils.HDel("stu-1", "username")

	name := redisUtils.HGetAll("stu-1")
	fmt.Println(name)

}
