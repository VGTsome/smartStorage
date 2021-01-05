package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageComLogRouter(Router *gin.RouterGroup) {
	SmartStorageComLogRouter := Router.Group("smartStorageComLog").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageComLogRouter.POST("createSmartStorageComLog", v1.CreateSmartStorageComLog)   // 新建SmartStorageComLog
		SmartStorageComLogRouter.DELETE("deleteSmartStorageComLog", v1.DeleteSmartStorageComLog) // 删除SmartStorageComLog
		SmartStorageComLogRouter.DELETE("deleteSmartStorageComLogByIds", v1.DeleteSmartStorageComLogByIds) // 批量删除SmartStorageComLog
		SmartStorageComLogRouter.PUT("updateSmartStorageComLog", v1.UpdateSmartStorageComLog)    // 更新SmartStorageComLog
		SmartStorageComLogRouter.GET("findSmartStorageComLog", v1.FindSmartStorageComLog)        // 根据ID获取SmartStorageComLog
		SmartStorageComLogRouter.GET("getSmartStorageComLogList", v1.GetSmartStorageComLogList)  // 获取SmartStorageComLog列表
	}
}
