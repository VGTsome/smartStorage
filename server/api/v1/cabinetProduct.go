package v1

import (
	"fmt"
	"gin-vue-admin/comlogic"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

// @Tags CabinetProduct
// @Summary 创建CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "创建CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sscp/createCabinetProduct [post]
func CreateCabinetProduct(c *gin.Context) {
	var sscp model.CabinetProduct
	_ = c.ShouldBindJSON(&sscp)
	err := service.CreateCabinetProduct(sscp)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags CabinetProduct
// @Summary 删除CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "删除CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sscp/deleteCabinetProduct [delete]
func DeleteCabinetProduct(c *gin.Context) {
	var sscp model.CabinetProduct
	_ = c.ShouldBindJSON(&sscp)
	err := service.DeleteCabinetProduct(sscp)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CabinetProduct
// @Summary 批量删除CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sscp/deleteCabinetProductByIds [delete]
func DeleteCabinetProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCabinetProductByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CabinetProduct
// @Summary 更新CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "更新CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sscp/updateCabinetProduct [put]
func UpdateCabinetProduct(c *gin.Context) {
	var sscp model.CabinetProduct
	_ = c.ShouldBindJSON(&sscp)

	err1, ssc := service.GetSmartStorageCabinetByCabinetID(sscp.CabinetId)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err1), c)
		return
	} else {
		//清零
		if sscp.ProductNumber == 0 {
			comlogic.SetZero(ssc.CabinetName)

		} else {
			//初始化
			comlogic.InitProduct(ssc.CabinetName, sscp.ProductNumber)

		}
	}

	err := service.UpdateCabinetProduct(&sscp)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags CabinetProduct
// @Summary 用id查询CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "用id查询CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sscp/findCabinetProduct [get]
func FindCabinetProduct(c *gin.Context) {
	var sscp model.CabinetProduct
	_ = c.ShouldBindQuery(&sscp)
	err, resscp := service.GetCabinetProduct(sscp.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resscp": resscp}, c)
	}
}

// @Tags CabinetProduct
// @Summary 分页获取CabinetProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CabinetProductSearch true "分页获取CabinetProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sscp/getCabinetProductList [get]
func GetCabinetProductList(c *gin.Context) {
	var pageInfo request.CabinetProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	//更新数量和单位重量
	if pageInfo.PageSize == 999 {
		pageInfo.PageSize = 10
		comlogic.UpdateAllProd()

	}
	err, list, total := service.GetCabinetProductInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
