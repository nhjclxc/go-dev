package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"
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
	// 服务器下发哪些app要被kill
	// ...

	pkgs := flag.String("pkgs", "cn.miguvideo.migutv,com.feedying.live.mix", "--pkgs=cn.miguvideo.migutv,com.feedying.live.mix")
	flag.Parse()

	if *pkgs == "" {
		fmt.Println("使用方法：./process --pkgs=pkg1,pkg2,pkg3")
	}
	pkgList := strings.Split(*pkgs, ",")
	fmt.Println("读取到的pkgList = ", pkgList)

	// 先检查是否有root权限
	if !hasRoot() {
		fmt.Println("当前盒子不具备root权限，进程关闭失败！！！")
	}

	//err := StopAppsRoot(pkgList)
	////err := KillAppsAsRoot(pkgList)
	//if err != nil {
	//	fmt.Println("关闭进程失败：", err)
	//	return
	//}

	//killList, err := StopAppsRoot2(pkgList)
	killList, err := StopAppsRoot3(pkgList)
	if err != nil {
		fmt.Println("关闭进程失败：", err)
		return
	}
	for _, status := range killList {
		fmt.Printf("执行结果：%v \n", status)
	}
	// 响应服务器杀死状态
	// ...

}

// StopAppsRoot3 一次进入 su 环境，批量 force-stop
func StopAppsRoot3(pkgList []string) ([]*KillAppStatus, error) {
	res := make([]*KillAppStatus, 0, len(pkgList))
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
		status := &KillAppStatus{Package: pkg}
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

type KillAppStatus struct {
	Package       string `json:"package"`
	CmdStatus     bool   `json:"cmdStatus"`
	CloseStatus   bool   `json:"closeStatus"`
	FailureReason string `json:"failureReason"`
}

// StopAppsRoot2 一次进入 su 环境，批量 force-stop
func StopAppsRoot2(pkgList []string) ([]*KillAppStatus, error) {
	res := make([]*KillAppStatus, 0)
	if len(pkgList) == 0 {
		return res, fmt.Errorf("包名列表为空")
	}

	// 构建多条命令
	for _, pkg := range pkgList {
		var cmdBuilder strings.Builder
		cmdBuilder.WriteString(fmt.Sprintf("echo '正在停止: %s'\n", pkg))
		cmdBuilder.WriteString(fmt.Sprintf("am force-stop %s\n", pkg))
		// 执行命令
		cmd := exec.Command("su")
		cmd.Stdin = strings.NewReader(cmdBuilder.String())

		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		err := cmd.Run()
		fmt.Println(out.String())

		r := &KillAppStatus{
			Package: pkg,
		}
		if err != nil {
			r.FailureReason = fmt.Sprintf("执行 su 停止失败: %s: %s", pkg, err.Error())
			r.CloseStatus = false
		} else {
			r.CloseStatus = true
		}
		fmt.Printf("执行 su 停止成功: %s \n", pkg)
		res = append(res, r)

	}

	return res, nil
}

// StopAppsRoot 一次进入 su 环境，批量 force-stop
func StopAppsRoot(pkgList []string) error {
	if len(pkgList) == 0 {
		return fmt.Errorf("包名列表为空")
	}

	// 构建多条命令
	var cmdBuilder strings.Builder
	for _, pkg := range pkgList {
		cmdBuilder.WriteString(fmt.Sprintf("echo '正在停止: %s'\n", pkg))
		cmdBuilder.WriteString(fmt.Sprintf("am force-stop %s\n", pkg))
	}
	cmdBuilder.WriteString("echo '所有任务执行完毕'\n")

	// 执行命令
	cmd := exec.Command("su")
	cmd.Stdin = strings.NewReader(cmdBuilder.String())

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	fmt.Println(out.String())

	if err != nil {
		return fmt.Errorf("执行 su 批量停止失败: %w", err)
	}

	return nil
}

// KillAppsAsRoot 以 root 权限批量停止多个包名，注意：Android盒子必须有root权限，可以使用"which su"检测root权限
//
// Param
//   - pkgList: 要杀死进程的包
func KillAppsAsRoot(pkgList []string) error {

	if len(pkgList) == 0 {
		return fmt.Errorf("没有指定要停止的包名")
	}

	// 拼接成多条 shell 命令
	var cmdBuilder strings.Builder
	for _, pkg := range pkgList {
		cmdBuilder.WriteString(fmt.Sprintf("am force-stop %s\n", pkg))
	}
	cmdBuilder.WriteString("exit\n") // su 模式下退出 shell

	cmdStr := cmdBuilder.String()

	// 启动 su 并一次性执行所有命令
	cmd := exec.Command("su")
	cmd.Stdin = strings.NewReader(cmdStr)
	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("批量停止失败: %v, 输出: %s", err, string(out))
	}

	fmt.Println("批量停止成功:\n" + string(out))
	return nil
}

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
