package main

import "fmt"

type OS string

var (
	OS_WIN OS = "Windows电脑"
	OS_LINUX OS = "Linux电脑"
	OS_MACOS OS = "MacOS电脑"
)

func IsValidOS(os OS) bool {
	switch os {
	case OS_WIN, OS_LINUX, OS_MACOS:
		return true
	default:
		return false
	}
}

type StatusEnum int

const (
	Unknown StatusEnum = iota
	Active
	Inactive
	Suspended
)

func (s StatusEnum) String() string {
	switch s {
	case Unknown:
		return "未知"
	case Active:
		return "激活"
	case Inactive:
		return "未激活"
	case Suspended:
		return "已暂停"
	default:
		return "非法状态"
	}
}




var OSValues = struct {
	WIN   OS
	LINUX OS
	MACOS OS
}{
	WIN:   "Windows电脑",
	LINUX: "Linux电脑",
	MACOS: "MacOS电脑",
}


func main() {

	fmt.Println(Unknown)
	fmt.Println(Unknown.String())
	fmt.Println(Active)
	fmt.Println(Active.String())
	fmt.Println(Inactive)
	fmt.Println(Inactive.String())
	fmt.Println()
	fmt.Println(OS_WIN)
	fmt.Println(OS_WIN == "Windows电脑")
	fmt.Println(OS_WIN == "Windows电111脑")
	fmt.Println("Windows电111脑" == OS_WIN)
	fmt.Println(OS_LINUX)
	fmt.Println(OS_MACOS)
	fmt.Println()
	fmt.Println(OSValues)
	fmt.Println(OSValues.WIN)
	fmt.Println(OSValues.LINUX)
	fmt.Println(OSValues.MACOS)


}
