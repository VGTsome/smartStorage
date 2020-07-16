// 自动生成模板SmartStorageCabinet
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageCabinet struct {
      gorm.Model
      CabinetId  int `json:"cabinetId" form:"cabinetId" gorm:"column:cabinet_id;comment:'';type:int(10)"`
      CabinetName  string `json:"cabinetName" form:"cabinetName" gorm:"column:cabinet_name;comment:'';type:varchar(2000)"`
      CabinetPosition  string `json:"cabinetPosition" form:"cabinetPosition" gorm:"column:cabinet_position;comment:'';type:varchar(2000)"`
      CabinetTotalNumber  int `json:"cabinetTotalNumber" form:"cabinetTotalNumber" gorm:"column:cabinet_total_number;comment:'';type:int(10)"` 
}


func (SmartStorageCabinet) TableName() string {
  return "smart_storage_cabinet"
}
