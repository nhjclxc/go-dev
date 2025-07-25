package dto

import (
	"fmt"
	"gin_generator/model"
	"github.com/jinzhu/copier"
	"testing"
	"time"
)


func TestName(t *testing.T) {

	genTableDto := GenTableDto{
		GenTable:  model.GenTable{ClassName: "六百六十六"},
		Keyword:   "1234",
		PageNum:   10,
		PageSize:  30,
		SatrtTime: time.Time{},
		EndTime:   time.Time{},
	}
	genTable := &model.GenTable{}
	_ = copier.Copy(&genTable, &genTableDto)

	fmt.Println(genTable) // nil


}