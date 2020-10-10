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

// @Tags SmartStorageWeightLog
// @Summary 创建SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "创建SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageWeightLog/createSmartStorageWeightLog [post]
func CreateSmartStorageWeightLog(c *gin.Context) {
	var smartStorageWeightLog model.SmartStorageWeightLog
	_ = c.ShouldBindJSON(&smartStorageWeightLog)
	err := service.CreateSmartStorageWeightLog(smartStorageWeightLog)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageWeightLog
// @Summary 删除SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "删除SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageWeightLog/deleteSmartStorageWeightLog [delete]
func DeleteSmartStorageWeightLog(c *gin.Context) {
	var smartStorageWeightLog model.SmartStorageWeightLog
	_ = c.ShouldBindJSON(&smartStorageWeightLog)
	err := service.DeleteSmartStorageWeightLog(smartStorageWeightLog)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageWeightLog
// @Summary 批量删除SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageWeightLog/deleteSmartStorageWeightLogByIds [delete]
func DeleteSmartStorageWeightLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageWeightLogByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageWeightLog
// @Summary 更新SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "更新SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageWeightLog/updateSmartStorageWeightLog [put]
func UpdateSmartStorageWeightLog(c *gin.Context) {
	var smartStorageWeightLog model.SmartStorageWeightLog
	_ = c.ShouldBindJSON(&smartStorageWeightLog)
	err := service.UpdateSmartStorageWeightLog(&smartStorageWeightLog)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageWeightLog
// @Summary 用id查询SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "用id查询SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageWeightLog/findSmartStorageWeightLog [get]
func FindSmartStorageWeightLog(c *gin.Context) {
	var smartStorageWeightLog model.SmartStorageWeightLog
	_ = c.ShouldBindQuery(&smartStorageWeightLog)
	err, resmartStorageWeightLog := service.GetSmartStorageWeightLog(smartStorageWeightLog.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageWeightLog": resmartStorageWeightLog}, c)
	}
}

// @Tags SmartStorageWeightLog
// @Summary 分页获取SmartStorageWeightLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageWeightLogSearch true "分页获取SmartStorageWeightLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageWeightLog/getSmartStorageWeightLogList [get]
func GetSmartStorageWeightLogList(c *gin.Context) {
	var pageInfo request.SmartStorageWeightLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageWeightLogInfoList(pageInfo)
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
