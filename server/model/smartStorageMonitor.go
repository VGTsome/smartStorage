// 自动生成模板SmartStorageMonitor
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageMonitor struct {
	gorm.Model
	OrderId       int    `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:int(11)"`
	CabinetId     int    `json:"cabinetId" form:"cabinetId" gorm:"column:cabinet_id;comment:'';type:int(11)"`
	MonitorImg    string `json:"monitorImg" form:"monitorImg" gorm:"column:monitor_img;comment:'';type:varchar(2000)"`
	ItemNumber    int    `json:"itemNumber" form:"itemNumber" gorm:"column:item_number;comment:'';type:int(11)"`
	ProductWeight int    `json:"productWeight" form:"productWeight" gorm:"column:product_weight;comment:'';type:int(11)"`
}

func (SmartStorageMonitor) TableName() string {
	return "smart_storage_monitor"
}
