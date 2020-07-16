package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageDoorRouter(Router *gin.RouterGroup) {
	SmartStorageDoorRouter := Router.Group("smartStorageDoor").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageDoorRouter.POST("createSmartStorageDoor", v1.CreateSmartStorageDoor)   // 新建SmartStorageDoor
		SmartStorageDoorRouter.DELETE("deleteSmartStorageDoor", v1.DeleteSmartStorageDoor) // 删除SmartStorageDoor
		SmartStorageDoorRouter.DELETE("deleteSmartStorageDoorByIds", v1.DeleteSmartStorageDoorByIds) // 批量删除SmartStorageDoor
		SmartStorageDoorRouter.PUT("updateSmartStorageDoor", v1.UpdateSmartStorageDoor)    // 更新SmartStorageDoor
		SmartStorageDoorRouter.GET("findSmartStorageDoor", v1.FindSmartStorageDoor)        // 根据ID获取SmartStorageDoor
		SmartStorageDoorRouter.GET("getSmartStorageDoorList", v1.GetSmartStorageDoorList)  // 获取SmartStorageDoor列表
	}
}
