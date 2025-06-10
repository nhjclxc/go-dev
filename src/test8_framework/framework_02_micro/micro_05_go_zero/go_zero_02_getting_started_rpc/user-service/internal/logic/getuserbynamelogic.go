package logic

import (
	"context"
	"fmt"

	"user-service/internal/svc"
	"user-service/proto/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.UserReq) (*user.UserResp, error) {
	// todo: add your logic here and delete this line

	fmt.Println("\n\nGetUserByName.rpc执行到了吗？？？\n\n", in.GetUserId())

	return &user.UserResp{
		UserId: in.UserId,
		Name: "go-zero-rpc geting-started",
		Age: 18,
	}, nil
}
