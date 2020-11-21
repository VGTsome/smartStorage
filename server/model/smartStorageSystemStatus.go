// 自动生成模板SmartStorageSystemStatus
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageSystemStatus struct {
      gorm.Model
      SystemStatus  int `json:"systemStatus" form:"systemStatus" gorm:"column:system_status;comment:'1:正常 2:管理员理货 3:取货中 4:异常状态';type:int(10)"` 
}


func (SmartStorageSystemStatus) TableName() string {
  return "smart_storage_system_status"
}
