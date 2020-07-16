package request

import "gin-vue-admin/model"

type SmartStorageDoorSearch struct{
    model.SmartStorageDoor
    PageInfo
}