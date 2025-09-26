package main

import (
	"fmt"
	"gin_casbin/config"
	"gin_casbin/handler"
	"gin_casbin/middleware"
	"gin_casbin/model"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// 初始化 DB
	config.InitDB()

	// GORM Adapter for MySQL
	// // 1. 用 GORM 的数据库连接创建 Adapter
	adapter, err := gormadapter.NewAdapterByDB(config.DB)
	if err != nil {
		fmt.Println("NewAdapterByDB", err)
		return
	}

	// 2. 创建 Enforcer，传入模型文件 + Adapter
	enforcer, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		fmt.Println("NewEnforcer", err)
		return
	}

	// 3. 从数据库加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		fmt.Println("LoadPolicy failed:", err)
		return
	}
	fmt.Println("LoadPolicy success")

	r := gin.Default()

	// 登录接口
	r.POST("/login", handler.Login)

	// 受保护的接口
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware(), middleware.CasbinMiddleware(enforcer))
	{
		// common
		auth.GET("/home", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /home \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})
		auth.GET("/home2", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /home2 \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})

		// admin
		auth.GET("/user", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /user \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})

		// superadmin
		auth.GET("/role", func(ctx *gin.Context) {
			ctx.JSON(200, getStr("/role", ctx))
		})

		// 新接口
		auth.POST("/goods/insert", func(ctx *gin.Context) { ctx.JSON(200, getStr("/goods/insert", ctx)) })
		auth.POST("/goods/update", func(ctx *gin.Context) { ctx.JSON(200, getStr("/goods/update", ctx)) })

		// 以下两个接口 /api/menu/insert、/api/casbin/insert 模拟新api来了和分配给角色的过程
		// INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', 'superadmin', '/api/menu/insert', 'POST');
		// INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', 'superadmin', '/api/casbin/insert', 'POST');

		/*
					http://localhost:8080/api/menu/insert

				{
				    "name": "新增商品",
				    "path": "/api/goods/insert",
				    "type": 2
				}
			{
			    "name": "修改商品",
			    "path": "/api/goods/update",
			    "type": 2
			}
		*/
		auth.POST("/menu/insert", func(ctx *gin.Context) {
			//config.DB
			type MenuInsertRequest struct {
				Name string `json:"name"`
				Path string `json:"path"`
				Type uint   `json:"type"` // 1=菜单,2=按钮/API
			}
			var req MenuInsertRequest
			ctx.ShouldBindJSON(&req)

			var menu model.Menu = model.Menu{
				Name:      req.Name,
				Path:      req.Path,
				Type:      req.Type,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			tx := config.DB.Model(model.Menu{}).Save(&menu)
			ctx.JSON(200, fmt.Sprintf("/menu/inser 影响行数：%d, err = %s", tx.RowsAffected, tx.Error))
		})

		/*
			http://localhost:8080/api/casbin/insert
				{
				    "ptype": "p",
				    "v0": "superadmin",
				    "v1": "/api/goods/insert",
				    "v2": "POST"
				}
		*/
		auth.POST("/casbin/insert", func(ctx *gin.Context) {
			type CasbinRuleReq struct {
				Ptype string `json:"ptype"`
				V0    string `json:"v0"`
				V1    string `json:"v1"`
				V2    string `json:"v2"`
			}
			var req CasbinRuleReq
			ctx.ShouldBindJSON(&req)

			var casbinRule model.CasbinRule = model.CasbinRule{
				Ptype: req.Ptype,
				V0:    req.V0,
				V1:    req.V1,
				V2:    req.V2,
			}

			tx := config.DB.Model(model.CasbinRule{}).Save(&casbinRule)

			// 数据库更新成功之后，刷新策略
			err := enforcer.LoadPolicy()
			if err != nil {
				fmt.Println("LoadPolicy error:", err)
				return
			}
			fmt.Println("Reload casbin policy success")

			ctx.JSON(200, fmt.Sprintf("/casbin/inser 影响行数：%d, err = %s", tx.RowsAffected, tx.Error))
		})

		auth.POST("/casbin/insert2", func(ctx *gin.Context) {
			type CasbinRuleReq struct {
				Ptype string `json:"ptype"`
				V0    string `json:"v0"`
				V1    string `json:"v1"`
				V2    string `json:"v2"`
			}
			var req CasbinRuleReq
			ctx.ShouldBindJSON(&req)

			// 如果你是通过 Casbin API 修改策略（比如 e.AddPolicy()、e.RemovePolicy()），Adapter 会自动写入数据库，并且内存里的策略已经是最新的 不需要刷新。
			//只有 你绕过 Casbin，直接 UPDATE/INSERT SQL 到 casbin_rule 表，才需要 LoadPolicy()。
			// 不直接插入数据库，而是调用策略的方法去更新策略的同时策略会帮我们更新数据库
			// 注意⚠️：AddPolicies添加的是策略p，而不是继承g，添加继承g的方法是（AddGroupingPolicy）
			policies, err := enforcer.AddPolicies([][]string{{req.V0, req.V1, req.V2}})
			//policies, err := enforcer.AddPolicy(req.V0, req.V1, req.V2)
			if err != nil {
				return
			}

			ctx.JSON(200, fmt.Sprintf("/casbin/inser 影响行数：%v, err = %s", policies, err))
		})

		// 自动同步策略：
		/*
			✅ 2. 自动同步（推荐，生产常用）

			Casbin 官方有个库 casbin/casbin-synced
			 和 Watcher 机制
			，可以在多节点部署时自动同步。

			如果你用的是 MySQL，可以用 casbin/redis-watcher
			，流程大致是：

			修改数据库策略 → 发布一个消息到 Redis

			其他节点收到消息 → 调用 LoadPolicy() 自动刷新

			示例：

			w, _ := rediswatcher.NewWatcher("127.0.0.1:6379")
			e.SetWatcher(w)

			w.SetUpdateCallback(func(msg string) {
			    fmt.Println("Watcher received update, reloading policy...")
			    e.LoadPolicy()
			})


			这样你不需要手动调用 LoadPolicy()，只要一个节点改了策略，其他节点会自动刷新。
		*/

	}

	open := r.Group("/openapi")
	{
		open.GET("/hi", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /hi \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})
	}

	r.Run(":8080")
}

func getStr(router string, ctx *gin.Context) string {
	username, _ := ctx.Get("username")
	str := fmt.Sprintf("%s 访问了 %s \n", username, router)
	fmt.Printf(str)
	return str
}
