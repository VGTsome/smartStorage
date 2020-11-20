package request

import "gin-vue-admin/model"

type SmartStorageCurrentOrderSearch struct{
    model.SmartStorageCurrentOrder
    PageInfo
}