import service from '@/utils/request'

// @Tags SmartStorageDoor
// @Summary 创建SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "创建SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageDoor/createSmartStorageDoor [post]
export const createSmartStorageDoor = (data) => {
     return service({
         url: "/smartStorageDoor/createSmartStorageDoor",
         method: 'post',
         data
     })
 }


// @Tags SmartStorageDoor
// @Summary 删除SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "删除SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageDoor/deleteSmartStorageDoor [delete]
 export const deleteSmartStorageDoor = (data) => {
     return service({
         url: "/smartStorageDoor/deleteSmartStorageDoor",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageDoor
// @Summary 删除SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageDoor/deleteSmartStorageDoor [delete]
 export const deleteSmartStorageDoorByIds = (data) => {
     return service({
         url: "/smartStorageDoor/deleteSmartStorageDoorByIds",
         method: 'delete',
         data
     })
 }

// @Tags SmartStorageDoor
// @Summary 更新SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "更新SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageDoor/updateSmartStorageDoor [put]
 export const updateSmartStorageDoor = (data) => {
     return service({
         url: "/smartStorageDoor/updateSmartStorageDoor",
         method: 'put',
         data
     })
 }


// @Tags SmartStorageDoor
// @Summary 用id查询SmartStorageDoor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageDoor true "用id查询SmartStorageDoor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageDoor/findSmartStorageDoor [get]
 export const findSmartStorageDoor = (params) => {
     return service({
         url: "/smartStorageDoor/findSmartStorageDoor",
         method: 'get',
         params
     })
 }


// @Tags SmartStorageDoor
// @Summary 分页获取SmartStorageDoor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageDoor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageDoor/getSmartStorageDoorList [get]
 export const getSmartStorageDoorList = (params) => {
     return service({
         url: "/smartStorageDoor/getSmartStorageDoorList",
         method: 'get',
         params
     })
 }