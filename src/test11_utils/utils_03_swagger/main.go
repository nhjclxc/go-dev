package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "go-dev/src/test11_utils/utils_03_swagger/docs"
	"net/http"
)

// go get github.com/swaggo/swag/cmd/swag
// go install github.com/swaggo/swag/cmd/swag
// go get github.com/swaggo/gin-swagger
// go get github.com/swaggo/files

// https://github.com/swaggo/gin-swagger
// https://swaggo.github.io/swaggo.io/declarative_comments_format/
func main() {
	r := gin.Default()
	r.POST("/login", login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8282")
}

// @登录
// @Description 登录接口
// @Accept  json
// @Produce json
// @Param   username   path    string     true        "登录用户名"
// @Param   password   path    string     true        "登录密码"
// @Success 200 {string} string    "ok"
// @Router /login [post]
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, "Hello world "+username+"_"+password)
}

/*
启动操作步骤：
	1、在执行完'go install github.com/swaggo/swag/cmd/swag'命令之后，在bin目录下会生成一个 'swag.exe'
	2、将上述 'swag.exe' 移动到和目前这个main.go的同级目录下
	3、在main.go的cmd里面执行 ‘swag init’ 命令，在这个目录下会生成一个 'doc'文件夹，里面包含docs.go、swagger.json、swagger.yaml
		swag init命令输出如下：[img.png](./img.png)
		```
			D:\code\go\go-dev\src\test11_utils\utils_03_swagger>swag init
			2025/04/19 21:00:56 Generate swagger docs....
			2025/04/19 21:00:56 Generate general API Info, search dir:./
			2025/04/19 21:01:13 create docs.go at docs/docs.go
			2025/04/19 21:01:13 create swagger.json at docs/swagger.json
			2025/04/19 21:01:13 create swagger.yaml at docs/swagger.yaml
		```
	4、go run main.go 启动，浏览器访问：http://127.0.0.1:8282/swagger/index.html
		发现页面显示：[img_1.png](./img_1.png)
			```
				Failed to load API definition.
				Fetch error
				Internal Server Error doc.json
			```
		这是因为本go文件没有导入第3步生成的docs文件夹里面的内容，将 _ "go-dev/src/test11_utils/utils_03_swagger/docs" 添加到imports下面即可
	5、重新启动 go run main.go，浏览器访问：http://127.0.0.1:8282/swagger/index.html，效果如[img_2.png](./img_2.png)
	注意：如果swagger注解内容发送了变化，那么必须重新执行第3步，重新生成swagger文档在重启才能生效
*/
