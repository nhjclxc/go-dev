package dto

import (
    "fmt"
    "gin_generator/model"
    "github.com/jinzhu/copier"
    "time"
)

// 定时任务调度对象 SysJobDto
// @author 
// @date 2025-07-25T10:33:22.528
type SysJobDto struct {

    model.SysJob

    Keyword  string `form:"keyword"` // 模糊搜索字段

    PageNum  int    `form:"pageNum"` // 页码
    PageSize int    `form:"pageSize"` // 页大小

    SatrtTime time.Time `form:"satrtTime" time_format:"2006-01-02 15:04:05"` // 开始时间
    EndTime   time.Time `form:"endTime" time_format:"2006-01-02 15:04:05"`   // 结束时间
}



// DtoToModel modelDto 转化为 model
func (this *SysJobDto) DtoToModel() (sysJob *model.SysJob, err error){
    // go get github.com/jinzhu/copier

    sysJob = &model.SysJob{} // copier.Copy 不会自动为其分配空间，所以初始化指针指向的结构体
    err = copier.Copy(&sysJob, &this)
    return
}


// ModelToDto model 转化为 modelDto
func (this *SysJobDto) ModelToDto(sysJob *model.SysJob) error {
    // go get github.com/jinzhu/copier

    err := copier.Copy(&this, &sysJob)
    if err != nil {
    fmt.Printf("DtoTo Copy error: %v", err)
        return err
    }
    return nil
}
