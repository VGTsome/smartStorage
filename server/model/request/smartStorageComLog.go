package request

import "gin-vue-admin/model"

type SmartStorageComLogSearch struct{
    model.SmartStorageComLog
    PageInfo
}