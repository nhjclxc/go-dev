package model_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero_19_db_mysql/internal/model"
	"testing"
	"time"
)


func main() {


	// 连接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 需要自行将 dsn 中的 host，账号 密码配置正确
	dsn := "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	//conn := sqlx.NewSqlConn("mysql", dsn)
	_ = conn

	var ctx context.Context = context.Background()

	// 开始crud

	//crud_01_insert(ctx, conn)

	//crud_02_update(ctx, conn)

	//crud_03_findOne(ctx, conn)

	//crud_03_find(ctx, conn)

	//crud_05_delete(ctx, conn)

	//crud_06_updateBySelective(ctx, conn)

	crud_07_pageList(ctx, conn)


}


func Test_crud_01_insert(t *testing.T)  {

}

func crud_07_pageList(ctx context.Context, conn sqlx.SqlConn) {
	userModel := model.NewTabUserModel(conn)

	pageReq := model.PageRequest{
		Page:     5,
		PageSize: 10,
	}

	list, total, err := userModel.PageList(ctx, pageReq)
	if err != nil {
		logx.Errorf("分页查询失败: %v", err)
	}
	fmt.Printf("共 %d 条数据，当前页 %d 条数据\n", total, len(list))

}

func crud_06_updateBySelective(ctx context.Context, conn sqlx.SqlConn) {
	userModel := model.NewTabUserModel(conn)

	u := model.TabUser{
		UserId: 26,
		Name:         "db_01_mysql",
		Email:        sql.NullString{"nhjclx111c@163.com", true},
		Age:          18,
	}

	err := userModel.UpdateBySelective(ctx, &u)
	if err != nil {
		fmt.Println("出现错误：", err)
		return
	}



}

func crud_05_delete(ctx context.Context, conn sqlx.SqlConn) {

	userModel := model.NewTabUserModel(conn)

	err := userModel.Delete(ctx, 27)
	if err != nil {
		return
	}


}

func crud_03_find(ctx context.Context, conn sqlx.SqlConn) {

	// 这个userModel类似于mapper
	userModel := model.NewTabUserModel(conn)

	u := model.TabUser{
		//UserId: 27,
		Name: "mike",
	}

	findList, err := userModel.Find(ctx, &u)
	if err != nil {
		fmt.Println("查询出错：", err)
		return
	}
	fmt.Println("\nquerySize = ", len(findList))
	for i, user := range findList {
		fmt.Printf("i = %d, userId = %d, name = %s \n", i, user.UserId, user.Name)
	}





}

func crud_03_findOne(ctx context.Context, conn sqlx.SqlConn) {
	userModel := model.NewTabUserModel(conn)

	one, err := userModel.FindOne(ctx, 27)
	if err != nil {
		return
	}

	fmt.Printf("%v \n", one)

	fmt.Println("one.Birthday", one.Birthday)
	fmt.Println("one.Birthday.Time", one.Birthday.Time)
	fmt.Println("one.Birthday.Valid", one.Birthday.Valid)


}

func crud_02_update(ctx context.Context, conn sqlx.SqlConn) {

	userModel := model.NewTabUserModel(conn)

	u := model.TabUser{
		UserId:       27,
		Name:         "db_01_mysql987654",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := userModel.Update(ctx, &u)
	if err != nil {
		return
	}


}

func crud_01_insert(ctx context.Context, conn sqlx.SqlConn) {

	userModel := model.NewTabUserModel(conn)

	u := model.TabUser{
		Name:         "db_01_mysql",
		Email:        sql.NullString{"nhjclxc@163.com", true},
		Age:          18,
		Birthday:     sql.NullTime{time.Now(), true},
		MemberNumber: sql.NullString{"110", true},
		Remark:       sql.NullString{"你好，世界", true},
		ActivatedAt:  sql.NullTime{time.Now(), true},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	insert, err := userModel.Insert(ctx, &u)
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


}
