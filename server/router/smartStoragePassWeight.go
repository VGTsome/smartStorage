package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStoragePassWeightRouter(Router *gin.RouterGroup) {
	SmartStoragePassWeightRouter := Router.Group("smartStoragePassWeight").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStoragePassWeightRouter.POST("createSmartStoragePassWeight", v1.CreateSmartStoragePassWeight)   // 新建SmartStoragePassWeight
		SmartStoragePassWeightRouter.DELETE("deleteSmartStoragePassWeight", v1.DeleteSmartStoragePassWeight) // 删除SmartStoragePassWeight
		SmartStoragePassWeightRouter.DELETE("deleteSmartStoragePassWeightByIds", v1.DeleteSmartStoragePassWeightByIds) // 批量删除SmartStoragePassWeight
		SmartStoragePassWeightRouter.PUT("updateSmartStoragePassWeight", v1.UpdateSmartStoragePassWeight)    // 更新SmartStoragePassWeight
		SmartStoragePassWeightRouter.GET("findSmartStoragePassWeight", v1.FindSmartStoragePassWeight)        // 根据ID获取SmartStoragePassWeight
		SmartStoragePassWeightRouter.GET("getSmartStoragePassWeightList", v1.GetSmartStoragePassWeightList)  // 获取SmartStoragePassWeight列表
	}
}
