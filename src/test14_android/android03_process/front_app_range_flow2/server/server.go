package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AppFlowStatus 用于记录每一个app每一次进出的流量
type AppFlowStatus struct {
	Package      string     `json:"package"`
	UID          string     `json:"uid"`
	EntryTime    *time.Time `json:"entryTime"`
	EntryRxTotal int64      `json:"entryRxTotal"`
	EntryTxTotal int64      `json:"entryTxTotal"`
	LeaveTime    *time.Time `json:"leaveTime"`
	LeaveRxTotal int64      `json:"leaveRxTotal"`
	LeaveTxTotal int64      `json:"leaveTxTotal"`
	RxAccum      int64      `json:"rxAccum"`
	TxAccum      int64      `json:"txAccum"`
}

/*
create table android_flow
(
    id                bigint auto_increment comment '主键ID' primary key,
    cinema_id         varchar(255) not null comment '影院id',
    cinema_name       varchar(255) not null comment '影院名称',
    proprietarycoding varchar(255) not null comment '专资',
    hall_id           int not null comment '影厅id',
    hall_name         varchar(255) not null comment '影厅名称',
    android_device_id varchar(100) not null comment '设备id',
    package           varchar(64)  not null comment '包名',
    app_name          varchar(64)  not null comment '应用名称',
    entry_rx_total    bigint       not null comment '进入时接收总计 rx',
    leave_rx_total    bigint       not null comment '离开时接收总计 rx',
    entry_tx_total    bigint       not null comment '进入时发送总计 tx',
    leave_tx_total    bigint       not null comment '离开时发送总计 tx',
    entry_time        timestamp    null comment '进入时间',
    leave_time        timestamp    null comment '离开时间',
    create_at         timestamp    null comment '创建时间'
)
    comment '安卓app流量监控';
*/

func (status *AppFlowStatus) ToString() string {
	enTime, leTime := "", ""
	if status.EntryTime != nil {
		enTime = status.EntryTime.Format("2006-01-02 15:04:05")
	}
	if status.LeaveTime != nil {
		leTime = status.LeaveTime.Format("2006-01-02 15:04:05")
	}
	return fmt.Sprintf("【%s】流量总计，接收总流量: %d, 发送总流量: %d, 进入时间: %s, 离开时间: %s \n",
		status.Package, status.RxAccum, status.TxAccum, enTime, leTime)
}

func main() {
	router := gin.Default()

	// POST /client/traffic
	router.POST("/client/traffic", func(c *gin.Context) {
		var data AppFlowStatus
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "invalid json: " + err.Error(),
			})
			return
		}

		fmt.Printf("当前[%s]流量总计: %s \n %v \n", data.Package, data.ToString(), data)
		fmt.Printf("Rx 总差值%d Byte, %d KB, %d MB  \n", data.RxAccum, data.RxAccum/1024, data.RxAccum/1024/1024)
		fmt.Printf("Tx 总差值%d Byte, %d KB, %d MB  \n", data.TxAccum, data.TxAccum/1024, data.TxAccum/1024/1024)

		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "received",
			"count": 1,
		})
	})

	log.Println("🚀 Traffic server started at :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

// 当前[cn.juqing.cesuwang_tv]流量总计: 【cn.juqing.cesuwang_tv】流量总计，接收总流量: 334196846, 发送总流量: 57394637, 进入时间: 2025-11-04 09:51:04, 离开时间: 2025-11-04 09:51:53
// 当前[cn.juqing.cesuwang_tv]流量总计: 【cn.juqing.cesuwang_tv】流量总计，接收总流量: 469812081, 发送总流量: 51687313, 进入时间: 2025-11-04 09:52:00, 离开时间: 2025-11-04 09:52:48

//fmt.Printf("Rx 总byte差值%d Byte,%d KB, %d MB  \n", 334196846, 334196846/1024, 334196846/1024/1024)
//fmt.Printf("Tx 总byte差值%d Byte, %d KB, %d MB  \n", 57394637, 57394637/1024, 57394637/1024/1024)
//fmt.Printf("Rx 总byte差值%d Byte,%d KB, %d MB  \n", 469812081, 469812081/1024, 469812081/1024/1024)
//fmt.Printf("Tx 总byte差值%d Byte, %d KB, %d MB  \n", 51687313, 51687313/1024, 51687313/1024/1024)
//Rx 总byte差值334196846 Byte,326364 KB, 318 MB
//Tx 总byte差值57394637 Byte, 56049 KB, 54 MB
//Rx 总byte差值469812081 Byte,458800 KB, 448 MB
//Tx 总byte差值51687313 Byte, 50475 KB, 49 MB
