package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitHkScanRouter(Router *gin.RouterGroup) {
	HkScanRouter := Router.Group("hkScan").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HkScanRouter.POST("createHkScan", v1.CreateHkScan)             // 新建HkScan
		HkScanRouter.DELETE("deleteHkScan", v1.DeleteHkScan)           // 删除HkScan
		HkScanRouter.DELETE("deleteHkScanByIds", v1.DeleteHkScanByIds) // 批量删除HkScan
		HkScanRouter.PUT("updateHkScan", v1.UpdateHkScan)              // 更新HkScan
		HkScanRouter.GET("findHkScan", v1.FindHkScan)                  // 根据ID获取HkScan
		HkScanRouter.GET("getHkScanList", v1.GetHkScanList)            // 获取HkScan列表
	}
}
