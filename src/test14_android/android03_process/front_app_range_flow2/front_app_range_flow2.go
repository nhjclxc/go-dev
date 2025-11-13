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
	"slices"
	"strconv"
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
		return "", fmt.Errorf("执行检测前台app命令失败: %v", err)
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

// AppFlowStatus 用于记录每一个app每一次进出的流量
type AppFlowStatus struct {
	Package      string     `json:"package"`
	UID          string     `json:"uid"`
	EntryTime    *time.Time `json:"entryTime"`
	EntryRxTotal int64      `json:"entryRxTotal"`
	EntryTxTotal int64      `json:"entryTxTotal"`
	LeaveTime    *time.Time `json:"leaveTime"`
	LeaveRxTotal int64      `json:"leaveRxTotal"`
	LeaveTxTotal int64      `json:"leaveTxTotal"`
	RxAccum      int64      `json:"rxAccum"`
	TxAccum      int64      `json:"txAccum"`
}

func (status *AppFlowStatus) ToString() string {
	enTime, leTime := "", ""
	if status.EntryTime != nil {
		enTime = status.EntryTime.Format("2006-01-02 15:04:05")
	}
	if status.LeaveTime != nil {
		leTime = status.LeaveTime.Format("2006-01-02 15:04:05")
	}
	return fmt.Sprintf("【%s】接收流量: %d, 发送流量: %d, 进入时间: %s, 离开时间: %s \n", status.Package, status.RxAccum, status.TxAccum, enTime, leTime)
}

func printMap(pkgStatusMapList map[string][]*AppFlowStatus) {
	fmt.Printf("输出记录信息：")
	for pkg, statusList := range pkgStatusMapList {
		fmt.Printf("\t[%s]的进出记录：\n", pkg)
		for i, status := range statusList {
			fmt.Printf("\t\t 第 %d 次, %s \n", i+1, status.ToString())
		}
	}
	fmt.Println()
}

// getUID getUID根据app package name 获取这个app对应的应用id（唯一id）用于查询流量
func getUID(pkg string) (string, error) {
	// dumpsys package 'cn.juqing.cesuwang_tv' | grep -E 'userId=|appId=|uid='
	// pm list packages -U | grep 'cn.juqing.cesuwang_tv'
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
//
// Param
//   - uid： pkg的唯一标识
//
// Returns
//
//   - rxTotal： 接收总字节数 Byte
//   - txTotal： 发送总字节数 Byte
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

// 批量删除list里面的元素
func removePkgList(originList, deleteList []string) []string {
	// 1. 把 deleteList 放入 map，方便快速查找
	delMap := make(map[string]struct{}, len(deleteList))
	for _, v := range deleteList {
		delMap[v] = struct{}{}
	}

	// 2. 构造新列表，只保留不在 delMap 里的元素
	result := originList[:0] // 复用底层数组，性能更好
	for _, v := range originList {
		if _, exists := delMap[v]; !exists {
			result = append(result, v)
		}
	}
	return result
}

// 监控指定pkg对应app进入app和离开app这一段时间内的流量
// CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o front_app_range_flow2 front_app_range_flow2.go
// adb -s 192.168.200.55:5555 push front_app_range_flow2 /data/local/tmp

func main() {

	// ./front_app_range_flow
	pkgs := flag.String("pkgs", "com.feedying.live.mix,cn.miguvideo.migutv,cn.juqing.cesuwang_tv", "Comma-separated package names (e.g. com.feedying.live.mix,cn.miguvideo.migutv)")
	reportURL := flag.String("server", "http://192.168.201.167:8080/client/traffic", "Server base URL for reporting")

	flag.Parse()

	if *pkgs == "" {
		fmt.Println("Usage: ./front_app_range_flow --pkgs=com.feedying.live.mix,cn.miguvideo.migutv")
		return
	}

	// 输出接收到的参数
	fmt.Println("接收到的参数：")
	pkgList := strings.Split(*pkgs, ",") // 如果需要，把包名解析成切片
	fmt.Println("解析后的包名列表:", pkgList)

	// 获取所有pkg的uid便于获取对应的流量
	var appUidMap map[string]string = make(map[string]string)
	removeList := make([]string, 0)
	for i := range pkgList {
		pkg := strings.TrimSpace(pkgList[i])
		pkgList[i] = pkg
		uid, err := getUID(pkg)
		if err != nil {
			fmt.Printf("[WARN] [ %s] 获取uid失败：%v\n", pkg, err)
			removeList = append(removeList, pkg)
			continue
		}
		appUidMap[pkg] = uid
		fmt.Printf("[INFO] %s → UID %s\n", pkg, uid)
	}
	pkgList = removePkgList(pkgList, removeList)
	if len(pkgList) == 0 {
		fmt.Printf("[ERR] 所有app获取uid都失败，程序退出。\n")
		return
	}

	// 每隔1s进行一个前台app监测，
	// 如果进入了要监测的app则记录该app当前流量作为开始节点，如果离开了要监测的app则该app当前流量为离开节点的流量，两者之差即为这个app这次使用的流量
	lastPackage := ""
	// 每一个app只保留最新的一次数据，离开app之后立即上报数据
	pkgStatusMap := make(map[string]*AppFlowStatus) // pkg -> *AppFlowStatus
	for range time.Tick(1 * time.Second) {
		pkg, err := GetForegroundApp()
		if err != nil {
			fmt.Println("前台app检测失败:", err)
			continue
		}
		//fmt.Println("当前前台应用包名:", pkg)

		if lastPackage == pkg {
			continue
		}

		flag1 := slices.Contains(pkgList, pkg)
		flag2 := slices.Contains(pkgList, lastPackage)

		if !flag1 && !flag2 {
			lastPackage = pkg
			continue
		}
		dealPkg := pkg
		if flag2 {
			dealPkg = lastPackage
		}

		t := NowInCST()
		if status, ok := pkgStatusMap[dealPkg]; ok {
			if flag2 {
				// 离开app 停止这一次的流量监控
				leaveAppFlow(appUidMap, dealPkg, t, status, *reportURL)
				// 处理完毕之后删除这个记录
				delete(pkgStatusMap, dealPkg)
			}
		} else {
			pkgStatusMap[dealPkg] = entryAppFlow(appUidMap, dealPkg, t)
		}
		lastPackage = pkg
	}

}

// leaveAppFlow 离开app时记录app流量
func leaveAppFlow(appUidMap map[string]string, dealPkg string, t time.Time, status *AppFlowStatus, reportURL string) {
	appUid := ""
	if uid, ok := appUidMap[dealPkg]; ok {
		appUid = uid
	} else {
		fmt.Printf("[ERR] appUidMap error 无法获取 [%s] 的uid: %v\n", dealPkg)
		return
	}

	rxTotal, txTotal, err := getTraffic(appUid)
	if err != nil {
		fmt.Printf("[ERR] getTraffic error %s: %v\n", dealPkg, err)
		//continue
	}

	status.LeaveTime = &t
	status.LeaveRxTotal = rxTotal
	status.LeaveTxTotal = txTotal
	status.RxAccum = status.LeaveRxTotal - status.EntryRxTotal
	status.TxAccum = status.LeaveTxTotal - status.EntryTxTotal

	fmt.Printf("离开[%s]时流量情况，rxTotal=%d, txTotal=%d, status=%v \n", dealPkg, rxTotal, txTotal, status)

	fmt.Printf("Rx 总byte差值%d Byte, %d KB, %d MB  \n", status.RxAccum, status.RxAccum/1024, status.RxAccum/1024/1024)
	fmt.Printf("Tx 总byte差值%d Byte, %d KB, %d MB  \n", status.TxAccum, status.TxAccum/1024, status.TxAccum/1024/1024)

	// 一次进出记录完毕之后上报数据
	go reportTraffic(reportURL, status)
}

// reportTraffic 上报流量数据到服务器
func reportTraffic(reportURL string, status *AppFlowStatus) {
	jsonData, _ := json.Marshal(status)
	fmt.Printf("[INFO] 上报参数: jsonData=%s\n", string(jsonData))
	resp, err := http.Post(reportURL, "application/json", bytes.NewBuffer(jsonData))
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

	fmt.Printf("[INFO] 上报结果: code=%d msg=%s count=%d \n", result.Code, result.Msg, result.Count)
}

// entryAppFlow 进入app时记录当前流量
func entryAppFlow(appUidMap map[string]string, dealPkg string, t time.Time) *AppFlowStatus {
	appUid := ""
	if uid, ok := appUidMap[dealPkg]; ok {
		appUid = uid
	} else {
		fmt.Printf("[ERR] appUidMap error 无法获取 [%s] 的uid: %v\n", dealPkg)
		return nil
	}

	rxTotal, txTotal, err := getTraffic(appUid)
	if err != nil {
		fmt.Printf("[ERR] getTraffic error %s: %v\n", dealPkg, err)
		//continue
	}
	fmt.Printf("进入[%s]时流量情况，rxTotal=%d, txTotal=%d \n", dealPkg, rxTotal, txTotal)

	newStatus := &AppFlowStatus{
		Package:      dealPkg,
		EntryTime:    &t,
		EntryRxTotal: rxTotal,
		EntryTxTotal: txTotal,
	}
	return newStatus
}

// [ERR] getTraffic error cn.juqing.cesuwang_tv: failed to run dumpsys netstats: exit status 1
