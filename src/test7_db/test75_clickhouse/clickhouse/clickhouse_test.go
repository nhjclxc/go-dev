package ch

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	err := ClickHouseCli.DB.Ping()
	if err != nil {
		return
	}

	stats := ClickHouseCli.DB.Stats()
	fmt.Println(stats.Idle)
	fmt.Println(stats.InUse)
	fmt.Println(stats.OpenConnections)
}
