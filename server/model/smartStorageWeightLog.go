// 自动生成模板SmartStorageWeightLog
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageWeightLog struct {
      gorm.Model
      OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:varchar(255)"`
      CabinetId  int `json:"cabinetId" form:"cabinetId" gorm:"column:cabinet_id;comment:'';type:int(10)"`
      ProductNumber  int `json:"productNumber" form:"productNumber" gorm:"column:product_number;comment:'';type:int(10)"`
      ProductWeight  int `json:"productWeight" form:"productWeight" gorm:"column:product_weight;comment:'';type:int(10)"` 
}


func (SmartStorageWeightLog) TableName() string {
  return "smart_storage_weight_log"
}
