非常好 👍
你希望**完全用 Go 实现**，不引入 Java 代码、也不依赖 gomobile bind。

不过——要先说明一个关键事实：

---

## ⚠️ 重要前提：Go 代码在 Android 上无法直接调用系统 API

在 Android 系统里：

* “当前前台应用” 属于 **系统运行状态信息**
* 这部分信息只能通过 **Android Framework API（Java 层）** 或 **系统命令（shell）** 获取
* Go 在用户态下运行，**无法直接访问 /proc、ActivityManager 等受限资源**

---

因此我们可以退一步，用纯 Go + shell 命令（无 root 兼容）实现“**获取前台应用包名**”，
特别是 **在 Android 盒子环境**（TV / Launcher / 定制系统）下是常用的方式。

---

## ✅ 一、核心思路

通过执行系统命令：

```bash
dumpsys activity activities | grep mResumedActivity
```

在大多数 Android 盒子上，这条命令输出类似：

```
mResumedActivity: ActivityRecord{a4b3f12 u0 com.netflix.ninja/.MainActivity t43}
```

包名就是 `com.netflix.ninja`。
Go 程序只要解析这个字符串即可。

---

## ✅ 二、Go 纯实现（无 Java / 无 root）

```go
package main

import (
    "bytes"
    "fmt"
    "os/exec"
    "regexp"
    "strings"
    "time"
)

func getForegroundApp() (string, error) {
    // 在 Android 盒子上执行 dumpsys 命令
    cmd := exec.Command("sh", "-c", "dumpsys activity activities | grep mResumedActivity")
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("执行命令失败: %v", err)
    }

    // 示例输出: mResumedActivity: ActivityRecord{a4b3f12 u0 com.netflix.ninja/.MainActivity t43}
    result := out.String()
    re := regexp.MustCompile(` ([a-zA-Z0-9_.]+)/`)
    match := re.FindStringSubmatch(result)
    if len(match) > 1 {
        return strings.TrimSpace(match[1]), nil
    }

    return "", fmt.Errorf("未找到前台应用包名")
}

func main() {
    fmt.Println("开始监控前台应用...")

    var last string
    for {
        pkg, err := getForegroundApp()
        if err != nil {
            fmt.Println("获取失败:", err)
        } else if pkg != last {
            fmt.Println("前台应用变化:", pkg)
            last = pkg
        }
        time.Sleep(2 * time.Second)
    }
}
```

---

## ✅ 三、运行环境要求

| 环境                     | 是否可行                |
| ---------------------- | ------------------- |
| Android 盒子 (TV / 定制固件) | ✅ 绝大多数可执行           |
| 普通手机（非 system app）     | ❌ `dumpsys` 受限（返回空） |
| adb shell 测试           | ✅                   |
| 无 root                 | ✅（盒子一般放开 dumpsys）   |

---

## ✅ 四、验证方法（手动）

你可以手动运行看看盒子返回什么：

```bash
adb shell dumpsys activity activities | grep mResumedActivity
```

或：

```bash
adb shell dumpsys window windows | grep mCurrentFocus
```

可能输出：

```
mCurrentFocus=Window{d2320f1 u0 com.google.android.tv.launcher/com.google.android.tv.launcher.MainActivity}
```

这时可以把上面 Go 代码中的命令改成：

```go
cmd := exec.Command("sh", "-c", "dumpsys window windows | grep mCurrentFocus")
```

---

## ✅ 五、优化版（兼容多种盒子）

可以多命令轮询匹配：

```go
func getForegroundApp() string {
    cmds := []string{
        "dumpsys activity activities | grep mResumedActivity",
        "dumpsys window windows | grep mCurrentFocus",
    }

    for _, c := range cmds {
        out, _ := exec.Command("sh", "-c", c).Output()
        re := regexp.MustCompile(` ([a-zA-Z0-9_.]+)/`)
        match := re.FindStringSubmatch(string(out))
        if len(match) > 1 {
            return match[1]
        }
    }
    return ""
}
```

---

## ✅ 六、监控循环

```go
func main() {
    last := ""
    for {
        pkg := getForegroundApp()
        if pkg != "" && pkg != last {
            fmt.Println("前台App变化:", pkg)
            last = pkg
        }
        time.Sleep(2 * time.Second)
    }
}
```

---

## ✅ 七、进阶：打包成后台服务

如果你要让它在 Android 上 **常驻运行（后台服务）**：

* 用 gomobile 构建成一个 `.apk`
* 启动后通过 `os/exec` 周期调用 `getForegroundApp()`
* 把结果上传到服务端或写入文件

示例命令：

```bash
gomobile build -target=android ./foreground_monitor
```

---

## ✅ 八、总结

| 目标            | 是否能用 Go 实现                     | 是否需 root | 稳定性       |
| ------------- | ------------------------------ | -------- | --------- |
| 获取当前前台 app 包名 | ✅ (通过 dumpsys)                 | ❌        | ✅（盒子普遍可行） |
| 获取全部前台 app 列表 | ⚠️（需 Java API）                 | ❌        | 不行        |
| 获取流量          | ✅ (/proc/net/xt_qtaguid/stats) | ❌        | ✅         |
| 获取运行进程列表      | ⚠️ (受限)                        | ❌        | 有限支持      |

---

如果你告诉我：

* 盒子的品牌型号或 Android 版本
* 是否能执行 `adb shell dumpsys activity activities`

我可以帮你定制出最兼容的 **Go-only 前台 App 监控脚本**（针对你盒子调优）。
是否要我帮你做这个针对性版本？
