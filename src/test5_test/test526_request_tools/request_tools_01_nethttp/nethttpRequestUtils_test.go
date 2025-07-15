package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
)

func Test1(t *testing.T) {

	params := url.Values{}
	params.Set("name", "alice")
	params.Set("age", "30")


	for key, val := range params {
		fmt.Println(key + "有参数" + val[0])
	}
	if params != nil && len(params) > 0 {
		fmt.Println("有参数")
	}

	finalURL := "baseURL" + "?" + params.Encode()

	fmt.Println(finalURL)
}



func TestGet(t *testing.T) {
	requestUrl := "http://localhost:8080/get"

	header := map[string]string {
		"Token": "TokenTokenToken",
		"Cookie": "CookieCookieCookie",
	}

	params := url.Values{}
	params.Set("name", "alice")
	params.Set("age", "30")

	responseBody, err := Get(requestUrl, RequestOptions{
		Headers: header,
		Params:  params,
		Timeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(requestUrl + " 请求出错 ", err)
		return
	}

	fmt.Println(requestUrl + " 请求成功！！！" + responseBody)

}


func TestPost(t *testing.T) {
	requestUrl := "http://localhost:8080/post"

	header := map[string]string {
		"Token": "TokenTokenToken-Post",
		"Cookie": "CookieCookieCookie-Post",
	}

	body := map[string]string {
		"name": "alice-Post",
		"age": "30-Post",
	}
	bodyByte, _ := json.Marshal(body)

	responseBody, err := Post(requestUrl, RequestOptions{
		Headers: header,
		JsonBody:  bytes.NewReader(bodyByte),
		Timeout: 5 * time.Second,
		ContentType: ContentTypeJSONUTF8,
	})
	if err != nil {
		fmt.Println(requestUrl + " 请求出错 ", err)
		return
	}

	fmt.Println(requestUrl + " 请求成功！！！" + responseBody)

}

func TestPut(t *testing.T) {
	requestUrl := "http://localhost:8080/put"

	header := map[string]string {
		"Token": "TokenTokenToken-put",
		"Cookie": "CookieCookieCookie-put",
	}

	body := map[string]string {
		"name": "alice-put",
		"age": "30-put",
	}
	bodyByte, _ := json.Marshal(body)

	responseBody, err := Put(requestUrl, RequestOptions{
		Headers: header,
		JsonBody:  bytes.NewReader(bodyByte),
		Timeout: 5 * time.Second,
		ContentType: ContentTypeJSONUTF8,
	})
	if err != nil {
		fmt.Println(requestUrl + " 请求出错 ", err)
		return
	}

	fmt.Println(requestUrl + " 请求成功！！！" + responseBody)

}

func TestDelete(t *testing.T) {
	requestUrl := "http://localhost:8080/delete"

	header := map[string]string {
		"Token": "TokenTokenToken-delete",
		"Cookie": "CookieCookieCookie-delete",
	}

	params := url.Values{}
	params.Set("name", "alice")
	params.Set("age", "30")

	responseBody, err := Delete(requestUrl, RequestOptions{
		Headers: header,
		Params:  params,
		Timeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println(requestUrl + " 请求出错 ", err)
		return
	}

	fmt.Println(requestUrl + " 请求成功！！！" + responseBody)

}


func TestPostForm(t *testing.T) {
	requestUrl := "http://localhost:8080/postForm"

	header := map[string]string {
		"Token": "TokenTokenToken-postForm",
		"Cookie": "CookieCookieCookie-postForm",
	}

	form := url.Values{}
	form.Set("name", "alice-postForm")
	form.Set("age", "30-postForm")

	responseBody, err := postForm(requestUrl, RequestOptions{
		Headers: header,
		Form:  form,
		Timeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(requestUrl + " 请求出错 ", err)
		return
	}

	fmt.Println(requestUrl + " 请求成功！！！" + responseBody)

}
func TestPostFormFile(t *testing.T) {
	requestUrl := "http://localhost:8080/postFile"

	header := map[string]string {
		"Token": "TokenTokenToken-postFile",
		"Cookie": "CookieCookieCookie-postFile",
	}

	form := url.Values{}
	form.Set("name", "alice-postFile")
	form.Set("age", "30-postFile")


	// 打开文件
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	responseBody, err := postForm(requestUrl, RequestOptions{
		Headers: header,
		Form:  form,
		Timeout: 5 * time.Second,
		FileFieldsName: "file",
		File: file,
	})
	if err != nil {
		fmt.Println(requestUrl + " 请求出错 ", err)
		return
	}

	fmt.Println(requestUrl + " 请求成功！！！" + responseBody)

}