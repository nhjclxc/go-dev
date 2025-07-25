package router

import (
    "gin_generator/api"
    "github.com/gin-gonic/gin"
)

// GenTableRouter 代码生成业务 路由Router层
type GenTableRouter struct {
    genTableApi api.GenTableApi
}

// InitGenTableRouter 初始化 GenTableRouter 路由
func (this *GenTableRouter) InitGenTableRouter(privateRouterOrigin *gin.RouterGroup, publicRouterOrigin *gin.RouterGroup) {
    privateRouter := privateRouterOrigin.Group("/gen/table")
    {
        // PrivateRouter 下是一些必须进行登录的接口
        // http://localhost:8080/private

        // http://127.0.0.1:8080/private/gen/table/6


        privateRouter.POST("", this.genTableApi.InsertGenTable)         // 新增代码生成业务
        privateRouter.PUT("", this.genTableApi.UpdateGenTable)          // 修改代码生成业务
        privateRouter.DELETE("/:tableIdList", this.genTableApi.DeleteGenTable)       // 删除代码生成业务
        privateRouter.GET("/:tableId", this.genTableApi.GetGenTableByTableId)  // 获取代码生成业务详细信息
        privateRouter.GET("/list", this.genTableApi.GetGenTableList)     // 查询代码生成业务列表
        privateRouter.GET("/pageList", this.genTableApi.GetGenTablePageList) // 分页查询代码生成业务列表
        privateRouter.GET("/export", this.genTableApi.ExportGenTable)       // 导出代码生成业务列表
    }

    //publicRouter := publicRouterOrigin.Group("/gen/table")
    //{
    //    // PublicRouter 下是一些无需登录的接口，可以直接访问，无须经过授权操作
    //    // http://localhost:8080/public
    //
    //    publicRouter.POST("", this.genTableApi.InsertGenTable)         // 新增代码生成业务
    //    publicRouter.PUT("", this.genTableApi.UpdateGenTable)          // 修改代码生成业务
    //    publicRouter.DELETE("/:tableId", this.genTableApi.DeleteGenTable)       // 删除代码生成业务
    //    publicRouter.GET("/:tableId", this.genTableApi.GetGenTableByTableId)  // 获取代码生成业务详细信息
    //    publicRouter.GET("/list", this.genTableApi.GetGenTableList)     // 查询代码生成业务列表
    //    publicRouter.GET("/pageList", this.genTableApi.GetGenTablePageList) // 分页查询代码生成业务列表
    //    publicRouter.GET("/export", this.genTableApi.ExportGenTable)       // 导出代码生成业务列表
    //}
}
