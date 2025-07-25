package dto

import (
    "fmt"
    "gin_generator/model"
    "github.com/jinzhu/copier"
    "time"
)

// 代码生成业务对象 GenTableDto
// @author
// @date 2025-07-25T09:17:40.679
type GenTableDto struct {

    model.GenTable

    Keyword  string `form:"keyword"` // 模糊搜索字段

    PageNum  int    `form:"pageNum"` // 页码
    PageSize int    `form:"pageSize"` // 页大小

    SatrtTime time.Time `form:"satrtTime" time_format:"2006-01-02 15:04:05"` // 开始时间
    EndTime   time.Time `form:"endTime" time_format:"2006-01-02 15:04:05"`   // 结束时间
}



// DtoToModel modelDto 转化为 model
func (this *GenTableDto) DtoToModel() (genTable *model.GenTable, err error){
    // go get github.com/jinzhu/copier

    genTable = &model.GenTable{}
    err = copier.Copy(&genTable, &this)
    return
}


// ModelToDto model 转化为 modelDto
func (this *GenTableDto) ModelToDto(genTable *model.GenTable) error {
    // go get github.com/jinzhu/copier

    err := copier.Copy(&this, &genTable)
    if err != nil {
        fmt.Printf("DtoTo Copy error: %v", err)
        return err
    }
    return nil
}
