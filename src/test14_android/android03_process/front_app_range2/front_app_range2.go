package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var cst *time.Location

func init() {
	var err error
	cst, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		cst = time.FixedZone("CST", 8*60*60)
	}
}

// NowInCST 获取东八区时间
func NowInCST() time.Time {
	return time.Now().In(cst)
}

// GetForegroundApp 检测当前属于前台的app的包名
func GetForegroundApp1() (string, error) {
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

// GetForegroundApp 检测当前属于前台的app的包名
//
// Returns:
//   - string: 当前前台app的包名
//   - error: 异常错误
func GetForegroundApp() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. 执行 dumpsys 命令
	// 优先尝试用 su -c 执行 dumpsys
	cmdStr := fmt.Sprintf("dumpsys activity activities | grep topResumedActivity")
	outBytes, err := exec.CommandContext(ctx, "su", "-c", cmdStr).CombinedOutput()
	if err != nil {
		// 如果 su 执行失败，尝试不使用 su（可能当前进程已是 root 或在某些环境里可以直接执行）
		out2, err2 := exec.CommandContext(ctx, "sh", "-c", cmdStr).CombinedOutput()
		if err2 != nil {
			return "", fmt.Errorf("android box run GetForegroundApp, su error: %v; fallback sh error: %v; last output: %s", err, err2, string(out2))
		}
		outBytes = out2
	}

	output := string(outBytes)
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

type AppEntryLeaveStatus struct {
	Package   string     `json:"package"`
	EntryTime *time.Time `json:"entryTime"`
	LeaveTime *time.Time `json:"leaveTime"`
}

func printMap(pkgMapList map[string][]*AppEntryLeaveStatus) {
	fmt.Printf("输出记录信息：\n")
	for pkg, statusList := range pkgMapList {
		fmt.Printf("\t[%s]的进出记录：\n", pkg)
		for i, status := range statusList {
			enTime, leTime := "", ""
			if status.EntryTime != nil {
				enTime = status.EntryTime.Format("2006-01-02 15:04:05")
			}
			if status.LeaveTime != nil {
				leTime = status.LeaveTime.Format("2006-01-02 15:04:05")
			}

			fmt.Printf("\t\t 第 %d 次进入时间: %s, 离开时间: %s \n", i+1, enTime, leTime)
		}
	}
	fmt.Println()

}

// 监控所有pkg对应app进入app和离开app的时间

// CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o front_app_range2 front_app_range2.go
// adb -s 192.168.200.55:5555 push front_app_range2 /data/local/tmp
// ./front_app_range2

func main() {

	// app进出记录
	pkgMapList := make(map[string][]*AppEntryLeaveStatus)

	count := 0
	lastPackage := ""
	lastReportTime := NowInCST()
	//pkgList := []string{"mix", "mix", "mix", "mix", "mix", "ceshu", "ceshu", "ceshu", "ceshu", "ceshu", "mix", "mix", "mix", "mix"}
	for range time.Tick(1 * time.Second) {

		pkg, err := GetForegroundApp()
		if err != nil {
			fmt.Println("检测失败:", err)
			continue
		}
		//pkh := pkgList[count]
		fmt.Println("当前前台应用包名:", pkg)

		// 第一次执行
		if lastPackage == "" {
			lastPackage = pkg
			dealEnterLeavePkg(pkgMapList, pkg, true)
			continue
		}

		// 后续的进出
		if lastPackage != pkg {
			// 处理离开的 lastPackage
			dealEnterLeavePkg(pkgMapList, lastPackage, false)

			// 处理进入的 pkg
			dealEnterLeavePkg(pkgMapList, pkg, true)

			lastPackage = pkg
		}

		if time.Since(lastReportTime) > 10*time.Second {
			fmt.Println("超过15秒了，数据上报")
			printMap(pkgMapList)

			// 刷新上报时间
			lastReportTime = NowInCST()
		}
		count++
	}

}

// 处理进出记录
//
// Params:
//   - pkgMapList: 存储基础记录的map
//   - dealPkg: 当前处理的包
//   - enterFlag: 进出的标记
func dealEnterLeavePkg(pkgMapList map[string][]*AppEntryLeaveStatus, dealPkg string, enterFlag bool) {
	t := NowInCST()
	if lis, ok := pkgMapList[dealPkg]; ok {
		lastEntity := lis[len(lis)-1]

		// 判断是进入还是离开
		if enterFlag {
			lis = append(lis, &AppEntryLeaveStatus{
				Package:   dealPkg,
				EntryTime: &t,
			})
			pkgMapList[dealPkg] = lis
		} else {
			lastEntity.LeaveTime = &t
		}
	} else {
		lis0 := append(make([]*AppEntryLeaveStatus, 0), &AppEntryLeaveStatus{
			Package:   dealPkg,
			EntryTime: &t,
		})
		pkgMapList[dealPkg] = lis0
	}
}
