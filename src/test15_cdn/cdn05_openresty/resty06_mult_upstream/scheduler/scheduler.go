package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpStream struct {
	Node   string `json:"node" form:"node"`
	Name   string `json:"name" form:"name"`
	Host   string `json:"host" form:"host"`
	Port   string `json:"port" form:"port"`
	Weight int    `json:"weight" form:"weight"`
}

var (
	data map[string][]UpStream = make(map[string][]UpStream)

	//defaultUpStream UpStream = UpStream{
	//	Node:   "default",
	//	Name:   "default",
	//	Addr:   "127.0.0.1:8080",
	//	Weight: 100,
	//}
)

func main() {

	e := gin.Default()
	api := e.Group("/api")

	// http://127.0.0.1:9090/api/upstreams/{nodename}
	api.GET("/upstreams/:node", func(c *gin.Context) {
		node := c.Param("node")
		lists := data[node]
		//if len(lists) == 0 {
		//	lists = append(lists, defaultUpStream)
		//}

		c.JSON(http.StatusOK, lists)
	})

	// http://127.0.0.1:9090/api/set?node=node1&name=u&host=upstream&port=8080&weight=10
	// http://127.0.0.1:9090/api/set?node=node1&name=u1&host=upstream1&port=8081&weight=20
	// http://127.0.0.1:9090/api/set?node=node1&name=u2&host=upstream2&port=8082&weight=30
	// http://127.0.0.1:9090/api/set?node=node1&name=u3&host=upstream3&port=8083&weight=50

	// http://127.0.0.1:9090/api/set?node=node1&name=u&host=127.0.0.1&port=8080&weight=10
	// http://127.0.0.1:9090/api/set?node=node1&name=u1&host=127.0.0.1&port=8081&weight=20

	// http://127.0.0.1:9090/api/set?node=node1&name=u&host=192.168.201.74&port=8080&weight=10
	// http://127.0.0.1:9090/api/set?node=node1&name=u1&host=192.168.201.74&port=8081&weight=20
	api.GET("/set", func(c *gin.Context) {
		var req UpStream
		err := c.ShouldBindQuery(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "请求参数错误",
			})
			return
		}

		data[req.Node] = append(data[req.Node], req)

		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("操作成功: %s:%d", req.Host, req.Port),
		})

	})

	// http://127.0.0.1:9090/api/delete/{nodename}
	api.GET("/delete", func(c *gin.Context) {
		node := c.Param("node")
		delete(data, node)

		c.JSON(http.StatusOK, gin.H{
			"msg": "操作成功",
		})

	})

	e.Run(":9090")
}
