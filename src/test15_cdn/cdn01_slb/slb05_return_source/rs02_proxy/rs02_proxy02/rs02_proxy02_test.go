package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	// 555M
	var cacheDir = "/Users/lxc20250729/cdn-cache/google_devices-tokay-12990991-8abda025.tgz" // 本地缓存目录
	// /Users/lxc20250729/cdn-cache
	fmt.Println(filepath.Dir(cacheDir))
	// google_devices-tokay-12990991-8abda025.tgz
	fmt.Println(filepath.Base(cacheDir))
}
func TestSplitFile(t *testing.T) {
	var cacheDir = "/Users/lxc20250729/cdn-cache/google_devices-tokay-12990991-8abda025.tgz" // 本地缓存目录
	// 100MB一个分片
	SplitFile(cacheDir, "", 1*1024*1024*100)

	// 分片完成之后，先移除两个分片，模拟当前丢失两个块，看看能不能实现丢失的从原站按Range下载
	// rm -rf google_devices-tokay-12990991-8abda025.tgz.part2 google_devices-tokay-12990991-8abda025.tgz.part3

}
func TestMergeFile(t *testing.T) {
	// 555M
	MergeFile("/Users/lxc20250729/cdn-cache", "google_devices-tokay-12990991-8abda025.tgz", "")
}
