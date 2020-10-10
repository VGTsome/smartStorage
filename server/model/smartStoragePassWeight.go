// 自动生成模板SmartStoragePassWeight
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStoragePassWeight struct {
      gorm.Model
      OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:varchar(255)"`
      CabinetId  int `json:"cabinetId" form:"cabinetId" gorm:"column:cabinet_id;comment:'';type:int(10)"`
      ProductNumber  int `json:"productNumber" form:"productNumber" gorm:"column:product_number;comment:'';type:int(10)"`
      ProductWeight  int `json:"productWeight" form:"productWeight" gorm:"column:product_weight;comment:'';type:int(10)"`
      Pass  int `json:"pass" form:"pass" gorm:"column:pass;comment:'0进1出';type:int(10)"` 
}


func (SmartStoragePassWeight) TableName() string {
  return "smart_storage_pass_weight"
}
