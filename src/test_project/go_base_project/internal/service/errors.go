package service

import "errors"

// 业务错误定义
var (
	ErrInvalidUsername = errors.New("无效的用户名")
	ErrUsernameExists  = errors.New("用户名已存在")
	ErrUserNotFound    = errors.New("用户不存在")
)
