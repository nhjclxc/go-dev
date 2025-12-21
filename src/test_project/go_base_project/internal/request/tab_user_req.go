package request

import (
	"time"
)

// TabUserReq 用户 对象请求结构体
type TabUserReq struct {
	Id int64 `json:"id" form:"id"` // 用户ID

	Username string `json:"username" form:"username"` // 用户名

	Password string `json:"password" form:"password"` // 密码（加密存储）

	Email string `json:"email" form:"email"` // 邮箱

	CreatedAt time.Time `json:"createdAt" form:"createdAt"` // 创建时间

	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"` // 更新时间

	Keyword string `form:"keyword"` // 模糊搜索字段

	PageNum  int `form:"pageNum"`  // 页码
	PageSize int `form:"pageSize"` // 页大小

	SatrtTime time.Time `form:"satrtTime" time_format:"2006-01-02 15:04:05"` // 开始时间
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02 15:04:05"`   // 结束时间

}
