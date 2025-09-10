package fingerprint_test

import (
	"github.com/gin-gonic/gin"
	"go-dev/src/test_best_practice/bp09_remote_login_notice/utils"
	"testing"
)

func TestName(t *testing.T) {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// 如果部署在反向代理后，确保设置可信代理（否则 c.ClientIP() 可能是代理）
	r.SetTrustedProxies([]string{"0.0.0.0/0"}) // 开发可用，生产请指定真实代理网段

	// 初始化 Fingerprinter，示例包含 proxy chain、query、并接收前端传来的 X-Client-FP
	fp := utils.NewFingerprinter(
		"your-secret-key-please-change",
		utils.WithIncludeProxyChain(true),
		utils.WithIncludeQuery(true),
		utils.WithCookies([]string{"sessionid"}),
		utils.WithExtraHeader("X-Client-FP"),
		utils.WithTruncate(32), // 截断为 32 字符 hex（可选）
		utils.WithDebug(false),
	)

	// 使用中间件：把指纹写到 context key "fingerprint" 并添加到响应 header
	r.Use(fp.Middleware("fingerprint", true))

	r.GET("/", func(c *gin.Context) {
		fpVal, _ := utils.GetFingerprintFromContext(c, "fingerprint")
		c.JSON(200, gin.H{
			"fingerprint": fpVal,
			"raw":         c.GetString("fingerprint_raw"), // 仅当 debug=true 时才有值
		})
	})

	r.Run(":8080")
}

func Test22(t *testing.T) {
	//assert.Equal()
}
