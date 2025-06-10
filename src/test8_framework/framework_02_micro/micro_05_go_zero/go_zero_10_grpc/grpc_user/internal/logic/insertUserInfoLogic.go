package logic

import (
	"context"

	"grpc_user/internal/svc"
	"grpc_user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增用户
func NewInsertUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserInfoLogic {
	return &InsertUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertUserInfoLogic) InsertUserInfo(req *types.InsertUserInfoReq) error {
	// todo: add your logic here and delete this line

	return nil
}
