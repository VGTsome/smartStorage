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

// @Tags SmartStorageProduct
// @Summary 创建SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "创建SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageProduct/createSmartStorageProduct [post]
func CreateSmartStorageProduct(c *gin.Context) {
	var smartStorageProduct model.SmartStorageProduct
	_ = c.ShouldBindJSON(&smartStorageProduct)

	err := service.CreateSmartStorageProduct(smartStorageProduct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SmartStorageProduct
// @Summary 删除SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "删除SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageProduct/deleteSmartStorageProduct [delete]
func DeleteSmartStorageProduct(c *gin.Context) {
	var smartStorageProduct model.SmartStorageProduct
	_ = c.ShouldBindJSON(&smartStorageProduct)
	err := service.DeleteSmartStorageProduct(smartStorageProduct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageProduct
// @Summary 批量删除SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageProduct/deleteSmartStorageProductByIds [delete]
func DeleteSmartStorageProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSmartStorageProductByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SmartStorageProduct
// @Summary 更新SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "更新SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageProduct/updateSmartStorageProduct [put]
func UpdateSmartStorageProduct(c *gin.Context) {
	var smartStorageProduct model.SmartStorageProduct
	_ = c.ShouldBindJSON(&smartStorageProduct)
	err := service.UpdateSmartStorageProduct(&smartStorageProduct)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SmartStorageProduct
// @Summary 用id查询SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "用id查询SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageProduct/findSmartStorageProduct [get]
func FindSmartStorageProduct(c *gin.Context) {
	var smartStorageProduct model.SmartStorageProduct
	_ = c.ShouldBindQuery(&smartStorageProduct)
	err, resmartStorageProduct := service.GetSmartStorageProduct(smartStorageProduct.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resmartStorageProduct": resmartStorageProduct}, c)
	}
}

// @Tags SmartStorageProduct
// @Summary 分页获取SmartStorageProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmartStorageProductSearch true "分页获取SmartStorageProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageProduct/getSmartStorageProductList [get]
func GetSmartStorageProductList(c *gin.Context) {
	var pageInfo request.SmartStorageProductSearch

	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageProductInfoList(pageInfo)
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
func GetSmartStorageProductValidList(c *gin.Context) {
	var pageInfo request.SmartStorageProductSearch

	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSmartStorageProductInfoValidList(pageInfo)
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
