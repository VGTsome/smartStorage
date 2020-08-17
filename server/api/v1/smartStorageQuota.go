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

// @Tags SmartStorageQuota
// @Summary 创建SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "创建SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssq/createSmartStorageQuota [post]
func CreateSmartStorageQuota(c *gin.Context) {
	var ssq model.SmartStorageQuota
	_ = c.ShouldBindJSON(&ssq)
	err := service.CreateSmartStorageQuota(ssq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageQuota
// @Summary 删除SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "删除SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssq/deleteSmartStorageQuota [delete]
func DeleteSmartStorageQuota(c *gin.Context) {
	var ssq model.SmartStorageQuota
	_ = c.ShouldBindJSON(&ssq)
	err := service.DeleteSmartStorageQuota(ssq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageQuota
// @Summary 批量删除SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssq/deleteSmartStorageQuotaByIds [delete]
func DeleteSmartStorageQuotaByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageQuotaByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageQuota
// @Summary 更新SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "更新SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ssq/updateSmartStorageQuota [put]
func UpdateSmartStorageQuota(c *gin.Context) {
	var ssq model.SmartStorageQuota
	_ = c.ShouldBindJSON(&ssq)
	err := service.UpdateSmartStorageQuota(&ssq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageQuota
// @Summary 用id查询SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "用id查询SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ssq/findSmartStorageQuota [get]
func FindSmartStorageQuota(c *gin.Context) {
	var ssq model.SmartStorageQuota
	_ = c.ShouldBindQuery(&ssq)
	err, ressq := service.GetSmartStorageQuota(ssq.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"ressq": ressq}, c)
	}
}

// @Tags SmartStorageQuota
// @Summary 分页获取SmartStorageQuota列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageQuotaSearch true "分页获取SmartStorageQuota列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssq/getSmartStorageQuotaList [get]
func GetSmartStorageQuotaList(c *gin.Context) {
	var pageInfo request.SmartStorageQuotaSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageQuotaInfoList(pageInfo)
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
