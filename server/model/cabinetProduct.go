// 自动生成模板CabinetProduct
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type CabinetProduct struct {
      gorm.Model
      CabinetId  int `json:"cabinetId" form:"cabinetId" gorm:"column:cabinet_id;comment:'';type:int(10)"`
      ProductId  int `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:int(10)"`
      Status  string `json:"status" form:"status" gorm:"column:status;comment:'';type:varchar(255)"` 
}


func (CabinetProduct) TableName() string {
  return "smart_storage_cabinet_product"
}
