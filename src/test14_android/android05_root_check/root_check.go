package main

import (
	"fmt"
	"os/exec"
	"strings"
)

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

// 检测Android盒子是不是有root	权限
/*
| 命令                   | 输出                | 含义                               |
| -------------------- | ----------------- | -------------------------------- |
| `su`                 | 提示符从 `$` → `#`    | ✅ 已成功进入 root 模式（# 表示 root shell） |
| `ls /system/xbin/su` | `/system/xbin/su` | ✅ 存在 su 可执行文件                    |
| `ls /system/bin/su`  | No such file      | ⚠️ 不影响，部分系统只放在 xbin              |
| `ls /sbin/su`        | No such file      | ⚠️ 也不影响                          |
| 当前提示符                | `#` 而不是 `$`       | ✅ 已是 root 用户                     |


FY928X-K:/data/local/tmp # which su
/system/xbin/su


FY928X-K:/data/local/tmp # id
uid=0(root) gid=0(root) groups=0(root),1004(input),1007(log),1011(adb),1015(sdcard_rw),1028(sdcard_r),1078(ext_data_rw),1079(ext_obb_rw),3001(net_bt_admin),3002(net_bt),3003(inet),3006(net_bw_stats),3009(readproc),3011(uhid),3012(readtracefs) context=u:r:su:s0
*/
func main() {
	if hasRoot() {
		fmt.Println("✅ 当前设备具有 root 权限")
	} else {
		fmt.Println("❌ 当前设备没有 root 权限")
	}
}
