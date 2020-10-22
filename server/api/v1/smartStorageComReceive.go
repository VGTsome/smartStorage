package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
)

// @Tags SmartStorageComReceive
// @Summary 创建SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "创建SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageComReceive/createSmartStorageComReceive [post]
func CreateSmartStorageComReceive(c *gin.Context) {
	var smartStorageComReceive model.SmartStorageComReceive
	_ = c.ShouldBindJSON(&smartStorageComReceive)

	utils.GetData("http://" + global.GVA_CONFIG.System.ComIPaddr + "/Service?comname=COM1&command=bf%20aa%20d4%20b4%20b5%20e7%20d7%20d3%20cd%20f8%203a%2077%2077%2077%202e%206f%2070%2065%206e%2065%2064%2076%202e%2063%206f%206d%200d%200a")
	_ = utils.NumberToHexstring(256)
	err := service.CreateSmartStorageComReceive(smartStorageComReceive)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageComReceive
// @Summary 删除SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "删除SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageComReceive/deleteSmartStorageComReceive [delete]
func DeleteSmartStorageComReceive(c *gin.Context) {
	var smartStorageComReceive model.SmartStorageComReceive
	_ = c.ShouldBindJSON(&smartStorageComReceive)
	err := service.DeleteSmartStorageComReceive(smartStorageComReceive)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageComReceive
// @Summary 批量删除SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageComReceive/deleteSmartStorageComReceiveByIds [delete]
func DeleteSmartStorageComReceiveByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageComReceiveByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageComReceive
// @Summary 更新SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "更新SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageComReceive/updateSmartStorageComReceive [put]
func UpdateSmartStorageComReceive(c *gin.Context) {
	var smartStorageComReceive model.SmartStorageComReceive
	_ = c.ShouldBindJSON(&smartStorageComReceive)
	err := service.UpdateSmartStorageComReceive(&smartStorageComReceive)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageComReceive
// @Summary 用id查询SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "用id查询SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageComReceive/findSmartStorageComReceive [get]
func FindSmartStorageComReceive(c *gin.Context) {
	var smartStorageComReceive model.SmartStorageComReceive
	_ = c.ShouldBindQuery(&smartStorageComReceive)
	err, resmartStorageComReceive := service.GetSmartStorageComReceive(smartStorageComReceive.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageComReceive": resmartStorageComReceive}, c)
	}
}

// @Tags SmartStorageComReceive
// @Summary 分页获取SmartStorageComReceive列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageComReceiveSearch true "分页获取SmartStorageComReceive列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageComReceive/getSmartStorageComReceiveList [get]
func GetSmartStorageComReceiveList(c *gin.Context) {
	var pageInfo request.SmartStorageComReceiveSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageComReceiveInfoList(pageInfo)
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
