package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 模拟代理回源，
// 当前服务去源站文件地址下载文件到当前服务器，同时向客户端提供服务
// 在rs02_proxy01.go的基础上实现“分块缓存策略”（注意：大文件分片是cdn内部行为，返回给用户的文件始终为用户可直接使用的文件，切勿将分片文件返回给用户）

var origin = "/Users/lxc20250729/cdn/origin"    // 源站目录
var originDir = "/Users/lxc20250729/cdn/origin" // 源站目录
var regionDir = "/Users/lxc20250729/cdn/region" // 区域目录
var cacheDir = "/Users/lxc20250729/cdn/cache"   // 缓存目录

func proxyHandler2(w http.ResponseWriter, r *http.Request) {

	// 解析查询参数
	filename := r.URL.Query().Get("filename") // 获取 filename 参数的值
	if filename == "" {
		http.Error(w, "filename parameter is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Requested filename: %s\n", filename)

	// 收集某该文件的所有分片
	partFiles, _ := collectFileAllPart(cacheDir, filename)
	if len(partFiles) > 0 {
		// 先把有的分片返回给客户端
	}

	// -------
	// 先检查本地是否有该文件
	localPath := filepath.Join(cacheDir, filename)
	fmt.Println("localPath", localPath)

	// 文件存在，提供服务
	if _, err := os.Stat(localPath); err == nil {
		log.Printf("Cache hit: %s", localPath)
		http.ServeFile(w, r, localPath)
		return
	}

	log.Printf("Cache miss, fetching from origin: %s", origin)

	// 目标文件不存在，去源站拉，同时缓存到本地和给客户端提供服务
	originUrl, _ := url.Parse(origin)
	client := &http.Client{Timeout: 10 * time.Minute}
	resp, err := client.Get(originUrl.String())
	if err != nil {
		http.Error(w, "Failed to fetch from origin", http.StatusBadGateway)
		log.Printf("Error fetching origin: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Origin return non-200", http.StatusBadGateway)
		log.Printf("Origin return non-200, now status: %d", resp.StatusCode)
		return
	}

	// 创建缓存目录
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		http.Error(w, "MkdirAll cacheDir err: "+err.Error(), http.StatusBadGateway)
		log.Printf("MkdirAll cacheDir err: %v", err)
		return
	}

	// 创建本地文件用于缓存
	createFile, err := os.Create(localPath)
	if err != nil {
		http.Error(w, "os.Create err: "+err.Error(), http.StatusBadGateway)
		log.Printf("os.Create err: %v", err)
		return
	}
	defer createFile.Close()

	// 边缓存边写
	// 边读取边写入本地 + 返回客户端（流式传输，大文件友好）
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// 创建一个 MultiWriter，写入到MultiWriter的数据会同时写入到MultiWriter的多个目标。
	// mw里面的数据同时会写入w 和 createFile
	mw := io.MultiWriter(w, createFile)
	// 将源站返回的数据写入mw
	if _, err := io.Copy(mw, resp.Body); err != nil {
		log.Printf("error while copying to client & cache: %v", err)
		return
	}

	log.Printf("Finished caching and serving: %s", filename)

}

// SplitFile 分片大文件，分片文件，文件分片
// filePath: 待分片文件路径
// chunkSize: 每个分片大小（单位：字节）
// outputDir: 分片输出目录，如果为空，则使用源文件同目录
func SplitFile(localFilepath, outputDir string, chunkSize int64) {

	// 判断文件是否存在
	fileInfo, err := os.Stat(localFilepath)
	if err != nil {
		fmt.Println("文件不存在，分片失败！！！", err)
		return
	}

	// 读取文件
	file, err := os.Open(localFilepath)
	if err != nil {
		fmt.Println("读取文件失败！！！", err)
		return
	}

	// 创建输出目录
	if outputDir == "" {
		outputDir = filepath.Dir(localFilepath)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("MkdirAll cacheDir err: %v \n", err)
		return
	}

	fileName := filepath.Base(localFilepath)

	// 写分片
	var partNum int
	buffer := make([]byte, chunkSize)
	for {
		// 读每一个分片的数据
		read, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("读取文件数据失败！！！", err)
			return
		}
		if read == 0 {
			break
		}

		// 写每一个分片的数据
		outputFilename := fmt.Sprintf("%s.part%d", fileName, partNum)
		outputFilePath := filepath.Join(outputDir, outputFilename)
		partFile, err := os.Create(outputFilePath)
		if err != nil {
			fmt.Println("分片文件创建失败！！！", err)
			return
		}
		// [:read]最后一次可能不足一个分片，因此最后一个有多少数据就写多少数据
		if _, err := partFile.Write(buffer[:read]); err != nil {
			partFile.Close()
			fmt.Printf("写入分片文件失败: %s \n", err)
		}
		partFile.Close()

		partNum++
	}

	fmt.Printf("文件 %s 分片完成，总分片数: %d, 每片大小: %d bytes, 文件总大小: %d bytes\n",
		fileInfo.Name(), partNum, chunkSize, fileInfo.Size())
}

// MergeFile 合并分片文件
// filePath: 待分片文件路径
// chunkSize: 每个分片大小（单位：字节）
// outputDir: 分片输出目录，如果为空，则使用源文件同目录
func MergeFile(partFileDir, originFilename, outputDir string) {

	partFiles, _ := collectFileAllPart(partFileDir, originFilename)
	if len(partFiles) == 0 {
		fmt.Printf("没有找到分片文件")
		return
	}

	// 按数字顺序排序所有分片文件
	sort.Slice(partFiles, func(i, j int) bool {
		// 获取 part 后的数字
		a := strings.TrimPrefix(partFiles[i], originFilename+".part")
		b := strings.TrimPrefix(partFiles[j], originFilename+".part")
		ai, _ := strconv.Atoi(a)
		bi, _ := strconv.Atoi(b)
		return ai < bi
	})

	// 创建输出目录
	if outputDir == "" {
		outputDir = partFileDir
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("MkdirAll cacheDir err: %v \n", err)
		return
	}

	// 创建输出文件
	outFile, err := os.Create(filepath.Join(outputDir, originFilename))
	if err != nil {
		fmt.Printf("创建输出文件失败: %s \n", err)
		return
	}

	// 依次写入所有分片数据
	for _, partFile := range partFiles {
		partFilePath := filepath.Join(partFileDir, partFile)
		partFile, err := os.Open(partFilePath)
		if err != nil {
			fmt.Printf("读取分片文件[%s]失败: %s \n", partFilePath, err)
			return
		}
		// 将分片数据写入输出文件

		if _, err = io.Copy(outFile, partFile); err != nil {
			fmt.Printf("写入分片文件[%s]失败: %s \n", partFilePath, err)
			return
		}
		partFile.Close()
	}

	fmt.Printf("合并完成，输出文件: %s\n", filepath.Join(outputDir, originFilename))
}

// collectFileAllPart 收集某个目录下某个文件的所有分片
func collectFileAllPart(partFileDir string, originFilename string) ([]string, error) {
	// 读取目录下所有文件
	files, err := os.ReadDir(partFileDir)
	if err != nil {
		fmt.Printf("读取目录失败: %s \n", err)
		return nil, err
	}

	// 收集该文件的所有分片文件
	var partFiles []string
	for _, f := range files {
		if strings.HasPrefix(f.Name(), originFilename+".part") {
			partFiles = append(partFiles, f.Name())
		}
	}
	return partFiles, err
}

func proxyHandler(c *gin.Context) {
	filename := c.Query("filename")
	log.Println("filename = ", filename)
}

func main() {

	// http://127.0.0.1:8081/?filename=google_devices-tokay-12990991-8abda025.tgz
	engine := gin.Default()
	engine.GET("", proxyHandler)
	log.Println("CDN Proxy on :8081")
	engine.Run(":8081")

}
