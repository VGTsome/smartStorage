// 自动生成模板SmartStorageComLog
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageComLog struct {
	gorm.Model
	Func       string `json:"func" form:"func" gorm:"column:func;comment:'';type:varchar(255)"`
	Isret      string `json:"isret" form:"isret" gorm:"column:isret;comment:'';type:varchar(255)"`
	Tatoo      string `json:"tatoo" form:"tatoo" gorm:"column:tatoo;comment:'';type:varchar(255)"`
	SendCmd    string `json:"sendCmd" form:"sendCmd" gorm:"column:sendCmd;comment:'';type:varchar(255)"`
	ReceiveCmd string `json:"receiveCmd" form:"receiveCmd" gorm:"column:receiveCmd;comment:'';type:varchar(255)"`
}

func (SmartStorageComLog) TableName() string {
	return "smart_storage_com_log"
}
