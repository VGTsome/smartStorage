package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageOrderRouter(Router *gin.RouterGroup) {
	SmartStorageOrderRouter := Router.Group("smartStorageOrder").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageOrderRouter.POST("createSmartStorageOrder", v1.CreateSmartStorageOrder)   // 新建SmartStorageOrder
		SmartStorageOrderRouter.DELETE("deleteSmartStorageOrder", v1.DeleteSmartStorageOrder) // 删除SmartStorageOrder
		SmartStorageOrderRouter.DELETE("deleteSmartStorageOrderByIds", v1.DeleteSmartStorageOrderByIds) // 批量删除SmartStorageOrder
		SmartStorageOrderRouter.PUT("updateSmartStorageOrder", v1.UpdateSmartStorageOrder)    // 更新SmartStorageOrder
		SmartStorageOrderRouter.GET("findSmartStorageOrder", v1.FindSmartStorageOrder)        // 根据ID获取SmartStorageOrder
		SmartStorageOrderRouter.GET("getSmartStorageOrderList", v1.GetSmartStorageOrderList)  // 获取SmartStorageOrder列表
	}
}
