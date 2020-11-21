package request

import "gin-vue-admin/model"

type SmartStorageSystemStatusSearch struct{
    model.SmartStorageSystemStatus
    PageInfo
}