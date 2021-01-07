import service from '@/utils/request'

// @Tags HkScan
// @Summary 创建HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "创建HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkScan/createHkScan [post]
export const createHkScan = (data) => {
     return service({
         url: "/hkScan/createHkScan",
         method: 'post',
         data
     })
 }


// @Tags HkScan
// @Summary 删除HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "删除HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkScan/deleteHkScan [delete]
 export const deleteHkScan = (data) => {
     return service({
         url: "/hkScan/deleteHkScan",
         method: 'delete',
         data
     })
 }

// @Tags HkScan
// @Summary 删除HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkScan/deleteHkScan [delete]
 export const deleteHkScanByIds = (data) => {
     return service({
         url: "/hkScan/deleteHkScanByIds",
         method: 'delete',
         data
     })
 }

// @Tags HkScan
// @Summary 更新HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "更新HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkScan/updateHkScan [put]
 export const updateHkScan = (data) => {
     return service({
         url: "/hkScan/updateHkScan",
         method: 'put',
         data
     })
 }


// @Tags HkScan
// @Summary 用id查询HkScan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkScan true "用id查询HkScan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkScan/findHkScan [get]
 export const findHkScan = (params) => {
     return service({
         url: "/hkScan/findHkScan",
         method: 'get',
         params
     })
 }


// @Tags HkScan
// @Summary 分页获取HkScan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取HkScan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkScan/getHkScanList [get]
 export const getHkScanList = (params) => {
     return service({
         url: "/hkScan/getHkScanList",
         method: 'get',
         params
     })
 }