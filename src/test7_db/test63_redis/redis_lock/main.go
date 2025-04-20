package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
)

// 创建redis客户端
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // Redis地址
	Password: "",               // 没有密码则留空
	DB:       0,                // 默认使用DB 0
})

// 加锁
func TryLock(key string, ttl time.Duration) (string, bool) {
	token := uuid.New().String()
	success, err := redisClient.SetNX(key, token, ttl).Result()
	if err != nil {
		return err.Error(), false
	}
	if success {
		return token, true
	}
	return "", false
}

// 释放锁
func Unlock(key, token string) bool {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		return false
	}
	if val == token {
		_, err := redisClient.Del(key).Result()
		return err == nil
	}
	return false
}

var unlockScript = redis.NewScript(
	`
    if redis.call("GET", KEYS[1]) == ARGV[1] then
        return redis.call("DEL", KEYS[1])
    else
        return 0
    end
	`)

// UnlockLua 因为 Get + Del 不是原子操作，可能导致误删别人的锁，因此使用 Lua 脚本来原子校验 + 删除。
func UnlockLua(key, token string) bool {
	result, err := unlockScript.Run(redisClient, []string{key}, token).Result()
	if err != nil {
		return false
	}
	return result.(int64) == 1
}

func main() {

	//// 创建redis客户端
	//redisClient := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379", // Redis地址
	//	Password: "",               // 没有密码则留空
	//	DB:       0,                // 默认使用DB 0
	//})
	//
	//// 延迟关闭
	//defer redisClient.Close()

	key := "my-lock"
	ttl := 5 * time.Second

	token, ok := TryLock(key, ttl)
	if !ok {
		fmt.Println("Failed to acquire lock")
		return
	}

	fmt.Println("Lock acquired:", token)

	// 做一些事情
	time.Sleep(2 * time.Second)

	//if Unlock(key, token) {
	//	fmt.Println("Lock released")
	//} else {
	//	fmt.Println("Failed to release lock")
	//}

	if UnlockLua(key, token) {
		fmt.Println("Lock released (Lua)")
	} else {
		fmt.Println("Failed to release lock (Lua)")
	}

	// 你也可以使用更高级的实现，比如：
	//	使用 RedLock 算法（多节点 Redis）
	//	用现成的库如 bsm/redislock [bsm/redislock](https://github.com/bsm/redislock)
}
