import service from '@/utils/request'

// @Tags SmartStorageCurrentOrder
// @Summary 创建SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "创建SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/createSmartStorageCurrentOrder [post]
export const createSmartStorageCurrentOrder = (data) => {
     return service({
         url: "/smartStorageCurrentOrder/createSmartStorageCurrentOrder",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageCurrentOrder
// @Summary 删除SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "删除SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCurrentOrder/deleteSmartStorageCurrentOrder [delete]
 export const deleteSmartStorageCurrentOrder = (data) => {
     return service({
         url: "/smartStorageCurrentOrder/deleteSmartStorageCurrentOrder",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageCurrentOrder
// @Summary 删除SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCurrentOrder/deleteSmartStorageCurrentOrder [delete]
 export const deleteSmartStorageCurrentOrderByIds = (data) => {
     return service({
         url: "/smartStorageCurrentOrder/deleteSmartStorageCurrentOrderByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageCurrentOrder
// @Summary 更新SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "更新SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageCurrentOrder/updateSmartStorageCurrentOrder [put]
 export const updateSmartStorageCurrentOrder = (data) => {
     return service({
         url: "/smartStorageCurrentOrder/updateSmartStorageCurrentOrder",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageCurrentOrder
// @Summary 用id查询SmartStorageCurrentOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCurrentOrder true "用id查询SmartStorageCurrentOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageCurrentOrder/findSmartStorageCurrentOrder [get]
 export const findSmartStorageCurrentOrder = (params) => {
     return service({
         url: "/smartStorageCurrentOrder/findSmartStorageCurrentOrder",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageCurrentOrder
// @Summary 分页获取SmartStorageCurrentOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageCurrentOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCurrentOrder/getSmartStorageCurrentOrderList [get]
 export const getSmartStorageCurrentOrderList = (params) => {
     return service({
         url: "/smartStorageCurrentOrder/getSmartStorageCurrentOrderList",
         method: 'get',
         params
     })
 }