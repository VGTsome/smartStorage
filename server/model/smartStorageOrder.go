// 自动生成模板SmartStorageOrder
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type SmartStorageOrder struct {
	gorm.Model
	OrderId             string              `json:"orderId" form:"orderId" gorm:"column:order_id;comment:'';type:varchar(255)"`
	UserId              int                 `json:"userId" form:"userId" gorm:"column:user_id;comment:'';type:int(11)"`
	SysUser             SysUser             `gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	ProductId           string              `json:"productId" form:"productId" gorm:"column:product_id;comment:'';type:varchar(255)"`
	SmartStorageProduct SmartStorageProduct `gorm:"ForeignKey:ProductId;AssociationForeignKey:ProductId"`
	OrderNumber         int                 `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:'';type:int(11)"`
	OrderStatus         int                 `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:'';type:int(11)"`
}

func (SmartStorageOrder) TableName() string {
	return "smart_storage_order"
}
