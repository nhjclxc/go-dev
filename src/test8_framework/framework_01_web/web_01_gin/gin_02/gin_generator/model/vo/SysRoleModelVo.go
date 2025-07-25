package vo

import (
    "fmt"
    "github.com/jinzhu/copier"
)

// 角色信息对象 SysRoleVo
// @author 
// @date 2025-07-25T11:09:41.517
type SysRoleVo struct {

    model.SysRole

    Foo string `form:"foo"` // foo
    Bar string `form:"bar"` // bar
    // ...
}


// ModelToVo model 转化为 modelVo
func (this *SysRoleVo) ModelToVo(sysRole *model.SysRole) error {
    // go get github.com/jinzhu/copier

    sysRole = &model.SysRole{} // copier.Copy 不会自动为其分配空间，所以初始化指针指向的结构体
    err := copier.Copy(&this, &sysRole)
    if err != nil {
        fmt.Printf("ModelToVo Copy error: %v", err)
        return err
    }
    return nil
}
