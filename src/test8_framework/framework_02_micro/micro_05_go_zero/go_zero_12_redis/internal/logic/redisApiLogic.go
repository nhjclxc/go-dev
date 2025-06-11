package logic

import (
	"context"
	"fmt"

	"go_zero_12_redis/internal/svc"
	"go_zero_12_redis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewRedisApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisApiLogic {
	return &RedisApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisApiLogic) RedisApi(req *types.RedisApiReq) (resp *types.RedisApiResp, err error) {
	// todo: add your logic here and delete this line

	redisCache := l.svcCtx.RedisCache

	// 设置 key，设置过期时间为60秒
	err1 := redisCache.Set("my-key", "my-value", 60)
	if err1 != nil {
		logx.Info("Setex.error = ", err1)
	}

	// 获取 key
	val, err2 := redisCache.Get("my-key")
	if err2 != nil {
		logx.Info("Get.error = ", err2)
	}
	fmt.Println("value:", val)

	// 删除 key
	_, err3 := redisCache.Delete("my-key")
	if err3 != nil {
		logx.Info("Del.error = ", err3)
	}

	return
}
