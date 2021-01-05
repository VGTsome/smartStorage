import service from '@/utils/request'

// @Tags CabinetProduct
// @Summary 创建CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "创建CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sscp/createCabinetProduct [post]
export const createCabinetProduct = (data) => {
	return service({
		url: '/sscp/createCabinetProduct',
		method: 'post',
		data,
	})
}
export const prepareShelf = (data) => {
	return service({
		url: '/sscp/prepareShelf',
		method: 'post',
		data,
	})
}
export const updateShelf = (data) => {
	return service({
		url: '/sscp/updateShelf',
		method: 'post',
		data,
	})
}

// @Tags CabinetProduct
// @Summary 删除CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "删除CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sscp/deleteCabinetProduct [delete]
export const deleteCabinetProduct = (data) => {
	return service({
		url: '/sscp/deleteCabinetProduct',
		method: 'delete',
		data,
	})
}

// @Tags CabinetProduct
// @Summary 删除CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sscp/deleteCabinetProduct [delete]
export const deleteCabinetProductByIds = (data) => {
	return service({
		url: '/sscp/deleteCabinetProductByIds',
		method: 'delete',
		data,
	})
}

// @Tags CabinetProduct
// @Summary 更新CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "更新CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sscp/updateCabinetProduct [put]
export const updateCabinetProduct = (data) => {
	return service({
		url: '/sscp/updateCabinetProduct',
		method: 'put',
		data,
	})
}

// @Tags CabinetProduct
// @Summary 用id查询CabinetProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CabinetProduct true "用id查询CabinetProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sscp/findCabinetProduct [get]
export const findCabinetProduct = (params) => {
	return service({
		url: '/sscp/findCabinetProduct',
		method: 'get',
		params,
	})
}

// @Tags CabinetProduct
// @Summary 分页获取CabinetProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CabinetProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sscp/getCabinetProductList [get]
export const getCabinetProductList = (params) => {
	return service({
		url: '/sscp/getCabinetProductList',
		method: 'get',
		params,
	})
}
