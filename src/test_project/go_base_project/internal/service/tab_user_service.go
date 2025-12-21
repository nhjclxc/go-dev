package service

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go_base_project/internal/model"
	"go_base_project/internal/repository"
	"go_base_project/internal/request"
	"go_base_project/internal/response"
	excelutils "go_base_project/internal/utils/excel_utils"
	"gorm.io/gorm"
	"math"
)

// TabUserService 用户  Service 层
type TabUserService struct {
	db                *gorm.DB
	tabUserRepository *repository.TabUserRepository
}

// NewTabUserService 创建 TabUser 用户  业务层对象
func NewTabUserService(db *gorm.DB, tabUserRepository *repository.TabUserRepository) *TabUserService {
	return &TabUserService{
		db:                db,
		tabUserRepository: tabUserRepository,
	}
}

// InsertTabUser 新增用户
func (tus *TabUserService) InsertTabUser(ctx context.Context, tabUser model.TabUser) (int, error) {
	dbData, err := tus.tabUserRepository.SelectTabUserUnique(ctx, tabUser)
	if dbData != nil {
		return 0, fmt.Errorf("xxx已被占用")
	}
	if err != nil {
		return 0, err
	}

	return tus.tabUserRepository.InsertTabUser(ctx, &tabUser)
}

// UpdateTabUser 修改用户
func (tus *TabUserService) UpdateTabUser(ctx context.Context, tabUser model.TabUser) (int, error) {
	if tabUser.Id == 0 {
		return 0, fmt.Errorf("TabUserService.UpdateTabUser Id 不能为空！！！: ")
	}

	dbData, err := tus.tabUserRepository.SelectTabUserUnique(ctx, tabUser)
	if dbData != nil && dbData.Id != tabUser.Id {
		return 0, fmt.Errorf("xxx已被占用")
	}
	if err != nil {
		return 0, err
	}

	return tus.tabUserRepository.UpdateTabUserById(ctx, &tabUser)
}

// DeleteTabUser 删除用户
func (tus *TabUserService) DeleteTabUser(ctx context.Context, idList []int64) (int, error) {

	return tus.tabUserRepository.BatchDeleteTabUser(ctx, idList)
}

// GetTabUserById 获取用户 业务详细信息
func (tus *TabUserService) GetTabUserById(ctx context.Context, id int64) (*model.TabUser, error) {

	tabUser, err := tus.tabUserRepository.FindTabUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return tabUser, nil
}

// GetTabUserList 查询用户 业务列表
func (tus *TabUserService) GetTabUserList(ctx context.Context, tabUserReq request.TabUserReq) (res any, err error) {

	var tabUser = model.TabUser{}
	_ = model.CopyAttribute(&tabUser, tabUserReq)
	tabUserList, err := tus.tabUserRepository.FindTabUserList(ctx, &tabUser, tabUserReq.SatrtTime, tabUserReq.EndTime)
	if err != nil {
		return nil, err
	}

	return tabUserList, nil
}

// GetTabUserPageList 分页查询用户 业务列表
func (tus *TabUserService) GetTabUserPageList(ctx context.Context, tabUserReq request.TabUserReq) (res []*response.TabUserResp, err error) {

	var tabUser = model.TabUser{}
	_ = model.CopyAttribute(&tabUser, tabUserReq)
	tabUserList, total, err := tus.tabUserRepository.FindTabUserPageList(ctx, &tabUser, tabUserReq.SatrtTime, tabUserReq.EndTime, tabUserReq.PageNum, tabUserReq.PageSize)
	if err != nil {
		return nil, err
	}
	_ = total

	rest := make([]*response.TabUserResp, len(tabUserList))
	for _, tabUser := range tabUserList {
		resp := response.TabUserResp{}
		_ = model.CopyAttribute(&resp, tabUser)
	}
	return rest, nil
}

// ExportTabUser 导出用户 业务列表
func (tus *TabUserService) ExportTabUser(ctx context.Context, tabUserReq request.TabUserReq) (res *excelize.File, err error) {

	var tabUser = model.TabUser{}
	_ = model.CopyAttribute(&tabUser, tabUserReq)
	tabUserPageList, total, err := tus.tabUserRepository.FindTabUserPageList(ctx, &tabUser, tabUserReq.SatrtTime, tabUserReq.EndTime, 1, math.MaxInt64)
	if err != nil {
		return nil, err
	}
	// 实现导出 ...
	_ = total
	rest := make([]*response.TabUserExport, len(tabUserPageList))
	for _, tabUser := range tabUserPageList {
		resp := response.TabUserExport{}
		_ = model.CopyAttribute(&resp, tabUser)
	}

	headerKeys := []string{"Id", "Username", "Password", "Email", "CreatedAt", "UpdatedAt"}
	headerValues := []string{"用户ID", "用户名", "密码", "邮箱", "创建时间", "更新时间"}

	excelfile := excelutils.ExportExcel[*response.TabUserExport](rest, 0, 1, headerKeys, headerValues)
	return excelfile, nil
}
