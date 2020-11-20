// 自动生成模板SmartStorageCurrentOrder
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageCurrentOrder struct {
	gorm.Model
	OrderId     string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:varchar(255)"`
	UserId      uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:'';type:int(10)"`
	ProductId   string `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:varchar(255)"`
	ProductName string `json:"productName" form:"productName" gorm:"column:product_name;comment:'';type:varchar(255)"`
	CabinetList string `json:"cabinetList" form:"cabinetList" gorm:"column:cabinet_list;comment:'';type:varchar(255)"`
	OrderNumber int    `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:'';type:int(10)"`
	OrderStatus int    `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:'';type:int(10)"`
}

func (SmartStorageCurrentOrder) TableName() string {
	return "smart_storage_current_order"
}
