package service

import (
    "gin_generator/internal/core"
    "gin_generator/model"
    "gin_generator/model/dto"
    "gin_generator/utils/commonUtils"
)

// SysCompanyService 公司 Service 层
type SysCompanyService struct {
}

// InsertSysCompany 新增公司
func (this *SysCompanyService) InsertSysCompany(sysCompany *model.SysCompany) (res any, err error) {

    return sysCompany.InsertSysCompany(core.GLOBAL_DB)
}

// UpdateSysCompany 修改公司
func (this *SysCompanyService) UpdateSysCompany(sysCompany *model.SysCompany) (res any, err error) {

    return sysCompany.UpdateSysCompanyById(core.GLOBAL_DB)
}

// DeleteSysCompany 删除公司
func (this *SysCompanyService) DeleteSysCompany(idList []int64) (res any, err error) {

    return (&model.SysCompany{}).DeleteSysCompany(core.GLOBAL_DB, idList)
}

// GetSysCompanyById 获取公司业务详细信息
func (this *SysCompanyService) GetSysCompanyById(id int64) (res any, err error) {

    sysCompany := model.SysCompany{}
    err = (&sysCompany).FindSysCompanyById(core.GLOBAL_DB, id)
    if err != nil {
        return nil, err
    }

    return sysCompany, nil
}

// GetSysCompanyList 查询公司业务列表
func (this *SysCompanyService) GetSysCompanyList(sysCompanyDto *dto.SysCompanyDto) (res any, err error) {

    sysCompany, err := sysCompanyDto.DtoToModel()
    sysCompanyList, err := sysCompany.FindSysCompanyList(core.GLOBAL_DB, sysCompanyDto.SatrtTime, sysCompanyDto.EndTime)
    if err != nil {
        return nil, err
    }

    return sysCompanyList, nil
}

// GetSysCompanyPageList 分页查询公司业务列表
func (this *SysCompanyService) GetSysCompanyPageList(sysCompanyDto *dto.SysCompanyDto) (res any, err error) {

    sysCompany, err := sysCompanyDto.DtoToModel()
    sysCompanyList, total, err := sysCompany.FindSysCompanyPageList(core.GLOBAL_DB, sysCompanyDto.SatrtTime, sysCompanyDto.EndTime, sysCompanyDto.PageNum, sysCompanyDto.PageSize)
    if err != nil {
        return nil, err
    }

    return commonUtils.BuildPageData[model.SysCompany](sysCompanyList, total, sysCompanyDto.PageNum, sysCompanyDto.PageSize), nil
}

// ExportSysCompany 导出公司业务列表
func (this *SysCompanyService) ExportSysCompany(sysCompanyDto *dto.SysCompanyDto) (res any, err error) {

    sysCompany, err := sysCompanyDto.DtoToModel()
    sysCompany.FindSysCompanyPageList(core.GLOBAL_DB, sysCompanyDto.SatrtTime, sysCompanyDto.EndTime, 1, 10000)
    // 实现导出 ...

    return nil, nil
}
