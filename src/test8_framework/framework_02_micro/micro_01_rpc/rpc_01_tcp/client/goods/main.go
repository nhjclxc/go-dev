package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"net/rpc"
	"strconv"
)

func main() {

	// 1、连接微服务
	conn, err := rpc.Dial("tcp", "127.0.0.1:8091")
	if err != nil {
		fmt.Println("微服务连接失败！！！", err)
		return
	}
	defer conn.Close()


	// 2、执行远程调用

	goods := Goods{
		Name: "商品111",
		Type: "吃的",
	}
	goods2 := Goods{
		Name: "成人用品",
		Type: "用的",
	}

	AddGoods(conn, goods)
	AddGoods(conn, goods2)
	GetGoodsList(conn)
	GetGoodsById(conn, 1)
	GetGoodsById(conn, 8)


}

func AddGoods(conn *rpc.Client, goods Goods) {

	body, err := json.Marshal(&goods)
	if err != nil {
		return
	}

	var response string

	err2 := conn.Call("GoodsService.AddGoods", string(body), &response)
	if err2 != nil {
		fmt.Println("GoodsService.AddGoods.error = ", err2)
		return
	}

	// 输出远程调用结果
	fmt.Println(response)
}

func GetGoodsList(conn *rpc.Client) {


	var response string

	err2 := conn.Call("GoodsService.GetGoodsList", "", &response)
	if err2 != nil {
		return
	}

	// 输出远程调用结果
	fmt.Println(response)
}


func GetGoodsById(conn *rpc.Client, id int) {

	idStr := strconv.Itoa(id)

	var response string

	err2 := conn.Call("GoodsService.GetGoodsById",  idStr, &response)
	if err2 != nil {
		return
	}

	// 输出远程调用结果
	fmt.Println(response)
}

// 定义商品结构体，便于json序列化
type Goods struct {
	Id   int
	Name string
	Type string
}