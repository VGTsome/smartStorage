package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitSmartStorageCurrentOrderRouter(Router *gin.RouterGroup) {
	SmartStorageCurrentOrderRouter := Router.Group("smartStorageCurrentOrder")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageCurrentOrderRouter.POST("createSmartStorageCurrentOrder", v1.CreateSmartStorageCurrentOrder)             // 新建SmartStorageCurrentOrder
		SmartStorageCurrentOrderRouter.DELETE("deleteSmartStorageCurrentOrder", v1.DeleteSmartStorageCurrentOrder)           // 删除SmartStorageCurrentOrder
		SmartStorageCurrentOrderRouter.DELETE("deleteSmartStorageCurrentOrderByIds", v1.DeleteSmartStorageCurrentOrderByIds) // 批量删除SmartStorageCurrentOrder
		SmartStorageCurrentOrderRouter.PUT("updateSmartStorageCurrentOrder", v1.UpdateSmartStorageCurrentOrder)              // 更新SmartStorageCurrentOrder
		SmartStorageCurrentOrderRouter.GET("findSmartStorageCurrentOrder", v1.FindSmartStorageCurrentOrder)                  // 根据ID获取SmartStorageCurrentOrder
		SmartStorageCurrentOrderRouter.GET("getSmartStorageCurrentOrderList", v1.GetSmartStorageCurrentOrderList)            // 获取SmartStorageCurrentOrder列表
	}
}
