package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

// 开一个 Goods 微服务
type GoodsService struct {
}

// 定义商品结构体，便于json序列化
type Goods struct {
	Id   int
	Name string
	Type string
}

var goodsMap map[int]*Goods = make(map[int]*Goods)

// AddGoods 新增商品详细
func (this *GoodsService) AddGoods(body string, response *string) error {

	var goods Goods = Goods{}
	// json 反序列化 body 参数
	err := json.Unmarshal([]byte(body), &goods)
	if err != nil {
		return err
	}

	fmt.Printf("获取到的参数：Name = %s, Tpye = %s \n", goods.Name, goods.Type)

	if goods.Name == "成人用品" {
		fmt.Println("不能添加18🈲用品！！！")
		return errors.New("不能添加18🈲用品！！！")
	}
	// 执行新增逻辑
	goods.Id = len(goodsMap) + 1
	goodsMap[goods.Id] = &goods
	fmt.Println("新增商品完成！！！")

	// 返回id
	marshal, err := json.Marshal(&goods)
	if err != nil {
		return err
	}
	*response = string(marshal)
	return nil
}

// GetGoodsList 查询商品列表
func (this *GoodsService) GetGoodsList(body string, response *string) error {

	// 反序列化参数

	// 查询指定参数的数据
	jsonByte, err := json.Marshal(goodsMap)
	if err != nil {
		return err
	}

	// 返回数据
	*response = string(jsonByte)
	return nil
}

// GetGoodsById 根据id查询商品详细信息
func (this *GoodsService) GetGoodsById(idStr string, response *string) error {
	id, _ := strconv.Atoi(idStr)

	goodsTemp := goodsMap[id]

	if goodsTemp == nil {
		return errors.New("指定id的商品不存在！！！")
	}

	marshal, err := json.Marshal(&goodsTemp)
	if err != nil {
		return err
	}

	*response = string(marshal)

	return nil
}

func main() {

	// 1、注册 RPC 服务，MicroRpcTcpServer 是服务名，在客户端调用的时候使用的就是这个
	err := rpc.RegisterName("GoodsService", new(GoodsService))
	if err != nil {
		return
	}

	// 2、设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8091")
	if err != nil {
		return
	}

	// 服务关闭之后，释放端口
	defer listener.Close()

	// 3、接受客户端连接
	for true {
		// 等待连接
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		// 开一个协程处理具体的业务
		// 会根据客户端的serviceMethod指定的方法，自动去分发处理方法
		go rpc.ServeConn(conn)
	}

}
