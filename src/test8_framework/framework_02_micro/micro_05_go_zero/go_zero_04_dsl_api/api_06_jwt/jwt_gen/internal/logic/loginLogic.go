package logic

import (
	"context"
	"errors"
	"jwt_gen/utils/jwt"

	"jwt_gen/internal/svc"
	"jwt_gen/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (this *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	// 登录
	// admin, admin123

	if "admin" != req.Username {
		return nil, errors.New("找不到该用户！！！")
	}

	if "admin123" != req.Password {
		return nil, errors.New("密码不正确！！！")
	}

	// 生成 token
	resp = &types.LoginResp{
		ID: "666",
	}

	auth := this.svcCtx.Config.Auth

	token, err := jwt.GenerateToken(resp.ID, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return nil, err
	}

	// 返回token
	resp.Token = token

	/*
	http://localhost:8090/user/login

	{
	"username": "admin",
	"password": "admin123"
	}

	*/
	return
}
