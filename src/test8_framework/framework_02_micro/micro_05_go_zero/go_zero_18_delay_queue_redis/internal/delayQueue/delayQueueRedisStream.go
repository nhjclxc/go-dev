package delayQueue

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go_zero_18_delay_queue_redis/common/utils/redisutils"
	"go_zero_18_delay_queue_redis/internal/types"
	"strconv"
	"time"
)

const StreamKey = "delay:stream"

type DelayQueueRedisStream struct {
	RedisCache *redisutils.RedisCache // 私有 redis 实例
}

func NewDelayQueueRedisStream(redisCache *redisutils.RedisCache) (*DelayQueueRedisStream) {
	return &DelayQueueRedisStream{RedisCache: redisCache}
}

func (dq *DelayQueueRedisStream) AddTask(task types.DelayTask) error {

	add, err := dq.RedisCache.XAdd(context.Background(), "delay:stream","qwertyu123456", "qqq", time.Second*10)
	if err != nil {
		return err
	}

	fmt.Printf("add = %s \n", add)

	return err
}

func (dq *DelayQueueRedisStream) ConsumeTasks(ctx context.Context) {
	group := "delayGroup"
	consumer := "consumer-1"
	key := "delay:stream"

	go func() {

		// 必须要先初始化
		err := dq.RedisCache.XGroupCreateMkStream(context.Background(), key, group, consumer)
		if err != nil {
			if "BUSYGROUP Consumer Group name already exists" != err.Error() {
				logx.Info(fmt.Sprintf("初始化失败！！！%s \n", err))
			}
		}

		for {


			fmt.Printf("\n\n\n拉取任务 \n\n")

			// 才能去定时获取数据
			// 每隔3秒去看看是不是有消息超时了
			timeoutResults, err := dq.RedisCache.XReadGroup(context.Background(), key, group, consumer, 10, 3*time.Second)
			if err != nil {
				return
			}

			fmt.Printf("\n\n\n拉取任务完成，开始消费 \n\n")

			if err != nil && err != redis.Nil {
				logx.Info(fmt.Sprintf("超时数据拉取失败！！！%s \n", err))
			}

			for _, stream := range timeoutResults {
				for _, message := range stream.Messages {
					fields := message.Values

					uuid := fields["uuid"].(string)
					payload := fields["payload"].(string)
					delayTime, _ := strconv.ParseInt(fields["delay_time"].(string), 10, 64)

					now := time.Now().Unix()
					if now >= delayTime {
						logx.Info(fmt.Sprintf("处理任务 id=%s, payload=%s\n", uuid, payload))


						// 处理后确认并删除
						dq.RedisCache.XAck(ctx, stream.Stream, group, message.ID)
						dq.RedisCache.XDel(ctx, stream.Stream, message.ID)

						logx.Info(fmt.Sprintf("任务删除完成 id=%s, payload=%s\n", uuid, payload))
					} else {
						logx.Info(fmt.Sprintf("任务未到期 id=%s，跳过...\n", uuid))
					}
				}
			}
		}

	}()

}
