package core

import (
	"fmt"
	"log_03_zap/global"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Cutter 实现 io.Writer 接口
// 用于日志切割, strings.Join([]string{director,layout, formats..., level+".log"}, os.PathSeparator)
type Cutter struct {
	level        string        // 日志级别(debug, info, warn, error, dpanic, panic, fatal)
	layout       string        // 时间格式 2006-01-02 15:04:05
	formats      []string      // 自定义参数([]string{Director,"2006-01-02", "business"(此参数可不写), level+".log"}
	director     string        // 日志文件夹
	retentionDay int           //日志保留天数
	file         *os.File      // 文件句柄
	mutex        *sync.RWMutex // 读写锁
}

type CutterOption func(*Cutter)

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

// CutterWithFormats 格式化参数
func CutterWithFormats(format ...string) CutterOption {
	return func(c *Cutter) {
		if len(format) > 0 {
			c.formats = format
		}
	}
}

func NewCutter(director string, level string, retentionDay int, options ...CutterOption) *Cutter {
	rotate := &Cutter{
		level:        level,
		director:     director,
		retentionDay: retentionDay,
		mutex:        new(sync.RWMutex),
	}
	for i := 0; i < len(options); i++ {
		options[i](rotate)
	}
	return rotate
}

// Write satisfies the io.Writer interface. It writes to the
// appropriate file handle that is currently being used.
// If we have reached rotation time, the target file gets
// automatically rotated, and also purged if necessary.
func (c *Cutter) Write0(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()
	length := len(c.formats)
	values := make([]string, 0, 3+length)
	values = append(values, c.director)
	if c.layout != "" {
		values = append(values, time.Now().Format(c.layout))
	}
	for i := 0; i < length; i++ {
		values = append(values, c.formats[i])
	}
	values = append(values, c.level+".log")
	filename := filepath.Join(values...)
	director := filepath.Dir(filename)
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err
	}
	err = removeNDaysFolders(c.director, c.retentionDay)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}



// 以下方法当单个log文件过最大log文件限制时，会自动在log后面加上标号1、2、3...
// 如果文件不需要加标号，则将上面的Write0改为Write，将本方法Write改为其他即可
// 注意：由于使用了读写锁，因此文件加标号会消费一定的性能
func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// 拼接基础路径
	length := len(c.formats)
	values := make([]string, 0, 3+length)
	values = append(values, c.director)
	if c.layout != "" {
		values = append(values, time.Now().Format(c.layout))
	}
	for i := 0; i < length; i++ {
		values = append(values, c.formats[i])
	}
	baseDir := filepath.Join(values...)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		return 0, err
	}

	// 清理历史目录
	if err := removeNDaysFolders(c.director, c.retentionDay); err != nil {
		return 0, err
	}

	// 自动编号文件路径
	logFilePath := ""
	//maxSize := int64(200 * 1024 * 1024) // 200MB
	maxSize0 := global.GlobalConfig.Zap.MaxSize
	if maxSize0 == 0 {
		maxSize0 = 200
	}
	maxSize := int64(maxSize0 * 1024 * 1024) // maxSize0单位为M
	for i := 1; ; i++ {
		tempPath := filepath.Join(baseDir, time.Now().Format("2006-01-02")+fmt.Sprintf(".%d.log", i))
		fi, err := os.Stat(tempPath)
		if err != nil { // 文件不存在
			logFilePath = tempPath
			break
		}
		if fi.Size() < maxSize {
			logFilePath = tempPath
			break
		}
	}

	// 打开并写入
	c.file, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = c.file.Close()
		c.file = nil
	}()

	return c.file.Write(bytes)
}


func (c *Cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.file != nil {
		return c.file.Sync()
	}
	return nil
}

// 增加日志目录文件清理 小于等于零的值默认忽略不再处理
func removeNDaysFolders(dir string, days int) error {
	if days <= 0 {
		return nil
	}
	cutoff := time.Now().AddDate(0, 0, -days)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.ModTime().Before(cutoff) && path != dir {
			err = os.RemoveAll(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
