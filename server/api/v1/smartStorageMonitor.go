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

// @Tags SmartStorageMonitor
// @Summary 创建SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "创建SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageMonitor/createSmartStorageMonitor [post]
func CreateSmartStorageMonitor(c *gin.Context) {
	var smartStorageMonitor model.SmartStorageMonitor
	_ = c.ShouldBindJSON(&smartStorageMonitor)
	err := service.CreateSmartStorageMonitor(smartStorageMonitor)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageMonitor
// @Summary 删除SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "删除SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageMonitor/deleteSmartStorageMonitor [delete]
func DeleteSmartStorageMonitor(c *gin.Context) {
	var smartStorageMonitor model.SmartStorageMonitor
	_ = c.ShouldBindJSON(&smartStorageMonitor)
	err := service.DeleteSmartStorageMonitor(smartStorageMonitor)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageMonitor
// @Summary 批量删除SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageMonitor/deleteSmartStorageMonitorByIds [delete]
func DeleteSmartStorageMonitorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageMonitorByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageMonitor
// @Summary 更新SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "更新SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageMonitor/updateSmartStorageMonitor [put]
func UpdateSmartStorageMonitor(c *gin.Context) {
	var smartStorageMonitor model.SmartStorageMonitor
	_ = c.ShouldBindJSON(&smartStorageMonitor)
	err := service.UpdateSmartStorageMonitor(&smartStorageMonitor)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageMonitor
// @Summary 用id查询SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "用id查询SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageMonitor/findSmartStorageMonitor [get]
func FindSmartStorageMonitor(c *gin.Context) {
	var smartStorageMonitor model.SmartStorageMonitor
	_ = c.ShouldBindQuery(&smartStorageMonitor)
	err, resmartStorageMonitor := service.GetSmartStorageMonitor(smartStorageMonitor.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageMonitor": resmartStorageMonitor}, c)
	}
}

// @Tags SmartStorageMonitor
// @Summary 分页获取SmartStorageMonitor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageMonitorSearch true "分页获取SmartStorageMonitor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageMonitor/getSmartStorageMonitorList [get]
func GetSmartStorageMonitorList(c *gin.Context) {
	var pageInfo request.SmartStorageMonitorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageMonitorInfoList(pageInfo)
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
