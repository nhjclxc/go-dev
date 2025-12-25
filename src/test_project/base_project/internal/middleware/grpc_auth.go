package middleware

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"base_project/config"
)

const (
	// AuthorizationHeader 鉴权头名称
	AuthorizationHeader = "authorization"
	// BearerPrefix Bearer 前缀
	BearerPrefix = "Bearer "
)

// AuthInterceptor Token 鉴权拦截器
type AuthInterceptor struct {
	config *config.GRPCAuthConfig
}

// NewAuthInterceptor 创建鉴权拦截器
func NewAuthInterceptor(cfg *config.GRPCAuthConfig) *AuthInterceptor {
	return &AuthInterceptor{
		config: cfg,
	}
}

// Unary 一元 RPC 鉴权拦截器
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 如果未启用鉴权，直接放行
		if !i.config.Enabled {
			return handler(ctx, req)
		}

		// 验证 token
		if err := i.authorize(ctx); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream 流式 RPC 鉴权拦截器
func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// 如果未启用鉴权，直接放行
		if !i.config.Enabled {
			return handler(srv, ss)
		}

		// 验证 token
		if err := i.authorize(ss.Context()); err != nil {
			return err
		}

		return handler(srv, ss)
	}
}

// authorize 验证请求中的 token
func (i *AuthInterceptor) authorize(ctx context.Context) error {
	// 从 metadata 中获取 token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.Warn("gRPC 鉴权失败: 缺少 metadata")
		return status.Error(codes.Unauthenticated, "missing metadata")
	}

	// 获取 authorization header
	values := md.Get(AuthorizationHeader)
	if len(values) == 0 {
		slog.Warn("gRPC 鉴权失败: 缺少 authorization header")
		return status.Error(codes.Unauthenticated, "missing authorization header")
	}

	// 解析 Bearer token
	authHeader := values[0]
	if !strings.HasPrefix(authHeader, BearerPrefix) {
		slog.Warn("gRPC 鉴权失败: 无效的 authorization 格式")
		return status.Error(codes.Unauthenticated, "invalid authorization format")
	}

	token := strings.TrimPrefix(authHeader, BearerPrefix)

	// 验证 token
	if token != i.config.Token {
		slog.Warn("gRPC 鉴权失败: token 无效")
		return status.Error(codes.Unauthenticated, "invalid token")
	}

	return nil
}
