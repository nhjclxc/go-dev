package utils

import (
	"encoding/json"
	"strings"
	"time"
)

// 定义时间格式
const DateFormat = "2006-01-02"
const DateTimeFormat = "2006-01-02 15:04:05"

// LocalTime 封装 time.Time，默认以 DateTimeFormat 格式序列化
type LocalTime struct {
	time.Time
}

// 实现 json.Marshaler 接口
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := t.Format(DateTimeFormat)
	return json.Marshal(formatted)
}

// 实现 json.Unmarshaler 接口
func (t *LocalTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	if str == "" || str == "null" {
		t.Time = time.Time{}
		return nil
	}
	// 兼容两种格式
	tt, err := time.ParseInLocation(DateTimeFormat, str, time.Local)
	if err != nil {
		tt, err = time.ParseInLocation(DateFormat, str, time.Local)
		if err != nil {
			return err
		}
	}
	t.Time = tt
	return nil
}

// 实现 fmt.Stringer 接口
func (t LocalTime) String() string {
	return t.Format(DateTimeFormat)
}
