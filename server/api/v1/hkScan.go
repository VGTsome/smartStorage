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

// @Tags HkScan
// @Summary 创建HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "创建HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkScan/createHkScan [post]
func CreateHkScan(c *gin.Context) {
	var hkScan model.HkScan
	_ = c.ShouldBindJSON(&hkScan)
	err := service.CreateHkScan(hkScan)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags HkScan
// @Summary 删除HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "删除HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkScan/deleteHkScan [delete]
func DeleteHkScan(c *gin.Context) {
	var hkScan model.HkScan
	_ = c.ShouldBindJSON(&hkScan)
	err := service.DeleteHkScan(hkScan)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HkScan
// @Summary 批量删除HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkScan/deleteHkScanByIds [delete]
func DeleteHkScanByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHkScanByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags HkScan
// @Summary 更新HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "更新HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkScan/updateHkScan [put]
func UpdateHkScan(c *gin.Context) {
	var hkScan model.HkScan
	_ = c.ShouldBindJSON(&hkScan)
	err := service.UpdateHkScan(&hkScan)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags HkScan
// @Summary 用id查询HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "用id查询HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkScan/findHkScan [get]
func FindHkScan(c *gin.Context) {
	var hkScan model.HkScan
	_ = c.ShouldBindQuery(&hkScan)
	err, rehkScan := service.GetHkScan(hkScan.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehkScan": rehkScan}, c)
	}
}

// @Tags HkScan
// @Summary 分页获取HkScan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HkScanSearch true "分页获取HkScan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkScan/getHkScanList [get]
func GetHkScanList(c *gin.Context) {
	var pageInfo request.HkScanSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHkScanInfoList(pageInfo)
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
