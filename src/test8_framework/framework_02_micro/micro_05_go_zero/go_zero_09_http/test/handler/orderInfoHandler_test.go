package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpc"
	"go_zero_09_http/internal/types"
	"io"
	"net/http"
	"testing"
)

var domain = "http://127.0.0.1:8899"

// http接口测试：test文件夹是本项目专门用于测试接口的文件夹，类似于SpringBoot的test文件夹
// 1、先启动主服务
// 2、启动测试地址

// size
func TestOrderInfoHandler(t *testing.T) {
	var req types.OrderInfoReq = types.OrderInfoReq{
		OrderId:     1024,
	}

	// http://127.0.0.1:8899/order/info?orderId=666
	resp, err := httpc.Do(context.Background(), http.MethodGet, domain+"/order/info", req)
	//http.MethodPost
	//http.MethodPut
	//http.MethodDelete
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	responseStr := string(bodyBytes)
	fmt.Println("响应内容:", responseStr)

	var result types.OrderInfoResp
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}
	fmt.Println("完整响应:", result)



}
