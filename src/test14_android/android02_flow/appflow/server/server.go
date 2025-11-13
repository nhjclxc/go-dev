package main

import (
	"fmt"
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

/*
create table android_app_flow
(
    id                bigint auto_increment comment '主键ID' primary key,
    proprietarycoding varchar(255) not null comment '专资',
    cinema_id         int          not null comment '影院id',
    hall_id           int          not null comment '影厅id',
    android_device_id varchar(100) not null comment '设备id',
    package           varchar(64)  not null comment '应用包名',
    app_name          varchar(64)  not null comment '应用名称',
    rx_total          bigint       not null comment '接收总计 rx',
    tx_total          bigint       not null comment '发送总计 tx',
    report_time       timestamp    null comment '上报时间',
    create_at         timestamp    null comment '创建时间'
)
    comment '安卓app流量监控';

安卓流量监控
android-traffic-monitor
*/

func main() {
	router := gin.Default()

	// pkg -> *TrafficData
	var totalMap map[string]*TrafficData = make(map[string]*TrafficData)

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

			var tt *TrafficData
			if t, ok := totalMap[d.Pkg]; ok {
				t.RxTraffic += d.RxTraffic
				t.TxTraffic += d.TxTraffic
				tt = t
			} else {
				t := TrafficData{
					Pkg:       d.Pkg,
					RxTraffic: d.RxTraffic,
					TxTraffic: d.TxTraffic,
				}
				totalMap[d.Pkg] = &t
				tt = &t
			}
			fmt.Printf("当前[%s]流量总计: rx = %d, tx = %d  \n", d.Pkg, tt.RxTraffic, tt.TxTraffic)
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

// 未关闭： rx = 530424421, tx = 46092222
// 中途退出：rx=，tx=
