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

// @Tags SmartStorageCurrentOrder
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
func CreateSmartStorageCurrentOrder(c *gin.Context) {
	var smartStorageCurrentOrder model.SmartStorageCurrentOrder
	_ = c.ShouldBindJSON(&smartStorageCurrentOrder)

	err := service.CreateSmartStorageCurrentOrder(smartStorageCurrentOrder)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}

// @Tags SmartStorageCurrentOrder
// @Summary 删除SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "删除SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCurrentOrder/deleteSmartStorageCurrentOrder [delete]
func DeleteSmartStorageCurrentOrder(c *gin.Context) {
	var smartStorageCurrentOrder model.SmartStorageCurrentOrder
	_ = c.ShouldBindJSON(&smartStorageCurrentOrder)
	err := service.DeleteSmartStorageCurrentOrder(smartStorageCurrentOrder)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageCurrentOrder
// @Summary 批量删除SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCurrentOrder/deleteSmartStorageCurrentOrderByIds [delete]
func DeleteSmartStorageCurrentOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageCurrentOrderByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageCurrentOrder
// @Summary 更新SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "更新SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageCurrentOrder/updateSmartStorageCurrentOrder [put]
func UpdateSmartStorageCurrentOrder(c *gin.Context) {
	var smartStorageCurrentOrder model.SmartStorageCurrentOrder
	_ = c.ShouldBindJSON(&smartStorageCurrentOrder)
	err := service.UpdateSmartStorageCurrentOrder(&smartStorageCurrentOrder)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageCurrentOrder
// @Summary 用id查询SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "用id查询SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageCurrentOrder/findSmartStorageCurrentOrder [get]
func FindSmartStorageCurrentOrder(c *gin.Context) {
	var smartStorageCurrentOrder model.SmartStorageCurrentOrder
	_ = c.ShouldBindQuery(&smartStorageCurrentOrder)
	err, resmartStorageCurrentOrder := service.GetSmartStorageCurrentOrder(smartStorageCurrentOrder.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageCurrentOrder": resmartStorageCurrentOrder}, c)
	}
}

// @Tags SmartStorageCurrentOrder
// @Summary 分页获取SmartStorageCurrentOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageCurrentOrderSearch true "分页获取SmartStorageCurrentOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/getSmartStorageCurrentOrderList [get]
func GetSmartStorageCurrentOrderList(c *gin.Context) {
	var pageInfo request.SmartStorageCurrentOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	cs := []string{"02", "03"}
	if pageInfo.PageSize == 999 {
		pageInfo.PageSize = 10
		comlogic.UpExitDoor("COM1", cs)
	}
	if pageInfo.PageSize == 99 {
		pageInfo.PageSize = 10
		comlogic.EnterScanFaceID("COM1", cs)
	}
	err, list, total := service.GetSmartStorageCurrentOrderInfoList(pageInfo)

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
