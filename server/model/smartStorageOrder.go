// 自动生成模板SmartStorageOrder
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageOrder struct {
      gorm.Model
      OrderId  int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:int(10)"`
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;comment:'';type:int(10)"`
      ProductId  int `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:int(10)"`
      OrderStatus  int `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:'';type:int(10)"` 
}


func (SmartStorageOrder) TableName() string {
  return "smart_storage_order"
}
