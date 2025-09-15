package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"Phone" validate:"required,min=3,max=5"`
	Email    string `json:"email" validate:"required,email"`
}

func parseBindingError2(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		messages := make([]string, 0)
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("%s 为必填项", e.Field()))
			case "email":
				messages = append(messages, fmt.Sprintf("%s 不是有效的邮箱", e.Field()))
			default:
				messages = append(messages, fmt.Sprintf("%s 校验失败: %s", e.Field(), e.Tag()))
			}
		}
		return fmt.Sprintf("%v", messages)
	}
	return err.Error()
}

func getValidationErrorMsg2(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		msgs := make([]string, 0)
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				msgs = append(msgs, fmt.Sprintf("%s 为必填项", e.Field()))
			case "email":
				msgs = append(msgs, fmt.Sprintf("%s 不是有效邮箱", e.Field()))
			case "min":
				msgs = append(msgs, fmt.Sprintf("%s 长度不能小于 %s", e.Field(), e.Param()))
			case "max":
				msgs = append(msgs, fmt.Sprintf("%s 长度不能大于 %s", e.Field(), e.Param()))
			default:
				msgs = append(msgs, fmt.Sprintf("%s 不合法", e.Field()))
			}
		}
		return fmt.Sprintf("%v", msgs)
	}
	return err.Error()
}

func main01() {
	r := gin.Default()

	r.POST("/json", func(c *gin.Context) {
		//// binding 校验
		//var req UserRequest
		//if err := c.ShouldBindJSON(&req); err != nil {
		//	c.JSON(400, gin.H{"error": parseBindingError(err)})
		//	return
		//}
		//
		//// validate 触发校验
		//if err := validator.New().Struct(req); err != nil {
		//	c.JSON(400, gin.H{"error": getValidationErrorMsg(err)})
		//	return
		//}

		var req UserRequest
		if !BindJSONAndValidate(c, &req) {
			return // 绑定/校验失败，已统一返回错误
		}

		fmt.Printf("val %#v", req)

		c.JSON(200, gin.H{"data": req})
	})

	r.Run(":8080")
}

var validate = validator.New()

// BindJSONAndValidate 封装 JSON 绑定 + 校验
func BindJSONAndValidate(c *gin.Context, obj any) bool {
	// JSON 绑定
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parseBindingError(err)})
		return false
	}
	// validator 校验
	if err := validate.Struct(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getValidationErrorMsg(err)})
		return false
	}
	return true
}

// 解析 Gin binding 错误（可自定义友好提示）
func parseBindingError(err error) string {

	if errs, ok := err.(validator.ValidationErrors); ok {
		messages := make([]string, 0)
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("%s 为必填项", e.Field()))
			case "email":
				messages = append(messages, fmt.Sprintf("%s 不是有效的邮箱", e.Field()))
			default:
				messages = append(messages, fmt.Sprintf("%s 校验失败: %s", e.Field(), e.Tag()))
			}
		}
		return fmt.Sprintf("%v", messages)
	}
	return err.Error()
}

// 解析 validator 错误为友好提示
func getValidationErrorMsg(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		msgs := make([]string, 0)
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				msgs = append(msgs, fmt.Sprintf("%s 为必填项", e.Field()))
			case "email":
				msgs = append(msgs, fmt.Sprintf("%s 不是有效邮箱", e.Field()))
			case "min":
				msgs = append(msgs, fmt.Sprintf("%s 长度不能小于 %s", e.Field(), e.Param()))
			case "max":
				msgs = append(msgs, fmt.Sprintf("%s 长度不能大于 %s", e.Field(), e.Param()))
			default:
				msgs = append(msgs, fmt.Sprintf("%s 不合法", e.Field()))
			}
		}
		return fmt.Sprintf("%v", msgs)
	}
	return err.Error()
}
