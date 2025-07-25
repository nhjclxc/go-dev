package service

import (
    "gin_generator/internal/core"
    "gin_generator/model"
    "gin_generator/model/dto"
)

// SysJobService 定时任务调度 Service 层
type SysJobService struct {
}

// InsertSysJob 新增定时任务调度
func (this *SysJobService) InsertSysJob(sysJob *model.SysJob) (res any, err error) {

    return sysJob.InsertSysJob(core.GLOBAL_DB)
}

// UpdateSysJob 修改定时任务调度
func (this *SysJobService) UpdateSysJob(sysJob *model.SysJob) (res any, err error) {

    return sysJob.UpdateSysJobByJobId(core.GLOBAL_DB)
}

// DeleteSysJob 删除定时任务调度
func (this *SysJobService) DeleteSysJob(jobIdList []int64) (res any, err error) {

    return (&model.SysJob{}).DeleteSysJob(core.GLOBAL_DB, jobIdList)
}

// GetSysJobByJobId 获取定时任务调度业务详细信息
func (this *SysJobService) GetSysJobByJobId(jobId int64) (res any, err error) {

    sysJob := model.SysJob{}
    err = (&sysJob).FindSysJobByJobId(core.GLOBAL_DB, jobId)
    if err != nil {
        return nil, err
    }

    return sysJob, nil
}

// GetSysJobList 查询定时任务调度业务列表
func (this *SysJobService) GetSysJobList(sysJobDto *dto.SysJobDto) (res any, err error) {

    sysJob, err := sysJobDto.DtoToModel()
    sysJobList, err := sysJob.FindSysJobList(core.GLOBAL_DB, sysJobDto.SatrtTime, sysJobDto.EndTime)
    if err != nil {
        return nil, err
    }

    return sysJobList, nil
}

// GetSysJobPageList 分页查询定时任务调度业务列表
func (this *SysJobService) GetSysJobPageList(sysJobDto *dto.SysJobDto) (res any, err error) {

    sysJob, err := sysJobDto.DtoToModel()
    sysJobList, total, err := sysJob.FindSysJobPageList(core.GLOBAL_DB, sysJobDto.SatrtTime, sysJobDto.EndTime, sysJobDto.PageSize, sysJobDto.PageNum)
    if err != nil {
        return nil, err
    }

    return BuildPageData[model.SysJob](sysJobList, total, sysJobDto.PageNum, sysJobDto.PageSize), nil
}

// ExportSysJob 导出定时任务调度业务列表
func (this *SysJobService) ExportSysJob(sysJobDto *dto.SysJobDto) (res any, err error) {

    sysJob, err := sysJobDto.DtoToModel()
    sysJob.FindSysJobPageList(core.GLOBAL_DB, sysJobDto.SatrtTime, sysJobDto.EndTime, 1, 10000)
    // 实现导出 ...

    return nil, nil
}
