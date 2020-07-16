package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags SmartStorageOrder
// @Summary 创建SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "创建SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageOrder/createSmartStorageOrder [post]
func CreateSmartStorageOrder(c *gin.Context) {
	var smartStorageOrder model.SmartStorageOrder
	_ = c.ShouldBindJSON(&smartStorageOrder)
	err := service.CreateSmartStorageOrder(smartStorageOrder)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageOrder
// @Summary 删除SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "删除SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageOrder/deleteSmartStorageOrder [delete]
func DeleteSmartStorageOrder(c *gin.Context) {
	var smartStorageOrder model.SmartStorageOrder
	_ = c.ShouldBindJSON(&smartStorageOrder)
	err := service.DeleteSmartStorageOrder(smartStorageOrder)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageOrder
// @Summary 批量删除SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageOrder/deleteSmartStorageOrderByIds [delete]
func DeleteSmartStorageOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageOrderByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageOrder
// @Summary 更新SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "更新SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageOrder/updateSmartStorageOrder [put]
func UpdateSmartStorageOrder(c *gin.Context) {
	var smartStorageOrder model.SmartStorageOrder
	_ = c.ShouldBindJSON(&smartStorageOrder)
	err := service.UpdateSmartStorageOrder(&smartStorageOrder)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageOrder
// @Summary 用id查询SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "用id查询SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageOrder/findSmartStorageOrder [get]
func FindSmartStorageOrder(c *gin.Context) {
	var smartStorageOrder model.SmartStorageOrder
	_ = c.ShouldBindQuery(&smartStorageOrder)
	err, resmartStorageOrder := service.GetSmartStorageOrder(smartStorageOrder.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageOrder": resmartStorageOrder}, c)
	}
}

// @Tags SmartStorageOrder
// @Summary 分页获取SmartStorageOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageOrderSearch true "分页获取SmartStorageOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageOrder/getSmartStorageOrderList [get]
func GetSmartStorageOrderList(c *gin.Context) {
	var pageInfo request.SmartStorageOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageOrderInfoList(pageInfo)
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
