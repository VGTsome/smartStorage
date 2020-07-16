// 自动生成模板SmartStorageMonitor
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageMonitor struct {
      gorm.Model
      OrderId  int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:int(10)"`
      MonitorId  int `json:"monitorId" form:"monitorId" gorm:"column:monitor_id;comment:'';type:int(10)"`
      MonitorImg  string `json:"monitorImg" form:"monitorImg" gorm:"column:monitor_img;comment:'';type:varchar(2000)"`
      OrderStatus  int `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:'';type:int(10)"` 
}


func (SmartStorageMonitor) TableName() string {
  return "smart_storage_monitor"
}
