package dto

import (
    "fmt"
    "gin_generator/model"
    "github.com/jinzhu/copier"
    "time"
)

// 角色信息对象 SysRoleDto
// @author 
// @date 2025-07-25T11:09:41.517
type SysRoleDto struct {

    model.SysRole

    Keyword  string `form:"keyword"` // 模糊搜索字段

    PageNum  int    `form:"pageNum"` // 页码
    PageSize int    `form:"pageSize"` // 页大小

    SatrtTime time.Time `form:"satrtTime" time_format:"2006-01-02 15:04:05"` // 开始时间
    EndTime   time.Time `form:"endTime" time_format:"2006-01-02 15:04:05"`   // 结束时间
}



// DtoToModel modelDto 转化为 model
func (this *SysRoleDto) DtoToModel() (sysRole *model.SysRole, err error){
    // go get github.com/jinzhu/copier

    sysRole = &model.SysRole{} // copier.Copy 不会自动为其分配空间，所以初始化指针指向的结构体
    err = copier.Copy(&sysRole, &this)
    return
}


// ModelToDto model 转化为 modelDto
func (this *SysRoleDto) ModelToDto(sysRole *model.SysRole) error {
    // go get github.com/jinzhu/copier

    err := copier.Copy(&this, &sysRole)
    if err != nil {
    fmt.Printf("DtoTo Copy error: %v", err)
        return err
    }
    return nil
}
