下面是使用 [**Resty**](https://github.com/go-resty/resty) 库实现的 Go 客户端，用于请求你提供的 6 个接口：

https://github.com/go-resty/resty


https://resty.dev/


---

### ✅ Resty 安装

先确保你已经安装了 Resty：

```bash
go get github.com/go-resty/resty/v2
```

---

### ✅ 客户端完整示例

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	baseURL := "http://localhost:8080"

	// GET 请求
	getResp, err := client.R().
		SetQueryParam("name", "chatgpt").
		Get(baseURL + "/get")
	handleResp("GET", getResp, err)

	// POST JSON 请求
	postResp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"id":   1,
			"name": "test",
		}).
		Post(baseURL + "/post")
	handleResp("POST JSON", postResp, err)

	// PUT 请求
	putResp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"id":   1,
			"data": "updated",
		}).
		Put(baseURL + "/put")
	handleResp("PUT", putResp, err)

	// DELETE 请求
	deleteResp, err := client.R().
		SetQueryParam("id", "123").
		Delete(baseURL + "/delete")
	handleResp("DELETE", deleteResp, err)

	// POST 表单请求（application/x-www-form-urlencoded）
	postFormResp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"username": "gopher",
			"age":      "25",
		}).
		Post(baseURL + "/postForm")
	handleResp("POST Form", postFormResp, err)

	// POST 上传文件请求（multipart/form-data）
	postFileResp, err := client.R().
		SetFile("file", "test.txt"). // 确保 test.txt 文件存在
		Post(baseURL + "/postFile")
	handleResp("POST File", postFileResp, err)
}

func handleResp(tag string, resp *resty.Response, err error) {
	if err != nil {
		log.Printf("[%s] 请求失败: %v\n", tag, err)
		return
	}
	fmt.Printf("[%s] 状态码: %d\n", tag, resp.StatusCode())
	fmt.Printf("[%s] 响应体: %s\n", tag, resp.String())
}
```

---

### 📌 注意事项

* `SetBody` 支持传入 `map`、`struct`、`[]byte` 等，Resty 会自动处理为 JSON。
* `SetFile("file", "test.txt")` 会上传 `test.txt` 文件，请确保该文件存在于当前目录。
* 可以设置公共 Header 或日志功能，例如：

  ```go
  client.SetHeader("User-Agent", "resty-client").EnableTrace()
  ```

---

需要我帮你实现这些接口的服务器端（`handleGet`, `handlePostJson` 等）吗？
