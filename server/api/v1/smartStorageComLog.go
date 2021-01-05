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

// @Tags SmartStorageComLog
// @Summary 创建SmartStorageComLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComLog true "创建SmartStorageComLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageComLog/createSmartStorageComLog [post]
func CreateSmartStorageComLog(c *gin.Context) {
	var smartStorageComLog model.SmartStorageComLog
	_ = c.ShouldBindJSON(&smartStorageComLog)
	err := service.CreateSmartStorageComLog(smartStorageComLog)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageComLog
// @Summary 删除SmartStorageComLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComLog true "删除SmartStorageComLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageComLog/deleteSmartStorageComLog [delete]
func DeleteSmartStorageComLog(c *gin.Context) {
	var smartStorageComLog model.SmartStorageComLog
	_ = c.ShouldBindJSON(&smartStorageComLog)
	err := service.DeleteSmartStorageComLog(smartStorageComLog)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageComLog
// @Summary 批量删除SmartStorageComLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageComLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageComLog/deleteSmartStorageComLogByIds [delete]
func DeleteSmartStorageComLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageComLogByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageComLog
// @Summary 更新SmartStorageComLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComLog true "更新SmartStorageComLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageComLog/updateSmartStorageComLog [put]
func UpdateSmartStorageComLog(c *gin.Context) {
	var smartStorageComLog model.SmartStorageComLog
	_ = c.ShouldBindJSON(&smartStorageComLog)
	err := service.UpdateSmartStorageComLog(&smartStorageComLog)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageComLog
// @Summary 用id查询SmartStorageComLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComLog true "用id查询SmartStorageComLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageComLog/findSmartStorageComLog [get]
func FindSmartStorageComLog(c *gin.Context) {
	var smartStorageComLog model.SmartStorageComLog
	_ = c.ShouldBindQuery(&smartStorageComLog)
	err, resmartStorageComLog := service.GetSmartStorageComLog(smartStorageComLog.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageComLog": resmartStorageComLog}, c)
	}
}

// @Tags SmartStorageComLog
// @Summary 分页获取SmartStorageComLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageComLogSearch true "分页获取SmartStorageComLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageComLog/getSmartStorageComLogList [get]
func GetSmartStorageComLogList(c *gin.Context) {
	var pageInfo request.SmartStorageComLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageComLogInfoList(pageInfo)
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
