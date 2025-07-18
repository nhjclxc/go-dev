// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type UserInsertReq struct {
	Name         string `json:"name"`         // A regular string field
	Email        string `json:"email"`        // A pointer to a string, allowing for null values
	MemberNumber string `json:"memberNumber"` // Uses sql.NullString to handle nullable strings
	Remark       string `json:"remark"`       // 备注
}

type UserUpdateReq struct {
	UserId       uint64 `json:"userId"`       // Standard field for the primary key
	Name         string `json:"name"`         // A regular string field
	Email        string `json:"email"`        // A pointer to a string, allowing for null values
	MemberNumber string `json:"memberNumber"` // Uses sql.NullString to handle nullable strings
	Remark       string `json:"remark"`       // 备注
}
type UserPageListReq struct {
	Name         string `json:"name"`         // A regular string field
	Email        string `json:"email"`        // A pointer to a string, allowing for null values
	MemberNumber string `json:"memberNumber"` // Uses sql.NullString to handle nullable strings
	Remark       string `json:"remark"`       // 备注

}

type MysqlApiReq struct {
	TaskId   int64  `form:"taskId,optional"`   //用户id
	TaskName string `form:"taskName,optional"` //用户名
}

type MysqlApiResp struct {
	TaskId   int64  `json:"taskId"`   //用户id
	TaskName string `json:"taskName"` //用户名
	Address  string `json:"address"`  //地址
	Age      int    `json:"age"`      //年龄
}
