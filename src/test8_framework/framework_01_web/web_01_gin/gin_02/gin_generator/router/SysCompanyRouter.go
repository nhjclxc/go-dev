package router

import (
    "gin_generator/api"
    "github.com/gin-gonic/gin"
)

// SysCompanyRouter 公司 路由Router层
type SysCompanyRouter struct {
    sysCompanyApi api.SysCompanyApi
}

// InitSysCompanyRouter 初始化 SysCompanyRouter 路由
func (this *SysCompanyRouter) InitSysCompanyRouter(privateRouterOrigin *gin.RouterGroup, publicRouterOrigin *gin.RouterGroup) {
    privateRouter := privateRouterOrigin.Group("/sys/company")
    {
        // PrivateRouter 下是一些必须进行登录的接口
        // http://localhost:8080/private

        privateRouter.POST("", this.sysCompanyApi.InsertSysCompany)         // 新增公司
        privateRouter.PUT("", this.sysCompanyApi.UpdateSysCompany)          // 修改公司
        privateRouter.DELETE("/:idList", this.sysCompanyApi.DeleteSysCompany)       // 删除公司
        privateRouter.GET("/:id", this.sysCompanyApi.GetSysCompanyById)  // 获取公司详细信息
        privateRouter.GET("/list", this.sysCompanyApi.GetSysCompanyList)     // 查询公司列表
        privateRouter.GET("/pageList", this.sysCompanyApi.GetSysCompanyPageList) // 分页查询公司列表
        privateRouter.GET("/export", this.sysCompanyApi.ExportSysCompany)       // 导出公司列表
    }

    //publicRouter := publicRouterOrigin.Group("/sys/company")
    //{
    //    // PublicRouter 下是一些无需登录的接口，可以直接访问，无须经过授权操作
    //    // http://localhost:8080/public
    //
    //    publicRouter.POST("", this.sysCompanyApi.InsertSysCompany)         // 新增公司
    //    publicRouter.PUT("", this.sysCompanyApi.UpdateSysCompany)          // 修改公司
    //    publicRouter.DELETE("/:id", this.sysCompanyApi.DeleteSysCompany)       // 删除公司
    //    publicRouter.GET("/:id", this.sysCompanyApi.GetSysCompanyById)  // 获取公司详细信息
    //    publicRouter.GET("/list", this.sysCompanyApi.GetSysCompanyList)     // 查询公司列表
    //    publicRouter.GET("/pageList", this.sysCompanyApi.GetSysCompanyPageList) // 分页查询公司列表
    //    publicRouter.GET("/export", this.sysCompanyApi.ExportSysCompany)       // 导出公司列表
    //}
}
