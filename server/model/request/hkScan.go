package request

import "gin-vue-admin/model"

type HkScanSearch struct{
    model.HkScan
    PageInfo
}