package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// S：表示该进程是一个“睡眠进程”，即该进程正在等待某些事件或资源。
// R：表示进程正在运行。
// Z：表示该进程是“僵尸进程”（已终止，但尚未被父进程清理）。
// T：表示该进程正在停止或挂起。

// ps中进程状态的含义
// | 状态字母  | 英文全称                  | 中文含义                | 是否“活跃执行中”   |
// | ----- | --------------------- | ------------------- | ----------- |
// | **R** | Running               | 正在运行（CPU 正在执行）      | ✅ 是         |
// | **S** | Sleeping              | 可中断睡眠（等待 I/O、事件）    | ⚠️ 常见但不算执行中 |
// | **D** | Uninterruptible Sleep | 不可中断睡眠（一般是系统调用、I/O） | ⚠️          |
// | **Z** | Zombie                | 僵尸进程（已结束但未回收）       | ❌           |
// | **T** | Stopped               | 暂停（被 SIGSTOP 或调试）   | ❌           |
// | **I** | Idle                  | 内核线程空闲              | ⚙️          |
var processPool map[string]string = map[string]string{
	"R": "正在运行",
	"S": "可中断睡眠",
	"D": "不可中断睡眠",
	"Z": "僵尸进程",
	"T": "进程挂起",
	"I": "内核线程空闲",
}

type ProcessStatus struct {
	Package      string `json:"package"`
	Status       string `json:"status"`
	ProcessFlag  string `json:"processFlag"`
	FrontAppFlag bool   `json:"frontAppFlag"`
}

func GetForegroundApp() (string, error) {
	// 1. 执行 dumpsys 命令
	// dumpsys activity activities | grep "topResumedActivity"
	// dumpsys window windows | grep "topResumedActivity"
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

	pkgs := flag.String("pkgs", "", "Comma-separated package names (e.g. com.feedying.live.mix,cn.miguvideo.migutv)")
	flag.Parse()

	if *pkgs == "" {
		fmt.Println("使用方法：./process --pkgs=pkg1,pkg2,pkg3")
	}

	pkgList := strings.Split(*pkgs, ",")
	fmt.Println("读取到的pkgList = ", pkgList)

	frontPkg, err := GetForegroundApp()
	if err != nil {
		fmt.Println("检测失败:", err)
		return
	}
	fmt.Println("当前前台应用包名:", frontPkg)

	processStatusList, err := getRunningProcess(*pkgs, frontPkg)
	if err != nil {
		fmt.Println("获取运行进程失败:", err)
		return
	}
	fmt.Println("所有进程状态：")

	// 向服务器推送当前的进程信息processStatusList
	for _, status := range processStatusList {
		fmt.Printf("\t进程状态: %v \n", *status)
	}

}

func getRunningProcess(pkgs string, frontPkg string) ([]*ProcessStatus, error) {
	//   ps -A | grep -E "com\.feedying\.live\.mix|cn\.miguvideo\.migutv|cn\.juqing\.cesuwang_tv"
	//cmd := exec.Command("sh", "-c", fmt.Sprintf("ps -A | grep -E '%s'", *pkgs))
	pattern := strings.ReplaceAll(pkgs, ",", "|")
	cmd := exec.Command("sh", "-c", fmt.Sprintf("ps -A | grep -E '%s'", pattern))

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr // 显示标准错误
	if err := cmd.Run(); err != nil {
		fmt.Printf("failed to run shell command [ps -A]: %w \n", err)
		return nil, fmt.Errorf("failed to run shell command [ps -A]: %w ", err)
	}

	res := make([]*ProcessStatus, 0)

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		//fmt.Printf("fields: %v \n", fields)

		pkg := fields[len(fields)-1]
		processFlag := ""
		if strings.Contains(pkg, ":") {
			processFlag = "子进程"
		} else {
			processFlag = "主进程"
		}
		//fmt.Printf("%s: %s\n", processFlag, pkg)
		status := fields[len(fields)-2]
		//fmt.Printf("fields[len(fields)-2]: %v \n", status)

		res = append(res, &ProcessStatus{
			Package:      pkg,
			Status:       processPool[status],
			ProcessFlag:  processFlag,
			FrontAppFlag: pkg == frontPkg,
		})
	}

	return res, nil
}

/*
ps -A | grep -E 'com.feedying.live.mix,cn.miguvideo.migutv,cn.juqing.cesuwang_tv'


./process --pkgs=com.feedying.live.mix,cn.miguvideo.migutv,cn.juqing.cesuwang_tv


*/
