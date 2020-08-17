package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmartStorageQuotaRouter(Router *gin.RouterGroup) {
	SmartStorageQuotaRouter := Router.Group("ssq").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageQuotaRouter.POST("createSmartStorageQuota", v1.CreateSmartStorageQuota)   // 新建SmartStorageQuota
		SmartStorageQuotaRouter.DELETE("deleteSmartStorageQuota", v1.DeleteSmartStorageQuota) // 删除SmartStorageQuota
		SmartStorageQuotaRouter.DELETE("deleteSmartStorageQuotaByIds", v1.DeleteSmartStorageQuotaByIds) // 批量删除SmartStorageQuota
		SmartStorageQuotaRouter.PUT("updateSmartStorageQuota", v1.UpdateSmartStorageQuota)    // 更新SmartStorageQuota
		SmartStorageQuotaRouter.GET("findSmartStorageQuota", v1.FindSmartStorageQuota)        // 根据ID获取SmartStorageQuota
		SmartStorageQuotaRouter.GET("getSmartStorageQuotaList", v1.GetSmartStorageQuotaList)  // 获取SmartStorageQuota列表
	}
}
