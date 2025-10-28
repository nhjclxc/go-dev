package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TrafficData 表示单个 App 的流量上报结构
type TrafficData struct {
	Pkg        string    `json:"pkg"`        // 包名
	RxTraffic  int64     `json:"rxtraffic"`  // 接收字节数
	TxTraffic  int64     `json:"txtraffic"`  // 发送字节数
	ReportTime time.Time `json:"reportTime"` // 上报时间（可选）
}

func main() {
	router := gin.Default()

	// POST /client/traffic
	router.POST("/client/traffic", func(c *gin.Context) {
		var data []TrafficData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "invalid json: " + err.Error(),
			})
			return
		}

		// 打印收到的数据
		for _, d := range data {
			log.Printf("[Traffic] pkg=%s rx=%dB tx=%dB time=%s",
				d.Pkg, d.RxTraffic, d.TxTraffic, d.ReportTime)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "received",
			"count": len(data),
		})
	})

	log.Println("🚀 Traffic server started at :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
