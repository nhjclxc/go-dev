package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 复习一下json序列化

	p1 := Person{
		Name:   "zhangsan",
		Age:    18,
		Salary: 123456.12,
	}

	fmt.Println(p1)

	// 序列化
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(bytes)
	fmt.Printf("jsonStr = %s \n", jsonStr)

	// 反序列化
	p2 := Person{}
	str := "{\"PersonName\":\"赵六\",\"AgeNumber\":12,\"Money\":123.12}"
	err2 := json.Unmarshal([]byte(str), &p2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(p2)
}

type Person struct {
	Name   string  `json:"PersonName"`
	Age    int     `json:"AgeNumber"`
	Salary float32 `json:"Money"`
}
