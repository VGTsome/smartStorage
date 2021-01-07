// 自动生成模板HkScan
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type HkScan struct {
      gorm.Model
      ScanId  string `json:"scanId" form:"scanId" gorm:"column:scan_id;comment:'员工扫描ID';type:varchar(255)"`
      ScanPic  string `json:"scanPic" form:"scanPic" gorm:"column:scan_pic;comment:'摄像图像';type:varchar(600)"`
      ScanStatus  string `json:"scanStatus" form:"scanStatus" gorm:"column:scan_status;comment:'摄像状态';type:varchar(255)"`
      DoorOpenStatus  string `json:"doorOpenStatus" form:"doorOpenStatus" gorm:"column:door_open_status;comment:'门开状态';type:varchar(255)"`
      DoorCloseStatus  string `json:"doorCloseStatus" form:"doorCloseStatus" gorm:"column:door_close_status;comment:'门关状态';type:varchar(255)"` 
}


func (HkScan) TableName() string {
  return "hk_scan"
}
