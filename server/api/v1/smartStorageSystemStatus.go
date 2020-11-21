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

// @Tags SmartStorageSystemStatus
// @Summary 创建SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "创建SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssss/createSmartStorageSystemStatus [post]
func CreateSmartStorageSystemStatus(c *gin.Context) {
	var ssss model.SmartStorageSystemStatus
	_ = c.ShouldBindJSON(&ssss)
	err := service.CreateSmartStorageSystemStatus(ssss)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageSystemStatus
// @Summary 删除SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "删除SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssss/deleteSmartStorageSystemStatus [delete]
func DeleteSmartStorageSystemStatus(c *gin.Context) {
	var ssss model.SmartStorageSystemStatus
	_ = c.ShouldBindJSON(&ssss)
	err := service.DeleteSmartStorageSystemStatus(ssss)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageSystemStatus
// @Summary 批量删除SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssss/deleteSmartStorageSystemStatusByIds [delete]
func DeleteSmartStorageSystemStatusByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageSystemStatusByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageSystemStatus
// @Summary 更新SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "更新SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ssss/updateSmartStorageSystemStatus [put]
func UpdateSmartStorageSystemStatus(c *gin.Context) {
	var ssss model.SmartStorageSystemStatus
	_ = c.ShouldBindJSON(&ssss)
	err := service.UpdateSmartStorageSystemStatus(&ssss)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageSystemStatus
// @Summary 用id查询SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "用id查询SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ssss/findSmartStorageSystemStatus [get]
func FindSmartStorageSystemStatus(c *gin.Context) {
	var ssss model.SmartStorageSystemStatus
	_ = c.ShouldBindQuery(&ssss)
	err, ressss := service.GetSmartStorageSystemStatus(ssss.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"ressss": ressss}, c)
	}
}

// @Tags SmartStorageSystemStatus
// @Summary 分页获取SmartStorageSystemStatus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageSystemStatusSearch true "分页获取SmartStorageSystemStatus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssss/getSmartStorageSystemStatusList [get]
func GetSmartStorageSystemStatusList(c *gin.Context) {
	var pageInfo request.SmartStorageSystemStatusSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageSystemStatusInfoList(pageInfo)
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
