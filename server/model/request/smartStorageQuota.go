package request

import "gin-vue-admin/model"

type SmartStorageQuotaSearch struct{
    model.SmartStorageQuota
    PageInfo
}