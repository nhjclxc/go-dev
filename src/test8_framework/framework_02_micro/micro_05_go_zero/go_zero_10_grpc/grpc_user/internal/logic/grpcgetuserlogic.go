package logic

import (
	"context"
	"fmt"

	"grpc_user/grpc/user"
	"grpc_user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrpcGetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrpcGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrpcGetUserLogic {
	return &GrpcGetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrpcGetUserLogic) GrpcGetUser(in *user.GrpcGetUserReq) (*user.GrpcGetUserResp, error) {
	// todo: add your logic here and delete this line

	fmt.Printf("\n\n\n GrpcGetUser 远程调用成功！！！ %#v \n\n\n", in)

	return &user.GrpcGetUserResp{
		UserId: in.UserId,
		Name:   "我是远程调用的用户",
		Age:    18,
	}, nil
}
