package main

import (
	"context"
	"fmt"
	"gorm_03/config"
	"testing"
	"time"
)

// Client 客户端连接 结构体
type Client struct {
	Id int64 `gorm:"column:id;type:bigint;primaryKey;autoIncrement;not null;comment:主键ID" json:"id" form:"id"` // 主键ID

	Hostname string `gorm:"column:hostname;type:varchar(64);not null;index:uk_hostname_port,priority:1;comment:分配的主机名" json:"hostname" form:"hostname"` // 分配的主机名

	Port int32 `gorm:"column:port;type:int;not null;index:uk_hostname_port,priority:2;comment:分配的端口号" json:"port" form:"port"` // 分配的端口号

	Status int8 `gorm:"column:status;type:tinyint;not null;default:1;comment:连接状态:1-已分配,2-已连接,3=未连接" json:"status" form:"status"` // 连接状态:1-已分配,2-已连接,3=未连接

	Provider string `gorm:"column:provider;type:varchar(64);not null;comment:客户端提供方" json:"provider" form:"provider"` // 客户端提供方

	ClientId string `gorm:"column:client_id;type:varchar(64);not null;comment:客户端唯一id" json:"clientId" form:"clientId"` // 客户端唯一id

	ClientIp string `gorm:"column:client_ip;type:varchar(16);not null;comment:客户端IP" json:"clientIp" form:"clientIp"` // 客户端IP

	ClientProvince string `gorm:"column:client_province;type:varchar(16);not null;comment:客户端省" json:"clientProvince" form:"clientProvince"` // 客户端省

	ClientCity string `gorm:"column:client_city;type:varchar(16);not null;comment:客户端市" json:"clientCity" form:"clientCity"` // 客户端市

	ClientIsp string `gorm:"column:client_isp;type:tinyint;not null;comment:客户端运营商:1=移动,2=联通,3=电信,4=广电" json:"clientIsp" form:"clientIsp"` // 客户端运营商:1=移动,2=联通,3=电信,4=广电

	Version string `gorm:"column:version;type:varchar(32);not null;default:'';comment:客户端版本号" json:"version" form:"version"` // 客户端版本号

	LastConnectTime *time.Time `gorm:"column:last_connect_time;type:timestamp;comment:上次连接时间" json:"lastConnectTime" form:"lastConnectTime"` // 上次连接时间

	LastDisconnectTime *time.Time `gorm:"column:last_disconnect_time;type:timestamp;comment:上次断开时间" json:"lastDisconnectTime" form:"lastDisconnectTime"` // 上次断开时间

}

// TableName 返回当前实体类的表名
func (c *Client) TableName() string {
	return "go_base_project_client"
}
func BatchUpdatego_base_projectClientSelective(updateColumns []string) {
	if len(updateColumns) == 0 {
		fmt.Println(111)
		updateColumns = []string{"hostname", "port", "status", "provider", "client_id", "client_ip", "client_province", "client_city", "client_isp", "last_connect_time", "last_disconnect_time"}
		return
	}
	fmt.Println(222)
}
func Test5(t *testing.T) {
	BatchUpdatego_base_projectClientSelective(nil)

	fmt.Println("----------")
	BatchUpdatego_base_projectClientSelective([]string{})

	fmt.Println("----------")
	BatchUpdatego_base_projectClientSelective(make([]string, 0))

}
func Test333(t *testing.T) {
	fmt.Println(config.DB)
	ctx := context.Background()
	client := Client{Id: 6, ClientIsp: "", LastDisconnectTime: nil}
	updateColums := []string{"id", "hostname", "port", "status", "provider", "client_id", "client_ip", "client_province", "client_city", "client_isp", "last_connect_time", "last_disconnect_time"}
	config.DB.WithContext(ctx).Model(&Client{}).
		Where("id = ?", client.Id).
		Select(updateColums).
		Updates(client)
}
func Test222(t *testing.T) {
	fmt.Println(config.DB)
	ctx := context.Background()
	config.DB.WithContext(ctx).Model(&Client{}).
		Where("id = ?", 6).Select("client_isp", "last_disconnect_time").Updates(Client{ClientIsp: "", LastDisconnectTime: nil})

	config.DB.WithContext(ctx).Model(Client{}).Where("id = ?", 6).Updates(Client{LastDisconnectTime: nil, ClientIsp: ""})

}
func Test111(t *testing.T) {

	fmt.Println(config.DB)
	ctx := context.Background()

	//config.DB.WithContext(ctx).Model(Client{}).Where("id = ?", 6).Updates(map[string]any{
	//	"last_disconnect_time": nil,
	//})

	//config.DB.WithContext(ctx).Model(Client{}).Where("id = ?", 6).Updates(Client{LastDisconnectTime: nil, ClientIsp: ""})

	updateColums := []string{"last_disconnect_time", "client_isp"}
	config.DB.WithContext(ctx).Model(Client{}).Where("id = ?", 6).
		Select(updateColums).Updates(Client{LastDisconnectTime: nil, ClientIsp: ""})

}
