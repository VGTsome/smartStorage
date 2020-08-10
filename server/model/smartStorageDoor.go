// 自动生成模板SmartStorageDoor
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageDoor struct {
	gorm.Model
	DoorId     int    `json:"doorId" form:"doorId" gorm:"column:door_id;comment:'';type:int(10)"`
	DoorName   string `json:"doorName" form:"doorName" gorm:"column:door_name;comment:'';type:varchar(2000)"`
	DoorStatus int    `json:"doorStatus" form:"doorStatus" gorm:"column:door_status;comment:'';type:tinyint(4)"`
}

func (SmartStorageDoor) TableName() string {
	return "smart_storage_door"
}
