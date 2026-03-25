package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"reflect"
)

// gin批量注册接口
func main() {

	x := math.NaN()
	fmt.Println(x == x)        // false
	fmt.Println(x != x)        // false
	fmt.Println(math.IsNaN(x)) // true

	zz := math.Inf(1)
	fmt.Println(math.IsNaN(zz)) // true

	fmt.Println(reflect.TypeOf(x))

	if true {
		return
	}

	e := gin.Default()

	for i := range 10 {
		path := fmt.Sprintf("/api%d", i)
		e.GET(path, func(c *gin.Context) {
			msg := c.FullPath()
			c.JSON(200, gin.H{
				"message": msg,
			})
		})
	}

	fmt.Println("Server start")
	err := e.Run(":8080")
	if err != nil {
		fmt.Println("server err: ", err)
		return
	}
}
