package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

type adbCmdType string

const (
	ROOT      adbCmdType = "root"
	SHELL     adbCmdType = "shell"
	PUSH      adbCmdType = "push"
	PULL      adbCmdType = "pull"
	LS        adbCmdType = "ls"
	CHMOD     adbCmdType = "chmod"
	INSTALL   adbCmdType = "install"
	UNINSTALL adbCmdType = "uninstall"
	LOGCAT    adbCmdType = "logcat"
	REBOOT    adbCmdType = "reboot"
	SCREENCAP adbCmdType = "screencap"
	EXECOUT   adbCmdType = "exec-out"
)

func main() {
	dest := "192.168.200.55:5555"

	_ = dest
	pwd, _ := pwd()

	SendPullCmd(dest, "/data/local/tmp/screen33.png", pwd+"/screen33.png")

	cmd, err := SendExecoutCmd(dest, SCREENCAP, pwd+"/screen1.png")
	if err != nil {
		fmt.Printf("exec-out screencap err: %s \n", err)
		return
	}
	fmt.Printf("cmd output: %s \n", cmd)
}

func SendExecoutCmd(target string, subCmd adbCmdType, outputPath string) (string, error) {
	// adb -s 192.168.200.55:5555 exec-out screencap -p > screen.png
	// adb -s 192.168.200.55:5555 exec-out screencap -p > /Users/lxc20250729/lxc/code/go-dev/src/test14_android/screen1.png
	out, err := SendCmd(target, EXECOUT, []string{string(subCmd), "-p"})
	if err != nil {
		return "", fmt.Errorf("exec-out screencap error: %w", err)
	}

	// out 就是 PNG 的二进制内容
	err = os.WriteFile(outputPath, out, 0644)
	if err != nil {
		return "", fmt.Errorf("exec-out screencap os.WriteFile error: %w", err)
	}

	return outputPath, err
}
func SendScreencapCmd(target string, outputPath string) (string, error) {
	// adb -s 192.168.200.55:5555 shell screencap -p /data/local/tmp/screen.png
	// adb -s 192.168.200.55:5555 exec-out screencap -p > screen.png
	out, err := SendCmd(target, SCREENCAP, []string{"-p", outputPath})
	return string(out), err
}
func SendUninstallCmd(target string, pkg string) (string, error) {
	// adb -s 192.168.200.55:5555 uninstall package...
	out, err := SendCmd(target, UNINSTALL, []string{pkg})
	return string(out), err
}
func SendInstallCmd(target string, options string, pkgFile string) (string, error) {
	// adb -s 192.168.200.55:5555 install -t apkName.apk
	out, err := SendCmd(target, INSTALL, []string{options, pkgFile})
	return string(out), err
}
func SendChmodCmd(target string, mode int, paths []string) (string, error) {
	// adb -s 192.168.200.55:5555 shell "chmod 755 /data/local/bin/android-traffic-monitor /data/local/bin/userinit.sh"
	args := make([]string, 0, 1+len(paths))
	args = append(args, strconv.Itoa(mode))
	for _, path := range paths {
		args = append(args, path)
	}
	out, err := SendCmd(target, CHMOD, args)
	return string(out), err
}
func SendLsCmd(target, path string) (string, error) {
	// adb -s 192.168.200.55:5555 ls /data/local/bin/
	out, err := SendCmd(target, LS, []string{path})
	return string(out), err
}
func SendPullCmd(target, remoteFile, localFile string) (string, error) {
	// adb -s 192.168.200.55:5555 push /data/local/bin/android_jlink ./go_base_project
	out, err := SendCmd(target, PULL, []string{remoteFile, localFile})
	return string(out), err
}

func SendPushCmd(target, localFile, remoteFile string) (string, error) {
	// adb -s 192.168.200.55:5555 push go_base_project /data/local/bin/android_jlink
	out, err := SendCmd(target, PUSH, []string{localFile, remoteFile})
	return string(out), err
}
func SendShellCmd(target, cmdStr string) (string, error) {
	// adb -s 192.168.200.55:5555 shell "ps -ef | grep -E 'traffic|jlink'"
	out, err := SendCmd(target, SHELL, []string{cmdStr})
	return string(out), err
}
func SendCmd(target string, adbCmd adbCmdType, cmdStrs []string) ([]byte, error) {
	// 构造命令
	args := append([]string{"-s", target, string(adbCmd)}, cmdStrs...)

	fmt.Printf("SendCmd: adb %s \n", args)
	cmd := exec.Command("adb", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("adb %s error: %s \n", args, err)
		return []byte{}, fmt.Errorf("%s: %s", err, string(out))
	}
	return out, nil
}
func SendRootCmd(target string, adbCmd adbCmdType, cmdStrs []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// adb -s 192.168.200.55:5555 shell su -c "screencap -p /data/local/tmp/screen.png"
	// adb -s 192.168.200.55:5555 shell su "screencap -p /data/local/tmp/screen.png"
	// adb -s 192.168.200.55:5555 shell "screencap -p /data/local/tmp/screen.png"

	// adb -s 192.168.200.55:5555 pull /data/local/tmp/screen.png ./

	// 优先尝试用 su -c 执行
	cmdStrsTemp := append([]string{"-s", target, string(adbCmd), "su"}, cmdStrs...)
	outBytes, err := exec.CommandContext(ctx, "adb", cmdStrsTemp...).CombinedOutput()
	if err != nil {
		// 如果 su 执行失败，尝试不使用 su（可能当前进程已是 root 或在某些环境里可以直接执行）
		cmdStrsTemp[1] = "sh"
		out2, err2 := exec.CommandContext(ctx, "adb", cmdStrsTemp...).CombinedOutput()
		if err2 != nil {
			return string(out2), fmt.Errorf("su error: %v; fallback sh error: %v; last output: %s", err, err2, string(out2))
		}
		outBytes = out2
	}

	return string(outBytes), nil
}

func pwd() (string, error) {
	cmd := exec.Command("pwd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("self run 'pwd' err: ", err)
		return "", fmt.Errorf("self run 'pwd' err: %s", err)
	}
	str := string(out)
	re := regexp.MustCompile(`[\n\t\r ]`) // 匹配换行、制表符、回车、空格
	str = re.ReplaceAllString(str, "")
	return str, err

}

func SendRoot(target string) (string, error) {
	// adb -s 192.168.200.55:5555 root
	out, err := SendCmd(target, ROOT, []string{"root"})
	return string(out), err
}

/*
adb命令大全
https://developer.android.com/tools/adb?hl=zh-cn


好的，我帮你整理一份 **常用 ADB 命令大全**，覆盖设备管理、文件传输、应用管理、日志调试等，方便你在 Android 开发和运维中使用。

---

## **1️⃣ 设备与连接管理**

| 命令                          | 说明                      |
| --------------------------- | ----------------------- |
| `adb devices`               | 列出所有连接设备（USB 或 TCP/IP）  |
| `adb -s <serial> <command>` | 指定设备执行命令，多设备时使用         |
| `adb connect <ip:port>`     | 通过 TCP/IP 连接 Android 设备 |
| `adb disconnect <ip:port>`  | 断开 TCP/IP 连接            |
| `adb usb`                   | 切换回 USB 模式              |
| `adb tcpip <port>`          | 切换到 TCP/IP 模式监听端口       |

---

## **2️⃣ Shell 与命令执行**

| 命令                                           | 说明                                          |
| -------------------------------------------- | ------------------------------------------- |
| `adb shell`                                  | 进入设备 shell                                  |
| `adb shell <command>`                        | 在设备上执行命令，例如 `adb shell ps -ef`              |
| `adb shell su`                               | 切换到 root 模式（设备已 root）                       |
| `adb shell top`                              | 查看实时 CPU/内存占用                               |
| `adb shell dumpsys`                          | 系统服务状态信息                                    |
| `adb shell getprop`                          | 查看系统属性，如 `getprop ro.build.version.release` |
| `adb shell pm list packages`                 | 列出已安装包                                      |
| `adb shell am start -n <package>/<activity>` | 启动应用的指定 Activity                            |
| `adb shell input keyevent <keycode>`         | 模拟按键事件，如 `KEYCODE_HOME`                     |

---

## **3️⃣ 文件传输**

| 命令                              | 说明       |
| ------------------------------- | -------- |
| `adb push <local> <remote>`     | 上传文件到设备  |
| `adb pull <remote> <local>`     | 从设备下载文件  |
| `adb shell ls <path>`           | 查看设备文件列表 |
| `adb shell rm <path>`           | 删除设备文件   |
| `adb shell mkdir <path>`        | 创建目录     |
| `adb shell chmod <mode> <path>` | 修改文件权限   |

---

## **4️⃣ 应用安装与管理**

| 命令                                      | 说明       |
| --------------------------------------- | -------- |
| `adb install <apk>`                     | 安装 APK   |
| `adb install -r <apk>`                  | 覆盖安装 APK |
| `adb uninstall <package>`               | 卸载应用     |
| `adb shell pm clear <package>`          | 清除应用数据   |
| `adb shell pm enable/disable <package>` | 启用/禁用应用  |

---

## **5️⃣ 日志与调试**

| 命令                                            | 说明          |
| --------------------------------------------- | ----------- |
| `adb logcat`                                  | 实时查看系统日志    |
| `adb logcat -v time`                          | 按时间戳查看日志    |
| `adb logcat -s <tag>`                         | 过滤指定日志 TAG  |
| `adb bugreport > bugreport.zip`               | 收集设备完整报错信息  |
| `adb forward tcp:<hostPort> tcp:<devicePort>` | 本地端口转发到设备端口 |
| `adb reverse tcp:<devicePort> tcp:<hostPort>` | 设备端口反向映射到主机 |

---

## **6️⃣ 系统管理**

| 命令                      | 说明                       |
| ----------------------- | ------------------------ |
| `adb reboot`            | 重启设备                     |
| `adb reboot bootloader` | 重启到 bootloader           |
| `adb reboot recovery`   | 重启到 recovery             |
| `adb root`              | 以 root 权限重新启动 adb daemon |
| `adb remount`           | 重新挂载系统分区（可写）             |

---

## **7️⃣ 高级调试**

| 命令                                          | 说明          |
| ------------------------------------------- | ----------- |
| `adb shell dumpsys cpuinfo`                 | 查看 CPU 使用情况 |
| `adb shell dumpsys meminfo <package>`       | 查看应用内存占用    |
| `adb shell dumpsys package <package>`       | 查看应用信息      |
| `adb shell dumpsys battery`                 | 查看电池状态      |
| `adb shell screencap -p /sdcard/screen.png` | 截屏          |
| `adb shell screenrecord /sdcard/demo.mp4`   | 屏幕录制        |

---

💡 **Tips**：

1. 多设备时必须加 `-s <serial>` 来指定。
2. 长命令可以放入 shell 脚本或者 Go 的 `exec.Command` 里执行。
3. ADB 命令几乎可以完成 **设备管理、文件操作、进程管理、日志调试、应用安装** 等所有运维需求。

---

如果你需要，我可以帮你写一个 **Go 工具库**，把常用 adb 命令封装成函数，直接在 Go 里调用，方便监控 Android 设备的 CPU/内存/网络等状态。

你希望我帮你写这个吗？

*/
