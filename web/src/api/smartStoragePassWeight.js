import service from '@/utils/request'

// @Tags SmartStoragePassWeight
// @Summary 创建SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "创建SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStoragePassWeight/createSmartStoragePassWeight [post]
export const createSmartStoragePassWeight = (data) => {
     return service({
         url: "/smartStoragePassWeight/createSmartStoragePassWeight",
         method: 'post',
         data
     })
 }


// @Tags SmartStoragePassWeight
// @Summary 删除SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "删除SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStoragePassWeight/deleteSmartStoragePassWeight [delete]
 export const deleteSmartStoragePassWeight = (data) => {
     return service({
         url: "/smartStoragePassWeight/deleteSmartStoragePassWeight",
         method: 'delete',
         data
     })
 }

// @Tags SmartStoragePassWeight
// @Summary 删除SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStoragePassWeight/deleteSmartStoragePassWeight [delete]
 export const deleteSmartStoragePassWeightByIds = (data) => {
     return service({
         url: "/smartStoragePassWeight/deleteSmartStoragePassWeightByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStoragePassWeight
// @Summary 更新SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "更新SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStoragePassWeight/updateSmartStoragePassWeight [put]
 export const updateSmartStoragePassWeight = (data) => {
     return service({
         url: "/smartStoragePassWeight/updateSmartStoragePassWeight",
         method: 'put',
         data
     })
 }


// @Tags SmartStoragePassWeight
// @Summary 用id查询SmartStoragePassWeight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStoragePassWeight true "用id查询SmartStoragePassWeight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStoragePassWeight/findSmartStoragePassWeight [get]
 export const findSmartStoragePassWeight = (params) => {
     return service({
         url: "/smartStoragePassWeight/findSmartStoragePassWeight",
         method: 'get',
         params
     })
 }


// @Tags SmartStoragePassWeight
// @Summary 分页获取SmartStoragePassWeight列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStoragePassWeight列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStoragePassWeight/getSmartStoragePassWeightList [get]
 export const getSmartStoragePassWeightList = (params) => {
     return service({
         url: "/smartStoragePassWeight/getSmartStoragePassWeightList",
         method: 'get',
         params
     })
 }