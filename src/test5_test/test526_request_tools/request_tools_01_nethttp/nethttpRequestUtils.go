package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

// ContentType 常量定义
type ContentType string

const (
	ContentTypeJSONUTF8           ContentType = "application/json; charset=utf-8"
	ContentTypeFormURLEncodedUTF8 ContentType = "application/x-www-form-urlencoded; charset=utf-8"
	ContentTypeXMLUTF8            ContentType = "application/xml; charset=utf-8"
	ContentTypeTextPlainUTF8      ContentType = "text/plain; charset=utf-8"

	// ⚠️ multipart/form-data 不需要 charset，使用 multipart.NewWriter 设置完整 Content-Type
	ContentTypeMultipartForm ContentType = "multipart/form-data"
)

// 请求可选参数结构体
type RequestOptions struct {
	Headers map[string]string // 请求头
	Params  url.Values        // GET 请求参数
	JsonBody    io.Reader         // POST/PUT 请求体（仅支持 JSON）
	Form           url.Values    // POST/PUT 请求体（表单）
	Timeout        time.Duration // 单个请求超时时间
	ContentType    ContentType   // 请求体 Content-Type
	FileFieldsName string        // postForm时传输文件的字段
	File           *os.File      // postForm时传输文件数据
}

// 通用请求函数
func doRequest(method, rawURL string, opts RequestOptions) (string, error) {
	// 拼接 URL 参数
	if opts.Params != nil && len(opts.Params) > 0 {
		rawURL += "?" + opts.Params.Encode()
	}

	// 构建请求
	req, err := http.NewRequest(method, rawURL, opts.JsonBody)
	if err != nil {
		return "", fmt.Errorf("构建请求失败: %w", err)
	}

	// 添加 Headers
	for key, val := range opts.Headers {
		req.Header.Set(key, val)
	}

	// 设置 Content-Type
	if opts.ContentType != "" {
		req.Header.Set("Content-Type", string(opts.ContentType))
	}

	// 设置超时
	client := &http.Client{
		Timeout: opts.Timeout,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}

	return string(body), nil
}

func Get(url string, opts RequestOptions) (string, error) {
	return doRequest(http.MethodGet, url, opts)
}

func Post(url string, opts RequestOptions) (string, error) {
	return doRequest(http.MethodPost, url, opts)
}

func Put(url string, opts RequestOptions) (string, error) {
	return doRequest(http.MethodPut, url, opts)
}

func Delete(url string, opts RequestOptions) (string, error) {
	return doRequest(http.MethodDelete, url, opts)
}

// post支持文件上传
func postForm(targetURL string, opts RequestOptions) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 写入字段
	for key, val := range opts.Form {
		_ = writer.WriteField(key, val[0])
	}

	// 写入文件（如果有）
	if opts.File != nil {
		fieldName := "file"
		if opts.FileFieldsName != "" {
			fieldName = opts.FileFieldsName
		}
		part, err := writer.CreateFormFile(fieldName, opts.File.Name())
		if err != nil {
			return "", fmt.Errorf("创建文件字段失败: %w", err)
		}
		if _, err := io.Copy(part, opts.File); err != nil {
			return "", fmt.Errorf("拷贝文件内容失败: %w", err)
		}
	}

	// ✅ 必须关闭 writer
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("关闭 multipart writer 失败: %w", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", targetURL, body)
	if err != nil {
		return "", fmt.Errorf("创建 POST 请求失败: %w", err)
	}

	// 设置 Content-Type（自动带 boundary）
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 添加自定义 Header
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: opts.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("POST 请求失败: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	return string(respBody), nil
}

