package model

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

/*
CREATE TABLE `tab_user`
(
    `user_id`       BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Standard field for the primary key',
    `name`          VARCHAR(255)     NOT NULL COMMENT 'A regular string field',
    `email`         VARCHAR(255)              DEFAULT NULL COMMENT 'A pointer to a string, allowing for null values',
    `age`           TINYINT UNSIGNED NOT NULL COMMENT 'An unsigned 8-bit integer',
    `birthday`      DATETIME                  DEFAULT NULL COMMENT 'A pointer to time.Time, can be null',
    `member_number` VARCHAR(255)              DEFAULT NULL COMMENT 'Uses sql.NullString to handle nullable strings',
    `remark` VARCHAR(128)              DEFAULT NULL COMMENT 'fields that aren't exported are ignored',
    `activated_at`  DATETIME                  DEFAULT NULL COMMENT 'Uses sql.NullTime for nullable time fields',
    `created_at`    DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Automatically managed by GORM for creation time',
    `updated_at`    DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Automatically managed by GORM for update time'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='tab_user';
*/

// 更多关于 GORM 模型的详细信息请查看：https://gorm.io/zh_CN/docs/models.html
// 用户结构体
type TabUser struct {
	// Standard field for the primary key
	UserId uint `gorm:"primaryKey;autoIncrement;column:user_id"`

	// A regular string field
	Name string `gorm:"column:name"`

	// A pointer to a string, allowing for null values
	Email string `gorm:"column:email"`

	// An unsigned 8-bit integer
	Age uint8 `gorm:"column:age;default:18"`

	// A pointer to time.Time, can be null
	Birthday time.Time `gorm:"column:birthday"`

	// Uses sql.NullString to handle nullable strings
	MemberNumber sql.NullString `gorm:"column:member_number"`

	// fields that aren't exported are ignored
	Remark string `gorm:"column:remark"`

	// GORM 约定使用 CreatedAt、UpdatedAt 追踪创建/更新时间。如果您定义了这种字段，GORM 在创建、更新时会自动填充 当前时间
	//
	//要使用不同名称的字段，您可以配置 autoCreateTime、autoUpdateTime 标签。
	//
	//如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time，您只需简单地将 time.Time 修改为 int 即可
	// Uses sql.NullTime for nullable time fields
	ActivatedAt sql.NullTime `gorm:"autoUpdateTime;column:activated_at"`

	// Automatically managed by GORM for creation time
	CreatedAt time.Time `gorm:"column:created_at"`

	// Automatically managed by GORM for update time
	UpdatedAt time.Time `gorm:"column:updated_at"`

	// 为了实现关联创建，这里必须把子表的结构体传入
	// foreignKey:UserId 里面的 UserId 是子表的 关联id
	// references:UserId 里面的 UserId 是主表的 关联id
	TabUserCard TabUserCard `gorm:"foreignKey:UserId;references:UserId"`
}

func (TabUser) TableName() string {
	return "tab_user"
}

// 字段标签 tag tags 的相关信息，可以查看[字段标签](https://gorm.io/zh_CN/docs/models.html)
// 如果是实体之间的关联标签 ，则移步至[实体关联](https://gorm.io/zh_CN/docs/associations.html#tags)查看更多信息

// 对象生命周期
//Hook 是在创建、查询、更新、删除等操作之前、之后调用的函数。
//如果您已经为模型定义了指定的方法，它会在创建、更新、查询、删除时自动被调用。如果任何回调返回错误，GORM 将停止后续的操作并回滚事务。
//钩子方法的函数签名应该是 func(*gorm.DB) error

/*
创建时可用的 hook

	// 开始事务
	BeforeSave
	BeforeCreate
	// 关联前的 save
	// 插入记录至 db
	// 关联后的 save
	AfterCreate
	AfterSave
	// 提交或回滚事务


*/

// BeforeCreate 插入数据前的一些操作
func (u *TabUser) BeforeCreate(tx *gorm.DB) (err error) {
	u.Remark = uuid.New().String()

	// 判断 u.Birthday 是否是默认值
	if u.Birthday.IsZero() {
		// 是默认值，则使用当前时间
		u.Birthday = time.Now()
	}

	fmt.Println("BeforeCreate Hook 被执行 ", u.Remark)

	//if u.Role == "admin" {
	//	return errors.New("invalid role")
	//}
	return
}

// AfterCreate 插入数据后的一些操作
func (u *TabUser) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate Hook 被执行 ", u.Remark)

	//if u.Role == "admin" {
	//	return errors.New("invalid role")
	//}
	return
}

/*
更新时可用的 hook

	// 开始事务
	BeforeSave
	BeforeUpdate
	// 关联前的 save
	// 更新 db
	// 关联后的 save
	AfterUpdate
	AfterSave
	// 提交或回滚事务

*/

/*
删除时可用的 hook

	// 开始事务
	BeforeDelete
	// 删除 db 中的数据
	AfterDelete
	// 提交或回滚事务

*/

/*
查询时可用的 hook

	// 从 db 中加载数据
	// Preloading (eager loading)
	AfterFind
*/
