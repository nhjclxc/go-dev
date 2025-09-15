package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"reflect"
)

var validate = validator.New()

func BindAndValidateMiddleware(objType reflect.Type) gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "读取请求体失败"})
			c.Abort()
			return
		}

		// 缓存 body，供 Handler 或其他中间件使用
		c.Set("rawBody", bodyBytes)

		// 重新设置 c.Request.Body，让 Handler 可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 创建结构体实例
		obj := reflect.New(objType).Interface()

		// 绑定 JSON
		if err := c.ShouldBindBodyWithJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": parseBindingError(err)})
			c.Abort()
			return
		}

		// validator 校验
		if err := validate.Struct(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": getValidationErrorMsg(err)})
			c.Abort()
			return
		}

		// 校验成功，把结构体放到 context
		c.Set("req", obj)
		c.Next()
	}
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"Phone" validate:"required,min=3,max=5"`
	Email    string `json:"email" validate:"required,email"`
}

func main() {

	r := gin.Default()

	r.POST("/json",
		BindAndValidateMiddleware(reflect.TypeOf(UserRequest{})), func(c *gin.Context) {
			var req UserRequest
			//err := c.ShouldBindJSON(&req)
			err := c.ShouldBindBodyWithJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 从 context 获取已经绑定并校验成功的结构体
			//req := c.MustGet("req").(*UserRequest)
			fmt.Printf("val %#v", req)

			c.JSON(200, gin.H{"data": req})
		})

	r.Run(":8080")

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
