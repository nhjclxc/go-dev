package logic

import (
	"context"

	"go_zero_19_db_mysql/internal/svc"
	"go_zero_19_db_mysql/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MysqlApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewMysqlApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MysqlApiLogic {
	return &MysqlApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MysqlApiLogic) MysqlApi(req *types.MysqlApiReq) (resp *types.MysqlApiResp, err error) {
	// todo: add your logic here and delete this line

	return
}
