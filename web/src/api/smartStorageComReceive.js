import service from '@/utils/request'

// @Tags SmartStorageComReceive
// @Summary 创建SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "创建SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageComReceive/createSmartStorageComReceive [post]
export const createSmartStorageComReceive = (data) => {
     return service({
         url: "/smartStorageComReceive/createSmartStorageComReceive",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageComReceive
// @Summary 删除SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "删除SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageComReceive/deleteSmartStorageComReceive [delete]
 export const deleteSmartStorageComReceive = (data) => {
     return service({
         url: "/smartStorageComReceive/deleteSmartStorageComReceive",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageComReceive
// @Summary 删除SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageComReceive/deleteSmartStorageComReceive [delete]
 export const deleteSmartStorageComReceiveByIds = (data) => {
     return service({
         url: "/smartStorageComReceive/deleteSmartStorageComReceiveByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageComReceive
// @Summary 更新SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "更新SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageComReceive/updateSmartStorageComReceive [put]
 export const updateSmartStorageComReceive = (data) => {
     return service({
         url: "/smartStorageComReceive/updateSmartStorageComReceive",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageComReceive
// @Summary 用id查询SmartStorageComReceive
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageComReceive true "用id查询SmartStorageComReceive"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageComReceive/findSmartStorageComReceive [get]
 export const findSmartStorageComReceive = (params) => {
     return service({
         url: "/smartStorageComReceive/findSmartStorageComReceive",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageComReceive
// @Summary 分页获取SmartStorageComReceive列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageComReceive列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageComReceive/getSmartStorageComReceiveList [get]
 export const getSmartStorageComReceiveList = (params) => {
     return service({
         url: "/smartStorageComReceive/getSmartStorageComReceiveList",
         method: 'get',
         params
     })
 }