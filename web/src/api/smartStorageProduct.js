import service from '@/utils/request'

// @Tags SmartStorageProduct
// @Summary 创建SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "创建SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageProduct/createSmartStorageProduct [post]
export const createSmartStorageProduct = (data) => {
     return service({
         url: "/smartStorageProduct/createSmartStorageProduct",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageProduct
// @Summary 删除SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "删除SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageProduct/deleteSmartStorageProduct [delete]
 export const deleteSmartStorageProduct = (data) => {
     return service({
         url: "/smartStorageProduct/deleteSmartStorageProduct",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageProduct
// @Summary 删除SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageProduct/deleteSmartStorageProduct [delete]
 export const deleteSmartStorageProductByIds = (data) => {
     return service({
         url: "/smartStorageProduct/deleteSmartStorageProductByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageProduct
// @Summary 更新SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "更新SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageProduct/updateSmartStorageProduct [put]
 export const updateSmartStorageProduct = (data) => {
     return service({
         url: "/smartStorageProduct/updateSmartStorageProduct",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageProduct
// @Summary 用id查询SmartStorageProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageProduct true "用id查询SmartStorageProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageProduct/findSmartStorageProduct [get]
 export const findSmartStorageProduct = (params) => {
     return service({
         url: "/smartStorageProduct/findSmartStorageProduct",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageProduct
// @Summary 分页获取SmartStorageProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageProduct/getSmartStorageProductList [get]
 export const getSmartStorageProductList = (params) => {
     return service({
         url: "/smartStorageProduct/getSmartStorageProductList",
         method: 'get',
         params
     })
 }