package logic

import (
	"context"
	"time"

	"go_zero_17_delay_queue/internal/svc"
	"go_zero_17_delay_queue/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelayQueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDelayQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelayQueueLogic {
	return &DelayQueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelayQueueLogic) DelayQueue(req *types.DelayQueueReq) error {
	// todo: add your logic here and delete this line


	msg := "data"

	// 1、5s后执行
	deplayResp, err := l.svcCtx.DqProducer.Delay([]byte(msg), time.Second*5)
	if err != nil {
		logx.Errorf("error from DqPusherClient Delay err : %v", err)
	}
	logx.Infof("resp : %s", deplayResp) // fmt.Sprintf("%s/%s/%d", p.endpoint, p.tube, id)

	// 2、在某个指定时间执行
	atResp, err := l.svcCtx.DqProducer.At([]byte(msg), time.Now())
	if err != nil {
		logx.Errorf("error from DqPusherClient Delay err : %v", err)
	}
	logx.Infof("resp : %s", atResp) // fmt.Sprintf("%s/%s/%d", p.endpoint, p.tube, id)

	return nil
}
