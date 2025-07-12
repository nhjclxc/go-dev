package main

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {


	data := map[string]map[string]int{
		"2025": {
			"Jan": 100,
			"Feb": 200,
		},
		"2026": {
			"Jan": 150,
		},
	}

	fmt.Println(Flatten(data))
}