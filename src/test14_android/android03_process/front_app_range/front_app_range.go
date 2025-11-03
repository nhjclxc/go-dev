package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"slices"
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

type AppEntryLeaveStatus struct {
	Package   string     `json:"package"`
	EntryTime *time.Time `json:"entryTime"`
	LeaveTime *time.Time `json:"leaveTime"`
}

func printMap(pkgMapList map[string][]*AppEntryLeaveStatus) {
	fmt.Printf("输出记录信息：")
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

// 监控指定pkg对应app进入app和离开app的时间
func main() {

	// ./front_app_range
	pkgs := flag.String("pkgs", "com.feedying.live.mix,cn.miguvideo.migutv,cn.juqing.cesuwang_tv", "Comma-separated package names (e.g. com.feedying.live.mix,cn.miguvideo.migutv)")
	flag.Parse()

	if *pkgs == "" {
		fmt.Println("Usage: ./front_app_range --pkgs=com.feedying.live.mix,cn.miguvideo.migutv")
		return
	}

	// 输出接收到的参数
	fmt.Println("接收到的参数：")
	pkgList := strings.Split(*pkgs, ",") // 如果需要，把包名解析成切片
	fmt.Println("解析后的包名列表:", pkgList)

	pkgMapList := make(map[string][]*AppEntryLeaveStatus)

	count := 0
	lastPackage := ""
	for range time.Tick(1 * time.Second) {
		pkg, err := GetForegroundApp()
		if err != nil {
			fmt.Println("检测失败:", err)
			return
		}
		fmt.Println("当前前台应用包名:", pkg)

		if lastPackage != pkg {

			flag1 := slices.Contains(pkgList, pkg)
			flag2 := slices.Contains(pkgList, lastPackage)
			if flag1 || flag2 {
				dealPkg := pkg
				if flag2 {
					dealPkg = lastPackage
				}

				t := NowInCST()
				if lis, ok := pkgMapList[dealPkg]; ok {
					lastEntity := lis[len(lis)-1]

					if flag1 {
						// 进入app
						newEntity := &AppEntryLeaveStatus{
							Package:   dealPkg,
							EntryTime: &t,
						}
						lis = append(lis, newEntity)
						pkgMapList[dealPkg] = lis
						// 开启流量监控
					}
					if flag2 {
						// 离开app
						lastEntity.LeaveTime = &t

						// 停止这一次的流量监控
					}
				} else {
					lis0 := make([]*AppEntryLeaveStatus, 0)
					lis0 = append(lis0, &AppEntryLeaveStatus{
						Package:   dealPkg,
						EntryTime: &t,
					})
					pkgMapList[dealPkg] = lis0
				}
			}
			lastPackage = pkg
		}
		count++

		if count%10 == 0 {
			printMap(pkgMapList)
		}
	}

}

// CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o front_app_range front_app_range.go
// adb -s 192.168.200.55:5555 push front_app_range /data/local/tmp
