// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: Group.proto

package userclassservice

import (
	"context"

	"go-dev/src/test8_framework/framework_02_micro/micro_05_go_zero/go_zero_05_dsl_proto/group/gen/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq            = user.LoginReq
	LoginResp           = user.LoginResp
	UserClassAddReq     = user.UserClassAddReq
	UserClassAddResp    = user.UserClassAddResp
	UserClassDeleteReq  = user.UserClassDeleteReq
	UserClassDeleteResp = user.UserClassDeleteResp
	UserClassInfoReq    = user.UserClassInfoReq
	UserClassInfoResp   = user.UserClassInfoResp
	UserClassListReq    = user.UserClassListReq
	UserClassListResp   = user.UserClassListResp
	UserClassUpdateReq  = user.UserClassUpdateReq
	UserClassUpdateResp = user.UserClassUpdateResp
	UserInfoReq         = user.UserInfoReq
	UserInfoResp        = user.UserInfoResp
	UserInfoUpdateReq   = user.UserInfoUpdateReq
	UserInfoUpdateResp  = user.UserInfoUpdateResp
	UserListReq         = user.UserListReq
	UserListResp        = user.UserListResp
	UserRoleAddReq      = user.UserRoleAddReq
	UserRoleAddResp     = user.UserRoleAddResp
	UserRoleDeleteReq   = user.UserRoleDeleteReq
	UserRoleDeleteResp  = user.UserRoleDeleteResp
	UserRoleInfoReq     = user.UserRoleInfoReq
	UserRoleInfoResp    = user.UserRoleInfoResp
	UserRoleListReq     = user.UserRoleListReq
	UserRoleListResp    = user.UserRoleListResp
	UserRoleUpdateReq   = user.UserRoleUpdateReq
	UserRoleUpdateResp  = user.UserRoleUpdateResp

	UserClassService interface {
		UserClassList(ctx context.Context, in *UserClassListReq, opts ...grpc.CallOption) (*UserClassListResp, error)
		UserClassUpdate(ctx context.Context, in *UserClassUpdateReq, opts ...grpc.CallOption) (*UserClassUpdateResp, error)
		UserClassInfo(ctx context.Context, in *UserClassInfoReq, opts ...grpc.CallOption) (*UserClassInfoResp, error)
		UserClassAdd(ctx context.Context, in *UserClassAddReq, opts ...grpc.CallOption) (*UserClassAddResp, error)
		UserClassDelete(ctx context.Context, in *UserClassDeleteReq, opts ...grpc.CallOption) (*UserClassDeleteResp, error)
	}

	defaultUserClassService struct {
		cli zrpc.Client
	}
)

func NewUserClassService(cli zrpc.Client) UserClassService {
	return &defaultUserClassService{
		cli: cli,
	}
}

func (m *defaultUserClassService) UserClassList(ctx context.Context, in *UserClassListReq, opts ...grpc.CallOption) (*UserClassListResp, error) {
	client := user.NewUserClassServiceClient(m.cli.Conn())
	return client.UserClassList(ctx, in, opts...)
}

func (m *defaultUserClassService) UserClassUpdate(ctx context.Context, in *UserClassUpdateReq, opts ...grpc.CallOption) (*UserClassUpdateResp, error) {
	client := user.NewUserClassServiceClient(m.cli.Conn())
	return client.UserClassUpdate(ctx, in, opts...)
}

func (m *defaultUserClassService) UserClassInfo(ctx context.Context, in *UserClassInfoReq, opts ...grpc.CallOption) (*UserClassInfoResp, error) {
	client := user.NewUserClassServiceClient(m.cli.Conn())
	return client.UserClassInfo(ctx, in, opts...)
}

func (m *defaultUserClassService) UserClassAdd(ctx context.Context, in *UserClassAddReq, opts ...grpc.CallOption) (*UserClassAddResp, error) {
	client := user.NewUserClassServiceClient(m.cli.Conn())
	return client.UserClassAdd(ctx, in, opts...)
}

func (m *defaultUserClassService) UserClassDelete(ctx context.Context, in *UserClassDeleteReq, opts ...grpc.CallOption) (*UserClassDeleteResp, error) {
	client := user.NewUserClassServiceClient(m.cli.Conn())
	return client.UserClassDelete(ctx, in, opts...)
}
