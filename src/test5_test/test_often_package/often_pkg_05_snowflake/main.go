package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func main() {

	// 全局雪花id的使用
	// https://github.com/bwmarrin/snowflake
	// go mod init often_pkg_05_snowflake
	// go get github.com/bwmarrin/snowflake

	// NewNode方法一般是在每个微服务启动的时候初始化，一个微服务就是一个结点node
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())




	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", node.Generate().Step())
	fmt.Printf("ID Step  : %d\n", node.Generate().Step())
	fmt.Printf("ID Step  : %d\n", node.Generate().Step())
	fmt.Printf("ID Step  : %d\n", node.Generate().Step())
	fmt.Printf("ID Step  : %d\n", node.Generate().Step())


}
