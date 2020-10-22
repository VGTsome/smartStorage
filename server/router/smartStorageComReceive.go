package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageComReceiveRouter(Router *gin.RouterGroup) {
	SmartStorageComReceiveRouter := Router.Group("smartStorageComReceive").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageComReceiveRouter.POST("createSmartStorageComReceive", v1.CreateSmartStorageComReceive)   // 新建SmartStorageComReceive
		SmartStorageComReceiveRouter.DELETE("deleteSmartStorageComReceive", v1.DeleteSmartStorageComReceive) // 删除SmartStorageComReceive
		SmartStorageComReceiveRouter.DELETE("deleteSmartStorageComReceiveByIds", v1.DeleteSmartStorageComReceiveByIds) // 批量删除SmartStorageComReceive
		SmartStorageComReceiveRouter.PUT("updateSmartStorageComReceive", v1.UpdateSmartStorageComReceive)    // 更新SmartStorageComReceive
		SmartStorageComReceiveRouter.GET("findSmartStorageComReceive", v1.FindSmartStorageComReceive)        // 根据ID获取SmartStorageComReceive
		SmartStorageComReceiveRouter.GET("getSmartStorageComReceiveList", v1.GetSmartStorageComReceiveList)  // 获取SmartStorageComReceive列表
	}
}
