package model

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/mitchellh/mapstructure"
    "gorm.io/gorm"
    "time"
)

// SysRole 角色信息 结构体
// @author 
// @date 2025-07-25T11:09:41.517
type SysRole struct {


    RoleId int64 `gorm:"column:role_id;primaryKey;auto_increment;not null;type:bigint" json:"roleId" form:"roleId"`// 角色ID

    RoleName string `gorm:"column:role_name;not null;type:varchar(30)" json:"roleName" form:"roleName"`// 角色名称

    RoleKey string `gorm:"column:role_key;not null;type:varchar(100)" json:"roleKey" form:"roleKey"`// 角色权限字符串

    RoleSort int `gorm:"column:role_sort;not null;type:int" json:"roleSort" form:"roleSort"`// 显示顺序

    DataScope string `gorm:"column:data_scope;type:char(1)" json:"dataScope" form:"dataScope"`// 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）

    MenuCheckStrictly int `gorm:"column:menu_check_strictly;type:tinyint(1)" json:"menuCheckStrictly" form:"menuCheckStrictly"`// 菜单树选择项是否关联显示

    DeptCheckStrictly int `gorm:"column:dept_check_strictly;type:tinyint(1)" json:"deptCheckStrictly" form:"deptCheckStrictly"`// 部门树选择项是否关联显示

    Status string `gorm:"column:status;not null;type:char(1)" json:"status" form:"status"`// 角色状态（0正常 1停用）

    DelFlag string `gorm:"column:del_flag;type:char(1)" json:"delFlag" form:"delFlag"`// 删除标志（0代表存在 2代表删除）

    CreateBy string `gorm:"column:create_by;type:varchar(64)" json:"createBy" form:"createBy"`// 创建者

    CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"createTime" form:"createTime"`// 创建时间

    UpdateBy string `gorm:"column:update_by;type:varchar(64)" json:"updateBy" form:"updateBy"`// 更新者

    UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"updateTime" form:"updateTime"`// 更新时间

    Remark string `gorm:"column:remark;type:varchar(500)" json:"remark" form:"remark"`// 备注

    // time_format:"2006-01-02 15:04:05"
}

// TableName 返回当前实体类的表名
func (this *SysRole) TableName() string {
    return "sys_role"
}


// 可用钩子函数包括：
// BeforeCreate / AfterCreate
// BeforeUpdate / AfterUpdate
// BeforeDelete / AfterDelete
func (this *SysRole) BeforeCreate(tx *gorm.DB) (err error) {
    this.CreateTime = time.Now()
    this.UpdateTime = time.Now()
    return
}

func (this *SysRole) BeforeUpdate(tx *gorm.DB) (err error) {
    this.UpdateTime = time.Now()
    return
}

// MapToSysRole map映射转化为当前结构体
func MapToSysRole(inputMap map[string]any) (*SysRole) {
    //go get github.com/mitchellh/mapstructure

    var sysRole SysRole
    err := mapstructure.Decode(inputMap, &sysRole)
    if err != nil {
        fmt.Printf("MapToStruct Decode error: %v", err)
        return nil
    }
    return &sysRole
}

// SysRoleToMap 当前结构体转化为map映射
func (this *SysRole) SysRoleToMap() (map[string]any) {
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

// InsertSysRole 新增角色信息
func (this *SysRole) InsertSysRole(DB *gorm.DB) (int, error) {
    fmt.Printf("InsertSysRole：%#v \n", this)

    // 先查询是否有相同 name 的数据存在
    temp := &SysRole{}
    // todo update name
    tx := DB.Where("role_name = ?", this.RoleName).First(temp)
    fmt.Printf("InsertSysRole.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
        return 0, errors.New("InsertSysRole.Where, 存在相同name: " + temp.RoleName)
    }

    // 执行 Insert
    err := DB.Create(&this).Error

    if err != nil {
        return 0, errors.New("InsertSysRole.DB.Create, 新增失败: " + err.Error())
    }
    return 1, nil
}

// BatchInsertSysRoles 批量新增角色信息
func (this *SysRole) BatchInsertSysRoles(DB *gorm.DB, tables []SysRole) (int, error) {

    result := DB.Create(&tables)

    if result.Error != nil {
        return 0, errors.New("BatchInsertSysRoles.DB.Create, 新增失败: " + result.Error.Error())
    }
    return int(result.RowsAffected), nil
}

// UpdateSysRoleByRoleId 根据主键修改角色信息的所有字段
func (this *SysRole) UpdateSysRoleByRoleId(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateSysRoleByRoleId：%#v \n", this)

    // 1、查询该id是否存在
    if this.RoleId == 0 {
        return 0, errors.New("RoleId 不能为空！！！: ")
    }

    // 2、再看看name是否重复
    temp := &SysRole{}
    // todo update name
    tx := DB.Where("role_name = ?", this.RoleName).First(temp)
    fmt.Printf("UpdateSysRoleByRoleId.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && temp.RoleId != this.RoleId {
        return 0, errors.New("UpdateSysRoleByRoleId.Where, 存在相同name: " + temp.RoleName)
    }

    // 3、执行修改
    //保存整个结构体（全字段更新）
    saveErr := DB.Save(this).Error
    if saveErr != nil {
        return 0, errors.New("UpdateSysRoleByRoleId.Save, 修改失败: " + saveErr.Error())
    }
    return 1, nil
}

// UpdateSysRoleSelective 修改角色信息不为默认值的字段
func (this *SysRole) UpdateSysRoleSelective(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateSysRoleSelective：%#v \n", this)

    // db.Model().Updates()：只更新指定字段
    err := DB.Model(this).
        Where("role_id = ?", this.RoleId).
        Updates(this).
        Error
    if err != nil {
        return 0, errors.New("UpdateSysRoleSelective.Updates, 选择性修改失败: " + err.Error())
    }

    return 1, nil
}

// DeleteSysRole 删除角色信息
func (this *SysRole) DeleteSysRole(DB *gorm.DB, roleIdList []int64) (int, error) {
    fmt.Printf("DeleteSysRole：%#v \n", roleIdList)

    // 以下使用的是软删除，以下必须有DeletedAt gorm.DeletedAt字段
    result := DB.Delete(&this, "role_id in ?", roleIdList)
    // result := DB.Model(&this).Where("role_id IN ?", tableIdList).Update("state", 0)
    if result.Error != nil {
        return 0, errors.New("DeleteSysRole.Delete, 删除失败: " + result.Error.Error())
    }

    //// 以下使用的是物理删除
    //result := DB.Unscoped().Delete(this, "role_id in ?", roleIdList)
    //if result.Error != nil {
    //	return 0, errors.New("DeleteSysRole.Delete, 删除失败: " + result.Error.Error())
    //}

    return int(result.RowsAffected), nil
}

// BatchDeleteSysRoles 根据主键批量删除角色信息
func (this *SysRole) BatchDeleteSysRoles(DB *gorm.DB, roleIdList []int64) error {
    return DB.Where("role_id IN ?", roleIdList).Delete(&this).Error
}

// FindSysRoleByRoleId 获取角色信息详细信息
func (this *SysRole) FindSysRoleByRoleId(DB *gorm.DB, roleId int64) (error) {
    fmt.Printf("DeleteSysRole：%#v \n", roleId)
    return DB.First(this, "role_id = ?", roleId).Error
}

// FindSysRolesByRoleIdList 根据主键批量查询角色信息详细信息
func FindSysRolesByRoleIdList(DB *gorm.DB, roleIdList []int64) ([]SysRole, error) {
    var result []SysRole
    err := DB.Where("id IN ?", roleIdList).Find(&result).Error
    return result, err
}

// FindSysRoleList 查询角色信息列表
func (this *SysRole) FindSysRoleList(DB *gorm.DB, satrtTime time.Time, endTime time.Time) ([]SysRole, error) {
    fmt.Printf("GetSysRoleList：%#v \n", this)

    var tables []SysRole
    query := DB.Model(this)

    // 构造查询条件
        if this.RoleId != 0 { query = query.Where("role_id = ?", this.RoleId ) }
        if this.RoleName != "" { query = query.Where("role_name LIKE ?", "%" + this.RoleName + "%") }
        if this.RoleKey != "" { query = query.Where("role_key = ?", this.RoleKey ) }
        if this.RoleSort != 0 { query = query.Where("role_sort = ?", this.RoleSort ) }
        if this.DataScope != "" { query = query.Where("data_scope = ?", this.DataScope ) }
        if this.MenuCheckStrictly != 0 { query = query.Where("menu_check_strictly = ?", this.MenuCheckStrictly ) }
        if this.DeptCheckStrictly != 0 { query = query.Where("dept_check_strictly = ?", this.DeptCheckStrictly ) }
        if this.Status != "" { query = query.Where("status = ?", this.Status ) }
        if this.DelFlag != "" { query = query.Where("del_flag = ?", this.DelFlag ) }
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
        if this.Remark != "" { query = query.Where("remark LIKE ?", "%" + this.Remark + "%") }

    if !satrtTime.IsZero() {
        query = query.Where("create_time >= ?", satrtTime)
    }
    if !endTime.IsZero() {
        query = query.Where("create_time <= ?", endTime)
    }

    // // 添加分页逻辑
    // if sysRole.PageNum > 0 && sysRole.PageSize > 0 {
    //     offset := (sysRole.PageNum - 1) * sysRole.PageSize
    //     query = query.Offset(offset).Limit(sysRole.PageSize)
    // }

    err := query.Find(&tables).Error
    return tables, err
}

// FindSysRolePageList 分页查询角色信息列表
func (this *SysRole) FindSysRolePageList(DB *gorm.DB, satrtTime time.Time, endTime time.Time, pageNum int, pageSize int) ([]SysRole, int64, error) {
    fmt.Printf("GetSysRolePageList：%#v \n", this)

    var (
        sysRoles []SysRole
        total     int64
    )

    query := DB.Model(&SysRole{})

// 构造查询条件
        if this.RoleId != 0 { query = query.Where("role_id = ?", this.RoleId ) }
        if this.RoleName != "" { query = query.Where("role_name LIKE ?", "%" + this.RoleName + "%") }
        if this.RoleKey != "" { query = query.Where("role_key = ?", this.RoleKey ) }
        if this.RoleSort != 0 { query = query.Where("role_sort = ?", this.RoleSort ) }
        if this.DataScope != "" { query = query.Where("data_scope = ?", this.DataScope ) }
        if this.MenuCheckStrictly != 0 { query = query.Where("menu_check_strictly = ?", this.MenuCheckStrictly ) }
        if this.DeptCheckStrictly != 0 { query = query.Where("dept_check_strictly = ?", this.DeptCheckStrictly ) }
        if this.Status != "" { query = query.Where("status = ?", this.Status ) }
        if this.DelFlag != "" { query = query.Where("del_flag = ?", this.DelFlag ) }
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
        if this.Remark != "" { query = query.Where("remark LIKE ?", "%" + this.Remark + "%") }

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
    // todo update create_time
    err := query.
        Limit(pageSize).Offset((pageNum - 1) * pageSize).
        Order("create_time desc").
        Find(&sysRoles).Error

    if err != nil {
        return nil, 0, err
    }

    return sysRoles, total, nil
}

