package response

import (
	"time"
)

// TabUserResp 用户 对象响应结构体
type TabUserResp struct {
	Id int64 `json:"id" form:"id"` // 用户ID

	Username string `json:"username" form:"username"` // 用户名

	Password string `json:"password" form:"password"` // 密码（加密存储）

	Email string `json:"email" form:"email"` // 邮箱

	CreatedAt time.Time `json:"createdAt" form:"createdAt"` // 创建时间

	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"` // 更新时间

	Foo string `form:"foo"` // foo
	Bar string `form:"bar"` // bar
	// ...
}

// TabUserExport 用于导出接口
type TabUserExport struct {
	Id int64 `json:"id" form:"id"` // 用户ID

	Username string `json:"username" form:"username"` // 用户名

	Password string `json:"password" form:"password"` // 密码（加密存储）

	Email string `json:"email" form:"email"` // 邮箱

	CreatedAt time.Time `json:"createdAt" form:"createdAt"` // 创建时间

	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"` // 更新时间

}
