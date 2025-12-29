package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_05_practice/config"
)

// 实现hash分表
func main() {

	// 安装userId来进行hash，以此进行数据的水平分表
	// Hash 分表（水平分表）
	/*
		CREATE TABLE user_log (
		  id BIGINT auto_increment PRIMARY KEY,
		  user_id BIGINT NOT NULL,
		  log_data json,
		  INDEX idx_user_id (user_id)
		);

		以上表创建8个，分别为 user_log_01，user_log_02，user_log_03 ... user_log_08
	*/

	/*
		同一规则，把一行数据“算”到唯一一张子表
		公式本质是：
			table_index = hash(shard_key) % N
				- shard_key：分表键（user_id / order_id / uid 等）
				- N：分表数量

	*/

	_ = config.DB

	//for i := range 100 {
	//	for j := range 10 {
	//
	//	}
	//}

}

func BatchInsert[T any](tableName string, batchSize int, data []*UserLog) error {
	if len(data) == 0 {
		return nil
	}

	return config.DB.Table(tableName).Session(&gorm.Session{
		PrepareStmt: true,
	}).CreateInBatches(data, batchSize).Error
}

func (*UserLog) TableName(userID int) string {
	idx := userID % 8
	return fmt.Sprintf("user_log_%d", idx)
}

type UserLog struct {
	Id      int    `gorm:"column:used;primary_key"`
	UserId  int    `gorm:"column:user_id"`
	LogData string `gorm:"column:log_data"`
}
