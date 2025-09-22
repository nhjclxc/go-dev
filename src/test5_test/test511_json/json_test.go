package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Aaa struct {
	Ids  []int `json:"ids[]"`
	Ids2 []int `json:"ids"`
}

func TestName(t *testing.T) {
	a := Aaa{
		Ids:  []int{1, 2, 3},
		Ids2: []int{1, 2, 3},
	}

	marshal, err := json.Marshal(&a)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
