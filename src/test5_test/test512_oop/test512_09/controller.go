package main

import "fmt"

type controller struct{}

func (c controller) SuccessResponse(data any) {
	res := make(map[string]any)
	res["code"] = 200
	res["success"] = true
	res["msg"] = "操作成功"
	res["data"] = data

	fmt.Println(res)
}
func (c controller) ErrorResponse(code int, msg string) {
	res := make(map[string]any)
	res["code"] = code
	res["success"] = false
	res["msg"] = fmt.Sprintf("操作失败：%s", msg)
	res["data"] = nil

	fmt.Println(res)
}

type TestController struct {
	controller
}

func (t TestController) Test1() {
	t.SuccessResponse("zz")
	t.ErrorResponse(500, "aaa")
}

// ==========

func SuccessResponse(data any) {
	res := make(map[string]any)
	res["code"] = 200
	res["success"] = true
	res["msg"] = "操作成功"
	res["data"] = data

	fmt.Println(res)
}
func ErrorResponse(code int, msg string) {
	res := make(map[string]any)
	res["code"] = code
	res["success"] = false
	res["msg"] = fmt.Sprintf("操作失败：%s", msg)
	res["data"] = nil

	fmt.Println(res)
}
