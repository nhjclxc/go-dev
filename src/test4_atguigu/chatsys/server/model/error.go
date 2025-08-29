package model

import (
	"errors"
)

// 自定义登录或注册是的错误
var (
	ErrUserNotExist  = errors.New("anonymous_user not exist")
	ErrInvalidPasswd = errors.New("Passwd or username not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("anonymous_user exist")
)
