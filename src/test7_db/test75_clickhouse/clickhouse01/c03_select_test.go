package main

import (
	"context"
	"fmt"
	ch "test75_clickhouse/clickhouse"
	"testing"
)

var (
	ctx  = context.Background()
	db   = ch.ClickHouseCli.DB
	conn = ch.ClickHouseCli.Conn
)

func TestSelect_00(t *testing.T) {
	/*
		SELECT *
		FROM my_first_table2
	*/

	rows, err := db.QueryContext(ctx, "SELECT * FROM my_first_table2")
	if err != nil {
		fmt.Println("QueryContext err: ", err)
		return
	}
	defer rows.Close()

	rows.Next()
	rows.Err()

}

func TestSelect_01(t *testing.T) {
	/*
		SELECT *
		FROM my_first_table2
		ORDER BY timestamp DESC
		LIMIT 10;
	*/
}
