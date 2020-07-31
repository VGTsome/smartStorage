package request

import "gin-vue-admin/model"

type SmartStorageOrderSearch struct {
	model.SmartStorageOrder
	PageInfo
}

type SmartStorageOrderListReq struct {
	SmartStorageOrderList []model.SmartStorageOrder `json:"fdata" form:"fdata"`
}
