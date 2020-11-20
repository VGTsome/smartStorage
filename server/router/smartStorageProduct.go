package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSmartStorageProductRouter(Router *gin.RouterGroup) {
	SmartStorageProductRouter := Router.Group("smartStorageProduct").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SmartStorageProductRouter.POST("createSmartStorageProduct", v1.CreateSmartStorageProduct)             // 新建SmartStorageProduct
		SmartStorageProductRouter.DELETE("deleteSmartStorageProduct", v1.DeleteSmartStorageProduct)           // 删除SmartStorageProduct
		SmartStorageProductRouter.DELETE("deleteSmartStorageProductByIds", v1.DeleteSmartStorageProductByIds) // 批量删除SmartStorageProduct
		SmartStorageProductRouter.PUT("updateSmartStorageProduct", v1.UpdateSmartStorageProduct)              // 更新SmartStorageProduct
		SmartStorageProductRouter.GET("findSmartStorageProduct", v1.FindSmartStorageProduct)                  // 根据ID获取SmartStorageProduct
		SmartStorageProductRouter.GET("getSmartStorageProductList", v1.GetSmartStorageProductList)            // 获取SmartStorageProduct列表
		SmartStorageProductRouter.GET("getSmartStorageProductValidList", v1.GetSmartStorageProductValidList)  // 获取可用SmartStorageProduct列表
	}
}
