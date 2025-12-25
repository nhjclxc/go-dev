package grpc

import (
	"io"
	"log/slog"
	"time"

	"base_project/internal/service"
)

// PcdnMetricsHandler PCDN 指标上报 gRPC 处理器
type PcdnMetricsHandler struct {
	schedulerv1.UnimplementedPcdnMetricsServiceServer
	pcdnReportService *service.PcdnReportService
}

// NewPcdnMetricsHandler 创建 PcdnMetricsHandler 实例
func NewPcdnMetricsHandler(pcdnReportService *service.PcdnReportService) *PcdnMetricsHandler {
	return &PcdnMetricsHandler{
		pcdnReportService: pcdnReportService,
	}
}

// ReportMetrics 处理指标上报双向流
func (h *PcdnMetricsHandler) ReportMetrics(stream schedulerv1.PcdnMetricsService_ReportMetricsServer) error {
	slog.Info("PCDN 指标上报连接建立")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			slog.Info("PCDN 指标上报连接关闭", "reason", "客户端主动关闭")
			return nil
		}
		if err != nil {
			slog.Error("PCDN 指标上报接收失败", "error", err)
			return err
		}

		// 打印请求基本信息
		slog.Info("PCDN 指标上报",
			"reporter_id", req.GetReporterId(),
			"timestamp", req.GetTimestamp(),
			"node_count", len(req.GetNodes()),
		)

		// 处理 PCDN 节点上报数据，写入数据库
		ctx := stream.Context()
		if h.pcdnReportService != nil {
			if err := h.pcdnReportService.ProcessPcdnReport(ctx, req.GetReporterId(), req.GetTimestamp(), req.GetNodes()); err != nil {
				slog.Error("处理 PCDN 节点上报数据失败", "error", err)
			}
		}

		// 发送响应
		resp := &schedulerv1.PcdnMetricsResponse{
			Success:    true,
			ServerTime: time.Now().Unix(),
		}
		if err := stream.Send(resp); err != nil {
			slog.Error("PCDN 指标上报响应发送失败", "error", err)
			return err
		}
	}
}
