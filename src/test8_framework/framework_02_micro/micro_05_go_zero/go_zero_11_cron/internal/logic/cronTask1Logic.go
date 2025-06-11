package logic

import (
	"context"
	"fmt"
	"time"

	"go_zero_11_cron/internal/svc"
	"go_zero_11_cron/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CronTask1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewCronTask1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *CronTask1Logic {
	return &CronTask1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CronTask1Logic) CronTask1(req *types.CronTaskReq) (resp *types.CronTaskResp, err error) {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n\n CronTask1 = %#v \n\n\n", req)

	CornDoSomething(l.svcCtx)

	return &types.CronTaskResp{
		TaskId:   req.TaskId,
		TaskName: "这是一个cron的定时任务哦",
		Address:  "阿斯顿法国红酒",
		Age:      10,
	}, nil
}

func CornDoSomething(ctx *svc.ServiceContext) {
	fmt.Println("执行定时任务逻辑  ", time.Now())
	// 在这里可以操作数据库、缓存、发送请求等
}