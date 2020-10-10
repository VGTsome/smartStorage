import service from '@/utils/request'

// @Tags SmartStorageWeightLog
// @Summary 创建SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "创建SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageWeightLog/createSmartStorageWeightLog [post]
export const createSmartStorageWeightLog = (data) => {
     return service({
         url: "/smartStorageWeightLog/createSmartStorageWeightLog",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageWeightLog
// @Summary 删除SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "删除SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageWeightLog/deleteSmartStorageWeightLog [delete]
 export const deleteSmartStorageWeightLog = (data) => {
     return service({
         url: "/smartStorageWeightLog/deleteSmartStorageWeightLog",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageWeightLog
// @Summary 删除SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageWeightLog/deleteSmartStorageWeightLog [delete]
 export const deleteSmartStorageWeightLogByIds = (data) => {
     return service({
         url: "/smartStorageWeightLog/deleteSmartStorageWeightLogByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageWeightLog
// @Summary 更新SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "更新SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageWeightLog/updateSmartStorageWeightLog [put]
 export const updateSmartStorageWeightLog = (data) => {
     return service({
         url: "/smartStorageWeightLog/updateSmartStorageWeightLog",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageWeightLog
// @Summary 用id查询SmartStorageWeightLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageWeightLog true "用id查询SmartStorageWeightLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageWeightLog/findSmartStorageWeightLog [get]
 export const findSmartStorageWeightLog = (params) => {
     return service({
         url: "/smartStorageWeightLog/findSmartStorageWeightLog",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageWeightLog
// @Summary 分页获取SmartStorageWeightLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageWeightLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageWeightLog/getSmartStorageWeightLogList [get]
 export const getSmartStorageWeightLogList = (params) => {
     return service({
         url: "/smartStorageWeightLog/getSmartStorageWeightLogList",
         method: 'get',
         params
     })
 }