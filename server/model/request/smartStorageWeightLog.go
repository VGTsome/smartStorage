package request

import "gin-vue-admin/model"

type SmartStorageWeightLogSearch struct{
    model.SmartStorageWeightLog
    PageInfo
}