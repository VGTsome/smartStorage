package request

import "gin-vue-admin/model"

type SmartStorageCabinetSearch struct{
    model.SmartStorageCabinet
    PageInfo
}