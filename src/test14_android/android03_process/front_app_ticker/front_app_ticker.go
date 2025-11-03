package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// GetForegroundApp 检测当前属于前台的app的包名
func GetForegroundApp() (string, error) {
	// 1. 执行 dumpsys 命令
	cmd := exec.Command("sh", "-c", "dumpsys activity activities | grep topResumedActivity")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("执行命令失败: %v", err)
	}

	output := out.String()
	if output == "" {
		return "", fmt.Errorf("未获取到前台活动信息")
	}

	// 2. 使用正则匹配包名
	re := regexp.MustCompile(`u\d+\s+([a-zA-Z0-9._]+)/`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return "", fmt.Errorf("未匹配到包名: %s", output)
	}

	pkg := strings.TrimSpace(matches[1])
	return pkg, nil
}

func main() {

	for range time.Tick(2 * time.Second) {
		pkg, err := GetForegroundApp()
		if err != nil {
			fmt.Println("检测失败:", err)
			return
		}
		fmt.Println("当前前台应用包名:", pkg)
	}

}
