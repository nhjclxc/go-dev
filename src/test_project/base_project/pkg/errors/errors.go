package errors

import (
	"errors"
	"fmt"
)

// AppError 应用错误，包含 HTTP 状态码和业务错误码
type AppError struct {
	HTTPCode int    `json:"-"`       // HTTP 状态码（不输出到 JSON）
	Code     int    `json:"code"`    // 业务错误码
	Message  string `json:"message"` // 用户可见消息
	Err      error  `json:"-"`       // 内部错误（不暴露给用户）
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Unwrap 实现 errors.Unwrap 接口
func (e *AppError) Unwrap() error {
	return e.Err
}

// WithError 包装内部错误
func (e *AppError) WithError(err error) *AppError {
	return &AppError{
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
		Err:      err,
	}
}

// WithMessage 自定义用户可见消息
func (e *AppError) WithMessage(message string) *AppError {
	return &AppError{
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  message,
		Err:      e.Err,
	}
}

// WithMessagef 自定义用户可见消息（格式化）
func (e *AppError) WithMessagef(format string, args ...interface{}) *AppError {
	return &AppError{
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  fmt.Sprintf(format, args...),
		Err:      e.Err,
	}
}

// Is 实现 errors.Is 接口，只比较错误码
func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// 错误码分段定义
// 10000-19999: 客户端错误（参数、验证等）
// 20000-29999: 认证和权限错误
// 30000-39999: 资源错误（不存在、冲突等）
// 50000-59999: 服务端错误

// 通用错误 - HTTP 400
var (
	ErrBadRequest      = &AppError{HTTPCode: 400, Code: 10001, Message: "请求参数错误"}
	ErrInvalidParam    = &AppError{HTTPCode: 400, Code: 10002, Message: "参数格式不正确"}
	ErrMissingParam    = &AppError{HTTPCode: 400, Code: 10003, Message: "缺少必要参数"}
	ErrValidationError = &AppError{HTTPCode: 400, Code: 10004, Message: "参数验证失败"}
)

// 用户相关错误 - HTTP 400
var (
	ErrInvalidUsername = &AppError{HTTPCode: 400, Code: 11001, Message: "无效的用户名"}
	ErrInvalidEmail    = &AppError{HTTPCode: 400, Code: 11002, Message: "无效的邮箱地址"}
	ErrUsernameExists  = &AppError{HTTPCode: 400, Code: 11003, Message: "用户名已存在"}
	ErrEmailExists     = &AppError{HTTPCode: 400, Code: 11004, Message: "邮箱已存在"}
)

// 认证错误 - HTTP 401
var (
	ErrUnauthorized    = &AppError{HTTPCode: 401, Code: 20001, Message: "未登录或登录已过期"}
	ErrInvalidToken    = &AppError{HTTPCode: 401, Code: 20002, Message: "无效的访问令牌"}
	ErrTokenExpired    = &AppError{HTTPCode: 401, Code: 20003, Message: "访问令牌已过期"}
	ErrInvalidPassword = &AppError{HTTPCode: 401, Code: 20004, Message: "密码错误"}
)

// 权限错误 - HTTP 403
var (
	ErrForbidden      = &AppError{HTTPCode: 403, Code: 20101, Message: "无权限访问"}
	ErrPermissionDeny = &AppError{HTTPCode: 403, Code: 20102, Message: "权限不足"}
)

// 资源错误 - HTTP 404
var (
	ErrNotFound     = &AppError{HTTPCode: 404, Code: 30001, Message: "资源不存在"}
	ErrUserNotFound = &AppError{HTTPCode: 404, Code: 30002, Message: "用户不存在"}
)

// 冲突错误 - HTTP 409
var (
	ErrConflict       = &AppError{HTTPCode: 409, Code: 30101, Message: "资源冲突"}
	ErrResourceExists = &AppError{HTTPCode: 409, Code: 30102, Message: "资源已存在"}
)

// 服务端错误 - HTTP 500（不暴露具体细节给用户）
var (
	ErrInternal   = &AppError{HTTPCode: 500, Code: 50000, Message: "服务器内部错误"}
	ErrDatabase   = &AppError{HTTPCode: 500, Code: 50001, Message: "服务器内部错误"}
	ErrRedis      = &AppError{HTTPCode: 500, Code: 50002, Message: "服务器内部错误"}
	ErrThirdParty = &AppError{HTTPCode: 500, Code: 50003, Message: "服务器内部错误"}
)

// As 是 errors.As 的快捷方式
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is 是 errors.Is 的快捷方式
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// New 创建一个新的错误
func New(message string) error {
	return errors.New(message)
}

// Wrap 包装错误
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}
