package response

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"

	"base_project/pkg/errors"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

// Success 成功响应
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
		TraceID: getTraceID(ctx),
	})
}

// Created 创建成功响应（HTTP 201）
func Created(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, Response{
		Code:    0,
		Message: "success",
		Data:    data,
		TraceID: getTraceID(ctx),
	})
}

// NoContent 无内容响应（HTTP 204）
func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// Error 错误响应
func Error(ctx *gin.Context, err error) {
	var appErr *errors.AppError

	if errors.As(err, &appErr) {
		// 已知业务错误
		// 如果是服务端错误，记录日志
		if appErr.HTTPCode >= 500 {
			slog.Error("服务端错误",
				"code", appErr.Code,
				"message", appErr.Message,
				"error", appErr.Err,
				"trace_id", getTraceID(ctx),
			)
		}

		ctx.JSON(appErr.HTTPCode, Response{
			Code:    appErr.Code,
			Message: appErr.Message,
			TraceID: getTraceID(ctx),
		})
		return
	}

	// 未知错误，统一返回 500，不暴露细节
	slog.Error("未处理的错误",
		"error", err,
		"trace_id", getTraceID(ctx),
	)

	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    errors.ErrInternal.Code,
		Message: errors.ErrInternal.Message,
		TraceID: getTraceID(ctx),
	})
}

// BadRequest 参数错误响应
func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    errors.ErrBadRequest.Code,
		Message: message,
		TraceID: getTraceID(ctx),
	})
}

// Unauthorized 未授权响应
func Unauthorized(ctx *gin.Context, message string) {
	if message == "" {
		message = errors.ErrUnauthorized.Message
	}
	ctx.JSON(http.StatusUnauthorized, Response{
		Code:    errors.ErrUnauthorized.Code,
		Message: message,
		TraceID: getTraceID(ctx),
	})
}

// Forbidden 禁止访问响应
func Forbidden(ctx *gin.Context, message string) {
	if message == "" {
		message = errors.ErrForbidden.Message
	}
	ctx.JSON(http.StatusForbidden, Response{
		Code:    errors.ErrForbidden.Code,
		Message: message,
		TraceID: getTraceID(ctx),
	})
}

// NotFound 资源不存在响应
func NotFound(ctx *gin.Context, message string) {
	if message == "" {
		message = errors.ErrNotFound.Message
	}
	ctx.JSON(http.StatusNotFound, Response{
		Code:    errors.ErrNotFound.Code,
		Message: message,
		TraceID: getTraceID(ctx),
	})
}

// getTraceID 从上下文获取 trace ID
func getTraceID(ctx *gin.Context) string {
	span := trace.SpanFromContext(ctx.Request.Context())
	if span.SpanContext().IsValid() {
		return span.SpanContext().TraceID().String()
	}
	return ""
}

// PageData 分页数据
type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// SuccessWithPage 分页成功响应
func SuccessWithPage(ctx *gin.Context, list interface{}, total int64, page, pageSize int) {
	Success(ctx, PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
