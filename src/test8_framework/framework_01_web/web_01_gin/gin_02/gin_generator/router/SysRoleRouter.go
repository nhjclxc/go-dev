package router

import (
    "gin_generator/api"
    "github.com/gin-gonic/gin"
)

// SysRoleRouter 角色信息 路由Router层
type SysRoleRouter struct {
    sysRoleApi api.SysRoleApi
}

// InitSysRoleRouter 初始化 SysRoleRouter 路由
func (this *SysRoleRouter) InitSysRoleRouter(privateRouterOrigin *gin.RouterGroup, publicRouterOrigin *gin.RouterGroup) {
    privateRouter := privateRouterOrigin.Group("/sys/role")
    {
        // PrivateRouter 下是一些必须进行登录的接口
        // http://localhost:8080/private

        privateRouter.POST("", this.sysRoleApi.InsertSysRole)         // 新增角色信息
        privateRouter.PUT("", this.sysRoleApi.UpdateSysRole)          // 修改角色信息
        privateRouter.DELETE("/:roleIdList", this.sysRoleApi.DeleteSysRole)       // 删除角色信息
        privateRouter.GET("/:roleId", this.sysRoleApi.GetSysRoleByRoleId)  // 获取角色信息详细信息
        privateRouter.GET("/list", this.sysRoleApi.GetSysRoleList)     // 查询角色信息列表
        privateRouter.GET("/pageList", this.sysRoleApi.GetSysRolePageList) // 分页查询角色信息列表
        privateRouter.GET("/export", this.sysRoleApi.ExportSysRole)       // 导出角色信息列表
    }

    //publicRouter := publicRouterOrigin.Group("/sys/role")
    //{
    //    // PublicRouter 下是一些无需登录的接口，可以直接访问，无须经过授权操作
    //    // http://localhost:8080/public
    //
    //    publicRouter.POST("", this.sysRoleApi.InsertSysRole)         // 新增角色信息
    //    publicRouter.PUT("", this.sysRoleApi.UpdateSysRole)          // 修改角色信息
    //    publicRouter.DELETE("/:roleId", this.sysRoleApi.DeleteSysRole)       // 删除角色信息
    //    publicRouter.GET("/:roleId", this.sysRoleApi.GetSysRoleByRoleId)  // 获取角色信息详细信息
    //    publicRouter.GET("/list", this.sysRoleApi.GetSysRoleList)     // 查询角色信息列表
    //    publicRouter.GET("/pageList", this.sysRoleApi.GetSysRolePageList) // 分页查询角色信息列表
    //    publicRouter.GET("/export", this.sysRoleApi.ExportSysRole)       // 导出角色信息列表
    //}
}
