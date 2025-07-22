package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	localConfig "gin_data_encryption/global"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ========== 响应包装器 ==========
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // 缓存响应体
	return w.ResponseWriter.Write(b)
}

func EncryptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		encryptHeader := c.GetHeader(localConfig.GlobalConfig.Request.Encrypt.EncryptHeader)

		// 没有加密头或为 false，跳过加解密
		enabled, err := strconv.ParseBool(encryptHeader)
		if encryptHeader == "" || err != nil || !enabled {
			c.Next()
			return
		}

		// 密钥（16字节 = 128位）
		sm4Key := []byte(localConfig.GlobalConfig.Request.Encrypt.Key)

		// 以下实现加密与解密

		// 1、解密前端传来的数据，get，post，put，delete，文件？
		/// ...
		// ========= 解密 URL 参数 =========
		query := c.Request.URL.Query()
		for k, vs := range query {
			for i, v := range vs {
				if v != "" {
					plain, err := decryptSM4Hex(v, sm4Key)
					if err == nil {
						vs[i] = plain
					}
				}
			}
			query[k] = vs
		}
		c.Request.URL.RawQuery = query.Encode()

		// ========= 解密 Body =========
		if c.Request.Body != nil && c.Request.ContentLength > 0 {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil && len(bodyBytes) > 0 {

				getContentType := c.Request.Header.Get("Content-Type")
				if strings.Contains(getContentType, "application/json") {
					var reqMap map[string]string
					if err := json.Unmarshal(bodyBytes, &reqMap); err == nil {
						for rkey, val := range reqMap {
							decrypted, err := decryptSM4Hex(val, sm4Key)
							if err == nil {
								reqMap[rkey] = decrypted
							}
						}
					}

					marshal, _ := json.Marshal(&reqMap)
					c.Request.Body = io.NopCloser(strings.NewReader(string(marshal)))
					c.Request.ContentLength = int64(len(marshal))
				} else if strings.Contains(getContentType, "multipart/form-data") {
					// ...
					doMultipartFormData(c, sm4Key)
				}
			}
		}

		// ========= 捕获响应 =========
		wrapperWriter := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = wrapperWriter // 替换默认的 Writer

		// 2、进入实际请求
		c.Next()

		// 此时 c.Writer 就是你自己定义的 bw
		fmt.Println("原始响应内容：", wrapperWriter.body.String())

		// 3、加密响应数据

		// 拦截响应数据
		contentType := c.Writer.Header().Get("Content-Type")
		originBody := wrapperWriter.body.Bytes()

		if enabled && c.Writer.Status() == http.StatusOK && strings.Contains(contentType, "application/json") {
			// 加密响应
			encryptedData, err := encryptSM4Hex(string(originBody), sm4Key)
			if err != nil {
				// 如果加密失败，返回错误
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"msg":   "加密失败",
					"error": err.Error(),
				})
				return
			}

			c.Writer.WriteHeader(http.StatusOK)
			c.Header("Access-Control-Expose-Headers", localConfig.GlobalConfig.Request.Encrypt.EncryptHeader)
			c.Header(localConfig.GlobalConfig.Request.Encrypt.EncryptHeader, "true")

			// 清空响应体，重写加密后的数据（防止重复输出）
			c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			c.Writer.WriteHeaderNow()             // 手动设置状态码
			c.Writer.Write([]byte(encryptedData)) // 返回加密内容

			//// 清空原始响应并重写
			c.Writer = wrapperWriter.ResponseWriter
		}

	}
}

func doMultipartFormData(c *gin.Context, sm4Key []byte) {
	// 1. 解析 multipart 请求
	err := c.Request.ParseMultipartForm(32 << 20) // 限制最大内存 32MB
	if err != nil {
		log.Println("解析 multipart/form-data 失败:", err)
		return
	}

	// 2. 遍历字段（普通表单字段）
	for key, vals := range c.Request.MultipartForm.Value {
		for i, val := range vals {
			decrypted, err := decryptSM4Hex(val, sm4Key)
			if err == nil {
				vals[i] = decrypted
			}
		}
		c.Request.MultipartForm.Value[key] = vals
	}

	// 3. 文件部分（如你想要解密上传的文件内容，也可以做）：
	for _, fileHeaders := range c.Request.MultipartForm.File {
		for _, fh := range fileHeaders {
			uploadedFile, err := fh.Open()
			if err != nil {
				continue
			}
			defer uploadedFile.Close()

			// 读取并解密文件内容（示例用 SM4 解密流）
			decryptedReader, err := decryptReader(uploadedFile, sm4Key)
			if err != nil {
				continue
			}

			// 保存为临时文件，替换原上传内容（示意）
			tmpfile, err := os.CreateTemp("", "decrypted-*")
			if err != nil {
				continue
			}
			defer tmpfile.Close()

			io.Copy(tmpfile, decryptedReader)

			// TODO：将 `tmpfile` 内容重新绑定到 `c.Request.MultipartForm.File[fieldName]` —— Gin 不支持直接修改上传文件内容，需要额外封装
		}
	}

}
