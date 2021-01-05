import service from '@/utils/request'

// @Tags SmartStorageOrder
// @Summary 创建SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "创建SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageOrder/createSmartStorageOrder [post]
export const createSmartStorageOrder = (data) => {
	return service({
		url: '/smartStorageOrder/createSmartStorageOrder',
		method: 'post',
		data,
	})
}

// @Tags SmartStorageOrder
// @Summary 删除SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "删除SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageOrder/deleteSmartStorageOrder [delete]
export const deleteSmartStorageOrder = (data) => {
	return service({
		url: '/smartStorageOrder/deleteSmartStorageOrder',
		method: 'delete',
		data,
	})
}

// @Tags SmartStorageOrder
// @Summary 删除SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /smartStorageOrder/deleteSmartStorageOrder [delete]
export const deleteSmartStorageOrderByIds = (data) => {
	return service({
		url: '/smartStorageOrder/deleteSmartStorageOrderByIds',
		method: 'delete',
		data,
	})
}

// @Tags SmartStorageOrder
// @Summary 更新SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "更新SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /smartStorageOrder/updateSmartStorageOrder [put]
export const updateSmartStorageOrder = (data) => {
	return service({
		url: '/smartStorageOrder/updateSmartStorageOrder',
		method: 'put',
		data,
	})
}

// @Tags SmartStorageOrder
// @Summary 用id查询SmartStorageOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SmartStorageOrder true "用id查询SmartStorageOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /smartStorageOrder/findSmartStorageOrder [get]
export const findSmartStorageOrder = (params) => {
	return service({
		url: '/smartStorageOrder/findSmartStorageOrder',
		method: 'get',
		params,
	})
}

// @Tags SmartStorageOrder
// @Summary 分页获取SmartStorageOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SmartStorageOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /smartStorageOrder/getSmartStorageOrderList [get]
export const getSmartStorageOrderList = (params) => {
	return service({
		url: '/smartStorageOrder/GetSmartStorageOrderInfoById',
		method: 'get',
		params,
	})
}
export const getAllOrderList = (params) => {
	return service({
		url: '/smartStorageOrder/getAllOrderList',
		method: 'get',
		params,
	})
}
export const getMyOrderList = (params) => {
	return service({
		url: '/smartStorageOrder/getMyOrderList',
		method: 'get',
		params,
	})
}
export const updateSmartStorageOrderStatus = (data) => {
	return service({
		url: '/smartStorageOrder/updateSmartStorageOrderStatus',
		method: 'put',
		data,
	})
}
