package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 如果要将时间格式化输出到前端，那么请不要使用 time.Time 类型，请使用 JSONTime 类型
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format("2006-01-02 15:04:05")
	return []byte(`"` + formatted + `"`), nil
}

// 让 model.JSONTime 实现 driver.Valuer 接口, 可以让你的自定义时间类型像 time.Time 一样被 GORM 自动识别和处理。
// 实现 fmt.Stringer 接口，方便输出格式化时间
func (t JSONTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// 实现 driver.Valuer 接口 —— 数据库写入时调用
func (t JSONTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// 实现 sql.Scanner 接口 —— 数据库读取时调用
func (t *JSONTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = JSONTime(vt)
	default:
		return fmt.Errorf("cannot convert %v to JSONTime", v)
	}
	return nil
}
