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

// @Tags SmartStorageCabinet
// @Summary 创建SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "创建SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCabinet/createSmartStorageCabinet [post]
func CreateSmartStorageCabinet(c *gin.Context) {
	var smartStorageCabinet model.SmartStorageCabinet
	_ = c.ShouldBindJSON(&smartStorageCabinet)
	err := service.CreateSmartStorageCabinet(smartStorageCabinet)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageCabinet
// @Summary 删除SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "删除SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCabinet/deleteSmartStorageCabinet [delete]
func DeleteSmartStorageCabinet(c *gin.Context) {
	var smartStorageCabinet model.SmartStorageCabinet
	_ = c.ShouldBindJSON(&smartStorageCabinet)
	err := service.DeleteSmartStorageCabinet(smartStorageCabinet)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageCabinet
// @Summary 批量删除SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCabinet/deleteSmartStorageCabinetByIds [delete]
func DeleteSmartStorageCabinetByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageCabinetByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageCabinet
// @Summary 更新SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "更新SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageCabinet/updateSmartStorageCabinet [put]
func UpdateSmartStorageCabinet(c *gin.Context) {
	var smartStorageCabinet model.SmartStorageCabinet
	_ = c.ShouldBindJSON(&smartStorageCabinet)
	err := service.UpdateSmartStorageCabinet(&smartStorageCabinet)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageCabinet
// @Summary 用id查询SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "用id查询SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageCabinet/findSmartStorageCabinet [get]
func FindSmartStorageCabinet(c *gin.Context) {
	var smartStorageCabinet model.SmartStorageCabinet
	_ = c.ShouldBindQuery(&smartStorageCabinet)
	err, resmartStorageCabinet := service.GetSmartStorageCabinet(smartStorageCabinet.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageCabinet": resmartStorageCabinet}, c)
	}
}

// @Tags SmartStorageCabinet
// @Summary 分页获取SmartStorageCabinet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageCabinetSearch true "分页获取SmartStorageCabinet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCabinet/getSmartStorageCabinetList [get]
func GetSmartStorageCabinetList(c *gin.Context) {
	var pageInfo request.SmartStorageCabinetSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageCabinetInfoList(pageInfo)
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
