package grpc

import (
	"base_project/internal/service"
	"google.golang.org/grpc"
	"io"
	"log/slog"
	"time"
)

// EdgeMetricsHandler Edge 指标上报 gRPC 处理器
type EdgeMetricsHandler struct {
	schedulerv1.UnimplementedEdgeMetricsServiceServer
	edgeReportService *service.EdgeReportService
}

// NewEdgeMetricsHandler 创建 EdgeMetricsHandler 实例
func NewEdgeMetricsHandler(edgeReportService *service.EdgeReportService) *EdgeMetricsHandler {
	return &EdgeMetricsHandler{
		edgeReportService: edgeReportService,
	}
}

// Connect 处理指标上报双向流
func (h *EdgeMetricsHandler) Connect(stream grpc.BidiStreamingServer[schedulerv1.EdgeMetricsRequest, schedulerv1.EdgeMetricsResponse]) error {

	slog.Info("edgeMetricsHandler Connect successful ")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			slog.Info("Edge 指标上报连接关闭", "reason", "客户端主动关闭")
			return nil
		}
		if err != nil {
			slog.Error("Edge 指标上报接收失败", "error", err)
			return err
		}

		// 打印请求基本信息
		slog.Info("Edge 指标上报",
			"node_id", req.GetHeartbeat().Data.NodeId,
			"timestamp", req.GetHeartbeat().GetTimestamp(),
		)

		// 处理 Edge 节点上报数据，写入数据库
		ctx := stream.Context()
		if h.edgeReportService != nil {
			if err := h.edgeReportService.ProcessEdgeReport(ctx, req.GetHeartbeat().Data.NodeId, req.GetHeartbeat().GetTimestamp(), req.GetHeartbeat().Data); err != nil {
				slog.Error("处理 Edge 节点上报数据失败", "error", err)
			}
		}

		// 发送响应
		resp := &schedulerv1.EdgeMetricsResponse{
			Payload: &schedulerv1.EdgeMetricsResponse_HeartbeatAck{
				HeartbeatAck: &schedulerv1.EdgeHeartbeatAck{
					HeartbeatInterval: 111,
					ServerTime:        time.Now().Unix(),
				},
			},
		}

		if err := stream.Send(resp); err != nil {
			slog.Error("Edge 指标上报响应发送失败", "error", err)
			return err
		}
	}
}
