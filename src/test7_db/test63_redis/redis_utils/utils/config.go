package utils

//
//import (
//	"github.com/go-redis/redis"
//)
//
//// 声明redis客户端
//var redisClient *redis.Client
//
//// 初始化包配置
//func init() {
//	redisClient = redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379", // 连接地址
//		Password: "",               // 密码
//		DB:       0,                // 数据库编号
//		//DialTimeout: 3 * time.Second,  // 链接超时
//	})
//}
//
//// GetRedisClient 获取 redisClient 指针
//func GetRedisClient() *redis.Client {
//	return redisClient
//}
