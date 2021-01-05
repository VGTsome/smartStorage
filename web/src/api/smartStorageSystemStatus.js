import service from '@/utils/request'

// @Tags SmartStorageSystemStatus
// @Summary 创建SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "创建SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssss/createSmartStorageSystemStatus [post]
export const createSmartStorageSystemStatus = (data) => {
     return service({
         url: "/ssss/createSmartStorageSystemStatus",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageSystemStatus
// @Summary 删除SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "删除SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssss/deleteSmartStorageSystemStatus [delete]
 export const deleteSmartStorageSystemStatus = (data) => {
     return service({
         url: "/ssss/deleteSmartStorageSystemStatus",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageSystemStatus
// @Summary 删除SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssss/deleteSmartStorageSystemStatus [delete]
 export const deleteSmartStorageSystemStatusByIds = (data) => {
     return service({
         url: "/ssss/deleteSmartStorageSystemStatusByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageSystemStatus
// @Summary 更新SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "更新SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ssss/updateSmartStorageSystemStatus [put]
 export const updateSmartStorageSystemStatus = (data) => {
     return service({
         url: "/ssss/updateSmartStorageSystemStatus",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageSystemStatus
// @Summary 用id查询SmartStorageSystemStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageSystemStatus true "用id查询SmartStorageSystemStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ssss/findSmartStorageSystemStatus [get]
 export const findSmartStorageSystemStatus = (params) => {
     return service({
         url: "/ssss/findSmartStorageSystemStatus",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageSystemStatus
// @Summary 分页获取SmartStorageSystemStatus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageSystemStatus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssss/getSmartStorageSystemStatusList [get]
 export const getSmartStorageSystemStatusList = (params) => {
     return service({
         url: "/ssss/getSmartStorageSystemStatusList",
         method: 'get',
         params
     })
 }