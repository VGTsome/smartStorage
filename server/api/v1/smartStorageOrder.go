package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"strconv"

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
	var SSOLRQ request.SmartStorageOrderListReq
	_ = c.ShouldBindJSON(&SSOLRQ)

	userId, err1 := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if err1 != nil {
		userId = 0
	}
	err := service.CreateSmartStorageOrder(SSOLRQ, userId)
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

func UpdateSmartStorageOrderStatus(c *gin.Context) {
	var smartStorageOrder model.SmartStorageOrder
	_ = c.ShouldBindJSON(&smartStorageOrder)
	if smartStorageOrder.OrderStatus == 999 {
		_, ssos := service.GetSmartStoragesOrderByOrderID(smartStorageOrder.OrderId)
		for _, sso := range ssos {
			sso.OrderStatus = 0
			service.UpdateSmartStorageOrder(&sso)
			response.OkWithMessage("更新成功", c)
			return
		}
	}
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
func GetSmartStorageOrderInfoById(c *gin.Context) {
	var pageInfo request.SmartStorageOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	userId, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	err, list, total := service.GetSmartStorageOrderInfoById(pageInfo, userId)
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

func GetAllOrderList(c *gin.Context) {
	var pageInfo request.SmartStorageOrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAllOrderList(pageInfo)
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
func GetMyOrderList(c *gin.Context) {
	var pageInfo request.SmartStorageOrderSearch
	userId, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetMyOrderList(pageInfo,userId)
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
