package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

func handleResp(tag string, resp *resty.Response, err error) {
	if err != nil {
		fmt.Printf("[%s] 请求失败: %v\n", tag, err)
		return
	}
	fmt.Printf("[%s] 状态码: %d\n", tag, resp.StatusCode())
	fmt.Printf("[%s] 响应体: %s\n", tag, resp.String())
}



var client = resty.New()

func TestGet(t *testing.T) {
	requestUrl := "http://localhost:8080/get"

	getResp, err := client.R().
		SetHeader("Token", "resty-Get-Token").
		SetHeader("Cookie", "resty-Get-Cookie").
		SetQueryParam("name", "resty-Get-name").
		SetQueryParam("age", "resty-Get-age").
		Get(requestUrl)
	handleResp("GET", getResp, err)

}


func TestPost(t *testing.T) {
	requestUrl := "http://localhost:8080/post"

	body := map[string]string {
		"name": "resty-Post-name",
		"age": "resty-Post-age",
	}

	getResp, err := client.R().
		SetHeader("Token", "resty-Post-Token").
		SetHeader("Cookie", "resty-Post-Cookie").
		SetBody(body).
		Post(requestUrl)
	handleResp("POST", getResp, err)

}

func TestPut(t *testing.T) {
	requestUrl := "http://localhost:8080/put"

	body := map[string]string {
		"name": "resty-Put-name",
		"age": "resty-Put-age",
	}

	getResp, err := client.R().
		SetHeader("Token", "resty-Put-Token").
		SetHeader("Cookie", "resty-Put-Cookie").
		SetBody(body).
		Put(requestUrl)
	handleResp("Put", getResp, err)

}

func TestDelete(t *testing.T) {
	requestUrl := "http://localhost:8080/delete"

	getResp, err := client.R().
		SetHeader("Token", "resty-Delete-Token").
		SetHeader("Cookie", "resty-Delete-Cookie").
		SetQueryParam("name", "resty-Delete-name").
		SetQueryParam("age", "resty-Delete-age").
		Delete(requestUrl)
	handleResp("Delete", getResp, err)

}


func TestPostForm(t *testing.T) {
	requestUrl := "http://localhost:8080/postForm"

	body := map[string]string {
		"name": "resty-postForm-name",
		"age": "resty-postForm-age",
	}

	// POST 表单请求（application/x-www-form-urlencoded）
	getResp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		//SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Token", "resty-postForm-Token").
		SetHeader("Cookie", "resty-postForm-Cookie").
		SetFormData(body).
		Post(requestUrl)

	handleResp("postForm", getResp, err)
}


func TestPostFormFile(t *testing.T) {
	requestUrl := "http://localhost:8080/postFile"

	// POST 上传文件请求（multipart/form-data）
	postFileResp, err := client.R().
		SetHeader("Token", "resty-postFile-Token").
		SetHeader("Cookie", "resty-postFile-Cookie").
		SetFile("file", "go.mod"). // 确保 * 文件存在
		Post(requestUrl)
	handleResp("postFile", postFileResp, err)

}

/*
client.R().Get(url)
client.R().Post(url)
client.R().Put(url)
client.R().Delete(url)
client.R().Patch(url)
client.R().Head(url)



 */