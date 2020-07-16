import service from '@/utils/request'

// @Tags SmartStorageMonitor
// @Summary 创建SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "创建SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageMonitor/createSmartStorageMonitor [post]
export const createSmartStorageMonitor = (data) => {
     return service({
         url: "/smartStorageMonitor/createSmartStorageMonitor",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageMonitor
// @Summary 删除SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "删除SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageMonitor/deleteSmartStorageMonitor [delete]
 export const deleteSmartStorageMonitor = (data) => {
     return service({
         url: "/smartStorageMonitor/deleteSmartStorageMonitor",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageMonitor
// @Summary 删除SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageMonitor/deleteSmartStorageMonitor [delete]
 export const deleteSmartStorageMonitorByIds = (data) => {
     return service({
         url: "/smartStorageMonitor/deleteSmartStorageMonitorByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageMonitor
// @Summary 更新SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "更新SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageMonitor/updateSmartStorageMonitor [put]
 export const updateSmartStorageMonitor = (data) => {
     return service({
         url: "/smartStorageMonitor/updateSmartStorageMonitor",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageMonitor
// @Summary 用id查询SmartStorageMonitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageMonitor true "用id查询SmartStorageMonitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageMonitor/findSmartStorageMonitor [get]
 export const findSmartStorageMonitor = (params) => {
     return service({
         url: "/smartStorageMonitor/findSmartStorageMonitor",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageMonitor
// @Summary 分页获取SmartStorageMonitor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageMonitor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageMonitor/getSmartStorageMonitorList [get]
 export const getSmartStorageMonitorList = (params) => {
     return service({
         url: "/smartStorageMonitor/getSmartStorageMonitorList",
         method: 'get',
         params
     })
 }