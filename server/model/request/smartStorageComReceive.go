package request

import "gin-vue-admin/model"

type SmartStorageComReceiveSearch struct{
    model.SmartStorageComReceive
    PageInfo
}