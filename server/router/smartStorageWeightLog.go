package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageWeightLogRouter(Router *gin.RouterGroup) {
	SmartStorageWeightLogRouter := Router.Group("smartStorageWeightLog").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageWeightLogRouter.POST("createSmartStorageWeightLog", v1.CreateSmartStorageWeightLog)   // 新建SmartStorageWeightLog
		SmartStorageWeightLogRouter.DELETE("deleteSmartStorageWeightLog", v1.DeleteSmartStorageWeightLog) // 删除SmartStorageWeightLog
		SmartStorageWeightLogRouter.DELETE("deleteSmartStorageWeightLogByIds", v1.DeleteSmartStorageWeightLogByIds) // 批量删除SmartStorageWeightLog
		SmartStorageWeightLogRouter.PUT("updateSmartStorageWeightLog", v1.UpdateSmartStorageWeightLog)    // 更新SmartStorageWeightLog
		SmartStorageWeightLogRouter.GET("findSmartStorageWeightLog", v1.FindSmartStorageWeightLog)        // 根据ID获取SmartStorageWeightLog
		SmartStorageWeightLogRouter.GET("getSmartStorageWeightLogList", v1.GetSmartStorageWeightLogList)  // 获取SmartStorageWeightLog列表
	}
}
