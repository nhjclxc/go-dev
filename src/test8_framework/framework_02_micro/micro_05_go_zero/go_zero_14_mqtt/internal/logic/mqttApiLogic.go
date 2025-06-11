package logic

import (
	"context"
	"fmt"

	"go_zero_14_mqtt/internal/svc"
	"go_zero_14_mqtt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MqttApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewMqttApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MqttApiLogic {
	return &MqttApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqttApiLogic) MqttApi(req *types.MqttApiReq) (err error) {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n MqttApi = %#v \n\n\n", req)

	err2 := l.svcCtx.MqttClient.PublishQos0(req.Topic, req.Msg)
	if err2 != nil {
		return err2
	}

	return
}
