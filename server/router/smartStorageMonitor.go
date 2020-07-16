package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageMonitorRouter(Router *gin.RouterGroup) {
	SmartStorageMonitorRouter := Router.Group("smartStorageMonitor").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageMonitorRouter.POST("createSmartStorageMonitor", v1.CreateSmartStorageMonitor)   // 新建SmartStorageMonitor
		SmartStorageMonitorRouter.DELETE("deleteSmartStorageMonitor", v1.DeleteSmartStorageMonitor) // 删除SmartStorageMonitor
		SmartStorageMonitorRouter.DELETE("deleteSmartStorageMonitorByIds", v1.DeleteSmartStorageMonitorByIds) // 批量删除SmartStorageMonitor
		SmartStorageMonitorRouter.PUT("updateSmartStorageMonitor", v1.UpdateSmartStorageMonitor)    // 更新SmartStorageMonitor
		SmartStorageMonitorRouter.GET("findSmartStorageMonitor", v1.FindSmartStorageMonitor)        // 根据ID获取SmartStorageMonitor
		SmartStorageMonitorRouter.GET("getSmartStorageMonitorList", v1.GetSmartStorageMonitorList)  // 获取SmartStorageMonitor列表
	}
}
