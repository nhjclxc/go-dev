package main

import "testing"



// 接口文档：https://pkg.go.dev/github.com/asaskevich/EventBus#Client
/*
type Client
	客户端——能够订阅远程事件总线的对象

	func NewClient(address, path string, eventBus Bus) *Client
		使用地址和服务器路径创建客户端对象

	func (client *Client) EventBus() Bus
	func (client *Client) Start() error
	func (client *Client) Stop()
	func (client *Client) Subscribe(topic string, fn interface{}, serverAddr, serverPath string)
	func (client *Client) SubscribeOnce(topic string, fn interface{}, serverAddr, serverPath string)

 */

// 基本使用
func TestClient1(t *testing.T) {

}

