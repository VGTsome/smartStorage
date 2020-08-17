import service from '@/utils/request'

// @Tags SmartStorageQuota
// @Summary 创建SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "创建SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssq/createSmartStorageQuota [post]
export const createSmartStorageQuota = (data) => {
     return service({
         url: "/ssq/createSmartStorageQuota",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageQuota
// @Summary 删除SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "删除SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssq/deleteSmartStorageQuota [delete]
 export const deleteSmartStorageQuota = (data) => {
     return service({
         url: "/ssq/deleteSmartStorageQuota",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageQuota
// @Summary 删除SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ssq/deleteSmartStorageQuota [delete]
 export const deleteSmartStorageQuotaByIds = (data) => {
     return service({
         url: "/ssq/deleteSmartStorageQuotaByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageQuota
// @Summary 更新SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "更新SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ssq/updateSmartStorageQuota [put]
 export const updateSmartStorageQuota = (data) => {
     return service({
         url: "/ssq/updateSmartStorageQuota",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageQuota
// @Summary 用id查询SmartStorageQuota
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageQuota true "用id查询SmartStorageQuota"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ssq/findSmartStorageQuota [get]
 export const findSmartStorageQuota = (params) => {
     return service({
         url: "/ssq/findSmartStorageQuota",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageQuota
// @Summary 分页获取SmartStorageQuota列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageQuota列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ssq/getSmartStorageQuotaList [get]
 export const getSmartStorageQuotaList = (params) => {
     return service({
         url: "/ssq/getSmartStorageQuotaList",
         method: 'get',
         params
     })
 }