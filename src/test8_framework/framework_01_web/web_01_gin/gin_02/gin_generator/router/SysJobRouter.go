package router

import (
    "gin_generator/api"
    "github.com/gin-gonic/gin"
)

// SysJobRouter 定时任务调度 路由Router层
type SysJobRouter struct {
    sysJobApi api.SysJobApi
}

// InitSysJobRouter 初始化 SysJobRouter 路由
func (this *SysJobRouter) InitSysJobRouter(privateRouterOrigin *gin.RouterGroup, publicRouterOrigin *gin.RouterGroup) {
    privateRouter := privateRouterOrigin.Group("/sys/job")
    {
        // PrivateRouter 下是一些必须进行登录的接口
        // http://localhost:8080/private

        privateRouter.POST("", this.sysJobApi.InsertSysJob)         // 新增定时任务调度
        privateRouter.PUT("", this.sysJobApi.UpdateSysJob)          // 修改定时任务调度
        privateRouter.DELETE("/:jobIdList", this.sysJobApi.DeleteSysJob)       // 删除定时任务调度
        privateRouter.GET("/:jobId", this.sysJobApi.GetSysJobByJobId)  // 获取定时任务调度详细信息
        privateRouter.GET("/list", this.sysJobApi.GetSysJobList)     // 查询定时任务调度列表
        privateRouter.GET("/pageList", this.sysJobApi.GetSysJobPageList) // 分页查询定时任务调度列表
        privateRouter.GET("/export", this.sysJobApi.ExportSysJob)       // 导出定时任务调度列表
    }

    publicRouter := publicRouterOrigin.Group("/sys/job")
    {
        // PublicRouter 下是一些无需登录的接口，可以直接访问，无须经过授权操作
        // http://localhost:8080/public

        publicRouter.POST("", this.sysJobApi.InsertSysJob)         // 新增定时任务调度
        publicRouter.PUT("", this.sysJobApi.UpdateSysJob)          // 修改定时任务调度
        publicRouter.DELETE("/:jobId", this.sysJobApi.DeleteSysJob)       // 删除定时任务调度
        publicRouter.GET("/:jobId", this.sysJobApi.GetSysJobByJobId)  // 获取定时任务调度详细信息
        publicRouter.GET("/list", this.sysJobApi.GetSysJobList)     // 查询定时任务调度列表
        publicRouter.GET("/pageList", this.sysJobApi.GetSysJobPageList) // 分页查询定时任务调度列表
        publicRouter.GET("/export", this.sysJobApi.ExportSysJob)       // 导出定时任务调度列表
    }
}
