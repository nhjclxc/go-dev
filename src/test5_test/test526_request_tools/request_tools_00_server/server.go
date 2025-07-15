package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 这里实现各类http接口的服务端
func main() {

	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePostJson)
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/delete", handleDelete)
	http.HandleFunc("/postForm", handlePostForm)
	http.HandleFunc("/postFile", handlePostFile)

	fmt.Println("🚀 HTTP Server running at http://localhost:8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func handlePostFile(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := request.Header.Get(`Token`)
	cookie := request.Header.Get(`Cookie`)
	fmt.Printf("handlePostFile, token = %s, cookie = %s \n", token, cookie)

	err := request.ParseMultipartForm(10 << 20) // 最大 10MB
	if err != nil {
		http.Error(response, "解析表单失败", http.StatusBadRequest)
		return
	}


	// 获取普通字段
	formData := map[string]string{}
	for key, values := range request.MultipartForm.Value {
		formData[key] = values[0]
	}

	// 读取文件
	file, header, err := request.FormFile("file")
	if err != nil {
		http.Error(response, "未找到文件字段: file", http.StatusBadRequest)
		return
	}

	// 保存上传的文件
	dst, err := os.Create("./uploaded_" + header.Filename)
	if err != nil {
		http.Error(response, "保存文件失败", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	fmt.Printf("handlePostFile 收到文件上传: %s\n", header.Filename)

	writeJSON(response, http.StatusOK, map[string]interface{}{
		"method": "POST",
		"success": "请求成功！！！222",
		"data": header.Filename,
	})


}

func handlePostForm(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Printf("Method Not Allowed \n")
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := request.Header.Get(`Token`)
	cookie := request.Header.Get(`Cookie`)
	fmt.Printf("handlePostForm, token = %s, cookie = %s \n", token, cookie)


	contentType := request.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		// ...从 r.MultipartForm 读取


		// 注意：提取form表单数据时，必须使用 request.ParseMultipartForm(10 << 20) 来解析表单数据
		// request.MultipartForm 默认是 nil，只有你手动调用 ParseMultipartForm() 后，它才会被解析并赋值。

		// 先解析 multipart/form-data（最大内存设为 10MB）
		//err := request.ParseMultipartForm(10 << 20) // 10MB
		err := request.ParseForm() // 👈 使用 ParseForm 而不是 ParseMultipartForm
		if err != nil {
			fmt.Printf("解析 multipart 表单失败: %s \n"+err.Error())
			http.Error(response, "解析 multipart 表单失败: "+err.Error(), http.StatusBadRequest)
			return
		}

		// 读取请求体
		formData := map[string]string{}
		for key, values := range request.MultipartForm.Value {
			formData[key] = values[0]
		}

		fmt.Printf("handlePostForm, formData = %#v \n", formData)
	} else {
		// ...从 r.Form 读取

		err := request.ParseForm()
		if err != nil {
			fmt.Printf("解析 application/x-www-form-urlencoded 表单失败: %s \n"+err.Error())
			http.Error(response, "解析 application/x-www-form-urlencoded 表单失败: "+err.Error(), http.StatusBadRequest)
			return
		}
		name := request.FormValue("name")
		age := request.FormValue("age")
		fmt.Printf("handlePostForm, name = %#v \n", name)
		fmt.Printf("handlePostForm, age = %#v \n", age)
	}

	writeJSON(response, http.StatusOK, map[string]interface{}{
		"method": "POST",
		"success": "请求成功！！！222",
		"data": "formData",
	})
}

func handleDelete(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodDelete {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := request.Header.Get(`Token`)
	cookie := request.Header.Get(`Cookie`)
	fmt.Printf("handleDelete, token = %s, cookie = %s \n", token, cookie)

	// 读取请求参数
	query := request.URL.Query()
	fmt.Printf("handlePut, query = %#v \n", query)

	writeJSON(response, http.StatusOK, map[string]interface{}{
		"method": "DELETE",
		"success": "请求成功！！！555",
		"data": query,
	})
}

func handlePut(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPut {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := request.Header.Get(`Token`)
	cookie := request.Header.Get(`Cookie`)
	fmt.Printf("handlePut, token = %s, cookie = %s \n", token, cookie)

	contentType := request.Header.Get(`Content-Type`)
	if !strings.Contains(contentType, "application/json"){
		http.Error(response, "contentType Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体
	var bodyMap map[string]any
	bodyByte, _ := io.ReadAll(request.Body)
	json.Unmarshal(bodyByte, &bodyMap)

	fmt.Printf("handlePut, bodyMap = %#v \n", bodyMap)


	writeJSON(response, http.StatusOK, map[string]interface{}{
		"method": "PUT",
		"success": "请求成功！！！333",
		"data": bodyMap,
	})

}

func handlePostJson(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := request.Header.Get(`Token`)
	cookie := request.Header.Get(`Cookie`)
	fmt.Printf("handlePostJson, token = %s, cookie = %s \n", token, cookie)

	contentType := request.Header.Get(`Content-Type`)
	if !strings.Contains(contentType, "application/json"){
		http.Error(response, "contentType Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体
	var bodyMap map[string]any
	bodyByte, _ := io.ReadAll(request.Body)
	json.Unmarshal(bodyByte, &bodyMap)

	fmt.Printf("handlePostJson, bodyMap = %#v \n", bodyMap)


	writeJSON(response, http.StatusOK, map[string]interface{}{
		"method": "POST",
		"success": "请求成功！！！222",
		"data": bodyMap,
	})
}

func handleGet(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token := request.Header.Get(`Token`)
	cookie := request.Header.Get(`Cookie`)
	fmt.Printf("handleGet, token = %s, cookie = %s \n", token, cookie)

	query := request.URL.Query()
	fmt.Printf("收到 GET 请求, 参数: %v\n", query)
	for key, val := range query {
		fmt.Printf("/get, key = %s, val = %v \n", key, val[0])
	}

	writeJSON(response, http.StatusOK, map[string]interface{}{
		"method": "GET",
		"success": "请求成功！！！111",
		"data": query,
	})
}


// 返回 JSON 响应的帮助函数
func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)


	fmt.Printf("\n\n")
}
