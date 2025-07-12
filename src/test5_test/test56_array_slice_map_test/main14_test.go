package main

import (
	"fmt"
	"testing"
)

func TestSumAmountByMonth(t *testing.T) {

	orders := []Order2{
		{"2025-07-01", 100},
		{"2025-07-15", 200},
		{"2025-08-01", 300},
		{"2025-08-20", 400},
	}
	result := SumAmountByMonth(orders)
	fmt.Println(result) // map[2025-07:300 2025-08:700]
}