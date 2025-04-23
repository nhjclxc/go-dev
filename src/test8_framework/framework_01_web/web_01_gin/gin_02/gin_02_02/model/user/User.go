package user

// 用户登录实体
type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
	Addr     string `json:"addr" form:"addr"`
}
