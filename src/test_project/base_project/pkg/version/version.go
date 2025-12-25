package version

import (
	"fmt"
	"runtime"
)

var (
	// Version 版本号（在编译时通过 -ldflags 注入）
	Version = "1.0.0"

	// BuildTime 构建时间（在编译时通过 -ldflags 注入）
	BuildTime = "unknown"

	// GitCommit Git 提交哈希（在编译时通过 -ldflags 注入）
	GitCommit = "unknown"
)

// Info 版本信息结构
type Info struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	GitCommit string `json:"git_commit"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// GetVersion 获取当前版本信息
func GetVersion() Info {
	return Info{
		Version:   Version,
		BuildTime: BuildTime,
		GitCommit: GitCommit,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// String 返回版本信息的字符串表示
func (i Info) String() string {
	return fmt.Sprintf("Version: %s\nBuildTime: %s\nGitCommit: %s\nGoVersion: %s\nPlatform: %s",
		i.Version, i.BuildTime, i.GitCommit, i.GoVersion, i.Platform)
}
