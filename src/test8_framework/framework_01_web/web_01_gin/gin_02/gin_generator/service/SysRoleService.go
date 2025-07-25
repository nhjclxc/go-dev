package service

import (
    "gin_generator/internal/core"
    "gin_generator/model"
    "gin_generator/model/dto"
    "gin_generator/utils/commonUtils"
)

// SysRoleService 角色信息 Service 层
type SysRoleService struct {
}

// InsertSysRole 新增角色信息
func (this *SysRoleService) InsertSysRole(sysRole *model.SysRole) (res any, err error) {

    return sysRole.InsertSysRole(core.GLOBAL_DB)
}

// UpdateSysRole 修改角色信息
func (this *SysRoleService) UpdateSysRole(sysRole *model.SysRole) (res any, err error) {

    return sysRole.UpdateSysRoleByRoleId(core.GLOBAL_DB)
}

// DeleteSysRole 删除角色信息
func (this *SysRoleService) DeleteSysRole(roleIdList []int64) (res any, err error) {

    return (&model.SysRole{}).DeleteSysRole(core.GLOBAL_DB, roleIdList)
}

// GetSysRoleByRoleId 获取角色信息业务详细信息
func (this *SysRoleService) GetSysRoleByRoleId(roleId int64) (res any, err error) {

    sysRole := model.SysRole{}
    err = (&sysRole).FindSysRoleByRoleId(core.GLOBAL_DB, roleId)
    if err != nil {
        return nil, err
    }

    return sysRole, nil
}

// GetSysRoleList 查询角色信息业务列表
func (this *SysRoleService) GetSysRoleList(sysRoleDto *dto.SysRoleDto) (res any, err error) {

    sysRole, err := sysRoleDto.DtoToModel()
    sysRoleList, err := sysRole.FindSysRoleList(core.GLOBAL_DB, sysRoleDto.SatrtTime, sysRoleDto.EndTime)
    if err != nil {
        return nil, err
    }

    return sysRoleList, nil
}

// GetSysRolePageList 分页查询角色信息业务列表
func (this *SysRoleService) GetSysRolePageList(sysRoleDto *dto.SysRoleDto) (res any, err error) {

    sysRole, err := sysRoleDto.DtoToModel()
    sysRoleList, total, err := sysRole.FindSysRolePageList(core.GLOBAL_DB, sysRoleDto.SatrtTime, sysRoleDto.EndTime, sysRoleDto.PageSize, sysRoleDto.PageNum)
    if err != nil {
        return nil, err
    }

    return commonUtils.BuildPageData[model.SysRole](sysRoleList, total, sysRoleDto.PageNum, sysRoleDto.PageSize), nil
}

// ExportSysRole 导出角色信息业务列表
func (this *SysRoleService) ExportSysRole(sysRoleDto *dto.SysRoleDto) (res any, err error) {

    sysRole, err := sysRoleDto.DtoToModel()
    sysRole.FindSysRolePageList(core.GLOBAL_DB, sysRoleDto.SatrtTime, sysRoleDto.EndTime, 1, 10000)
    // 实现导出 ...

    return nil, nil
}
