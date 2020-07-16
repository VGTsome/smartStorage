package request

import "gin-vue-admin/model"

type CabinetProductSearch struct{
    model.CabinetProduct
    PageInfo
}