package request

import "gin-vue-admin/model"

type SmartStoragePassWeightSearch struct{
    model.SmartStoragePassWeight
    PageInfo
}