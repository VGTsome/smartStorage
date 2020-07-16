package request

import "gin-vue-admin/model"

type SmartStorageOrderSearch struct{
    model.SmartStorageOrder
    PageInfo
}