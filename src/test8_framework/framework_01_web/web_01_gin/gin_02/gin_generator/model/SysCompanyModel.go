package model

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/mitchellh/mapstructure"
    "gorm.io/gorm"
    "time"
)

// SysCompany 公司 结构体
// @author 
// @date 2025-07-25T13:56:40.647
type SysCompany struct {


    Id int64 `gorm:"column:id;primaryKey;auto_increment;not null;type:bigint" json:"id" form:"id"`// 公司id

    CompanyName string `gorm:"column:company_name;not null;type:varchar(30)" json:"companyName" form:"companyName"`// 公司名称

    OrderNum int `gorm:"column:order_num;type:int" json:"orderNum" form:"orderNum"`// 显示顺序

    Status int `gorm:"column:status;type:int" json:"status" form:"status"`// 公司状态（1正常,0停用）

    DelFlag int `gorm:"column:del_flag;type:int" json:"delFlag" form:"delFlag"`// 删除标志（0代表存在,1代表删除）

    CompanyConfig string `gorm:"column:company_config;type:varchar(4096)" json:"companyConfig" form:"companyConfig"`// 公司配置信息

    CreateBy string `gorm:"column:create_by;not null;type:varchar(64)" json:"createBy" form:"createBy"`// 创建者

    CreateTime time.Time `gorm:"column:create_time;not null;type:timestamp" json:"createTime" form:"createTime"`// 创建时间

    UpdateBy string `gorm:"column:update_by;not null;type:varchar(64)" json:"updateBy" form:"updateBy"`// 更新者

    UpdateTime time.Time `gorm:"column:update_time;not null;type:timestamp" json:"updateTime" form:"updateTime"`// 更新时间


    // time_format:"2006-01-02 15:04:05"
}

// TableName 返回当前实体类的表名
func (this *SysCompany) TableName() string {
    return "sys_company"
}


// 可用钩子函数包括：
// BeforeCreate / AfterCreate
// BeforeUpdate / AfterUpdate
// BeforeDelete / AfterDelete
func (this *SysCompany) BeforeCreate(tx *gorm.DB) (err error) {
    this.CreateTime = time.Now()
    this.UpdateTime = time.Now()
    return
}

func (this *SysCompany) BeforeUpdate(tx *gorm.DB) (err error) {
    this.UpdateTime = time.Now()
    return
}

// MapToSysCompany map映射转化为当前结构体
func MapToSysCompany(inputMap map[string]any) (*SysCompany) {
    //go get github.com/mitchellh/mapstructure

    var sysCompany SysCompany
    err := mapstructure.Decode(inputMap, &sysCompany)
    if err != nil {
        fmt.Printf("MapToStruct Decode error: %v", err)
        return nil
    }
    return &sysCompany
}

// SysCompanyToMap 当前结构体转化为map映射
func (this *SysCompany) SysCompanyToMap() (map[string]any) {
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

// InsertSysCompany 新增公司
func (this *SysCompany) InsertSysCompany(DB *gorm.DB) (int, error) {
    fmt.Printf("InsertSysCompany：%#v \n", this)

    // 先查询是否有相同 name 的数据存在
    temp := &SysCompany{}
    tx := DB.Where("company_name = ?", this.CompanyName).First(temp)
    fmt.Printf("InsertSysCompany.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
        return 0, errors.New("InsertSysCompany.Where, 存在相同 company_name: " + temp.CompanyName)
    }

    // 执行 Insert
    err := DB.Create(&this).Error

    if err != nil {
        return 0, errors.New("InsertSysCompany.DB.Create, 新增失败: " + err.Error())
    }
    return 1, nil
}

// BatchInsertSysCompanys 批量新增公司
func (this *SysCompany) BatchInsertSysCompanys(DB *gorm.DB, tables []SysCompany) (int, error) {

    result := DB.Create(&tables)

    if result.Error != nil {
        return 0, errors.New("BatchInsertSysCompanys.DB.Create, 新增失败: " + result.Error.Error())
    }
    return int(result.RowsAffected), nil
}

// UpdateSysCompanyById 根据主键修改公司的所有字段
func (this *SysCompany) UpdateSysCompanyById(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateSysCompanyById：%#v \n", this)

    // 1、查询该id是否存在
    if this.Id == 0 {
        return 0, errors.New("Id 不能为空！！！: ")
    }

    // 2、再看看name是否重复
    temp := &SysCompany{}
    tx := DB.Where("company_name = ?", this.CompanyName).First(temp)
    fmt.Printf("UpdateSysCompanyById.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && temp.Id != this.Id {
        return 0, errors.New("UpdateSysCompanyById.Where, 存在相同 company_name: " + temp.CompanyName)
    }

    // 3、执行修改
    //保存整个结构体（全字段更新）
    saveErr := DB.Save(this).Error
    if saveErr != nil {
        return 0, errors.New("UpdateSysCompanyById.Save, 修改失败: " + saveErr.Error())
    }
    return 1, nil
}

// UpdateSysCompanySelective 修改公司不为默认值的字段
func (this *SysCompany) UpdateSysCompanySelective(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateSysCompanySelective：%#v \n", this)

    // db.Model().Updates()：只更新指定字段
    err := DB.Model(this).
        Where("id = ?", this.Id).
        Updates(this).
        Error
    if err != nil {
        return 0, errors.New("UpdateSysCompanySelective.Updates, 选择性修改失败: " + err.Error())
    }

    return 1, nil
}

// DeleteSysCompany 删除公司
func (this *SysCompany) DeleteSysCompany(DB *gorm.DB, idList []int64) (int, error) {
    fmt.Printf("DeleteSysCompany：%#v \n", idList)

    // 当存在DeletedAt gorm.DeletedAt字段时为软删除，否则为物理删除
    result := DB.Delete(&this, "id in ?", idList)
    // result := DB.Model(&this).Where("id IN ?", tableIdList).Update("state", 0)
    if result.Error != nil {
        return 0, errors.New("DeleteSysCompany.Delete, 删除失败: " + result.Error.Error())
    }

    //// 以下使用的是物理删除
    //result := DB.Unscoped().Delete(this, "id in ?", idList)
    //if result.Error != nil {
    //	return 0, errors.New("DeleteSysCompany.Delete, 删除失败: " + result.Error.Error())
    //}

    return int(result.RowsAffected), nil
}

// BatchDeleteSysCompanys 根据主键批量删除公司
func (this *SysCompany) BatchDeleteSysCompanys(DB *gorm.DB, idList []int64) error {
    return DB.Where("id IN ?", idList).Delete(&this).Error
}

// FindSysCompanyById 获取公司详细信息
func (this *SysCompany) FindSysCompanyById(DB *gorm.DB, id int64) (error) {
    fmt.Printf("DeleteSysCompany：%#v \n", id)
    return DB.First(this, "id = ?", id).Error
}

// FindSysCompanysByIdList 根据主键批量查询公司详细信息
func FindSysCompanysByIdList(DB *gorm.DB, idList []int64) ([]SysCompany, error) {
    var result []SysCompany
    err := DB.Where("id IN ?", idList).Find(&result).Error
    return result, err
}

// FindSysCompanyList 查询公司列表
func (this *SysCompany) FindSysCompanyList(DB *gorm.DB, satrtTime time.Time, endTime time.Time) ([]SysCompany, error) {
    fmt.Printf("GetSysCompanyList：%#v \n", this)

    var tables []SysCompany
    query := DB.Model(this)

        // 构造查询条件
        if this.Id != 0 { query = query.Where("id = ?", this.Id ) }
        if this.CompanyName != "" { query = query.Where("company_name LIKE ?", "%" + this.CompanyName + "%") }
        if this.OrderNum != 0 { query = query.Where("order_num = ?", this.OrderNum ) }
        if this.Status != 0 { query = query.Where("status = ?", this.Status ) }
        if this.DelFlag != 0 { query = query.Where("del_flag = ?", this.DelFlag ) }
        if this.CompanyConfig != "" { query = query.Where("company_config = ?", this.CompanyConfig ) }
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

    if !satrtTime.IsZero() {
        query = query.Where("create_time >= ?", satrtTime)
    }
    if !endTime.IsZero() {
        query = query.Where("create_time <= ?", endTime)
    }

    // // 添加分页逻辑
    // if sysCompany.PageNum > 0 && sysCompany.PageSize > 0 {
    //     offset := (sysCompany.PageNum - 1) * sysCompany.PageSize
    //     query = query.Offset(offset).Limit(sysCompany.PageSize)
    // }

    err := query.Find(&tables).Error
    return tables, err
}

// FindSysCompanyPageList 分页查询公司列表
func (this *SysCompany) FindSysCompanyPageList(DB *gorm.DB, satrtTime time.Time, endTime time.Time, pageNum int, pageSize int) ([]SysCompany, int64, error) {
    fmt.Printf("GetSysCompanyPageList：%#v \n", this)

    var (
        sysCompanys []SysCompany
        total     int64
    )

    query := DB.Model(&SysCompany{})

// 构造查询条件
        if this.Id != 0 { query = query.Where("id = ?", this.Id ) }
        if this.CompanyName != "" { query = query.Where("company_name LIKE ?", "%" + this.CompanyName + "%") }
        if this.OrderNum != 0 { query = query.Where("order_num = ?", this.OrderNum ) }
        if this.Status != 0 { query = query.Where("status = ?", this.Status ) }
        if this.DelFlag != 0 { query = query.Where("del_flag = ?", this.DelFlag ) }
        if this.CompanyConfig != "" { query = query.Where("company_config = ?", this.CompanyConfig ) }
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
        Find(&sysCompanys).Error

    if err != nil {
        return nil, 0, err
    }

    return sysCompanys, total, nil
}

