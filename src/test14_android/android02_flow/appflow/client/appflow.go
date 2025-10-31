package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// AppStat 用于记录每一个app的流量
type AppStat struct {
	Package      string
	UID          string
	RxPrev       int64
	TxPrev       int64
	RxAccum      int64 // 5分钟累计
	TxAccum      int64 // 5分钟累计
	LastReportTs time.Time
}

// TrafficData 上报数据结果
type TrafficData struct {
	Pkg        string    `json:"pkg"`        // 包名
	RxTraffic  int64     `json:"rxtraffic"`  // 接收字节数
	TxTraffic  int64     `json:"txtraffic"`  // 发送字节数
	ReportTime time.Time `json:"reportTime"` // 上报时间（可选）
}

// getUID根据app package name 获取这个app对应的应用id（唯一id）用于查询流量
func getUID(pkg string) (string, error) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("dumpsys package %s | grep -E 'userId=|appId=|uid='", pkg))
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get uid for %s: %v", pkg, err)
	}
	re := regexp.MustCompile(`(\d{4,6})`)
	m := re.FindStringSubmatch(string(out))
	if len(m) > 1 {
		return m[1], nil
	}
	return "", fmt.Errorf("no uid found for %s", pkg)
}

// getTraffic 根据 UID 获取总接收字节数和总发送字节数
// rxTotal 接收总字节数 Byte
// txTotal 发送总字节数 Byte
func getTraffic(uid string) (rxTotal, txTotal int64, err error) {
	// 直接过滤 UID，减少输出
	//cmd := exec.Command("sh", "-c", fmt.Sprintf("dumpsys netstats | grep 'uid=%s'", uid))
	cmd := exec.Command("sh", "-c", fmt.Sprintf("dumpsys netstats | sed 's/^[ \\t]*//' | grep '^%s '", uid))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr // 显示标准错误
	if err := cmd.Run(); err != nil {
		return 0, 0, fmt.Errorf("failed to run dumpsys netstats: %v", err)
	}

	rxTotal = 0
	txTotal = 0
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		if fields[0] == uid {
			// 格式1: UID 开头的行 → UID rxBytes rxPackets txBytes txPackets
			// [10096 2170864477 1600852 231078772 252088]
			rx, _ := strconv.ParseInt(fields[1], 10, 64)
			tx, _ := strconv.ParseInt(fields[3], 10, 64)
			rxTotal += rx
			txTotal += tx
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("failed to run scanner : %v", err)
	}

	return rxTotal, txTotal, nil
}

// 上报流量数据到服务器
func reportTraffic(reportURL string, data []*TrafficData) {
	jsonData, _ := json.Marshal(data)
	fmt.Printf("[INFO] 上报参数: jsonData=%s\n", string(jsonData))
	resp, err := http.Post(reportURL+"/client/traffic", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("[ERROR] 上报失败: %v", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[ERROR] 读取响应失败: %v\n", err)
		return
	}

	fmt.Printf("[INFO] 上报完成，HTTP %d, 响应: %s\n", resp.StatusCode, string(body))

	var result struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Count int    `json:"count"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("[WARN] JSON解析失败: %v\n", err)
		return
	}

	fmt.Printf("[INFO] 上报结果: code=%d msg=%s count=%d\n", result.Code, result.Msg, result.Count)
}

func main() {
	pkgList := flag.String("pkgs", "", "Comma-separated package names (e.g. com.feedying.live.mix,cn.miguvideo.migutv)")
	interval := flag.Int("interval", 3, "Sampling interval in seconds")
	reportURL := flag.String("server", "http://192.168.201.167:8080", "Server base URL for reporting")
	reportCycle := flag.Int("report-cycle", 300, "Reporting interval in seconds (default 300 = 5min)")
	flag.Parse()

	if *pkgList == "" {
		fmt.Println("Usage: ./appflow --pkgs=com.feedying.live.mix,cn.miguvideo.migutv [--server=http://192.168.201.167:8080]")
		return
	}

	// 输出接收到的参数
	fmt.Println("接收到的参数：")
	fmt.Println("Packages:", *pkgList)
	fmt.Println("Sampling interval (s):", *interval)
	fmt.Println("Server URL:", *reportURL)
	fmt.Println("Report cycle (s):", *reportCycle)
	pkgs := strings.Split(*pkgList, ",") // 如果需要，把包名解析成切片
	fmt.Println("解析后的包名列表:", pkgs)

	var apps []AppStat
	for _, pkg := range pkgs {
		pkg = strings.TrimSpace(pkg)
		uid, err := getUID(pkg)
		if err != nil {
			fmt.Printf("[WARN] %v\n", err)
			continue
		}
		apps = append(apps, AppStat{Package: pkg, UID: uid, LastReportTs: time.Now()})
		fmt.Printf("[INFO] %s → UID %s\n", pkg, uid)
	}

	if len(apps) == 0 {
		fmt.Println("No valid package found.")
		return
	}

	fmt.Printf("\n Start monitoring every %d seconds, reporting every %d seconds...\n\n", *interval, *reportCycle)

	first := true
	count := 1
	lastReportTs := time.Now()
	for {
		for i := range apps {
			// 获取引用
			app := &apps[i]
			// 获取这个app的总计流量
			rx, tx, err := getTraffic(app.UID)
			if err != nil {
				fmt.Printf("[ERR] getTraffic error %s: %v\n", app.Package, err)
				continue
			}
			fmt.Printf("[INFO] [%s] now traffic total, ↓ %v Byte  ↑ %v Byte. \n", app.Package, rx, tx)

			// 记录第一次的值，便于下次计算差值
			if first {
				app.RxPrev = rx
				app.TxPrev = tx
				continue
			}

			// 3s的差值
			rxDelta := rx - app.RxPrev
			txDelta := tx - app.TxPrev
			fmt.Printf("[%s] rxDelta ↓ %v B, txDelta ↑ %v B\n", app.Package, rxDelta, txDelta)

			// 这三秒的速率
			rxRate := float64(rxDelta) / 1024.0 / float64(*interval)
			txRate := float64(txDelta) / 1024.0 / float64(*interval)
			fmt.Printf("[%s] rate ↓ %.2f KB/s  ↑ %.2f KB/s\n", app.Package, rxRate, txRate)

			// 差值累计
			app.RxAccum += rxDelta
			app.TxAccum += txDelta
			fmt.Printf("[%s] 差值累计, app.RxAccum ↓ %v Byte  app.TxAccum ↑ %v Byte. \n", app.Package, app.RxAccum, app.TxAccum)

			// 记录当前值，用于下次计算
			app.RxPrev = rx
			app.TxPrev = tx
		}
		first = false

		// 检查是否到上报时间
		// lastReportTs到当前时间的差值是多少秒
		if time.Since(lastReportTs) >= time.Duration(*reportCycle)*time.Second {
			var reportData []*TrafficData
			for i := range apps {
				app := &apps[i] // 获取引用
				reportData = append(reportData, &TrafficData{
					Pkg:        app.Package,
					RxTraffic:  app.RxAccum,
					TxTraffic:  app.TxAccum,
					ReportTime: time.Now(),
					//ReportTime: time.Now().Format(time.RFC3339),
				})
				// 清空累计数据
				app.RxAccum = 0
				app.TxAccum = 0
			}

			reportTraffic(*reportURL, reportData)
			lastReportTs = time.Now()
		}

		count++
		fmt.Printf("---------- 循环 %d ----------\n", count)
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

/*
监控每一个app每隔5分钟的流量统计，并且上报服务器

1、获取包名
HYLive，com.feedying.live.mix
咪视界，cn.miguvideo.migutv
测速网tv，cn.juqing.cesuwang_tv

2、获取某个app的uid
dumpsys package com.feedying.live.mix | grep -E "userId=|uid=|appId="
dumpsys package cn.miguvideo.migutv | grep -E "userId=|uid=|appId="
dumpsys package cn.juqing.cesuwang_tv | grep -E "userId=|uid=|appId="

3、根据uid获取流量
dumpsys netstats | grep 10106

---------------------------------------------------------------

GOOS=android GOARCH=arm CGO_ENABLED=1 go build -o appflow appflow.go
adb push appflow /data/local/tmp/
adb shell
cd /data/local/tmp
chmod +x appflow
./appflow --pkgs=cn.miguvideo.migutv,com.feedying.live.mix \
         --server=http://192.168.201.167:8080 \
         --interval=3 --report-cycle=30

./appflow --interval=3 --report-cycle=15 --pkgs=cn.miguvideo.migutv,com.feedying.live.mix,cn.juqing.cesuwang_tv --server=http://192.168.201.167:8080
./appflow --interval=3 --report-cycle=15 --pkgs=cn.juqing.cesuwang_tv --server=http://192.168.201.167:8080
./appflow --interval=15 --report-cycle=15 --pkgs=cn.juqing.cesuwang_tv --server=http://192.168.201.167:8080


```130|FY928X-K:/data/local/tmp $ dumpsys netstats | grep '10096'
    uid=10096 set=1
    10096 2170864477 1600852 231078772 252088
    16 wlan0 0x0 10096 1 20986 251 826245 570```每一列分别表示什么？

UID + rxBytes + rxPackets + txBytes + txPackets → 累计总流量
| 列号 | 值示例        | 含义                   |
| -- | ---------- | -------------------- |
| 1  | 10096      | UID                  |
| 2  | 2170864477 | **总接收字节数（rxBytes）**  |
| 3  | 1600852    | **总接收包数（rxPackets）** |
| 4  | 231078772  | **总发送字节数（txBytes）**  |
| 5  | 252088     | **总发送包数（txPackets）** |


接口 + UID + tag + set + rx/tx bytes + rx/tx packets → 可以区分前台/后台流量
| 列号 | 值示例    | 含义                                             |
| -- | ------ | ---------------------------------------------- |
| 1  | 16     | **Index/iface ID**（内部统计序号）                     |
| 2  | wlan0  | **网络接口名**（iface）                               |
| 3  | 0x0    | **tag**（应用流量标签，一般是 0x0）                        |
| 4  | 10096  | **UID**（对应应用的用户 ID）                            |
| 5  | 1      | **set**（0=DEFAULT, 1=FOREGROUND, 2=BACKGROUND） |
| 6  | 20986  | **rxBytes**（接收字节数）                             |
| 7  | 251    | **rxPackets**（接收数据包数）                          |
| 8  | 826245 | **txBytes**（发送字节数）                             |
| 9  | 570    | **txPackets**（发送数据包数）                          |



*/
