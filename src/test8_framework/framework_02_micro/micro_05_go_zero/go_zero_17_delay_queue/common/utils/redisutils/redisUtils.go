package redisutils

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

// 来源：src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_12_redis/common/utils/redisutils/redisUtils.go

// RedisCache 封装 go-zero 的 redis.Redis 实例
type RedisCache struct {
	Client         *redis.Redis
	Ctx            context.Context
	RedisInitOnce  sync.Once // 用于确保只初始化一次
	RedisInitError error     // 保存初始化时的错误
	RedisConf redis.RedisConf
}

// NewRedisCache 创建 Redis 客户端封装
func NewRedisCache(redisConf redis.RedisConf) *RedisCache {
	r := RedisCache{}
	r.RedisConf = redisConf

	// todo GetRedisClient 懒加载 Redis 客户端 
	r.Client, r.RedisInitError = redis.NewRedis(r.RedisConf)
	if r.RedisInitError != nil {
		logx.Errorf("Redis 初始化失败: %v", r.RedisInitError)
	}

	return &r
}



// ====== String 操作 ======

// Set 设置 key 值（可选过期时间）
func (r *RedisCache) Set(key string, value string, expireSeconds int) error {
	if expireSeconds > 0 {
		return r.Client.Setex(key, value, expireSeconds)
	}
	return r.Client.Set(key, value)
}

// Get 获取字符串值
func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(key)
}

// Exists 判断 key 是否存在
func (r *RedisCache) Exists(key string) (bool, error) {
	return r.Client.Exists(key)
}

// Delete 删除一个或多个 key
func (r *RedisCache) Delete(keys ...string) (int, error) {
	return r.Client.Del(keys...)
}

// ====== Hash 操作 ======

// HSet 设置哈希字段值
func (r *RedisCache) HSet(key, field, value string) error {
	return r.Client.HsetCtx(r.Ctx, key, field, value)
}

// HGet 获取哈希字段值
func (r *RedisCache) HGet(key, field string) (string, error) {
	return r.Client.HgetCtx(r.Ctx, key, field)
}

// HGetAll 获取整个哈希
func (r *RedisCache) HGetAll(key string) (map[string]string, error) {
	return r.Client.HgetallCtx(r.Ctx, key)
}

// HDel 删除哈希字段
func (r *RedisCache) HDel(key string, fields ...string) (bool, error) {
	//args := append([]interface{}{key}, stringSliceToInterfaceSlice(fields)...)
	return r.Client.HdelCtx(r.Ctx, key, fields...)
}

// ====== List 操作 ======

// LPush 左插入
func (r *RedisCache) LPush(key string, values ...string) (int, error) {
	//args := append([]interface{}{key}, stringSliceToInterfaceSlice(values)...)
	return r.Client.LpushCtx(r.Ctx, key, values)
}

// RPush 右插入
func (r *RedisCache) RPush(key string, values ...string) (int, error) {
	//args := append([]interface{}{key}, stringSliceToInterfaceSlice(values)...)
	return r.Client.RpushCtx(r.Ctx, key, values)
}

// LRange 获取列表范围
func (r *RedisCache) LRange(key string, start, stop int) ([]string, error) {
	return r.Client.LrangeCtx(r.Ctx, key, start, stop)
}

// ====== Set 集合操作 ======

// SAdd 添加元素到集合
func (r *RedisCache) SAdd(key string, members ...string) (int, error) {
	//args := append([]interface{}{key}, stringSliceToInterfaceSlice(members)...)
	return r.Client.SaddCtx(r.Ctx, key, members)
}

// SMembers 获取集合所有元素
func (r *RedisCache) SMembers(key string) ([]string, error) {
	return r.Client.SmembersCtx(r.Ctx, key)
}

// ====== 分布式锁 ======

// TryLock 使用 SETNX 实现分布式锁（单位：秒）
func (r *RedisCache) TryLock(key, value string, expireSeconds int) (bool, error) {
	ok, err := r.Client.SetnxEx(key, value, expireSeconds)
	return ok, err
}

// Unlock 解锁（建议加锁时记录 value，校验解锁身份）
func (r *RedisCache) Unlock(key, expectedValue string) (bool, error) {
	lua := `
			if redis.call("get", KEYS[1]) == ARGV[1]
			then
				return redis.call("del", KEYS[1])
			else
				return 0
			end
			`
	reply, err := r.Client.Eval(lua, []string{key}, []string{expectedValue})
	if err != nil {
		return false, err
	}
	return reply.(int64) == 1, nil
}

//// ====== 辅助工具 ======
//
//func stringSliceToInterfaceSlice(s []string) []interface{} {
//	result := make([]interface{}, len(s))
//	for i, v := range s {
//		result[i] = v
//	}
//	return result
//}
