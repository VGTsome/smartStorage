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

// @Tags SmartStorageDoor
// @Summary 创建SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "创建SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageDoor/createSmartStorageDoor [post]
func CreateSmartStorageDoor(c *gin.Context) {
	var smartStorageDoor model.SmartStorageDoor
	_ = c.ShouldBindJSON(&smartStorageDoor)
	err := service.CreateSmartStorageDoor(smartStorageDoor)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageDoor
// @Summary 删除SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "删除SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageDoor/deleteSmartStorageDoor [delete]
func DeleteSmartStorageDoor(c *gin.Context) {
	var smartStorageDoor model.SmartStorageDoor
	_ = c.ShouldBindJSON(&smartStorageDoor)
	err := service.DeleteSmartStorageDoor(smartStorageDoor)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageDoor
// @Summary 批量删除SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageDoor/deleteSmartStorageDoorByIds [delete]
func DeleteSmartStorageDoorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageDoorByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageDoor
// @Summary 更新SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "更新SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageDoor/updateSmartStorageDoor [put]
func UpdateSmartStorageDoor(c *gin.Context) {
	var smartStorageDoor model.SmartStorageDoor
	_ = c.ShouldBindJSON(&smartStorageDoor)
	err := service.UpdateSmartStorageDoor(&smartStorageDoor)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageDoor
// @Summary 用id查询SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "用id查询SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageDoor/findSmartStorageDoor [get]
func FindSmartStorageDoor(c *gin.Context) {
	var smartStorageDoor model.SmartStorageDoor
	_ = c.ShouldBindQuery(&smartStorageDoor)
	err, resmartStorageDoor := service.GetSmartStorageDoor(smartStorageDoor.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageDoor": resmartStorageDoor}, c)
	}
}

// @Tags SmartStorageDoor
// @Summary 分页获取SmartStorageDoor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageDoorSearch true "分页获取SmartStorageDoor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageDoor/getSmartStorageDoorList [get]
func GetSmartStorageDoorList(c *gin.Context) {
	var pageInfo request.SmartStorageDoorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageDoorInfoList(pageInfo)
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
