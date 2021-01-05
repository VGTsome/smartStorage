package v1

import (
	"fmt"
	"gin-vue-admin/comlogic"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"strconv"

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
func UpdateShelf(c *gin.Context) {
	var she model.Shelf
	_ = c.ShouldBindJSON(&she)
	if she.ShelfID != "0" {
		shelfNum, _ := strconv.Atoi(she.ShelfID)
		comlogic.UpdateAllProd("01", shelfNum)
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(fmt.Sprintf("更新失败"), c)
	}

}
func PrepareShelf(c *gin.Context) {
	var she model.Shelf
	_ = c.ShouldBindJSON(&she)
	if she.ShelfID != "0" {
		shelfNum, _ := strconv.Atoi(she.ShelfID)
		comlogic.UpdateAllProd("00", shelfNum)
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(fmt.Sprintf("更新失败"), c)
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
			systemStatus := service.GetSystemStatus()
			if systemStatus < 3 {
				if systemStatus == 1 {
					service.SetSystemStatus(2)
				}
				comlogic.SetZero(ssc.CabinetName)
				response.OkWithMessage("更新成功", c)
				return
			} else {
				response.FailWithMessage(fmt.Sprintf("更新失败，%v", "有用户在取货"), c)
				return
			}

		} else if sscp.ProductNumber == 1000 {
			systemStatus := service.GetSystemStatus()
			if systemStatus < 3 {
				if systemStatus == 1 {
					service.SetSystemStatus(2)
				}
				comlogic.Set1000(ssc.CabinetName)
				response.OkWithMessage("更新成功", c)
				return
			} else {
				response.FailWithMessage(fmt.Sprintf("更新失败，%v", "有用户在取货"), c)
				return
			}
		} else if sscp.ProductNumber > 0 {
			//初始化
			systemStatus := service.GetSystemStatus()
			if systemStatus < 3 {
				if systemStatus == 1 {
					service.SetSystemStatus(2)
				}
				comlogic.InitProduct(ssc.CabinetName, sscp.ProductNumber)
				response.OkWithMessage("更新成功", c)
				return
			} else {
				response.FailWithMessage(fmt.Sprintf("更新失败，%v", "有用户在取货"), c)
				return
			}
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

	// //更新数量和单位重量
	// if pageInfo.PageSize == 919 {
	// 	//发送预备动作
	// 	comlogic.UpdateAllProd("00", pageInfo.Page)
	// 	pageInfo.PageSize = 10
	// 	pageInfo.Page = 1
	// }
	// if pageInfo.PageSize == 999 {
	// 	//更新货架
	// 	comlogic.UpdateAllProd("01", pageInfo.Page)
	// 	pageInfo.PageSize = 10
	// 	pageInfo.Page = 1
	// }

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
