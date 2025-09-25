package middleware

import (
	"fmt"
	"gin_casbin/config"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"testing"
)

func TestName(t *testing.T) {

	// 初始化 DB
	config.InitDB()

	// GORM Adapter for MySQL
	a, err := gormadapter.NewAdapterByDB(config.DB)
	if err != nil {
		fmt.Println("NewAdapterByDB", err)
		return
	}

	e, err := casbin.NewEnforcer("../model.conf", a)
	if err != nil {
		fmt.Println("NewEnforcer", err)
		return
	}
	if err := e.LoadPolicy(); err != nil {
		fmt.Println("LoadPolicy failed:", err)
		return
	}
	fmt.Println("LoadPolicy success")

	policy, err := e.GetPolicy()
	if err != nil {
		return
	}
	fmt.Println("Policies:", policy)
	groupingPolicy, err := e.GetGroupingPolicy()
	if err != nil {
		return
	}
	fmt.Println("Groupings:", groupingPolicy)

	ok, err := e.Enforce("common", "/home", "GET")
	if err != nil {
		fmt.Println("Enforce error:", err)
	}
	fmt.Println("common can access /home GET ?", ok)

	ok, err = e.Enforce("common", "/api/home", "GET")
	if err != nil {
		fmt.Println("Enforce error:", err)
	}
	fmt.Println("common can access /api/home GET ?", ok)

	ok, err = e.Enforce("admin", "/api/home", "GET")
	if err != nil {
		fmt.Println("Enforce error:", err)
	}
	fmt.Println("admin can access /api/home GET ?", ok)

	ok, err = e.Enforce("superadmin", "/api/home", "GET")
	if err != nil {
		fmt.Println("Enforce error:", err)
	}
	fmt.Println("superadmin can access /api/home GET ?", ok)

}
