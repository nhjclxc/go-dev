package logic

import (
	"context"
	"fmt"
	"time"

	"go_zero_18_delay_queue_redis/internal/svc"
	"go_zero_18_delay_queue_redis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelayQueueRedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDelayQueueRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelayQueueRedisLogic {
	return &DelayQueueRedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelayQueueRedisLogic) DelayQueueRedis(req *types.DelayQueueRedisReq) error {
	// todo: add your logic here and delete this line


	add, err := l.svcCtx.RedisCache.XAdd(context.Background(), "delay:stream","qwertyu123456", "qqq", time.Second*10)
	if err != nil {
		return err
	}

	fmt.Printf("add = %s \n", add)


	return nil
}
