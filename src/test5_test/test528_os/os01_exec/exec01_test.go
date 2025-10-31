package os01_exec

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func Test01(t *testing.T) {

	// 定义命令
	cmd := exec.Command("ps")

	// 设置命令的输出
	var buf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf

	// 执行命令
	err := cmd.Run()
	if err != nil {
		sss := errBuf.String()
		fmt.Printf("命令执行失败: %s！！！%v \n", sss, err)
		return
	}

	// 命令执行成功读取输出数据
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fmt.Println("fields", fields)
	}

}
