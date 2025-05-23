// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: Group.proto

package userroleservice

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

	UserRoleService interface {
		UserRoleList(ctx context.Context, in *UserRoleListReq, opts ...grpc.CallOption) (*UserRoleListResp, error)
		UserRoleUpdate(ctx context.Context, in *UserRoleUpdateReq, opts ...grpc.CallOption) (*UserRoleUpdateResp, error)
		UserRoleInfo(ctx context.Context, in *UserRoleInfoReq, opts ...grpc.CallOption) (*UserRoleInfoResp, error)
		UserRoleAdd(ctx context.Context, in *UserRoleAddReq, opts ...grpc.CallOption) (*UserRoleAddResp, error)
		UserRoleDelete(ctx context.Context, in *UserRoleDeleteReq, opts ...grpc.CallOption) (*UserRoleDeleteResp, error)
	}

	defaultUserRoleService struct {
		cli zrpc.Client
	}
)

func NewUserRoleService(cli zrpc.Client) UserRoleService {
	return &defaultUserRoleService{
		cli: cli,
	}
}

func (m *defaultUserRoleService) UserRoleList(ctx context.Context, in *UserRoleListReq, opts ...grpc.CallOption) (*UserRoleListResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.UserRoleList(ctx, in, opts...)
}

func (m *defaultUserRoleService) UserRoleUpdate(ctx context.Context, in *UserRoleUpdateReq, opts ...grpc.CallOption) (*UserRoleUpdateResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.UserRoleUpdate(ctx, in, opts...)
}

func (m *defaultUserRoleService) UserRoleInfo(ctx context.Context, in *UserRoleInfoReq, opts ...grpc.CallOption) (*UserRoleInfoResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.UserRoleInfo(ctx, in, opts...)
}

func (m *defaultUserRoleService) UserRoleAdd(ctx context.Context, in *UserRoleAddReq, opts ...grpc.CallOption) (*UserRoleAddResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.UserRoleAdd(ctx, in, opts...)
}

func (m *defaultUserRoleService) UserRoleDelete(ctx context.Context, in *UserRoleDeleteReq, opts ...grpc.CallOption) (*UserRoleDeleteResp, error) {
	client := user.NewUserRoleServiceClient(m.cli.Conn())
	return client.UserRoleDelete(ctx, in, opts...)
}
