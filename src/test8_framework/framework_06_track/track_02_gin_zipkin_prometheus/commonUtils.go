package main


import (
	"context"
	"go.opentelemetry.io/otel/trace"
)

// 从上下文中提取 traceId 和 spanId
func GetTraceInfo(ctx context.Context) (traceId, spanId string) {
	span := trace.SpanFromContext(ctx)
	spanCtx := span.SpanContext()

	if spanCtx.HasTraceID() && spanCtx.HasSpanID() {
		return spanCtx.TraceID().String(), spanCtx.SpanID().String()
	}
	return "", ""
}
