package request

import "gin-vue-admin/model"

type SmartStorageProductSearch struct{
    model.SmartStorageProduct
    PageInfo
}