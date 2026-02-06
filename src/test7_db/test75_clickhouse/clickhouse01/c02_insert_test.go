package main

import (
	"context"
	"fmt"
	"golang.org/x/exp/rand"
	"log"
	ch "test75_clickhouse/clickhouse"
	"testing"
	"time"
)

// 将数据写入 ClickHouse: https://clickhouse.com/docs/zh/guides/inserting-data

type MyFirstTable2 struct {
	UserId    int
	Message   string
	Timestamp int64
	Metric    int
}

func genDatas(i int) []*MyFirstTable2 {
	result := make([]*MyFirstTable2, i)
	rand.Seed(uint64(time.Now().UnixNano()))
	nowts := time.Now().Unix()
	for j := 0; j < i; j++ {
		result[j] = &MyFirstTable2{
			UserId:    rand.Intn(100),
			Message:   fmt.Sprintf("<Message> %d", rand.Intn(100)),
			Timestamp: nowts - rand.Int63n(1000),
			Metric:    rand.Intn(100),
		}
	}
	return result
}

func TestInsert01(t *testing.T) {
	ctx := context.Background()
	_ = ctx
	// 这个函数的批插入是伪批插入，模仿了mysql的批插入

	// 数据准备
	datas := genDatas(10)

	// 一般情况下clickhouse的数据都要求进行批插入
	// 先在内存中收集保存批量数据，到达批插入阈值的时候才进行批插入，【切勿一条一条数据进行插入这样会拉低性能】

	// 开启批插入事务
	tx, err := ch.ClickHouseCli.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("开始事务失败: %w", err)
		return
	}
	commited := false
	defer func() {
		v := recover()
		if v != nil {
			if !commited {
				fmt.Println("事务提交失败，执行回滚")
				err := tx.Rollback()
				if err != nil {
					fmt.Println("<tx.Rollback()>: %w", err)
					return
				}
			}
		}
	}()

	// 准备 batch
	stmt, err := tx.PrepareContext(ctx,
		`
			INSERT INTO my_first_table2 (user_id, message, timestamp, metric)
			VALUES (?, ?, ?, ?)
			`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 执行sql
	for _, data := range datas {
		_, err := stmt.ExecContext(ctx, data.UserId, data.Message, data.Timestamp, data.Metric)
		if err != nil {
			log.Fatal("批量插入err：", err.Error())
			return
		}

		// 模拟panic
		//if data.UserId%3 == 0 {
		//	x := 0
		//	fmt.Println("模拟panic：", data.UserId/x)
		//}

		fmt.Printf("数据插入成功：%#v \n", data)
		fmt.Println()
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		log.Fatal("事务提交失败：", err.Error())
		return
	}
	commited = true

	fmt.Println("事务提交成功！！！")

}

func TestInsert02(t *testing.T) {
	ctx := context.Background()
	_ = ctx
	conn := ch.ClickHouseCli.Conn
	_ = conn

	// 数据准备
	datas := genDatas(5)

	// INSERT INTO my_first_table2 (user_id, message, timestamp, metric)
	batch, err := conn.PrepareBatch(ctx,
		"INSERT INTO my_first_table2 (user_id, message, timestamp, metric)",
	)
	if err != nil {
		log.Fatal("PrepareBatch 失败：", err)
		return
	}
	defer batch.Close()

	// 追加数据
	for _, data := range datas {
		err := batch.Append(data.UserId, data.Message, data.Timestamp, data.Metric)
		if err != nil {
			log.Fatal("batch.Append 失败：", err)
			return
		}
	}

	// 提交数据
	if err := batch.Send(); err != nil {
		log.Fatal("batch.Send() 失败：", err)
		return
	}
	fmt.Println("<数据插入成功>")
}
