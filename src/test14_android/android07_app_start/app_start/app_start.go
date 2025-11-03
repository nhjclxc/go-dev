package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// 检查主进程PID是否存在
func getPid(pkg string) string {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("pidof %s", pkg))
	out, _ := cmd.Output()
	return strings.TrimSpace(string(out))
}

// 检查Activity是否仍在系统中活跃
func isActivityAlive(pkg string) bool {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("dumpsys activity activities | grep %s", pkg))
	out, _ := cmd.Output()
	return strings.Contains(string(out), pkg)
}

func detectAppStatus(pkg string) string {
	pid := getPid(pkg)
	activityAlive := isActivityAlive(pkg)
	fmt.Println("detectAppStatus", pid, activityAlive)

	switch {
	case pid == "" && !activityAlive:
		return "🟥 已完全退出"
	case pid != "" && !activityAlive:
		return "🟧 残留进程（主进程或子进程存在，但界面已退出）"
	case pid != "" && activityAlive:
		return "🟩 正在运行"
	default:
		return "⚪ 状态未知"
	}
}

func main() {
	pkg := "com.feedying.live.mix"
	lastStatus := ""

	for range time.Tick(2 * time.Second) {
		status := detectAppStatus(pkg)
		fmt.Println("结果：", status, lastStatus)
		if status != lastStatus {
			fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), status)
			lastStatus = status
		}
	}
}

// com.feedying.live.mix; am force-stop cn.miguvideo.migutv; am force-stop cn.juqing.cesuwang_tv
// com.feedying.live.mix
// cn.miguvideo.migutv
// cn.juqing.cesuwang_tv
// ps | grep com.feedying.live.mix | grep -v grep
// adb shell ps | com.feedying.live.mix
// pidof cn.miguvideo.migutv
