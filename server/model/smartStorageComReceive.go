// 自动生成模板SmartStorageComReceive
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageComReceive struct {
      gorm.Model
      Com  string `json:"com" form:"com" gorm:"column:com;comment:'';type:varchar(255)"`
      Command  string `json:"command" form:"command" gorm:"column:command;comment:'';type:varchar(255)"` 
}


func (SmartStorageComReceive) TableName() string {
  return "smart_storage_com_receive"
}
