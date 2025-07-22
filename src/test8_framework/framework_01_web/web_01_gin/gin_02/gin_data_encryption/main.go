package main

import (
	"errors"
	"fmt"
	"gin_data_encryption/core"
	"gin_data_encryption/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	localConfig "gin_data_encryption/global"
)

// gin实现前后端接口对称加密(SM4)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !validateToken(token) {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
func validateToken(token string) bool {
	fmt.Printf("AuthMiddleware.validateToken = %s \n", token)
	return true
}


func main() {

	localConfig.GlobalViper = core.Viper() // 初始化Viper

	fmt.Printf("全局配置：%#v \n", localConfig.GlobalConfig)


	r := gin.Default()

	// 启用跨域支持
	r.Use(cors.Default())


	authGroup := r.Group("/api", AuthMiddleware())
	encryptionAuthGroup := authGroup.Group("/encryption", middleware.EncryptionMiddleware())
	encryptionAuthGroup.GET("/getUser", getUser)
	encryptionAuthGroup.POST("/postUser", postUser)
	encryptionAuthGroup.PUT("/putUser", putUser)
	encryptionAuthGroup.DELETE("/deleteUser/:ids", deleteUser)
	encryptionAuthGroup.POST("/uploadFile", uploadFile)


	r.Run(":8080")
}

func uploadFile(c *gin.Context) {
	_, header, _ := c.Request.FormFile("file")
	file, _ := UploadFile(header) // 文件上传后拿到文件路径

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo uploadFile " + file},
		},
	)
}
func  UploadFile(header *multipart.FileHeader) (string, error) {
	filePath, key, uploadErr := doUploadFile(header)
	fmt.Printf("key = %v, key = %v \n", filePath, key)
	return filePath, uploadErr
}

func doUploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll("file", os.ModePerm)
	if mkdirErr != nil {
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := "file" + "/" + filename
	filepath := "file" + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}


func getUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	username := c.Query("username")

	fmt.Printf("getUser，authorization = %s, username = %s \n", authorization, username)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo getUser " + time.Now().String()},
		},
	)
}

func postUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	var dto UserDto
	c.ShouldBindJSON(&dto)

	fmt.Printf("postUser，authorization = %s, dto = %v \n", authorization, dto)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo postUser"},
		},
	)
}

func putUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	var dto UserDto
	c.ShouldBindJSON(&dto)

	fmt.Printf("putUser，authorization = %s, dto = %v \n", authorization, dto)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo putUser"},
		},
	)
}

func deleteUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	idsStr := c.Param("ids")         // 获取路径参数字符串，如 "123,234,345,567"
	ids := strings.Split(idsStr, ",") // 切分成字符串数组

	fmt.Printf("deleteUser，authorization = %s, userIds = %v \n", authorization, ids)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  LogoutVo{Foo: "Foo"},
		},
	)
}

// 接口统一响应结构
type JsonResponse struct {
	Code    int    `json:"code"`    // 响应码
	Msg     string `json:"msg"`     // 失败消息
	Success bool   `json:"success"` // 是否操作成功，操作成功返回true，否则返回false
	Data    any    `json:"data"`    // 响应数据
}

type LogoutVo struct {
	Foo string `json:"foo"` // 测试的字段
}

type UserDto struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Foo      string `json:"foo"`      // 测试的字段
}
