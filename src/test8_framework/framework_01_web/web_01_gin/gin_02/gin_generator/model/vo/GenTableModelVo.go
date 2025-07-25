package vo

import (
    "fmt"
    "gin_generator/model"
    "github.com/jinzhu/copier"
)

// 代码生成业务对象 GenTableVo
// @author
// @date 2025-07-25T09:17:40.679
type GenTableVo struct {

    model.GenTable

    Foo string `form:"foo"` // foo
    Bar string `form:"bar"` // bar
    // ...
}


// ModelToVo model 转化为 modelVo
func (this *GenTableVo) ModelToVo(genTable *model.GenTable) error {
    // go get github.com/jinzhu/copier

    genTable = &model.GenTable{}
    err := copier.Copy(&this, &genTable)
    if err != nil {
        fmt.Printf("ModelToVo Copy error: %v", err)
        return err
    }
    return nil
}
