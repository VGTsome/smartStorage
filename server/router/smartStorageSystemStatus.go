package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitSmartStorageSystemStatusRouter(Router *gin.RouterGroup) {
	SmartStorageSystemStatusRouter := Router.Group("ssss")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		// SmartStorageSystemStatusRouter.POST("createSmartStorageSystemStatus", v1.CreateSmartStorageSystemStatus)             // 新建SmartStorageSystemStatus
		// SmartStorageSystemStatusRouter.DELETE("deleteSmartStorageSystemStatus", v1.DeleteSmartStorageSystemStatus)           // 删除SmartStorageSystemStatus
		// SmartStorageSystemStatusRouter.DELETE("deleteSmartStorageSystemStatusByIds", v1.DeleteSmartStorageSystemStatusByIds) // 批量删除SmartStorageSystemStatus
		// SmartStorageSystemStatusRouter.PUT("updateSmartStorageSystemStatus", v1.UpdateSmartStorageSystemStatus)              // 更新SmartStorageSystemStatus
		SmartStorageSystemStatusRouter.GET("findSmartStorageSystemStatus", v1.FindSmartStorageSystemStatus)       // 根据ID获取SmartStorageSystemStatus
		SmartStorageSystemStatusRouter.GET("getSmartStorageSystemStatusList", v1.GetSmartStorageSystemStatusList) // 获取SmartStorageSystemStatus列表
	}
}
