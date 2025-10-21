package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	uploadDir = "./upload"
)

// gin实现流式请求分片上传视频
func main() {

	engine := gin.Default()
	// 添加 CORS 中间件 跨域 gin实现跨域 gin跨域
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		//AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowHeaders: []string{"*"},
		//ExposeHeaders:    []string{"Content-Length"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.POST("/upload", Upload)
	engine.GET("/download/:fileName", Download)

	engine.Run(":8080")
}

// Download 分片下载
func Download(ctx *gin.Context) {
	// HTTP Range 请求下载（推荐）
	//浏览器 / 客户端 发送带 Range: bytes=start-end 的请求，后端根据 Range 返回对应的字节流。
	//这是最标准的“断点续传、分片下载”方式，适合在线播放 / 大文件下载。
	fileName := ctx.Param("fileName")
	filePath := filepath.Join(uploadDir, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(404, gin.H{"msg": "文件不存在"})
		return
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	fileSize := fileStat.Size()

	// 检查是否带 Range 请求
	rangeHeader := ctx.GetHeader("Range")
	if rangeHeader == "" {
		// 没有 Range，则直接返回整个文件
		ctx.Header("Content-Disposition", "attachment; filename="+fileName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Length", strconv.FormatInt(fileSize, 10))
		io.Copy(ctx.Writer, file)
		return
	}

	// 解析 Range 格式: "bytes=start-end"
	var start, end int64
	fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
	//if end == 0 || end >= fileSize {
	//	end = fileSize - 1
	//}
	if start == 0 && end == 0 {
		// 前端第一次请求是：bytes=0-0，表示这个是为了获取文件大小的请求，实际开发中，这个可以不要，文件实际大小可有其他前置接口返回
		ctx.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
		ctx.Status(206) // Partial Content
		return
	}
	if end >= fileSize {
		end = fileSize - 1
	}
	if start > end || start >= fileSize {
		ctx.JSON(416, gin.H{"msg": "无效的 Range"})
		return
	}
	time.Sleep(1 * time.Second)

	chunkSize := end - start + 1
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Accept-Ranges", "bytes")
	ctx.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	ctx.Header("Content-Length", strconv.FormatInt(chunkSize, 10))
	ctx.Status(206) // Partial Content

	// 移动到 start 位置并读取
	file.Seek(start, 0)
	io.CopyN(ctx.Writer, file, chunkSize)
}

// Upload 分片上传
func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Printf("未读取到任何文件：%v \n", err)
		c.JSON(200, gin.H{"msg": "未读取到任何文件：" + err.Error()})
		return
	}

	fileName := c.PostForm("fileName")
	chunkIndex, _ := strconv.Atoi(c.PostForm("chunkIndex"))
	totalChunks, _ := strconv.Atoi(c.PostForm("totalChunks"))

	fmt.Printf("/upload, file: %v, fileHeader: %v, fileName: %s, chunkIndex: %d, totalChunks: %d \n", file, fileHeader.Size, fileName, chunkIndex, totalChunks)

	// 临时目录 判断目录是否存在

	tempDir := "./upload/tmp/" + md5Value(fileName)
	if _, err = os.Stat(tempDir); os.IsNotExist(err) {
		if err = os.MkdirAll(tempDir, 0755); err != nil {
			c.JSON(200, gin.H{"msg": "创建目录失败：" + err.Error()})
			return
		}
	}

	filePath := filepath.Join(tempDir, fileName+"."+strconv.Itoa(chunkIndex)+".tmp")
	fmt.Printf("上传的文件：%s \n", filePath)

	// 保存分片到临时文件
	writeFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(200, gin.H{"msg": "创建新文件失败Create：" + err.Error()})
		return
	}
	defer writeFile.Close()
	_, err = io.Copy(writeFile, file)
	if err != nil {
		c.JSON(200, gin.H{"msg": "创建新文件失败Copy：" + err.Error()})
		return
	}

	// 还不是最后一个分片则接收完这一个分片就可以退出了，如果是最后一个分片就还需要合并所有分片
	if chunkIndex < totalChunks-1 {
		fmt.Printf("分片[%d/%d]上传完毕! \n", chunkIndex, totalChunks)
		c.JSON(200, gin.H{
			"msg": "分片上传成功！",
		})
		// 测试断点续传
		//if chunkIndex == 3 {
		//	os.Exit(0)
		//}
		return
	}

	// 如果当前是最后一个分片，合并文件
	finnalFile, err := os.Create(filepath.Join(uploadDir, fileName))
	if err != nil {
		c.JSON(200, gin.H{"msg": "创建文件失败Create：" + err.Error()})
		return
	}
	defer finnalFile.Close()

	// 		按顺序合并
	for i := 0; i < totalChunks; i++ {
		mFilePath := filepath.Join(tempDir, fileName+"."+strconv.Itoa(i)+".tmp")
		if _, err := os.Stat(mFilePath); os.IsNotExist(err) {
			c.JSON(200, gin.H{"msg": fmt.Sprintf("缺少分片: %d", i)})
			return
		}
		tmpFile, err := os.Open(mFilePath)
		if err != nil {
			c.JSON(200, gin.H{"msg": "打开文件失败 Open：" + err.Error()})
			return
		}
		io.Copy(finnalFile, tmpFile)
		tmpFile.Close()
		fmt.Printf("分片[%d/%d]合并完毕! %s \n", i, totalChunks, mFilePath)
	}
	fmt.Printf("所有分片[%d/%d]合并完毕!!! \n", totalChunks, totalChunks)

	// 		清理分片文件
	//err = os.RemoveAll(tempDir)
	//if err != nil {
	//	fmt.Printf("临时文件移除失败! %v \n", err)
	//}
	c.JSON(200, gin.H{
		"msg": "文件上传成功！",
	})
	return
}

// 计算字符串的MD5
func md5Value(fileName string) string {
	hash := md5.Sum([]byte(fileName))
	return hex.EncodeToString(hash[:])
}
