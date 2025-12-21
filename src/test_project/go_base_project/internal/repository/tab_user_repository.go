package repository

import (
	"context"
	"errors"
	"fmt"
	"go_base_project/internal/model"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

// TabUserRepository 用户  结构体持久层
type TabUserRepository struct {
	db *gorm.DB
}

// NewTabUserRepository 创建 TabUser 用户  持久层对象
func NewTabUserRepository(db *gorm.DB) *TabUserRepository {
	return &TabUserRepository{
		db: db,
	}
}

// 由于有时需要开启事务，因此 DB *gorm.DB 可以选择从外部传入

// InsertTabUser 新增用户
func (repo *TabUserRepository) InsertTabUser(ctx context.Context, tabUser *model.TabUser) (int, error) {
	slog.Info("TabUserRepository.InsertTabUser：", slog.Any("tabUser", tabUser))

	result := repo.db.WithContext(ctx).Model(&model.TabUser{}).Create(&tabUser)

	if result.Error != nil {
		return 0, fmt.Errorf("TabUserRepository.InsertTabUserCreate, 新增失败: %w ", result.Error)
	}
	return int(result.RowsAffected), nil
}

// BatchInsertTabUsers 批量新增用户
func (repo *TabUserRepository) BatchInsertTabUsers(ctx context.Context, tabUsers []*model.TabUser) (int, error) {
	slog.Info("TabUserRepository.BatchInsertTabUsers：", slog.Any("tabUsers", tabUsers))

	result := repo.db.WithContext(ctx).Model(&model.TabUser{}).Create(&tabUsers)
	if result.Error != nil {
		return 0, fmt.Errorf("TabUserRepository.BatchInsertTabUsers.Create, 新增失败: %w ", result.Error)
	}
	return int(result.RowsAffected), nil
}

// UpdateTabUserById 根据主键修改用户 的所有字段
func (repo *TabUserRepository) UpdateTabUserById(ctx context.Context, tabUser *model.TabUser) (int, error) {
	slog.Info("TabUserRepository.UpdateTabUserById：", slog.Any("tabUser", tabUser))

	//保存整个结构体（全字段更新）
	result := repo.db.WithContext(ctx).Model(&model.TabUser{}).Save(tabUser)
	if result.Error != nil {
		return 0, fmt.Errorf("TabUserRepository.UpdateTabUserById.Save, 修改失败: %w ", result.Error)
	}
	return int(result.RowsAffected), nil
}

// UpdateTabUserSelective 修改用户 不为默认值的字段
func (repo *TabUserRepository) UpdateTabUserSelective(ctx context.Context, updateColumns []string, tabUser *model.TabUser) (int, error) {
	slog.Info("TabUserRepository.UpdateTabUserSelective：", slog.Any("tabUser", tabUser))

	if len(updateColumns) == 0 {
		updateColumns = []string{"username", "password", "email", "created_at", "updated_at"}
	}
	// repo.db.WithContext(ctx).Model(&model.TabUser{}).Updates()：只更新指定字段
	result := repo.db.WithContext(ctx).Model(&model.TabUser{}).
		Where("id = ?", tabUser.Id).
		Select(updateColumns).
		Updates(tabUser)
	if result.Error != nil {
		return 0, fmt.Errorf("TabUserRepository.UpdateTabUserSelective.Updates, 选择性修改失败: %w ", result.Error)
	}

	return int(result.RowsAffected), nil
}

// BatchUpdateTabUserSelective 批量修改用户
func (repo *TabUserRepository) BatchUpdateTabUserSelective(ctx context.Context, updateColumns []string, tabUsers []model.TabUser) error {
	if len(updateColumns) == 0 {
		updateColumns = []string{"username", "password", "email", "created_at", "updated_at"}
	}
	return repo.db.WithContext(ctx).Model(&model.TabUser{}).Transaction(func(tx *gorm.DB) error {
		for _, v := range tabUsers {
			err := tx.Model(&model.TabUser{}).Where("id = ?", v.Id).Select(updateColumns).Updates(v).Error
			if err != nil {
				return err // 触发回滚
			}
		}
		return nil // 提交事务
	})
}

// BatchDeleteTabUserByState 批量软删除 用户
func (repo *TabUserRepository) BatchDeleteTabUserByState(ctx context.Context, idList []int64) error {
	slog.Info("TabUserRepository.BatchDeleteTabUserByState：", slog.Any("idList", idList))

	return repo.db.WithContext(ctx).Model(&model.TabUser{}).Model(&model.TabUser{}).
		Where("id IN ?", idList).
		Updates(map[string]any{
			"state":      0,
			"deleted_at": time.Now(),
		}).Error
}

// BatchDeleteTabUser 根据主键批量删除 用户
func (repo *TabUserRepository) BatchDeleteTabUser(ctx context.Context, idList []int64) (int, error) {
	slog.Info("TabUserRepository.BatchDeleteTabUser：", slog.Any("idList", idList))

	// 当存在DeletedAt gorm.DeletedAt字段时为软删除，否则为物理删除
	result := repo.db.WithContext(ctx).Model(&model.TabUser{}).Where("id IN ?", idList).Delete(&model.TabUser{})
	// result := repo.db.WithContext(ctx).Model(&model.TabUser{}).Model(&model.TabUser{}).Where("id IN ?", idList).Update("state", 0)
	if result.Error != nil {
		return 0, fmt.Errorf("TabUserRepository.BatchDeleteTabUser.Delete, 删除失败: %w ", result.Error)
	}

	//// 以下使用的是物理删除
	//result := repo.db.WithContext(ctx).Model(&model.TabUser{}).Unscoped().Delete(&model.TabUser{}, "id IN ?", idList)
	//if result.Error != nil {
	//	return 0, fmt.Errorf("TabUserRepository.BatchDeleteTabUser.Delete, 删除失败: %w ", result.Error)
	//}

	return int(result.RowsAffected), nil
}

// FindTabUserById 获取用户 详细信息
func (repo *TabUserRepository) FindTabUserById(ctx context.Context, id int64) (*model.TabUser, error) {
	slog.Info("TabUserRepository.FindTabUserById：", slog.Any("id", id))

	tabUser := model.TabUser{}
	err := repo.db.WithContext(ctx).Model(&model.TabUser{}).First(&tabUser, "id = ?", id).Error
	return &tabUser, err
}

// FindTabUsersByIdList 根据主键批量查询用户 详细信息
func (repo *TabUserRepository) FindTabUsersByIdList(ctx context.Context, idList []int64) ([]*model.TabUser, error) {
	slog.Info("TabUserRepository.FindTabUsersByIdList：", slog.Any("idList", idList))

	var result []*model.TabUser
	err := repo.db.WithContext(ctx).Model(&model.TabUser{}).Where("id IN ?", idList).Find(&result).Error
	return result, err
}

// SelectTabUserUnique 查询用户 的唯一记录
func (repo *TabUserRepository) SelectTabUserUnique(ctx context.Context, tabUser model.TabUser) (*model.TabUser, error) {
	slog.Info("TabUserRepository.SelectTabUserUnique：", slog.Any("tabUser", tabUser))

	var temp model.TabUser
	query := repo.db.WithContext(ctx).Model(&model.TabUser{})

	// 构造查询条件
	if tabUser.Id != 0 {
		query = query.Where("id = ?", tabUser.Id)
	}
	if tabUser.Username != "" {
		query = query.Where("username LIKE ?", "%"+tabUser.Username+"%")
	}
	if tabUser.Password != "" {
		query = query.Where("password = ?", tabUser.Password)
	}
	if tabUser.Email != "" {
		query = query.Where("email = ?", tabUser.Email)
	}
	if !tabUser.CreatedAt.IsZero() {
		query = query.Where("created_at = ?", tabUser.CreatedAt)
		// query = query.Where("DATE(created_at) = ?", tabUser.$column.goField.Format("2006-01-02"))
	}
	if !tabUser.UpdatedAt.IsZero() {
		query = query.Where("updated_at = ?", tabUser.UpdatedAt)
		// query = query.Where("DATE(updated_at) = ?", tabUser.$column.goField.Format("2006-01-02"))
	}

	tx := query.First(&temp)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			// 没有数据，返回 nil
			return nil, nil
		}
		// 其他错误，直接返回
		return nil, tx.Error
	}

	// 找到数据，返回对象
	return &temp, nil
}

// FindTabUserList 查询用户 列表
func (repo *TabUserRepository) FindTabUserList(ctx context.Context, tabUser *model.TabUser, startTime time.Time, endTime time.Time) ([]*model.TabUser, error) {
	slog.Info("TabUserRepository.FindTabUserList：", slog.Any("tabUser", tabUser))

	var tabUsers []*model.TabUser
	query := repo.db.WithContext(ctx).Model(&model.TabUser{})

	// 构造查询条件
	if tabUser.Id != 0 {
		query = query.Where("id = ?", tabUser.Id)
	}
	if tabUser.Username != "" {
		query = query.Where("username LIKE ?", "%"+tabUser.Username+"%")
	}
	if tabUser.Password != "" {
		query = query.Where("password = ?", tabUser.Password)
	}
	if tabUser.Email != "" {
		query = query.Where("email = ?", tabUser.Email)
	}
	if !tabUser.CreatedAt.IsZero() {
		query = query.Where("created_at = ?", tabUser.CreatedAt)
		// query = query.Where("DATE(created_at) = ?", tabUser.$column.goField.Format("2006-01-02"))
	}
	if !tabUser.UpdatedAt.IsZero() {
		query = query.Where("updated_at = ?", tabUser.UpdatedAt)
		// query = query.Where("DATE(updated_at) = ?", tabUser.$column.goField.Format("2006-01-02"))
	}

	if !startTime.IsZero() {
		query = query.Where("created_at >= ?", startTime)
	}
	if !endTime.IsZero() {
		query = query.Where("created_at <= ?", endTime)
	}

	err := query.Find(&tabUsers).Error
	return tabUsers, err
}

// FindTabUserPageList 分页查询用户 列表
func (repo *TabUserRepository) FindTabUserPageList(ctx context.Context, tabUser *model.TabUser, startTime time.Time, endTime time.Time, pageNum int, pageSize int) ([]*model.TabUser, int64, error) {
	slog.Info("TabUserRepository.FindTabUserPageList：", slog.Any("tabUser", tabUser))

	var (
		tabUsers []*model.TabUser
		total    int64
	)

	query := repo.db.WithContext(ctx).Model(&model.TabUser{})

	// 构造查询条件
	if tabUser.Id != 0 {
		query = query.Where("id = ?", tabUser.Id)
	}
	if tabUser.Username != "" {
		query = query.Where("username LIKE ?", "%"+tabUser.Username+"%")
	}
	if tabUser.Password != "" {
		query = query.Where("password = ?", tabUser.Password)
	}
	if tabUser.Email != "" {
		query = query.Where("email = ?", tabUser.Email)
	}
	if !tabUser.CreatedAt.IsZero() {
		query = query.Where("created_at = ?", tabUser.CreatedAt)
		// query = query.Where("DATE(created_at) = ?", tabUser.$column.goField.Format("2006-01-02"))
	}
	if !tabUser.UpdatedAt.IsZero() {
		query = query.Where("updated_at = ?", tabUser.UpdatedAt)
		// query = query.Where("DATE(updated_at) = ?", tabUser.$column.goField.Format("2006-01-02"))
	}

	if !startTime.IsZero() {
		query = query.Where("created_at >= ?", startTime)
	}
	if !endTime.IsZero() {
		query = query.Where("created_at <= ?", endTime)
	}

	// 分页参数默认值
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 先统计总数，Session()创建新会话防止Count()破坏Order()、Preload()、Group()等等
	err := query.Session(&gorm.Session{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 再进行分页查询
	err = query.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("id desc").Find(&tabUsers).Error
	if err != nil {
		return nil, 0, err
	}

	return tabUsers, total, nil
}
