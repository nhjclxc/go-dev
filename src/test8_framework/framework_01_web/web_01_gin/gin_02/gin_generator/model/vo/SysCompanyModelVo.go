package vo

import (
    "fmt"
    "gin_generator/model"
    "github.com/jinzhu/copier"
)

// 公司对象 SysCompanyVo
// @author 
// @date 2025-07-25T13:56:40.647
type SysCompanyVo struct {

    model.SysCompany

    Foo string `form:"foo"` // foo
    Bar string `form:"bar"` // bar
    // ...
}


// ModelToVo model 转化为 modelVo
func (this *SysCompanyVo) ModelToVo(sysCompany *model.SysCompany) error {
    // go get github.com/jinzhu/copier

    sysCompany = &model.SysCompany{} // copier.Copy 不会自动为其分配空间，所以初始化指针指向的结构体
    err := copier.Copy(&this, &sysCompany)
    if err != nil {
        fmt.Printf("ModelToVo Copy error: %v", err)
        return err
    }
    return nil
}
