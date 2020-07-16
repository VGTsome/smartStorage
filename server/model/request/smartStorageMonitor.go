package request

import "gin-vue-admin/model"

type SmartStorageMonitorSearch struct{
    model.SmartStorageMonitor
    PageInfo
}