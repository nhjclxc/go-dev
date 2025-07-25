package model

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/mitchellh/mapstructure"
    "gorm.io/gorm"
    "time"
)

// SysJob 定时任务调度 结构体
// @author 
// @date 2025-07-25T10:33:22.528
type SysJob struct {


    JobId int64 `gorm:"column:job_id;primaryKey;auto_increment;not null;type:bigint" json:"jobId" form:"jobId"`// 任务ID

    JobName string `gorm:"column:job_name;primaryKey;not null;type:varchar(64)" json:"jobName" form:"jobName"`// 任务名称

    JobGroup string `gorm:"column:job_group;primaryKey;not null;type:varchar(64)" json:"jobGroup" form:"jobGroup"`// 任务组名

    InvokeTarget string `gorm:"column:invoke_target;not null;type:varchar(500)" json:"invokeTarget" form:"invokeTarget"`// 调用目标字符串

    CronExpression string `gorm:"column:cron_expression;type:varchar(255)" json:"cronExpression" form:"cronExpression"`// cron执行表达式

    MisfirePolicy string `gorm:"column:misfire_policy;type:varchar(20)" json:"misfirePolicy" form:"misfirePolicy"`// 计划执行错误策略（1立即执行 2执行一次 3放弃执行）

    Concurrent string `gorm:"column:concurrent;type:char(1)" json:"concurrent" form:"concurrent"`// 是否并发执行（0允许 1禁止）

    Status string `gorm:"column:status;type:char(1)" json:"status" form:"status"`// 状态（0正常 1暂停）

    CreateBy string `gorm:"column:create_by;type:varchar(64)" json:"createBy" form:"createBy"`// 创建者

    CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"createTime" form:"createTime"`// 创建时间

    UpdateBy string `gorm:"column:update_by;type:varchar(64)" json:"updateBy" form:"updateBy"`// 更新者

    UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"updateTime" form:"updateTime"`// 更新时间

    Remark string `gorm:"column:remark;type:varchar(500)" json:"remark" form:"remark"`// 备注信息

    // time_format:"2006-01-02 15:04:05"
}

// TableName 返回当前实体类的表名
func (this *SysJob) TableName() string {
    return "sys_job"
}


// 可用钩子函数包括：
// BeforeCreate / AfterCreate
// BeforeUpdate / AfterUpdate
// BeforeDelete / AfterDelete
func (this *SysJob) BeforeCreate(tx *gorm.DB) (err error) {
    this.CreateTime = time.Now()
    this.UpdateTime = time.Now()
    return
}

func (this *SysJob) BeforeUpdate(tx *gorm.DB) (err error) {
    this.UpdateTime = time.Now()
    return
}

// MapToStruct map映射转化为当前结构体
func MapToSysJob(inputMap map[string]any) (*SysJob) {
    //go get github.com/mitchellh/mapstructure

    var sysJob SysJob
    err := mapstructure.Decode(inputMap, &sysJob)
    if err != nil {
        fmt.Printf("MapToStruct Decode error: %v", err)
        return nil
    }
    return &sysJob
}

// StructToMap 当前结构体转化为map映射
func (this *SysJob) SysJobToMap() (map[string]any) {
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

// InsertSysJob 新增定时任务调度
func (this *SysJob) InsertSysJob(DB *gorm.DB) (int, error) {
    fmt.Printf("InsertSysJob：%#v \n", this)

    // 先查询是否有相同 name 的数据存在
    temp := &SysJob{}
    // todo update name
    tx := DB.Where("job_name = ?", this.JobName).First(temp)
    fmt.Printf("InsertSysJob.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
        return 0, errors.New("InsertSysJob.Where, 存在相同name: " + temp.JobName)
    }

    // 执行 Insert
    err := DB.Create(&this).Error

    if err != nil {
        return 0, errors.New("InsertSysJob.DB.Create, 新增失败: " + err.Error())
    }
    return 1, nil
}

// UpdateSysJobByJobId 根据主键修改代码生成的所有字段
func (this *SysJob) UpdateSysJobByJobId(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateSysJobByJobId：%#v \n", this)

    // 1、查询该id是否存在
    if this.JobId == 0 {
        return 0, errors.New("JobId 不能为空！！！: ")
    }

    // 2、再看看name是否重复
    temp := &SysJob{}
    // todo update name
    tx := DB.Where("job_name = ?", this.JobName).First(temp)
    fmt.Printf("UpdateSysJobByJobId.Where：%#v \n", temp)
    if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && temp.JobId != this.JobId {
        return 0, errors.New("UpdateSysJobByJobId.Where, 存在相同name: " + temp.JobName)
    }

    // 3、执行修改
    //保存整个结构体（全字段更新）
    saveErr := DB.Save(this).Error
    if saveErr != nil {
        return 0, errors.New("UpdateSysJobByJobId.Save, 修改失败: " + saveErr.Error())
    }
    return 1, nil
}

// UpdateSysJobSelective 修改代码生成不为默认值的字段
func (this *SysJob) UpdateSysJobSelective(DB *gorm.DB) (int, error) {
    fmt.Printf("UpdateSysJobSelective：%#v \n", this)

    // db.Model().Updates()：只更新指定字段
    err := DB.Model(this).
        Where("job_id = ?", this.JobId).
        Updates(this).
        Error
    if err != nil {
        return 0, errors.New("UpdateSysJobSelective.Updates, 选择性修改失败: " + err.Error())
    }

    return 1, nil
}

// DeleteSysJob 删除代码生成业务
func (this *SysJob) DeleteSysJob(DB *gorm.DB, jobIdList []int64) (int, error) {
    fmt.Printf("DeleteSysJob：%#v \n", jobIdList)

    // 以下使用的是软删除，以下必须有DeletedAt gorm.DeletedAt字段
    result := DB.Delete(&this, "job_id in ?", jobIdList)
    // result := DB.Model(&this).Where("job_id IN ?", tableIdList).Update("state", 0)
    if result.Error != nil {
        return 0, errors.New("DeleteSysJob.Delete, 删除失败: " + result.Error.Error())
    }

    //// 以下使用的是物理删除
    //result := DB.Unscoped().Delete(this, "job_id in ?", jobIdList)
    //if result.Error != nil {
    //	return 0, errors.New("DeleteSysJob.Delete, 删除失败: " + result.Error.Error())
    //}

    return int(result.RowsAffected), nil
}

// FindSysJobByJobId 获取代码生成业务业务详细信息
func (this *SysJob) FindSysJobByJobId(DB *gorm.DB, jobId int64) (error) {
    fmt.Printf("DeleteSysJob：%#v \n", jobId)
    return DB.First(this, "job_id = ?", jobId).Error
}

// FindSysJobList 查询代码生成业务业务列表
func (this *SysJob) FindSysJobList(DB *gorm.DB, satrtTime time.Time, endTime time.Time) ([]SysJob, error) {
    fmt.Printf("GetSysJobList：%#v \n", this)

    var tables []SysJob
    query := DB.Model(this)

    // 构造查询条件
        if this.JobId != 0 { query = query.Where("job_id = ?", this.JobId ) }
        if this.JobName != "" { query = query.Where("job_name LIKE ?", "%" + this.JobName + "%") }
        if this.JobGroup != "" { query = query.Where("job_group = ?", this.JobGroup ) }
        if this.InvokeTarget != "" { query = query.Where("invoke_target = ?", this.InvokeTarget ) }
        if this.CronExpression != "" { query = query.Where("cron_expression = ?", this.CronExpression ) }
        if this.MisfirePolicy != "" { query = query.Where("misfire_policy = ?", this.MisfirePolicy ) }
        if this.Concurrent != "" { query = query.Where("concurrent = ?", this.Concurrent ) }
        if this.Status != "" { query = query.Where("status = ?", this.Status ) }
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
    // if sysJob.PageNum > 0 && sysJob.PageSize > 0 {
    //     offset := (sysJob.PageNum - 1) * sysJob.PageSize
    //     query = query.Offset(offset).Limit(sysJob.PageSize)
    // }

    err := query.Find(&tables).Error
    return tables, err
}

// FindSysJobPageList 分页查询代码生成业务业务列表
func (this *SysJob) FindSysJobPageList(DB *gorm.DB, satrtTime time.Time, endTime time.Time, pageNum int, pageSize int) ([]SysJob, int64, error) {
    fmt.Printf("GetSysJobPageList：%#v \n", this)

    var (
        sysJobs []SysJob
        total     int64
    )

    query := DB.Model(&SysJob{})

// 构造查询条件
        if this.JobId != 0 { query = query.Where("job_id = ?", this.JobId ) }
        if this.JobName != "" { query = query.Where("job_name LIKE ?", "%" + this.JobName + "%") }
        if this.JobGroup != "" { query = query.Where("job_group = ?", this.JobGroup ) }
        if this.InvokeTarget != "" { query = query.Where("invoke_target = ?", this.InvokeTarget ) }
        if this.CronExpression != "" { query = query.Where("cron_expression = ?", this.CronExpression ) }
        if this.MisfirePolicy != "" { query = query.Where("misfire_policy = ?", this.MisfirePolicy ) }
        if this.Concurrent != "" { query = query.Where("concurrent = ?", this.Concurrent ) }
        if this.Status != "" { query = query.Where("status = ?", this.Status ) }
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
    // todo update created_time
    err := query.
        Limit(pageSize).Offset((pageNum - 1) * pageSize).
        Order("created_time desc").
        Find(&sysJobs).Error

    if err != nil {
        return nil, 0, err
    }

    return sysJobs, total, nil
}

