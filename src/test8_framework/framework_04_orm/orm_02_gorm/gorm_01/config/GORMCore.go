package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 定义公开的连接信息
var DB *gorm.DB
var err error

func init() {
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "用户名:密码@tcp(IP:Port)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印所有日志
		Logger: logger.Default.LogMode(logger.Info),
		// 显示出查询的所有字段
		QueryFields: true,
		// 禁用默认事务，true 表示不开起事务；不写或写false表示开启事务
		//SkipDefaultTransaction: true,
	})

	/*
		// Config GORM 配置
		type Config struct {
			// SkipDefaultTransaction：GORM 默认每个单独的增删改操作都会开启事务以保证数据一致性。
			// 如果你不想每次都开启事务，可以将此选项设置为 true。
			SkipDefaultTransaction bool // 是否跳过默认事务处理

			// NamingStrategy：定义表名、列名等命名策略。
			NamingStrategy schema.Namer // 命名策略（表名、列名等）

			// FullSaveAssociations：保存时是否级联保存所有关联数据。
			FullSaveAssociations bool // 是否保存所有关联关系

			// Logger：GORM 使用的日志接口，可以配置日志级别、格式等。
			Logger logger.Interface // 日志记录器

			// NowFunc：创建时间戳时使用的函数，默认是 time.Now。
			NowFunc func() time.Time // 获取当前时间的函数

			// DryRun：开启后不会执行 SQL，只会生成 SQL 语句（调试用）。
			DryRun bool // 仅生成 SQL，不执行（干跑模式）

			// PrepareStmt：是否启用预编译语句缓存，提高性能。
			PrepareStmt bool // 是否启用预编译语句缓存

			// DisableAutomaticPing：是否禁用自动 PING 检查数据库连接状态。
			DisableAutomaticPing bool // 是否禁用自动 ping 数据库

			// DisableForeignKeyConstraintWhenMigrating：迁移表时是否禁用外键约束。
			DisableForeignKeyConstraintWhenMigrating bool // 迁移时禁用外键约束

			// IgnoreRelationshipsWhenMigrating：迁移表时是否忽略关联关系。
			IgnoreRelationshipsWhenMigrating bool // 迁移时忽略关系映射

			// DisableNestedTransaction：是否禁用嵌套事务。
			DisableNestedTransaction bool // 禁用嵌套事务支持

			// AllowGlobalUpdate：是否允许不带 WHERE 条件的全表更新或删除操作。
			AllowGlobalUpdate bool // 允许全表更新或删除

			// QueryFields：查询时是否使用 SELECT * 指定所有字段。
			QueryFields bool // 查询时是否显式列出所有字段

			// CreateBatchSize：批量插入时的默认批次大小。
			CreateBatchSize int // 批量插入的默认大小

			// TranslateError：是否启用错误翻译（例如外键冲突等错误信息更友好）。
			TranslateError bool // 启用错误翻译

			// PropagateUnscoped：是否将 Unscoped（不包含软删除过滤）传播到所有子语句中。
			PropagateUnscoped bool // 是否传播 Unscoped 到子语句

			// ClauseBuilders：用于自定义 SQL 子句构建器。
			ClauseBuilders map[string]clause.ClauseBuilder // SQL 子句构建器

			// ConnPool：数据库连接池接口。
			ConnPool ConnPool // 数据库连接池

			// Dialector：数据库方言（MySQL、Postgres、SQLite 等）。
			Dialector // 数据库方言（用于连接不同数据库）

			// Plugins：注册的插件集合。
			Plugins map[string]Plugin // 注册的插件
		}

	*/

	if err != nil {
		panic("数据库连接失败！！！")
	}

	// &{0xc00019e6c0 <nil> 0 0xc000134000 1}
}
