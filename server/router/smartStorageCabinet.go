package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSmartStorageCabinetRouter(Router *gin.RouterGroup) {
	SmartStorageCabinetRouter := Router.Group("smartStorageCabinet").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageCabinetRouter.POST("createSmartStorageCabinet", v1.CreateSmartStorageCabinet)             // 新建SmartStorageCabinet
		SmartStorageCabinetRouter.DELETE("deleteSmartStorageCabinet", v1.DeleteSmartStorageCabinet)           // 删除SmartStorageCabinet
		SmartStorageCabinetRouter.DELETE("deleteSmartStorageCabinetByIds", v1.DeleteSmartStorageCabinetByIds) // 批量删除SmartStorageCabinet
		SmartStorageCabinetRouter.PUT("updateSmartStorageCabinet", v1.UpdateSmartStorageCabinet)              // 更新SmartStorageCabinet
		SmartStorageCabinetRouter.GET("findSmartStorageCabinet", v1.FindSmartStorageCabinet)                  // 根据ID获取SmartStorageCabinet
		SmartStorageCabinetRouter.GET("getSmartStorageCabinetList", v1.GetSmartStorageCabinetList)            // 获取SmartStorageCabinet列表
	}
}
