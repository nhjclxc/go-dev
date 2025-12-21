package utils

import "time"

var location = time.FixedZone("CST", 8*60*60)
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

func NowInCSTPtr() *time.Time {
	t := NowInCST()
	return &t
}
