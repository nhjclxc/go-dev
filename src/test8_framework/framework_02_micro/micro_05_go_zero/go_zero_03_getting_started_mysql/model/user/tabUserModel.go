package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TabUserModel = (*customTabUserModel)(nil)

type (
	// TabUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTabUserModel.
	TabUserModel interface {
		tabUserModel
		withSession(session sqlx.Session) TabUserModel
	}

	customTabUserModel struct {
		*defaultTabUserModel
	}
)

// NewTabUserModel returns a model for the database table.
func NewTabUserModel(conn sqlx.SqlConn) TabUserModel {
	return &customTabUserModel{
		defaultTabUserModel: newTabUserModel(conn),
	}
}

func (m *customTabUserModel) withSession(session sqlx.Session) TabUserModel {
	return NewTabUserModel(sqlx.NewSqlConnFromSession(session))
}
