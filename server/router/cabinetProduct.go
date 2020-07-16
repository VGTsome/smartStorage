package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCabinetProductRouter(Router *gin.RouterGroup) {
	CabinetProductRouter := Router.Group("sscp").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CabinetProductRouter.POST("createCabinetProduct", v1.CreateCabinetProduct)   // 新建CabinetProduct
		CabinetProductRouter.DELETE("deleteCabinetProduct", v1.DeleteCabinetProduct) // 删除CabinetProduct
		CabinetProductRouter.DELETE("deleteCabinetProductByIds", v1.DeleteCabinetProductByIds) // 批量删除CabinetProduct
		CabinetProductRouter.PUT("updateCabinetProduct", v1.UpdateCabinetProduct)    // 更新CabinetProduct
		CabinetProductRouter.GET("findCabinetProduct", v1.FindCabinetProduct)        // 根据ID获取CabinetProduct
		CabinetProductRouter.GET("getCabinetProductList", v1.GetCabinetProductList)  // 获取CabinetProduct列表
	}
}
