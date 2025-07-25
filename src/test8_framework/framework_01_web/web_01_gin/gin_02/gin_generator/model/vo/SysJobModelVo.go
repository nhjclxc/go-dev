package vo

import (
    "fmt"
    "gin_generator/model"
    "github.com/jinzhu/copier"
)

// 定时任务调度对象 SysJobVo
// @author 
// @date 2025-07-25T10:33:22.528
type SysJobVo struct {

    model.SysJob

    Foo string `form:"foo"` // foo
    Bar string `form:"bar"` // bar
    // ...
}


// ModelToVo model 转化为 modelVo
func (this *SysJobVo) ModelToVo(sysJob *model.SysJob) error {
    // go get github.com/jinzhu/copier

    sysJob = &model.SysJob{} // copier.Copy 不会自动为其分配空间，所以初始化指针指向的结构体
    err := copier.Copy(&this, &sysJob)
    if err != nil {
        fmt.Printf("ModelToVo Copy error: %v", err)
        return err
    }
    return nil
}
