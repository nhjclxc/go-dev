package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go_zero_19_db_mysql/internal/svc"
	"go_zero_19_db_mysql/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_19_db_mysql/internal/model"
)

type MysqlInsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewMysqlInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MysqlInsertLogic {
	return &MysqlInsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MysqlInsertLogic) MysqlInsert(req *types.UserInsertReq) (resp *types.MysqlApiResp, err error) {
	// todo: add your logic here and delete this line

	fmt.Printf("mysqlInsertLogic.MysqlInsert = %#v \n\n", req)

	userModel := model.NewTabUserModel(l.svcCtx.SqlConn)

	u := model.TabUser{
		Name:         req.Name,
		Email:        sql.NullString{req.Email, true},
		Age:          18,
		Birthday:     sql.NullTime{time.Now(), true},
		MemberNumber: sql.NullString{req.MemberNumber, true},
		Remark:       sql.NullString{String: req.Remark, Valid: true},
		ActivatedAt:  sql.NullTime{time.Now(), true},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	insert, err := userModel.Insert(l.ctx, &u)
	if err != nil {
		fmt.Printf("插入失败！！！", err)
		return
	}

	id, err := insert.LastInsertId()
	if err != nil {
		return
	}

	affected, err := insert.RowsAffected()
	if err != nil {
		return
	}

	fmt.Println("id = ", id)
	fmt.Println("affected = ", affected)



	return
}
