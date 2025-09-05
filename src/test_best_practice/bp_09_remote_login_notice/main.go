package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-dev/src/test_best_practice/bp_09_remote_login_notice/utils"
	"net/http"
	"time"
)

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "操作成功！",
		"data": data,
	})
}
func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  fmt.Sprintf("操作失败: %s", err.Error()),
	})
}

type LoginInfoReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type LoginInfo struct {
	Username      string `json:"username" form:"username"`
	LastIP        string
	LastUserAgent string
	LastLogin     time.Time
}

// map[username]password
var userDB map[string]string = map[string]string{
	"user1": "pwd1",
	"user2": "pwd2",
	"a":     "a",
}

// map[username]ip
var userLastLoginInfo map[string]LoginInfo = make(map[string]LoginInfo)

// 异地登陆通知功能
func main() {

	r := gin.Default()

	// 使用一个get方法，便于在多设备测试，其他设备发送get不好操作
	//r.POST("/login", func(c *gin.Context) {
	r.GET("/login", func(c *gin.Context) {
		var req LoginInfoReq
		if err := c.ShouldBindQuery(&req); err != nil {
			ErrorResponse(c, 500, err)
			return
		}

		pwd, ok := userDB[req.Username]
		if !ok {
			ErrorResponse(c, 500, fmt.Errorf("用户不存在"))
			return
		}
		if pwd != req.Password {
			ErrorResponse(c, 500, fmt.Errorf("密码错误"))
			return
		}

		fmt.Println("GetUserAgent", utils.GetUserAgent(c))
		fmt.Println("GetClientIP", utils.GetClientIP(c))
		fmt.Println("GetProxyChain", utils.GetProxyChain(c))

		// 判断是不是异地登陆
		loginUserAgent := utils.GetUserAgent(c)
		loginIp := utils.GetClientIP(c)
		lastLoginInfo, ok := userLastLoginInfo[req.Username]
		if ok {
			// 有上次的登陆信息，要判断是不是异地登陆，如果没有则表示第一次登陆
			if loginUserAgent != lastLoginInfo.LastUserAgent {
				// 是异地登陆，提醒
				// send user email to notice ...
				ErrorResponse(c, 500, fmt.Errorf("当前账号是新设备登陆"))
				return
			}
			if loginIp != lastLoginInfo.LastIP {
				// 是异地登陆，提醒
				// send user email to notice ...
				ErrorResponse(c, 500, fmt.Errorf("当前账号是异地登陆"))
				return
			}
		}

		// 将本次登陆的ip存起来，便于下次登陆的时候判断是不是异地登陆
		userLastLoginInfo[req.Username] = LoginInfo{
			Username:      req.Username,
			LastIP:        loginIp,
			LastUserAgent: loginUserAgent,
			LastLogin:     time.Now(),
		}

		token, err := utils.GenerateToken2(req.Username, time.Duration(30)*time.Minute)
		if err != nil {
			ErrorResponse(c, 500, fmt.Errorf("token生成失败"))
			return
		}
		res := make(map[string]any)
		res["username"] = req.Username
		res["token"] = token
		SuccessResponse(c, res)
	})
	r.Run(":8080")
}

// http://192.168.200.149:8080/login?username=user1&password=pwd1
// http://127.0.0.1:8080/login?username=user1&password=pwd1
// http://localhost:8080/login?username=user1&password=pwd1
