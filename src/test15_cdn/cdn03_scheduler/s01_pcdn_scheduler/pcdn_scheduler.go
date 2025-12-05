package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// CDNDomainDetail 单个域名的配置详情
type CDNDomainDetail struct {
	DomainName  string       `json:"domain_name"`  // 域名
	DomainType  string       `json:"domain_type"`  // web/download/video
	CNAMEStatus CNAMEStatus  `json:"cname_status"` // 同步中/已生效
	CNAME       string       `json:"cname"`        // xxx.aliyuncdn.com
	Status      string       `json:"status"`       // online/offline/configuring
	HTTPS       bool         `json:"https"`        // 是否开启HTTPS
	Origins     []OriginInfo `json:"origins"`      // 源站信息
	CreatedAt   string       `json:"created_at"`   // 创建时间
	Tags        []string     `json:"tags"`         // 标签
}

// OriginInfo 源站信息
type OriginInfo struct {
	Address string `json:"address"` // 源站地址
	Type    string `json:"type"`    // IP/DOMAIN/OSS
	Weight  int    `json:"weight"`  // 权重
	Backup  bool   `json:"backup"`  // 是否备源
}

type CNAMEStatus int

const (
	Syncing   CNAMEStatus = 1
	Activated CNAMEStatus = 2
	SyncError CNAMEStatus = 3
)

var domainList []CDNDomainDetail

func init() {

	domainList = []CDNDomainDetail{
		{
			DomainName:  "example.com",
			DomainType:  "web",
			CNAMEStatus: Activated,
			CNAME:       "example.com.w.kunlunsl.com",
			Status:      "online",
			HTTPS:       true,
			Origins: []OriginInfo{
				{Address: "1.1.1.1", Type: "ip", Weight: 10, Backup: false},
				{Address: "2.2.2.2", Type: "ip", Weight: 10, Backup: true},
			},
			CreatedAt: time.Now().AddDate(0, -1, 0).Format(time.RFC3339),
			Tags:      []string{"prod", "global"},
		},
		{
			DomainName:  "cdn.test.com",
			DomainType:  "download",
			CNAMEStatus: Syncing,
			CNAME:       "cdn.test.com.w.kunlunsl.com",
			Status:      "configuring",
			HTTPS:       false,
			Origins: []OriginInfo{
				{Address: "origin.test.com", Type: "domain", Weight: 20, Backup: false},
			},
			CreatedAt: time.Now().Format(time.RFC3339),
			Tags:      []string{"test"},
		},
	}
}

func main() {

	e := gin.Default()

	// 获取域名列表
	// http://127.0.0.1:8899/api/v1/domain/getList
	e.GET("/api/v1/domain/getList", func(c *gin.Context) {

		log.Println("请求")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": domainList,
		})
	})

	e.Run(":8899")

}
