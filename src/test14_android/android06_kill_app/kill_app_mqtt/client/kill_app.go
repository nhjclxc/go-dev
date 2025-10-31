package main

import (
	"android06_kill_app/kill_app_mqtt/dto"
	"android06_kill_app/kill_app_mqtt/mqttcore"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os/exec"
	"strings"
	"time"
)

// 测试杀死某个进程
// am force-stop每次只能传入一个参数
// am force-stop com.feedying.live.mix
// am force-stop cn.miguvideo.migutv
// am force-stop cn.juqing.cesuwang_tv
// 或者一行执行
// am force-stop com.feedying.live.mix; am force-stop cn.miguvideo.migutv; am force-stop cn.juqing.cesuwang_tv
// 或者是以下
// for pkg in com.feedying.live.mix cn.miguvideo.migutv cn.juqing.cesuwang_tv; do
//
//	am force-stop $pkg
//
// done
func main() {

	// CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o kill_app kill_app.go
	// adb -s 192.168.200.55:5555 push kill_app /data/local/tmp

	// 服务器下发哪些app要被kill
	// #表示deviceId
	// cmd/+/kill，服务端通知#指定的设备进行杀进程
	// cmd/+/process，服务端请求#指定的设备的当前运行的进程信息
	// cmd/+/flow，服务端通知#指定的设备进行流量统计
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mqttClient *mqttcore.MqttClient

	deviceId, err := getDeviceSerialNumber()
	if err != nil {
		log.Printf("getDeviceSerialNumber 获取设备ID出错", "err", err)
		deviceId = fmt.Sprintf("%s-%d", "kill-app-client", time.Now().Unix())
	}

	topic := "cmd/" + deviceId + "/kill"
	subscribeTopics := make([]*mqttcore.SubscribeTopic, 0)
	subscribeTopics = append(subscribeTopics, &mqttcore.SubscribeTopic{
		Topic: topic,
		Callback: func(client mqtt.Client, msg mqtt.Message) {
			log.Printf("[Message - [%s]消息回调处理器正在处理消息] topic=%s payload=%s", topic, msg.Topic(), msg.Payload())

			// 执行命令
			statusList, err := appKill(msg.Payload())

			// 结果回传
			cmdStatus := true
			reason := ""
			if err != nil {
				fmt.Println("关闭进程失败：", err)
				cmdStatus = false
				reason = fmt.Sprintf("关闭进程失败：%s", err)
			}
			var resp = dto.KillAppCmdResp{
				Reason:        reason,
				CmdStatus:     cmdStatus,
				KillAppStatus: statusList,
			}
			data, err := json.Marshal(resp)
			if err != nil {
				fmt.Println("json序列化失败：", err)
			}

			err = mqttClient.PublishWithAck("cmd/"+deviceId+"/kill"+"/callback", 1, true, data)
			if err != nil {
				fmt.Println("消息回调失败失败：", err)
			}
		},
	})

	// Broker 地址（tcp://host:port）
	mqttCfg := mqttcore.MqttConfig{
		Broker:          "tcp://192.168.201.167:1883",
		ClientId:        "killapp-client-",
		Username:        "admin",
		Password:        "public",
		SubscribeTopics: subscribeTopics,
	}

	mqttClient = mqttcore.NewMqttClient(&mqttCfg)

	mqttClient.Run(ctx)

	//// 持续运行等待消息
	//select {}
}

func appKill(msgByte []byte) ([]*dto.KillAppStatus, error) {
	var req = dto.KillAppCmdReq{}
	json.Unmarshal(msgByte, &req)
	fmt.Println("读取到的pkgList = ", req.PackageList)

	// 先检查是否有root权限
	if !hasRoot() {
		fmt.Println("当前盒子不具备root权限，进程关闭失败！！！")
	}

	killList, err := StopAppsRoot(req.PackageList)
	if err != nil {
		fmt.Println("关闭进程失败：", err)
		return nil, err
	}
	for _, status := range killList {
		fmt.Printf("执行结果：%v \n", status)
	}

	return killList, nil
}

// StopAppsRoot 一次进入 su 环境，批量 force-stop
func StopAppsRoot(pkgList []string) ([]*dto.KillAppStatus, error) {
	res := make([]*dto.KillAppStatus, 0, len(pkgList))
	if len(pkgList) == 0 {
		return res, fmt.Errorf("包名列表为空")
	}

	// 构建单次执行的 shell 脚本
	var cmdBuilder strings.Builder
	for _, pkg := range pkgList {
		cmdBuilder.WriteString(fmt.Sprintf(`echo "[开始停止] %s"
			if ps -A | grep -q "%s"; then
			  am force-stop %s
			  echo "[成功] %s"
			else
			  echo "[跳过] %s 未在运行"
			fi
			`, pkg, pkg, pkg, pkg, pkg))
	}
	cmdBuilder.WriteString("echo '[任务执行完毕]'\n")

	// 启动 su
	cmd := exec.Command("su")
	cmd.Stdin = strings.NewReader(cmdBuilder.String())

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	output := out.String()
	fmt.Println(output)

	// 解析输出，生成结果列表
	for _, pkg := range pkgList {
		status := &dto.KillAppStatus{Package: pkg}
		if strings.Contains(output, fmt.Sprintf("[成功] %s", pkg)) {
			status.CmdStatus = true
			status.CloseStatus = true
		} else if strings.Contains(output, fmt.Sprintf("[跳过] %s", pkg)) {
			status.CmdStatus = false
			status.CloseStatus = true
			status.FailureReason = "未运行"
		} else {
			status.CmdStatus = false
			status.CloseStatus = false
			status.FailureReason = "执行失败或未知错误"
		}
		res = append(res, status)
	}

	if err != nil {
		return res, fmt.Errorf("执行 su 批量停止失败: %w", err)
	}
	return res, nil
}

//
//type KillAppStatus struct {
//	Package       string `json:"package"`
//	CmdStatus     bool   `json:"cmdStatus"`
//	CloseStatus   bool   `json:"closeStatus"`
//	FailureReason string `json:"failureReason"`
//}

// hasRoot 检查Android盒子是否有root权限
//
// return
//   - bool: true表示有root权限
func hasRoot() bool {
	cmd := exec.Command("sh", "-c", "which su")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	path := strings.TrimSpace(string(out))
	return path != ""
}

// getDeviceSerialNumber 通过 getprop 命令获取设备序列号
// 注意：程序运行在 Android 设备上，直接执行 getprop 命令即可
// 返回值:
//   - string: 设备序列号
//   - error: 错误信息
func getDeviceSerialNumber() (string, error) {
	// 执行 getprop ro.serialno 命令
	cmd := exec.Command("getprop", "ro.serialno")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("执行 getprop 命令失败: %w", err)
	}

	// 去除前后空白字符
	serialNumber := strings.TrimSpace(string(output))

	return serialNumber, nil
}
