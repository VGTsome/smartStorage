import service from '@/utils/request'

// @Tags SmartStorageCabinet
// @Summary 创建SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "创建SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCabinet/createSmartStorageCabinet [post]
export const createSmartStorageCabinet = (data) => {
     return service({
         url: "/smartStorageCabinet/createSmartStorageCabinet",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageCabinet
// @Summary 删除SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "删除SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCabinet/deleteSmartStorageCabinet [delete]
 export const deleteSmartStorageCabinet = (data) => {
     return service({
         url: "/smartStorageCabinet/deleteSmartStorageCabinet",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageCabinet
// @Summary 删除SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageCabinet/deleteSmartStorageCabinet [delete]
 export const deleteSmartStorageCabinetByIds = (data) => {
     return service({
         url: "/smartStorageCabinet/deleteSmartStorageCabinetByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageCabinet
// @Summary 更新SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "更新SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageCabinet/updateSmartStorageCabinet [put]
 export const updateSmartStorageCabinet = (data) => {
     return service({
         url: "/smartStorageCabinet/updateSmartStorageCabinet",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageCabinet
// @Summary 用id查询SmartStorageCabinet
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageCabinet true "用id查询SmartStorageCabinet"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageCabinet/findSmartStorageCabinet [get]
 export const findSmartStorageCabinet = (params) => {
     return service({
         url: "/smartStorageCabinet/findSmartStorageCabinet",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageCabinet
// @Summary 分页获取SmartStorageCabinet列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageCabinet列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageCabinet/getSmartStorageCabinetList [get]
 export const getSmartStorageCabinetList = (params) => {
     return service({
         url: "/smartStorageCabinet/getSmartStorageCabinetList",
         method: 'get',
         params
     })
 }