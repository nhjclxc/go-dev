package model

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/mitchellh/mapstructure"
    "gorm.io/gorm"
    "time"
)

// GenTable 代码生成业务 结构体
// @author
// @date 2025-07-25T10:06:35.433
type GenTable struct {


    TableId int64 `gorm:"column:table_id;primaryKey;auto_increment;not null;type:bigint" json:"tableId" form:"tableId"`// 编号

    TableComment string `gorm:"column:table_comment;type:varchar(500)" json:"tableComment" form:"tableComment"`// 表描述

    SubTableName string `gorm:"column:sub_table_name;type:varchar(64)" json:"subTableName" form:"subTableName"`// 关联子表的表名

    SubTableFkName string `gorm:"column:sub_table_fk_name;type:varchar(64)" json:"subTableFkName" form:"subTableFkName"`// 子表关联的外键名

    ClassName string `gorm:"column:class_name;type:varchar(100)" json:"className" form:"className"`// 实体类名称

    TplCategory string `gorm:"column:tpl_category;type:varchar(200)" json:"tplCategory" form:"tplCategory"`// 使用的模板（crud单表操作 tree树表操作）

    TplWebType string `gorm:"column:tpl_web_type;type:varchar(30)" json:"tplWebType" form:"tplWebType"`// 前端模板类型（element-ui模版 element-plus模版）

    PackageName string `gorm:"column:package_name;type:varchar(100)" json:"packageName" form:"packageName"`// 生成包路径

    ModuleName string `gorm:"column:module_name;type:varchar(30)" json:"moduleName" form:"moduleName"`// 生成模块名

    BusinessName string `gorm:"column:business_name;type:varchar(30)" json:"businessName" form:"businessName"`// 生成业务名

    FunctionName string `gorm:"column:function_name;type:varchar(50)" json:"functionName" form:"functionName"`// 生成功能名

    FunctionAuthor string `gorm:"column:function_author;type:varchar(50)" json:"functionAuthor" form:"functionAuthor"`// 生成功能作者

    GenType string `gorm:"column:gen_type;type:char(1)" json:"genType" form:"genType"`// 生成代码方式（0zip压缩包 1自定义路径）

    GenPath string `gorm:"column:gen_path;type:varchar(200)" json:"genPath" form:"genPath"`// 生成路径（不填默认项目路径）

    Options string `gorm:"column:options;type:varchar(1000)" json:"options" form:"options"`// 其它生成选项

    CreateBy string `gorm:"column:create_by;type:varchar(64)" json:"createBy" form:"createBy"`// 创建者

    CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"createTime" form:"createTime"`// 创建时间

    UpdateBy string `gorm:"column:update_by;type:varchar(64)" json:"updateBy" form:"updateBy"`// 更新者

    UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"updateTime" form:"updateTime"`// 更新时间

    Remark string `gorm:"column:remark;type:varchar(500)" json:"remark" form:"remark"`// 备注

    // time_format:"2006-01-02 15:04:05"
}

// TableName 返回当前实体类的表名
func (this *GenTable) TableName() string {
    return "gen_table"
}


// 可用钩子函数包括：
// BeforeCreate / AfterCreate
// BeforeUpdate / AfterUpdate
// BeforeDelete / AfterDelete
func (this *GenTable) BeforeCreate(tx *gorm.DB) (err error) {
    this.CreateTime = time.Now()
    this.UpdateTime = time.Now()
    return
}

func (this *GenTable) BeforeUpdate(tx *gorm.DB) (err error) {
    this.UpdateTime = time.Now()
    return
}

// MapToStruct map映射转化为当前结构体
func MapToGenTable(inputMap map[string]any) (*GenTable) {
    //go get github.com/mitchellh/mapstructure

    var genTable GenTable
    err := mapstructure.Decode(inputMap, &genTable)
    if err != nil {
        fmt.Printf("MapToStruct Decode error: %v", err)
        return nil
    }
    return &genTable
}

// StructToMap 当前结构体转化为map映射
func (this *GenTable) GenTableToMap() (map[string]any) {
    var m map[string]any
    bytes, err := json.Marshal(this)
    if err != nil {
        fmt.Printf("StructToMap marshal error: %v", err)
        return nil
    }

    err = json.Unmarshal(bytes, &m)
    if err != nil {
        fmt.Printf("StructToMap unmarshal error: %v", err)
        return nil
    }
    return m
}



// 由于有时需要开启事务，因此 DB *gorm.DB 选择从外部传入

// InsertGenTable 新增代码生成业务
func (this *GenTable) InsertGenTable(DB *gorm.DB) (int, error) {
    fmt.Printf("InsertGenTable：%#v \n", this)

    // 先查询是否有相同 name 的数据存在
    temp := &GenTable{}
    // todo name
    tx := DB.Where("job_name = ?", this.ClassName).First(temp)
    fmt.Printf("InsertGenTable.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
        return 0, errors.New("InsertGenTable.Where, 存在相同name: " + temp.ClassName)
    }

    // 执行 Insert
    err := DB.Create(&this).Error

    if err != nil {
        return 0, errors.New("InsertGenTable.DB.Create, 新增失败: " + err.Error())
    }
    return 1, nil
}

// BatchInsertGenTables 批量新增代码生成业务
func BatchInsertGenTables(DB *gorm.DB, tables []GenTable) (int, error) {

    result := DB.Create(&tables)

    if result.Error != nil {
        return 0, errors.New("InsertGenTable.DB.Create, 新增失败: " + result.Error.Error())
    }
    return int(result.RowsAffected), nil
}

// UpdateGenTableByTableId 根据主键修改代码生成的所有字段
func (this *GenTable) UpdateGenTableByTableId(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateGenTableByTableId：%#v \n", this)

    // 1、查询该id是否存在
    if this.TableId == 0 {
        return 0, errors.New("TableId 不能为空！！！: ")
    }

    // 2、再看看name是否重复
    temp := &GenTable{}
    // todo name
    tx := DB.Where("job_name = ?", this.ClassName).First(temp)
    fmt.Printf("UpdateGenTableByTableId.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && temp.TableId != this.TableId {
        return 0, errors.New("UpdateGenTableByTableId.Where, 存在相同name: " + temp.ClassName)
    }

    // 3、执行修改
    //保存整个结构体（全字段更新）
    saveErr := DB.Save(this).Error
    if saveErr != nil {
        return 0, errors.New("UpdateGenTableByTableId.Save, 修改失败: " + saveErr.Error())
    }
    return 1, nil
}

// FindGenTablesByIDList 根据主键批量查询代码生成
func FindGenTablesByIDList(DB *gorm.DB, idList []int64) ([]GenTable, error) {
    var result []GenTable
    err := DB.Where("id IN ?", idList).Find(&result).Error
    return result, err
}

// BatchDeleteGenTables 根据主键修改代码生成的所有字段
func BatchDeleteGenTables(DB *gorm.DB, idList []int64) error {
    return DB.Where("id IN ?", idList).Delete(&GenTable{}).Error
}

// UpdateGenTableSelective 修改代码生成不为默认值的字段
func (this *GenTable) UpdateGenTableSelective(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateGenTableSelective：%#v \n", this)

    // db.Model().Updates()：只更新指定字段
    err := DB.Model(this).
        Where("table_id = ?", this.TableId).
        Updates(this).
        Error
    if err != nil {
        return 0, errors.New("UpdateGenTableSelective.Updates, 选择性修改失败: " + err.Error())
    }

    return 1, nil
}

// DeleteGenTable 删除代码生成业务
func (this *GenTable) DeleteGenTable(DB *gorm.DB, tableIdList []int64) (int, error) {
    fmt.Printf("DeleteGenTable：%#v \n", tableIdList)

    // 以下使用的是软删除，以下必须有DeletedAt gorm.DeletedAt字段
    result := DB.Delete(&this, "table_id in ?", tableIdList)
    // result := DB.Model(&this).Where("table_id IN ?", tableIdList).Update("state", 0)
    if result.Error != nil {
        return 0, errors.New("DeleteGenTable.Delete, 删除失败: " + result.Error.Error())
    }

    //// 以下使用的是物理删除
    //result := DB.Unscoped().Delete(this, "table_id in ?", tableIdList)
    //if result.Error != nil {
    //	return 0, errors.New("DeleteGenTable.Delete, 删除失败: " + result.Error.Error())
    //}

    return int(result.RowsAffected), nil
}

// GetGenTableByTableId 获取代码生成业务业务详细信息
func (this *GenTable) GetGenTableByTableId(DB *gorm.DB, tableId int64) (error) {
    fmt.Printf("DeleteGenTable：%#v \n", tableId)
    return DB.First(this, "table_id = ?", tableId).Error
}

// FindGenTableList 查询代码生成业务业务列表
func (this *GenTable) FindGenTableList(DB *gorm.DB, satrtTime time.Time, endTime time.Time) ([]GenTable, error) {
    fmt.Printf("GetGenTableList：%#v \n", this)

    var tables []GenTable
    query := DB.Model(this)

    // 构造查询条件

    if this.TableId != 0 { query = query.Where("table_id = ?", this.TableId ) }


    if this.TableComment != "" { query = query.Where("table_comment LIKE ?", "%" + this.TableComment + "%") }

    if this.SubTableName != "" { query = query.Where("sub_table_name LIKE ?", "%" + this.SubTableName + "%") }

    if this.SubTableFkName != "" { query = query.Where("sub_table_fk_name LIKE ?", "%" + this.SubTableFkName + "%") }

    if this.ClassName != "" { query = query.Where("class_name LIKE ?", "%" + this.ClassName + "%") }

    if this.TplCategory != "" { query = query.Where("tpl_category = ?", this.TplCategory ) }

    if this.TplWebType != "" { query = query.Where("tpl_web_type = ?", this.TplWebType ) }

    if this.PackageName != "" { query = query.Where("package_name LIKE ?", "%" + this.PackageName + "%") }

    if this.ModuleName != "" { query = query.Where("module_name LIKE ?", "%" + this.ModuleName + "%") }

    if this.BusinessName != "" { query = query.Where("business_name LIKE ?", "%" + this.BusinessName + "%") }

    if this.FunctionName != "" { query = query.Where("function_name LIKE ?", "%" + this.FunctionName + "%") }

    if this.FunctionAuthor != "" { query = query.Where("function_author = ?", this.FunctionAuthor ) }

    if this.GenType != "" { query = query.Where("gen_type = ?", this.GenType ) }

    if this.GenPath != "" { query = query.Where("gen_path = ?", this.GenPath ) }

    if this.Options != "" { query = query.Where("options = ?", this.Options ) }

    if this.CreateBy != "" { query = query.Where("create_by = ?", this.CreateBy ) }
    if !this.CreateTime.IsZero() {
        query = query.Where("create_time = ?", this.CreateTime)
        // query = query.Where("DATE(create_time) = ?", this.$column.goField.Format("2006-01-02"))
    }

    if this.UpdateBy != "" { query = query.Where("update_by = ?", this.UpdateBy ) }
    if !this.UpdateTime.IsZero() {
        query = query.Where("update_time = ?", this.UpdateTime)
        // query = query.Where("DATE(update_time) = ?", this.$column.goField.Format("2006-01-02"))
    }

    if this.Remark != "" { query = query.Where("remark = ?", this.Remark ) }

    if !satrtTime.IsZero() {
        query = query.Where("create_time >= ?", satrtTime)
    }
    if !endTime.IsZero() {
        query = query.Where("create_time <= ?", endTime)
    }

    // // 添加分页逻辑
    // if genTable.PageNum > 0 && genTable.PageSize > 0 {
    //     offset := (genTable.PageNum - 1) * genTable.PageSize
    //     query = query.Offset(offset).Limit(genTable.PageSize)
    // }

    err := query.Find(&tables).Error
    return tables, err
}

// FindGenTablePageList 分页查询代码生成业务业务列表
func (this *GenTable) FindGenTablePageList(DB *gorm.DB, satrtTime time.Time, endTime time.Time, pageNum int, pageSize int) ([]GenTable, int64, error) {
    fmt.Printf("GetGenTablePageList：%#v \n", this)

    var (
        genTables []GenTable
        total     int64
    )

    query := DB.Model(&GenTable{})

    // 构造查询条件
    if this.TableId != 0 { query = query.Where("table_id = ?", this.TableId ) }
    if this.TableComment != "" { query = query.Where("table_comment LIKE ?", "%" + this.TableComment + "%") }
    if this.SubTableName != "" { query = query.Where("sub_table_name LIKE ?", "%" + this.SubTableName + "%") }
    if this.SubTableFkName != "" { query = query.Where("sub_table_fk_name LIKE ?", "%" + this.SubTableFkName + "%") }
    if this.ClassName != "" { query = query.Where("class_name LIKE ?", "%" + this.ClassName + "%") }
    if this.TplCategory != "" { query = query.Where("tpl_category = ?", this.TplCategory ) }
    if this.TplWebType != "" { query = query.Where("tpl_web_type = ?", this.TplWebType ) }
    if this.PackageName != "" { query = query.Where("package_name LIKE ?", "%" + this.PackageName + "%") }
    if this.ModuleName != "" { query = query.Where("module_name LIKE ?", "%" + this.ModuleName + "%") }
    if this.BusinessName != "" { query = query.Where("business_name LIKE ?", "%" + this.BusinessName + "%") }
    if this.FunctionName != "" { query = query.Where("function_name LIKE ?", "%" + this.FunctionName + "%") }
    if this.FunctionAuthor != "" { query = query.Where("function_author = ?", this.FunctionAuthor ) }
    if this.GenType != "" { query = query.Where("gen_type = ?", this.GenType ) }
    if this.GenPath != "" { query = query.Where("gen_path = ?", this.GenPath ) }
    if this.Options != "" { query = query.Where("options = ?", this.Options ) }
    if this.CreateBy != "" { query = query.Where("create_by = ?", this.CreateBy ) }
    if !this.CreateTime.IsZero() {
        query = query.Where("create_time = ?", this.CreateTime)
        // query = query.Where("DATE(create_time) = ?", this.$column.goField.Format("2006-01-02"))
    }
    if this.UpdateBy != "" { query = query.Where("update_by = ?", this.UpdateBy ) }
    if !this.UpdateTime.IsZero() {
        query = query.Where("update_time = ?", this.UpdateTime)
        // query = query.Where("DATE(update_time) = ?", this.$column.goField.Format("2006-01-02"))
    }
    if this.Remark != "" { query = query.Where("remark = ?", this.Remark ) }

    if !satrtTime.IsZero() {
        query = query.Where("create_time >= ?", satrtTime)
    }
    if !endTime.IsZero() {
        query = query.Where("create_time <= ?", endTime)
    }

    // 查询总数
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 分页参数默认值
    if pageNum <= 0 {
        pageNum = 1
    }
    if pageSize <= 0 {
        pageSize = 10
    }

    // 分页数据
    err := query.
        Limit(pageSize).Offset((pageNum - 1) * pageSize).
        Order("create_time desc").
        Find(&genTables).Error

    if err != nil {
        return nil, 0, err
    }

    return genTables, total, nil
}

