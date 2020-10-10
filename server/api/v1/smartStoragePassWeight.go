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

// @Tags SmartStoragePassWeight
// @Summary 创建SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "创建SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStoragePassWeight/createSmartStoragePassWeight [post]
func CreateSmartStoragePassWeight(c *gin.Context) {
	var smartStoragePassWeight model.SmartStoragePassWeight
	_ = c.ShouldBindJSON(&smartStoragePassWeight)
	err := service.CreateSmartStoragePassWeight(smartStoragePassWeight)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStoragePassWeight
// @Summary 删除SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "删除SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStoragePassWeight/deleteSmartStoragePassWeight [delete]
func DeleteSmartStoragePassWeight(c *gin.Context) {
	var smartStoragePassWeight model.SmartStoragePassWeight
	_ = c.ShouldBindJSON(&smartStoragePassWeight)
	err := service.DeleteSmartStoragePassWeight(smartStoragePassWeight)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStoragePassWeight
// @Summary 批量删除SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStoragePassWeight/deleteSmartStoragePassWeightByIds [delete]
func DeleteSmartStoragePassWeightByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStoragePassWeightByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStoragePassWeight
// @Summary 更新SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "更新SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStoragePassWeight/updateSmartStoragePassWeight [put]
func UpdateSmartStoragePassWeight(c *gin.Context) {
	var smartStoragePassWeight model.SmartStoragePassWeight
	_ = c.ShouldBindJSON(&smartStoragePassWeight)
	err := service.UpdateSmartStoragePassWeight(&smartStoragePassWeight)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStoragePassWeight
// @Summary 用id查询SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "用id查询SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStoragePassWeight/findSmartStoragePassWeight [get]
func FindSmartStoragePassWeight(c *gin.Context) {
	var smartStoragePassWeight model.SmartStoragePassWeight
	_ = c.ShouldBindQuery(&smartStoragePassWeight)
	err, resmartStoragePassWeight := service.GetSmartStoragePassWeight(smartStoragePassWeight.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStoragePassWeight": resmartStoragePassWeight}, c)
	}
}

// @Tags SmartStoragePassWeight
// @Summary 分页获取SmartStoragePassWeight列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStoragePassWeightSearch true "分页获取SmartStoragePassWeight列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStoragePassWeight/getSmartStoragePassWeightList [get]
func GetSmartStoragePassWeightList(c *gin.Context) {
	var pageInfo request.SmartStoragePassWeightSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStoragePassWeightInfoList(pageInfo)
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
