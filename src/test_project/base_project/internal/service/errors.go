package service

import (
	"base_project/pkg/errors"
)

// 业务错误定义 - 直接引用 pkg/errors 中的错误
var (
	ErrInvalidUsername = errors.ErrInvalidUsername
	ErrUsernameExists  = errors.ErrUsernameExists
	ErrUserNotFound    = errors.ErrUserNotFound

	// 节点相关错误
	ErrNodeNotFound     = errors.ErrNotFound.WithMessage("节点不存在")
	ErrInvalidNodeID    = errors.ErrInvalidParam.WithMessage("无效的节点 ID")
	ErrConnectionClosed = errors.ErrBadRequest.WithMessage("连接已关闭")
)
