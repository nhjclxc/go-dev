package model_utils

import (
	"fmt"
	"testing"
)

type BaseStruct struct {
	ID   int
	Name string
}

type ReqStruct struct {
	BaseStruct
	Age int // 这个字段不会被拷贝
}

type ResStruct struct {
	BaseStruct
	City string // 这个字段不会被填充
}

func TestModelUtils(t *testing.T) {

	req := ReqStruct{
		BaseStruct: BaseStruct{
			ID:   1,
			Name: "Alice",
		},
		Age: 25,
	}
	var res ResStruct

	CopyStructFields[ReqStruct, ResStruct](req, &res)

	fmt.Printf("%+v\n", res) // Output: {ID:1 Name:Alice City:}

}
